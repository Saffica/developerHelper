package vcsrepository

type vcsResponse struct {
	Data []vcsDataResponse `json:"data"`
}

type vcsDataResponse struct {
	TableName string `json:"table_name"`
	RecordID  string `json:"record_id"`
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

type tableResponse struct {
	Data []tableDataResponse `json:"data"`
}
type tableDataResponse struct {
	TableName string `json:"name"`
	TableID   string `json:"sys_id"`
}
