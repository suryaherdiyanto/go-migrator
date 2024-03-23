package gomigrator

import (
	"bytes"
	"testing"
)

func TestVarcharColumn(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Varchar("first_name", 50, nil)
	}, MYSQL)

	stmt := table.Columns[0].ParseColumn()

	expected := "first_name varchar(50)"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestIntColumnWithAutoIncrement(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Int("ID", &NumericColumnProps{AutoIncrement: true})
	}, MYSQL)

	stmt := table.Columns[0].ParseColumn()

	expected := "ID int AUTO_INCREMENT PRIMARY KEY"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestUnsignedInt(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Int("ID", &NumericColumnProps{Unsigned: true})
	}, MYSQL)
	stmt := table.Columns[0].ParseColumn()

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
			case VARCHAR:
				t.Varchar("name", s.Size, nil)
			case DATE:
				t.Date("name", nil)
			case DATETIME:
				t.DateTime("name", nil)
			}
		}, MYSQL)
		stmt := table.Columns[0].ParseColumn()

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
			case TINYINT:
				t.Tinyint("ID", nil)
			case MEDIUMINT:
				t.Mediumint("ID", nil)
			case BIGINT:
				t.Bigint("ID", nil)
			case BOOL:
				t.Boolean("ID", nil)
			}
		}, MYSQL)

		stmt := table.Columns[0].ParseColumn()

		if stmt != s.Expected {
			t.Errorf("Expected: %s, but got %q", s.Expected, stmt)
		}
	}
}

func TestEnumWithDefault(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Enum("role", []string{"admin", "employee", "supervisor"}, &EnumColumnProps{Default: "admin"})
	}, MYSQL)

	stmt := table.Columns[0].ParseColumn()

	expected := "role enum('admin', 'employee', 'supervisor') DEFAULT 'admin'"

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
	}, MYSQL)

	stmt := columns.Columns[0].ParseColumn()

	expected := "mark float(53) NULL"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestDoubleWithNullable(t *testing.T) {
	table := CreateTable("users", func(t *Table) {
		t.Double("mark", &NumericColumnProps{Nullable: true, Size: 53})
	}, MYSQL)

	stmt := table.Columns[0].ParseColumn()

	expected := "mark double(53) NULL"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestIncrement(t *testing.T) {
	table := CreateTable("users", func(table *Table) {
		table.Increment("ID")
	}, MYSQL)

	stmt := table.Columns[0].ParseColumn()
	expected := "ID int AUTO_INCREMENT PRIMARY KEY"

	if condition := stmt != expected; condition {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}

	table = CreateTable("users", func(table *Table) {
		table.Increment("ID")
	}, POSTGRES)

	stmt = table.Columns[0].ParseColumn()
	expected = "ID serial PRIMARY KEY"

	if condition := stmt != expected; condition {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestBigIncrement(t *testing.T) {
	table := CreateTable("users", func(table *Table) {
		table.BigIncrement("ID")
	}, MYSQL)

	stmt := table.Columns[0].ParseColumn()
	expected := "ID bigint AUTO_INCREMENT PRIMARY KEY"

	if condition := stmt != expected; condition {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}

	table = CreateTable("users", func(table *Table) {
		table.BigIncrement("ID")
	}, POSTGRES)

	stmt = table.Columns[0].ParseColumn()
	expected = "ID bigserial PRIMARY KEY"

	if condition := stmt != expected; condition {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestUuid(t *testing.T) {
	table := CreateTable("users", func(table *Table) {
		table.Uuid("ID", &TextColumnProps{PrimaryKey: true})
	}, POSTGRES)

	stmt := table.Columns[0].ParseColumn()
	expected := "ID uuid PRIMARY KEY DEFAULT gen_random_uuid()"
	if condition := stmt != expected; condition {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}

	table = CreateTable("users", func(table *Table) {
		table.Uuid("ID", &TextColumnProps{PrimaryKey: true})
	}, MYSQL)

	stmt = table.Columns[0].ParseColumn()
	expected = "ID varchar(36) PRIMARY KEY DEFAULT uuid()"

	if condition := stmt != expected; condition {
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
	}, MYSQL)

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
