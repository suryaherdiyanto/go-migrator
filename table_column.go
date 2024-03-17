package gomigrator

import (
	"fmt"
	"slices"
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
	UUID             SQLDataType = "uuid"
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

func (c *TableColumn) ParseColumn() string {
	col := &TableColumn{
		Name:     c.Name,
		Property: c.Property,
	}

	return columnParser(col)
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

func IsTextColumn(t SQLDataType) bool {
	types := []SQLDataType{
		CHAR,
		VARCHAR,
		TEXT,
		DATE,
		DATETIME,
	}

	return slices.Index(types, t) >= 0
}

func columnParser(col *TableColumn) string {
	stmt := col.Name + " " + string(col.Property.Type)

	if size := col.Property.Size; size > 0 {
		if precision := col.Property.Precision; precision > 0 {
			stmt += fmt.Sprintf("(%d, %d)", size, precision)
		} else {
			stmt += fmt.Sprintf("(%d)", size)
		}
	}

	if col.Property.Type == ENUM {
		stmt += "(" + col.Property.PrintEnumValues() + ")"
	}

	if col.Property.AutoIncrement {
		stmt += " AUTO_INCREMENT PRIMARY KEY"
	}

	if col.Property.Unsigned {
		stmt += " UNSIGNED"
	}

	if col.Property.Unique {
		stmt += " UNIQUE"
	}

	if col.Property.PrimaryKey {
		stmt += " PRIMARY KEY"
	}

	if col.Property.Nullable {
		stmt += " NULL"
	}

	if col.Property.Default != nil {
		stmt += " DEFAULT " + fmt.Sprintf("%v", col.Property.Default)
	}

	return stmt
}

func fillProps(t *SQLTableProp, props interface{}) error {
	switch p := props.(type) {
	case *TextColumnProps:
		t.Unique = p.Unique
		t.Default = p.Default
		t.PrimaryKey = p.PrimaryKey
		t.Nullable = p.Nullable
		return nil
	case *NumericColumnProps:
		t.Unique = p.Unique
		t.Default = p.Default
		t.PrimaryKey = p.PrimaryKey
		t.Nullable = p.Nullable
		t.AutoIncrement = p.AutoIncrement
		t.Unsigned = p.Unsigned
		t.Precision = p.Precision
		t.Size = p.Size
		return nil
	case *EnumColumnProps:
		t.Default = p.Default
		t.Nullable = p.Nullable
		return nil
	}

	return fmt.Errorf("invalid type %v", props)
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
	}

	if props != nil {
		fillProps(&dataType, props)
	}

	if t.Dialect == POSTGRES {
		enumType := t.Name + "_" + name + "_type"
		t.EnumStatements = append(t.EnumStatements, t.CreateEnum(enumType, options))
		dataType.Type = SQLDataType(enumType)
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

func (t *Table) Increment(name string) {
	dataType := SQLTableProp{
		Type:       SERIAL,
		PrimaryKey: true,
	}

	if t.Dialect == MYSQL {
		dataType.Type = INT
		dataType.PrimaryKey = false
		dataType.AutoIncrement = true
	}

	t.AddColumn(name, dataType)
}

func (t *Table) BigIncrement(name string) {
	dataType := SQLTableProp{
		Type:       BIGSERIAL,
		PrimaryKey: true,
	}

	if t.Dialect == MYSQL {
		dataType.Type = BIGINT
		dataType.PrimaryKey = false
		dataType.AutoIncrement = true
	}

	t.AddColumn(name, dataType)
}

func (t *Table) Uuid(name string, props *TextColumnProps) {
	dataType := SQLTableProp{
		Type:       UUID,
		Default:    "gen_random_uuid()",
		PrimaryKey: props.PrimaryKey,
		Unique:     props.Unique,
	}

	if t.Dialect == MYSQL {
		dataType.Type = VARCHAR
		dataType.Default = "uuid()"
		dataType.Size = 36
	}

	t.AddColumn(name, dataType)
}
