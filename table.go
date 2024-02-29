package gomigrator

type MysqlDataType struct {
	Name   string
	Length interface{}
}

type MysqlTable struct {
	Name string
	Type MysqlDataType
}
