package gomigrator

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

const (
	VARCHAR          SQLDataType = "varchar"
	CHAR             SQLDataType = "char"
	BIGINT           SQLDataType = "bigint"
	INT              SQLDataType = "int"
	TINYINT          SQLDataType = "tinyint"
	MEDIUMINT        SQLDataType = "mediumint"
	SERIAL           SQLDataType = "serial"
	BIGSERIAL        SQLDataType = "bigserial"
	BOOL             SQLDataType = "bool"
	FLOAT            SQLDataType = "float"
	DOUBLE           SQLDataType = "double"
	REAL             SQLDataType = "real"
	DOUBLE_PRECISION SQLDataType = "double precision"
	TEXT             SQLDataType = "text"
	DATE             SQLDataType = "date"
	ENUM             SQLDataType = "enum"
	DATETIME         SQLDataType = "datetime"
	TIMESTAMP        SQLDataType = "timestamp"
)

type TableColumn struct {
	Name     string
	Property *SQLTableProp
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

type EnumColumnProps struct {
	Default  interface{}
	Nullable bool
	Dialect  SQLDialect
}

func (c *TableColumn) ParseColumn() (string, error) {
	col := &TableColumn{
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

func (t *Table) AddColumn(name string, props SQLTableProp) {
	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &props,
	})
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
		REAL,
		SERIAL,
		BIGSERIAL,
		DOUBLE_PRECISION,
	}

	return slices.Index(types, t) >= 0
}

func parseColumnTemplate(w io.Writer, data *TableColumn) error {
	templatePath := "./template/types/" + string(data.Property.Type) + ".go.tmpl"
	templateName := string(data.Property.Type) + ".go.tmpl"

	if IsNumericColumn(data.Property.Type) {
		templatePath = "./template/types/int.go.tmpl"
		templateName = "int.go.tmpl"

		if data.Property.Type == FLOAT || data.Property.Type == DOUBLE {
			templatePath = "./template/types/float.go.tmpl"
			templateName = "float.go.tmpl"
		}
	}

	if IsTextColumn(data.Property.Type) {
		templatePath = "./template/types/varchar.go.tmpl"
		templateName = "varchar.go.tmpl"
	}

	return parseTemplate(w, data, templateName, templatePath)
}

func fillProps(t *SQLTableProp, props interface{}) error {
	switch p := props.(type) {
	case *TextColumnProps:
		t.Unique = p.Unique
		t.Default = p.Default
		t.PrimaryKey = p.PrimaryKey
		t.Nullable = p.Nullable
	case *NumericColumnProps:
		t.Unique = p.Unique
		t.Default = p.Default
		t.PrimaryKey = p.PrimaryKey
		t.Nullable = p.Nullable
		t.AutoIncrement = p.AutoIncrement
		t.Unsigned = p.Unsigned
		t.Precision = p.Precision
		t.Size = p.Size
	case *EnumColumnProps:
		t.Default = p.Default
		t.Nullable = p.Nullable
		t.Dialect = p.Dialect
	default:
		return errors.New(fmt.Sprintf("Invalid type %v", props))
	}

	return nil
}
func (t *Table) Varchar(name string, length int, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: VARCHAR,
		Size: length,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Char(name string, length int, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: CHAR,
		Size: length,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Text(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: TEXT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Date(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: DATE,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Timestamp(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: TIMESTAMP,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) DateTime(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: DATETIME,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Enum(name string, options []string, props *EnumColumnProps) {
	dataType := SQLTableProp{
		Type:        ENUM,
		EnumOptions: options,
		Dialect:     MYSQL,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Int(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: INT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Serial(name string) {
	dataType := SQLTableProp{
		Type: SERIAL,
	}

	t.AddColumn(name, dataType)
}

func (t *Table) BigSerial(name string) {
	dataType := SQLTableProp{
		Type:     BIGSERIAL,
		Unsigned: true,
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Tinyint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: TINYINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (t *Table) Mediumint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: MEDIUMINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (t *Table) Bigint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: BIGINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.Columns = append(t.Columns, TableColumn{
		Name:     name,
		Property: &dataType,
	})
}

func (t *Table) Boolean(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: BOOL,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Float(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: FLOAT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Double(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: DOUBLE,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Real(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: REAL,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}

func (t *Table) DoublePrecision(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: DOUBLE_PRECISION,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	t.AddColumn(name, dataType)
}
