package gomigrator

type MysqlDataType struct {
	Type string
	Size interface{}
}

type MysqlTable struct {
	Name string
	Type MysqlDataType
}
