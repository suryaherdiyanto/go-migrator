package gomigrator

import (
	"bytes"
	"io"
	"slices"
	"text/template"
)

type SQLDataType string

const (
	VARCHAR   SQLDataType = "varchar"
	BIGINT    SQLDataType = "bigint"
	INT       SQLDataType = "int"
	TINYINT   SQLDataType = "tinyint"
	MEDIUMINT SQLDataType = "mediumint"
	BOOL      SQLDataType = "bool"
	FLOAT     SQLDataType = "float"
	DOUBLE    SQLDataType = "double"
	TEXT      SQLDataType = "text"
	DATE      SQLDataType = "date"
	DATETIME  SQLDataType = "datetime"
)

type MysqlDataType struct {
	Type          SQLDataType
	Size          int
	Default       interface{}
	Nullable      bool
	AutoIncrement bool
	Unsigned      bool
}

type MysqlColumn struct {
	Name     string
	Property *MysqlDataType
}

type MysqlTable struct {
	Name    string
	Columns []MysqlColumn
}

func (c *MysqlColumn) ParseColumn() (string, error) {
	col := &MysqlColumn{
		Name:     c.Name,
		Property: c.Property,
	}

	w := new(bytes.Buffer)
	err := ParseColumnTemplate(w, col)
	if err != nil {
		return "", err
	}

	return w.String(), nil
}

func CreateColumn(columnName string, property *MysqlDataType) *MysqlColumn {
	return &MysqlColumn{
		Name:     columnName,
		Property: property,
	}
}

func ParseColumnTemplate(w io.Writer, data *MysqlColumn) error {
	templatePath := "./template/types/" + string(data.Property.Type) + ".go.tmpl"
	templateName := string(data.Property.Type) + ".go.tmpl"

	if IsIntegerColumn(data.Property.Type) {
		templatePath = "./template/types/int.go.tmpl"
		templateName = "int.go.tmpl"
	}

	tmpl, err := template.New(templateName).ParseFiles(templatePath)

	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		return err
	}

	return nil
}

func IsIntegerColumn(t SQLDataType) bool {
	types := []SQLDataType{
		INT,
		TINYINT,
		MEDIUMINT,
		BIGINT,
		BOOL,
	}

	return slices.Index(types, t) >= 0
}
