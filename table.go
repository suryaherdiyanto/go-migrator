package gomigrator

import (
	"io"
	"slices"
	"strings"
	"text/template"
)

type SQLDataType string

type MysqlDataType struct {
	Type          SQLDataType
	Size          int
	Default       interface{}
	Nullable      bool
	AutoIncrement bool
	Unsigned      bool
	EnumOptions   []string
	Unique        bool
	PrimaryKey    bool
	Precision     int
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

func CreateTableFunc(name string, tableColumns func(cols *MysqlColumns)) *MysqlTable {
	columns := NewMysqlColumns()
	tableColumns(columns)

	return &MysqlTable{
		Name:    name,
		Columns: columns.Columns,
	}
}

func CreateIndex(indexName string, table string, columns []string) string {
	return "CREATE INDEX " + indexName + " ON " + table + "(" + strings.Join(columns, ", ") + ");"
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
