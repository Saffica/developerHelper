package client

type Record struct {
	FileName  string
	Value     string
	TableName string
	SysID     string
}

type recordReponse struct {
	Data []map[string]interface{} `json:"data"`
}

type Table struct {
	Column   []string
	RecordID []string
}
