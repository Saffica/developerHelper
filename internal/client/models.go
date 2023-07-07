package client

type SimpleColumn struct {
	name string
}

type SimpleTable struct {
	name    string
	columns []SimpleColumn
}

func CreateColumn(n string) SimpleColumn {
	newColumn := SimpleColumn{name: n}
	return newColumn
}

func CreateTable(n string, c []SimpleColumn) SimpleTable {
	newTable := SimpleTable{name: n, columns: c}
	return newTable
}
