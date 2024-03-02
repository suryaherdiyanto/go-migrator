package gomigrator

import (
	"bytes"
	"io"
	"text/template"
)

type MysqlDataType struct {
	Type          string
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
	tmpl, err := template.New(data.Property.Type + ".go.tmpl").ParseFiles("./template/types/" + data.Property.Type + ".go.tmpl")

	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		return err
	}

	return nil
}
