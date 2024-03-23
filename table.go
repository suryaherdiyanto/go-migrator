package gomigrator

import (
	"database/sql"
	"strings"
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
}

const (
	POSTGRES SQLDialect = "postgres"
	MYSQL    SQLDialect = "mysql"
)

type Table struct {
	Name           string
	Columns        []TableColumn
	EnumStatements []string
	Dialect        SQLDialect
}

func (mt *Table) ColumnLength() int {
	return len(mt.Columns) - 1
}

func (o *SQLTableProp) PrintEnumValues() string {
	return "'" + strings.Join(o.EnumOptions, "', '") + "'"
}

func CreateTable(name string, tableColumns func(table *Table), dialect SQLDialect) *Table {
	table := &Table{Name: name, Dialect: dialect}
	tableColumns(table)

	return table
}

func CreateIndex(indexName string, table string, columns []string) string {
	return "CREATE INDEX " + indexName + " ON " + table + "(" + strings.Join(columns, ", ") + ");"
}

func (t *Table) Run(db *sql.DB) error {
	stmt := parseTableTemplate(t)

	if len(t.EnumStatements) > 0 {
		for _, enum := range t.EnumStatements {
			_, err := db.Exec(enum)

			if err != nil {
				return err
			}
		}
	}

	_, err := db.Exec(stmt)

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

func (t *Table) CreateEnum(name string, options []string) string {
	return "DROP TYPE IF EXISTS " + name + "; CREATE TYPE " + name + " AS ENUM('" + strings.Join(options, "', '") + "');"
}

func parseTableTemplate(data *Table) string {
	stmt := "CREATE TABLE IF NOT EXISTS"
	stmt += " " + data.Name + "("
	for i, column := range data.Columns {
		stmt += column.ParseColumn() + IfNe(i != data.ColumnLength(), ",")
	}

	stmt = stmt + ")"

	return stmt
}

func (t *Table) AddColumn(name string, props SQLTableProp) {
	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &props,
	})
}
