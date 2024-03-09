package gomigrator

import (
	"bytes"
	"io"
	"strings"
)

type MysqlColumn struct {
	Name     string
	Property *MysqlDataType
}

func (c *MysqlColumn) ParseColumn() (string, error) {
	col := &MysqlColumn{
		Name:     c.Name,
		Property: c.Property,
	}

	w := new(bytes.Buffer)
	err := parseColumnTemplate(w, col)
	if err != nil {
		return "", err
	}

	return strings.Replace(w.String(), "\n", "", 1), nil
}

func parseColumnTemplate(w io.Writer, data *MysqlColumn) error {
	templatePath := "./template/types/" + string(data.Property.Type) + ".go.tmpl"
	templateName := string(data.Property.Type) + ".go.tmpl"

	if IsNumericColumn(data.Property.Type) {
		templatePath = "./template/types/int.go.tmpl"
		templateName = "int.go.tmpl"
	}

	if IsTextColumn(data.Property.Type) {
		templatePath = "./template/types/varchar.go.tmpl"
		templateName = "varchar.go.tmpl"
	}

	return parseTemplate(w, data, templateName, templatePath)
}

func CreateColumn(columnName string, property *MysqlDataType) *MysqlColumn {
	return &MysqlColumn{
		Name:     columnName,
		Property: property,
	}
}
