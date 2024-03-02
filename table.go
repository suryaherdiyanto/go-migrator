package gomigrator

type MysqlDataType struct {
	Type          string
	Size          int
	Default       interface{}
	Nullable      bool
	AutoIncrement bool
}

type MysqlColumn struct {
	Name     string
	Property MysqlDataType
}

type MysqlTable struct {
	Name    string
	Columns []MysqlColumn
}
