package gomigrator

import (
	"bytes"
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
	if IsNumericColumn(VARCHAR) {
		t.Errorf("Expected %s to be not an integer column", VARCHAR)
	}
}

func TestFloatWithNullable(t *testing.T) {
	column := CreateColumn("mark", &MysqlDataType{
		Type:     FLOAT,
		Nullable: true,
		Size:     53,
	})

	stmt, err := column.ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "mark float(53) NULL"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestDoubleWithNullable(t *testing.T) {
	column := CreateColumn("mark", &MysqlDataType{
		Type:     DOUBLE,
		Nullable: true,
		Size:     53,
	})

	stmt, err := column.ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "mark double(53) NULL"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestCreateTableParsing(t *testing.T) {
	table := CreateTable("users", func() []MysqlColumn {
		return []MysqlColumn{
			*CreateColumn("ID", &MysqlDataType{
				Type:          INT,
				AutoIncrement: true,
			}),
			*CreateColumn("first_name", &MysqlDataType{
				Type: VARCHAR,
				Size: 50,
			}),
			*CreateColumn("last_name", &MysqlDataType{
				Type:     VARCHAR,
				Size:     50,
				Nullable: true,
			}),
			*CreateColumn("dob", &MysqlDataType{
				Type: DATE,
			}),
			*CreateColumn("bio", &MysqlDataType{
				Type: TEXT,
			}),
		}
	})

	buff := new(bytes.Buffer)
	err := parseTableTemplate(buff, table)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateTable(t *testing.T) {
	db, err := NewConnection("root:root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		t.Error(err)
	}

	table := CreateTable("users", func() []MysqlColumn {
		return []MysqlColumn{
			*CreateColumn("ID", &MysqlDataType{
				Type:          INT,
				AutoIncrement: true,
			}),
			*CreateColumn("first_name", &MysqlDataType{
				Type: VARCHAR,
				Size: 50,
			}),
			*CreateColumn("last_name", &MysqlDataType{
				Type:     VARCHAR,
				Size:     50,
				Nullable: true,
			}),
			*CreateColumn("dob", &MysqlDataType{
				Type: DATE,
			}),
			*CreateColumn("bio", &MysqlDataType{
				Type: TEXT,
			}),
			*CreateColumn("sex", &MysqlDataType{
				Type:        ENUM,
				EnumOptions: []string{"l", "p"},
				Default:     "p",
			}),
		}
	})
	defer db.Close()

	buff := new(bytes.Buffer)
	err = parseTableTemplate(buff, table)

	if err != nil {
		t.Error(err)
	}

	_, err = db.Exec(buff.String())

	if err != nil {
		t.Error(err)
	}

}
