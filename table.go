package gomigrator

import (
	"database/sql"
	"io"
	"slices"
	"strings"
	"text/template"
)

type SQLDataType string

type SQLDialect string

type SQLTableProp struct {
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
	Dialect       SQLDialect
}

const (
	POSTGRES SQLDialect = "postgres"
	MYSQL    SQLDialect = "mysql"
)

type Table struct {
	Name           string
	Columns        []TableColumn
	EnumStatements []string
}

func (mt *Table) ColumnLength() int {
	return len(mt.Columns) - 1
}

func (o *SQLTableProp) PrintEnumValues() string {
	return "'" + strings.Join(o.EnumOptions, "', '") + "'"
}

func CreateTable(name string, tableColumns func(table *Table)) *Table {
	table := &Table{Name: name}
	tableColumns(table)

	return table
}

func CreateIndex(indexName string, table string, columns []string) string {
	return "CREATE INDEX " + indexName + " ON " + table + "(" + strings.Join(columns, ", ") + ");"
}

func (t *Table) Run(db *sql.DB) error {
	buff := new(strings.Builder)
	err := parseTableTemplate(buff, t)

	if err != nil {
		return err
	}

	if len(t.EnumStatements) > 0 {
		for _, enum := range t.EnumStatements {
			_, err = db.Exec(enum)

			if err != nil {
				return err
			}
		}
	}
	_, err = db.Exec(buff.String())
	defer db.Close()

	return err
}

func (t *Table) CreateEnum(name string, options []string) string {
	return "DROP TYPE IF EXISTS " + name + "; CREATE TYPE " + name + " AS ENUM('" + strings.Join(options, "', '") + "');"
}

func parseTableTemplate(w io.Writer, data *Table) error {
	templatePath := "./template/mysql/create-table.go.tmpl"
	return parseTemplate(w, data, "create-table.go.tmpl", templatePath)
}

func (t *Table) AddColumn(name string, props SQLTableProp) {
	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &props,
	})
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
