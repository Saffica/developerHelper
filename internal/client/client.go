package client

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"

	"main/pkg/config"
)

const (
	pathToConfig = "./configs/config.json"
)

var (
	cfg                 = getConfig(pathToConfig)
	authorizationHeader = getAuthorizationHeader(cfg)
	tr                  = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client              = &http.Client{Transport: tr, Timeout: 10 * time.Second}
)

func GetRecordsByVCS(vcs map[string]table) ([]record, error) {
	r := []record{}

	for tableName, v := range vcs {
		if len(v.RecordID) == 0 {
			continue
		}

		p := fmt.Sprintf("sysparm_query=sys_idIN%s&sysparm_fields=%s&sysparm_limit=0", strings.Join(v.RecordID, "@"), strings.Join(v.Column, ","))
		body, err := fetchData(tableName, p)
		if err != nil {
			return nil, err
		}

		currentStruct := recordReponse{}
		err = json.Unmarshal(body, &currentStruct)
		if err != nil {
			return nil, err
		}

		for i := range currentStruct.Data {
			r = append(r, prepareRecords(currentStruct.Data[i], v.Column, tableName)...)
		}
	}
	return r, nil

}

func GetVCS() ([]vcsDataResponse, error) {
	p := fmt.Sprintf("sysparm_query=is_current=1^local_pack_id=%s&sysparm_fields=table_name,record_id&sysparm_limit=0", cfg.GetString("LOCAL_PACK_ID"))
	body, err := fetchData("sys_vcs_record", p)
	if err != nil {
		return nil, err
	}

	vcs := vcsResponse{}
	err = json.Unmarshal(body, &vcs)
	if err != nil {
		return nil, err
	}

	return vcs.Data, nil
}

func GetScriptedColumns() (map[string][]string, error) {
	body, err := fetchData("sys_db_column", "sysparm_fields=column_name,table_id&sysparm_query=column_type_id=29&sysparm_limit=0")
	if err != nil {
		return nil, err
	}

	tables, err := getTables()
	if err != nil {
		return nil, err
	}

	scr := scriptColumnResponse{}
	err = json.Unmarshal(body, &scr)
	if err != nil {
		return nil, err
	}

	r := map[string][]string{
		"sys_widget":    {"template"},
		"sys_ui_action": {"condition"},
	}
	for i := range scr.Data {
		tableName := tables[scr.Data[i].TableID.Value]
		r[tableName] = append(r[tableName], scr.Data[i].ColumnName)
	}
	return r, nil

}

func getTables() (map[string]string, error) {
	body, err := fetchData("sys_db_table", "sysparm_fields=sys_id,name&sysparm_limit=0")
	if err != nil {
		return nil, err
	}

	tds := tableResponse{}
	err = json.Unmarshal(body, &tds)
	if err != nil {
		return nil, err
	}

	r := map[string]string{}
	for _, val := range tds.Data {
		r[val.TableID] = val.TableName
	}
	return r, nil
}

func fetchData(t string, p string) ([]byte, error) {
	url := fmt.Sprintf("%s/rest/v1/table/%s?%s", cfg.GetString("INSTANCE_URL"), t, p)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", authorizationHeader)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getConfig(p string) *viper.Viper {
	cfg, err := config.LoadConfig(p)
	if err != nil {
		log.Panic(err.Error())
	}
	return cfg
}

func getAuthorizationHeader(c *viper.Viper) string {
	authorizationHeader := fmt.Sprintf("Bearer %s", c.GetString("TOKEN"))
	return authorizationHeader
}

func prepareRecords(inputRecord map[string]interface{}, columns []string, tableName string) []record {
	m := map[string]string{}
	var sysID string
	r := []record{}

	for j := range columns {
		if columns[j] == "sys_id" {
			sysID = inputRecord[columns[j]].(string)
			continue
		}

		if entry, ok := inputRecord[columns[j]]; ok {
			m[columns[j]] = entry.(string)
		}
	}

	for columnName := range m {
		r = append(r, record{
			FileName:  fmt.Sprintf("%s_%s.js", sysID, columnName),
			Value:     m[columnName],
			TableName: tableName,
			SysID:     sysID,
		})
	}
	return r
}
