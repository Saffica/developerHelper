package client

type record struct {
	FileName  string
	Value     string
	TableName string
	SysID     string
}

type recordReponse struct {
	Data []map[string]interface{} `json:"data"`
}

type table struct {
	Column   []string
	RecordID []string
}

type vcsResponse struct {
	Data []vcsDataResponse `json:"data"`
}

type vcsDataResponse struct {
	TableName string `json:"table_name"`
	RecordID  string `json:"record_id"`
}

type tableResponse struct {
	Data []tableDataResponse `json:"data"`
}
type tableDataResponse struct {
	TableName string `json:"name"`
	TableID   string `json:"sys_id"`
}

type scriptColumnResponse struct {
	Data []scriptColumnData `json:"data"`
}

type scriptColumnData struct {
	ColumnName string  `json:"column_name"`
	TableID    tableID `json:"table_id"`
}

type tableID struct {
	Value string `json:"value"`
}
