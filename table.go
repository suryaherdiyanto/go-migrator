package gomigrator

import (
	"io"
	"slices"
	"strings"
	"text/template"
)

type SQLDataType string

const (
	VARCHAR   SQLDataType = "varchar"
	CHAR      SQLDataType = "char"
	BIGINT    SQLDataType = "bigint"
	INT       SQLDataType = "int"
	TINYINT   SQLDataType = "tinyint"
	MEDIUMINT SQLDataType = "mediumint"
	BOOL      SQLDataType = "bool"
	FLOAT     SQLDataType = "float"
	DOUBLE    SQLDataType = "double"
	TEXT      SQLDataType = "text"
	DATE      SQLDataType = "date"
	ENUM      SQLDataType = "enum"
	DATETIME  SQLDataType = "datetime"
	TIMESTAMP SQLDataType = "timestamp"
)

type MysqlDataType struct {
	Type          SQLDataType
	Size          int
	Default       interface{}
	Nullable      bool
	AutoIncrement bool
	Unsigned      bool
	EnumOptions   []string
}

type MysqlTable struct {
	Name    string
	Columns []MysqlColumn
}

func (mt *MysqlTable) ColumnLength() int {
	return len(mt.Columns) - 1
}

func (o *MysqlDataType) PrintEnumValues() string {
	return "'" + strings.Join(o.EnumOptions, "', '") + "'"
}

func CreateTable(name string, tableColumns func() []MysqlColumn) *MysqlTable {
	return &MysqlTable{
		Name:    name,
		Columns: tableColumns(),
	}
}

func parseTableTemplate(w io.Writer, data *MysqlTable) error {
	templatePath := "./template/mysql/create-table.go.tmpl"
	return parseTemplate(w, data, "create-table.go.tmpl", templatePath)
}

func parseTemplate(w io.Writer, data any, name string, path string) error {
	tmpl, err := template.New(name).ParseFiles(path)

	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}

func IsNumericColumn(t SQLDataType) bool {
	types := []SQLDataType{
		INT,
		TINYINT,
		MEDIUMINT,
		BIGINT,
		BOOL,
		FLOAT,
		DOUBLE,
	}

	return slices.Index(types, t) >= 0
}

func IsTextColumn(t SQLDataType) bool {
	types := []SQLDataType{
		CHAR,
		VARCHAR,
		TEXT,
		DATE,
		DATETIME,
	}

	return slices.Index(types, t) >= 0
}
