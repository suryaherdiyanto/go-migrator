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

type ForeignKeyOptions struct {
	OnDelete string
	OnUpdate string
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

func (t *Table) ForeignKey(column string, reference string, options *ForeignKeyOptions) string {
	stmt := "ALTER TABLE " + t.Name + " ADD CONSTRAINT fk_" + t.Name + "_" + reference + " FOREIGN KEY (" + column + ") REFERENCES " + reference + ";"

	if options != nil {
		if options.OnDelete != "" {
			stmt += " ON DELETE " + options.OnDelete
		}

		if options.OnUpdate != "" {
			stmt += " ON UPDATE " + options.OnUpdate
		}
	}

	stmt += ";"

	return stmt
}

func parseTableTemplate(t *Table) string {
	stmt := "CREATE TABLE IF NOT EXISTS"
	stmt += " " + t.Name + "("
	for i, column := range t.Columns {
		stmt += column.ParseColumn() + IfNe(i, t.ColumnLength(), ",")
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
