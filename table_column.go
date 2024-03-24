package gomigrator

import (
	"fmt"
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
}

type UUIDColumnProps struct {
	PrimaryKey bool
	Unique     bool
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
		switch col.Property.Default.(type) {
		case string:
			if strings.Contains(fmt.Sprintf("%s", col.Property.Default), "()") {
				stmt += " DEFAULT " + fmt.Sprintf("%v", col.Property.Default)
			} else {
				stmt += " DEFAULT " + fmt.Sprintf("'%v'", col.Property.Default)
			}
		default:
			stmt += " DEFAULT " + fmt.Sprintf("%v", col.Property.Default)
		}
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
	case *UUIDColumnProps:
		t.PrimaryKey = p.PrimaryKey
		t.Unique = p.Unique
	}

	return fmt.Errorf("invalid type %v", props)
}
