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
	ReferenceTable  string
	ReferenceColumn string
	OnDelete        string
	OnUpdate        string
}

const (
	POSTGRES SQLDialect = "postgres"
	MYSQL    SQLDialect = "mysql"
)

type Table struct {
	Name                 string
	Columns              []TableColumn
	EnumStatements       []string
	ForeignKeyStatements []string
	IndexStatements      []string
	Dialect              SQLDialect
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

func (t *Table) CreateIndex(columns []string) {
	indexName := strings.Join(columns, "_")
	stmt := "CREATE INDEX " + t.Name + "_" + indexName + "_idx" + " ON " + t.Name + "(" + strings.Join(columns, ", ") + ");"

	t.IndexStatements = append(t.IndexStatements, stmt)
}

func (t *Table) Run(db *sql.DB) error {
	stmt := parseTableTemplate(t)

	if len(t.EnumStatements) > 0 {
		execStatements(t.EnumStatements, db)
	}

	_, err := db.Exec(stmt)

	if len(t.ForeignKeyStatements) > 0 {
		execStatements(t.ForeignKeyStatements, db)
	}

	if len(t.IndexStatements) > 0 {
		execStatements(t.IndexStatements, db)
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *Table) CreateEnum(name string, options []string) string {
	return "DROP TYPE IF EXISTS " + name + "; CREATE TYPE " + name + " AS ENUM('" + strings.Join(options, "', '") + "');"
}

func (t *Table) ForeignKey(column string, options *ForeignKeyOptions) {
	stmt := "ALTER TABLE " + t.Name + " ADD FOREIGN KEY (" + column + ") REFERENCES " + options.ReferenceTable + "(" + options.ReferenceColumn + ")"

	if options.OnDelete != "" {
		stmt += " ON DELETE " + options.OnDelete
	}

	if options.OnUpdate != "" {
		stmt += " ON UPDATE " + options.OnUpdate
	}

	stmt += ";"

	t.ForeignKeyStatements = append(t.ForeignKeyStatements, stmt)
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

func execStatements(statements []string, db *sql.DB) error {
	for _, stmt := range statements {
		_, err := db.Exec(stmt)

		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Table) AddColumn(name string, props SQLTableProp) {
	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &props,
	})
}
