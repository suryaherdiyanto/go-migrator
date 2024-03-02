package gomigrator

type MysqlDataType struct {
	Name          string
	Type          string
	Size          interface{}
	Default       interface{}
	Nullable      bool
	AutoIncrement bool
}

type MysqlTable struct {
	Attributes []MysqlDataType
}
