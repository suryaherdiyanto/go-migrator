package gomigrator

import (
	"bytes"
	"testing"
)

func TestVarcharColumn(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Varchar("first_name", 50, nil)
	})

	stmt, err := table.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "first_name varchar(50)"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestIntColumnWithAutoIncrement(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Int("ID", &NumericColumnProps{AutoIncrement: true})
	})

	stmt, err := table.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "ID int AUTO_INCREMENT PRIMARY KEY"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestUnsignedInt(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Int("ID", &NumericColumnProps{Unsigned: true})
	})
	stmt, err := table.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "ID int UNSIGNED"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestIntegerColumn(t *testing.T) {
	if !IsNumericColumn(TINYINT) {
		t.Errorf("Expected %s to be a integer column", TINYINT)
	}
}
func TestNotIntegerColumn(t *testing.T) {
	if IsNumericColumn(VARCHAR) {
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
		table := CreateTable("users", func(t *Table) {
			switch s.Type {
			case CHAR:
				t.Char("name", s.Size, nil)
				break
			case VARCHAR:
				t.Varchar("name", s.Size, nil)
				break
			case DATE:
				t.Date("name", nil)
				break
			case DATETIME:
				t.DateTime("name", nil)
				break
			}
		})
		stmt, err := table.Columns[0].ParseColumn()

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
		table := CreateTable("users", func(t *Table) {
			switch s.Type {
			case INT:
				t.Int("ID", nil)
				break
			case TINYINT:
				t.Tinyint("ID", nil)
				break
			case MEDIUMINT:
				t.Mediumint("ID", nil)
				break
			case BIGINT:
				t.Bigint("ID", nil)
				break
			case BOOL:
				t.Boolean("ID", nil)
				break
			}
		})

		stmt, err := table.Columns[0].ParseColumn()

		if err != nil {
			t.Error(err)
		}

		if stmt != s.Expected {
			t.Errorf("Expected: %s, but got %q", s.Expected, stmt)
		}
	}
}

func TestEnumWithDefault(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Enum("role", []string{"admin", "employee", "supervisor"}, &EnumColumnProps{Default: "admin"})
	})

	stmt, err := table.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "role enum('admin', 'employee', 'supervisor') DEFAULT 'admin'\n\n"

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
	if IsNumericColumn(VARCHAR) {
		t.Errorf("Expected %s to be not an integer column", VARCHAR)
	}
}

func TestFloatWithNullable(t *testing.T) {
	columns := CreateTable("users", func(t *Table) {
		t.Float("mark", &NumericColumnProps{Nullable: true, Size: 53})
	})

	stmt, err := columns.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "mark float(53) NULL"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestDoubleWithNullable(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Double("mark", &NumericColumnProps{Nullable: true, Size: 53})
	})

	stmt, err := table.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "mark double(53) NULL"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestCreateTableParsing(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Int("ID", &NumericColumnProps{AutoIncrement: true})
		t.Varchar("first_name", 50, nil)
		t.Varchar("last_name", 50, &TextColumnProps{Nullable: true})
		t.Date("dob", nil)
		t.Text("bio", nil)
	})

	buff := new(bytes.Buffer)
	err := parseTableTemplate(buff, table)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateIndex(t *testing.T) {
	index := CreateIndex("idx_users", "users", []string{"first_name", "last_name"})

	expected := "CREATE INDEX idx_users ON users(first_name, last_name);"

	if index != expected {
		t.Errorf("Expected: %s, and got %q", expected, index)
	}
}
