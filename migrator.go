package gomigrator

func Run() {

}

func CreateTable(tableName string, columns []MysqlColumn) *MysqlTable {
	return &MysqlTable{
		Name:    tableName,
		Columns: columns,
	}
}

func CreateColumn(columnName string, property MysqlDataType) *MysqlColumn {
	return &MysqlColumn{
		Name:     columnName,
		Property: property,
	}
}
