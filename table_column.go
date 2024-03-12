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

type TableColumn struct {
	Name     string
	Property *SQLTableProp
}

type TableColumns struct {
	Columns []TableColumn
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

func NewTableColumns() *TableColumns {
	return &TableColumns{}
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
	default:
		return errors.New(fmt.Sprintf("Invalid type %v", props))
	}

	return nil
}
func (c *TableColumns) Varchar(name string, length int, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: VARCHAR,
		Size: length,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Char(name string, length int, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: CHAR,
		Size: length,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Text(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: TEXT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Date(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: DATE,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Timestamp(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: TIMESTAMP,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) DateTime(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type: DATETIME,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Enum(name string, options []string, props *EnumColumnProps) {
	dataType := SQLTableProp{
		Type:        ENUM,
		EnumOptions: options,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Int(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: INT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Tinyint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: TINYINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}
	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Mediumint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: MEDIUMINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Bigint(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: BIGINT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Boolean(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: BOOL,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Float(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: FLOAT,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}

func (c *TableColumns) Double(name string, props *NumericColumnProps) {
	dataType := SQLTableProp{
		Type: DOUBLE,
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	col := &TableColumn{
		Name:     name,
		Property: &dataType,
	}

	c.Columns = append(c.Columns, *col)
}
