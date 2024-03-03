package gomigrator

import (
	"testing"
)

func TestVarcharColumn(t *testing.T) {
	column := CreateColumn("first_name", &MysqlDataType{
		Type: "varchar",
		Size: 50,
	})
	stmt, err := column.ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "first_name varchar(50)"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestIntColumnWithAutoIncrement(t *testing.T) {
	column := CreateColumn("ID", &MysqlDataType{
		Type:          "int",
		AutoIncrement: true,
	})

	stmt, err := column.ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "ID int AUTO_INCREMENT"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestUnsignedInt(t *testing.T) {
	column := CreateColumn("ID", &MysqlDataType{
		Type:     INT,
		Unsigned: true,
	})

	stmt, err := column.ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "ID int UNSIGNED"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestIntegerColumn(t *testing.T) {
	if !IsIntegerColumn(TINYINT) {
		t.Errorf("Expected %s to be a integer column", TINYINT)
	}
}
func TestNotIntegerColumn(t *testing.T) {
	if IsIntegerColumn(VARCHAR) {
		t.Errorf("Expected %s to be not an integer column", TINYINT)
	}
}

func TestTextsColumn(t *testing.T) {
	type sample struct {
		Type     SQLDataType
		Expected string
		Size     int
	}
	samples := []sample{
		{
			Type:     CHAR,
			Expected: "name char(10)",
			Size:     10,
		},
		{
			Type:     VARCHAR,
			Expected: "name varchar(50)",
			Size:     50,
		},
		{
			Type:     DATE,
			Expected: "name date",
		},
		{
			Type:     DATETIME,
			Expected: "name datetime",
		},
	}

	for _, s := range samples {
		column := CreateColumn("name", &MysqlDataType{
			Type: SQLDataType(s.Type),
			Size: s.Size,
		})

		stmt, err := column.ParseColumn()

		if err != nil {
			t.Error(err)
		}

		if stmt != s.Expected {
			t.Errorf("Expected: %s, but got %q", s.Expected, stmt)
		}
	}
}

func TestIntegersColumn(t *testing.T) {
	type sample struct {
		Type     SQLDataType
		Expected string
	}
	samples := []sample{
		{
			Type:     INT,
			Expected: "ID int",
		},
		{
			Type:     TINYINT,
			Expected: "ID tinyint",
		},
		{
			Type:     MEDIUMINT,
			Expected: "ID mediumint",
		},
		{
			Type:     BIGINT,
			Expected: "ID bigint",
		},
		{
			Type:     BOOL,
			Expected: "ID bool",
		},
	}

	for _, s := range samples {
		column := CreateColumn("ID", &MysqlDataType{
			Type: SQLDataType(s.Type),
		})

		stmt, err := column.ParseColumn()

		if err != nil {
			t.Error(err)
		}

		if stmt != s.Expected {
			t.Errorf("Expected: %s, but got %q", s.Expected, stmt)
		}
	}
}

func TestEnumWithDefault(t *testing.T) {
	column := CreateColumn("role", &MysqlDataType{
		Type:        ENUM,
		EnumOptions: []string{"admin", "employee", "supervisor"},
		Default:     "admin",
	})

	stmt, err := column.ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "role enum(admin,employee,supervisor) DEFAULT 'admin'"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestCharColumn(t *testing.T) {
	if !IsTextColumn(VARCHAR) {
		t.Errorf("Expected %s to be a text column", VARCHAR)
	}
}

func TestNotVarcharColumn(t *testing.T) {
	if IsIntegerColumn(VARCHAR) {
		t.Errorf("Expected %s to be not an integer column", VARCHAR)
	}
}
