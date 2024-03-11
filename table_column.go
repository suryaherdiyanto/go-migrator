package gomigrator

import (
	"bytes"
	"io"
	"slices"
	"strings"
)

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

type MysqlColumn struct {
	Name     string
	Property *MysqlDataType
}

type MysqlColumns struct {
	Columns []MysqlColumn
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

func NewMysqlColumns() *MysqlColumns {
	return &MysqlColumns{}
}

func (c *MysqlColumns) Varchar(name string, length int, props *MysqlDataType) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       VARCHAR,
			Size:       length,
			Unique:     props.Unique,
			Nullable:   props.Nullable,
			PrimaryKey: props.PrimaryKey,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Char(name string, length int, props *MysqlDataType) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       CHAR,
			Size:       length,
			Unique:     props.Unique,
			Nullable:   props.Nullable,
			PrimaryKey: props.PrimaryKey,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Integer(name string, length int, props *MysqlDataType) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:          INT,
			Unique:        props.Unique,
			Nullable:      props.Nullable,
			Unsigned:      props.Unsigned,
			AutoIncrement: props.AutoIncrement,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func CreateColumn(columnName string, property *MysqlDataType) *MysqlColumn {
	return &MysqlColumn{
		Name:     columnName,
		Property: property,
	}
}
