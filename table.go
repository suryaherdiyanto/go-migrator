package gomigrator

type MysqlDataType struct {
	Name string
	Size interface{}
}

type MysqlTable struct {
	Name string
	Type MysqlDataType
}
