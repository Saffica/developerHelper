package model

type Record struct {
	FileName  string
	Value     string
	TableName string
	SysID     string
}
type Table struct {
	Column   []string
	RecordID []string
}
