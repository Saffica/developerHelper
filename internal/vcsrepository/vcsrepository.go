package vcsrepository

import (
	"encoding/json"
	"fmt"
	"main/internal/model"
)

type Client interface {
	FetchData(t string, p string) ([]byte, error)
}

type Cfg interface {
	GetString(string) string
}

type app struct {
	client Client
	cfg    Cfg
}

func New(cfg Cfg, client Client) *app {
	return &app{
		client: client,
		cfg:    cfg,
	}
}

func (a *app) FindByLocalPackID(localPackID string) (map[string]model.Table, error) {
	p := fmt.Sprintf("sysparm_query=is_current=1^local_pack_id=%s&sysparm_fields=table_name,record_id&sysparm_limit=0", localPackID)
	body, err := a.client.FetchData("sys_vcs_record", p)
	if err != nil {
		return nil, err
	}

	vcs := vcsResponse{}
	err = json.Unmarshal(body, &vcs)
	if err != nil {
		return nil, err
	}

	r, err := a.prepareVCS(vcs.Data)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (a *app) prepareVCS(vcs []vcsDataResponse) (map[string]model.Table, error) {
	c, err := a.getScriptedColumns()
	if err != nil {
		return nil, err
	}

	r := make(map[string]model.Table, len(c))

	for k, v := range c {
		if _, ok := r[k]; !ok {
			r[k] = model.Table{Column: append(v, "sys_id")}
		}
	}

	for _, v := range vcs {
		if entry, ok := r[v.TableName]; ok {
			entry.RecordID = append(entry.RecordID, v.RecordID)
			r[v.TableName] = entry
		}
	}

	return r, nil
}

func (a *app) getScriptedColumns() (map[string][]string, error) {
	body, err := a.client.FetchData("sys_db_column", "sysparm_fields=column_name,table_id&sysparm_query=column_type_id=29&sysparm_limit=0")
	if err != nil {
		return nil, err
	}

	tables, err := a.getTables()
	if err != nil {
		return nil, err
	}

	scr := scriptColumnResponse{}
	err = json.Unmarshal(body, &scr)
	if err != nil {
		return nil, err
	}

	r := map[string][]string{
		"sys_widget":    {"template", "css"},
		"sys_ui_action": {"condition"},
	}
	for i := range scr.Data {
		tableName := tables[scr.Data[i].TableID.Value]
		r[tableName] = append(r[tableName], scr.Data[i].ColumnName)
	}
	return r, nil

}

func (a *app) getTables() (map[string]string, error) {
	body, err := a.client.FetchData("sys_db_table", "sysparm_fields=sys_id,name&sysparm_limit=0")
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
