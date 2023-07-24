package client

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"

	"main/internal/model"
)

var (
	tr     = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client = &http.Client{Transport: tr, Timeout: 10 * time.Second}
)

type app struct {
	cfg        *viper.Viper
	authHeader string
}

func New(cfg *viper.Viper) *app {
	return &app{
		cfg:        cfg,
		authHeader: fmt.Sprintf("Bearer %s", cfg.GetString("TOKEN")),
	}
}

func (a *app) GetRecordsByVCS(vcs map[string]model.Table) ([]model.Record, error) {
	r := []model.Record{}

	for tableName, v := range vcs {
		if len(v.RecordID) == 0 {
			continue
		}

		p := fmt.Sprintf("sysparm_query=sys_idIN%s&sysparm_fields=%s&sysparm_limit=0", strings.Join(v.RecordID, "@"), strings.Join(v.Column, ","))
		body, err := a.FetchData(tableName, p)
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

func (a *app) FetchData(t string, p string) ([]byte, error) {
	url := fmt.Sprintf("%s/rest/v1/table/%s?%s", a.cfg.GetString("INSTANCE_URL"), t, p)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", a.authHeader)
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

func prepareRecords(inputRecord map[string]interface{}, columns []string, tableName string) []model.Record {
	m := map[string]string{}
	var sysID string
	r := []model.Record{}

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
		extension := geteFileExtension(tableName, columnName)
		r = append(r, model.Record{
			FileName:  getFileName(sysID, tableName, columnName, extension),
			Value:     m[columnName],
			TableName: tableName,
			SysID:     sysID,
		})
	}
	return r
}

func getFileName(sysID string, tableName string, columnName string, extension string) string {
	tableWithShortName := []string{"sys_widget"}

	for _, v := range tableWithShortName {
		if v == tableName {
			return fmt.Sprintf("%s.%s", columnName, extension)
		}
	}

	return fmt.Sprintf("%s_%s.%s", sysID, columnName, extension)
}

func geteFileExtension(tableName string, columnName string) string {
	key := fmt.Sprintf("%s.%s", tableName, columnName)
	extensions := map[string]string{
		"sys_widget.css":      "css",
		"sys_widget.template": "html",
	}

	if entry, ok := extensions[key]; ok {
		return entry
	}

	return "js"
}
