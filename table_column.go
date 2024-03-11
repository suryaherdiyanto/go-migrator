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

type TextColumnProps struct {
	Unique     bool
	Nullable   bool
	Default    interface{}
	PrimaryKey bool
	Size       int
}

type NumericColumnProps struct {
	Unique        bool
	Nullable      bool
	Default       interface{}
	PrimaryKey    bool
	Unsigned      bool
	AutoIncrement bool
	Precision     int
	Size          int
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

func (c *MysqlColumns) Varchar(name string, length int, props *TextColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       VARCHAR,
			Size:       length,
			Unique:     props.Unique,
			Nullable:   props.Nullable,
			PrimaryKey: props.PrimaryKey,
			Default:    props.Default,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Char(name string, length int, props *TextColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       CHAR,
			Size:       length,
			Unique:     props.Unique,
			Nullable:   props.Nullable,
			PrimaryKey: props.PrimaryKey,
			Default:    props.Default,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Text(name string, props *TextColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       TEXT,
			Unique:     props.Unique,
			Nullable:   props.Nullable,
			PrimaryKey: props.PrimaryKey,
			Default:    props.Default,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Date(name string, props *TextColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       DATE,
			Nullable:   props.Nullable,
			PrimaryKey: props.PrimaryKey,
			Default:    props.Default,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Timestamp(name string, props *TextColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type: DATE,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) DateTime(name string, props *TextColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       DATETIME,
			Unique:     props.Unique,
			Nullable:   props.Nullable,
			PrimaryKey: props.PrimaryKey,
			Default:    props.Default,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Enum(name string, options []string, props *TextColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:        ENUM,
			Unique:      props.Unique,
			Nullable:    props.Nullable,
			Default:     props.Default,
			EnumOptions: options,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Int(name string, props *NumericColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:          INT,
			Unique:        props.Unique,
			Nullable:      props.Nullable,
			PrimaryKey:    props.PrimaryKey,
			Default:       props.Default,
			AutoIncrement: props.AutoIncrement,
			Unsigned:      props.Unsigned,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Tinyint(name string, props *NumericColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:          TINYINT,
			Unique:        props.Unique,
			Nullable:      props.Nullable,
			PrimaryKey:    props.PrimaryKey,
			Default:       props.Default,
			AutoIncrement: props.AutoIncrement,
			Unsigned:      props.Unsigned,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Smallint(name string, props *NumericColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:          MEDIUMINT,
			Unique:        props.Unique,
			Nullable:      props.Nullable,
			PrimaryKey:    props.PrimaryKey,
			Default:       props.Default,
			AutoIncrement: props.AutoIncrement,
			Unsigned:      props.Unsigned,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Boolean(name string, props *NumericColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:          BOOL,
			Nullable:      props.Nullable,
			Default:       props.Default,
			AutoIncrement: props.AutoIncrement,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Float(name string, props *NumericColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       FLOAT,
			Nullable:   props.Nullable,
			Default:    props.Default,
			PrimaryKey: props.PrimaryKey,
		},
	}

	c.Columns = append(c.Columns, *col)
}

func (c *MysqlColumns) Double(name string, props *NumericColumnProps) {
	col := &MysqlColumn{
		Name: name,
		Property: &MysqlDataType{
			Type:       FLOAT,
			Size:       props.Size,
			Precision:  props.Precision,
			Nullable:   props.Nullable,
			Default:    props.Default,
			PrimaryKey: props.PrimaryKey,
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
