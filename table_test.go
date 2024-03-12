package gomigrator

import (
	"bytes"
	"testing"
)

func TestVarcharColumn(t *testing.T) {
	columns := NewTableColumns()
	columns.Varchar("first_name", 50, nil)

	stmt, err := columns.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "first_name varchar(50)"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestIntColumnWithAutoIncrement(t *testing.T) {
	columns := NewTableColumns()
	columns.Int("ID", &NumericColumnProps{AutoIncrement: true})

	stmt, err := columns.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "ID int AUTO_INCREMENT PRIMARY KEY"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestUnsignedInt(t *testing.T) {
	columns := NewTableColumns()
	columns.Int("ID", &NumericColumnProps{Unsigned: true})
	stmt, err := columns.Columns[0].ParseColumn()

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
		columns := NewTableColumns()
		switch s.Type {
		case CHAR:
			columns.Char("name", s.Size, nil)
			break
		case VARCHAR:
			columns.Varchar("name", s.Size, nil)
			break
		case DATE:
			columns.Date("name", nil)
			break
		case DATETIME:
			columns.DateTime("name", nil)
			break
		}
		stmt, err := columns.Columns[0].ParseColumn()

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
		columns := NewTableColumns()
		switch s.Type {
		case INT:
			columns.Int("ID", nil)
			break
		case TINYINT:
			columns.Tinyint("ID", nil)
			break
		case MEDIUMINT:
			columns.Mediumint("ID", nil)
			break
		case BIGINT:
			columns.Bigint("ID", nil)
			break
		case BOOL:
			columns.Boolean("ID", nil)
			break
		}

		stmt, err := columns.Columns[0].ParseColumn()

		if err != nil {
			t.Error(err)
		}

		if stmt != s.Expected {
			t.Errorf("Expected: %s, but got %q", s.Expected, stmt)
		}
	}
}

func TestEnumWithDefault(t *testing.T) {
	columns := NewTableColumns()
	columns.Enum("role", []string{"admin", "employee", "supervisor"}, &EnumColumnProps{Default: "admin"})

	stmt, err := columns.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

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
	columns := NewTableColumns()
	columns.Float("mark", &NumericColumnProps{Nullable: true, Size: 53})

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
	columns := NewTableColumns()
	columns.Double("mark", &NumericColumnProps{Nullable: true, Size: 53})

	stmt, err := columns.Columns[0].ParseColumn()

	if err != nil {
		t.Error(err)
	}

	expected := "mark double(53) NULL"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}

func TestCreateTableParsing(t *testing.T) {
	table := CreateTable("users", func(cols *TableColumns) {
		cols.Int("ID", &NumericColumnProps{AutoIncrement: true})
		cols.Varchar("first_name", 50, nil)
		cols.Varchar("last_name", 50, &TextColumnProps{Nullable: true})
		cols.Date("dob", nil)
		cols.Text("bio", nil)
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

	table := CreateTable("users", func(cols *TableColumns) {
		cols.Int("ID", &NumericColumnProps{AutoIncrement: true})
		cols.Varchar("first_name", 50, nil)
		cols.Varchar("last_name", 50, &TextColumnProps{Nullable: true})
		cols.Date("dob", nil)
		cols.Text("bio", nil)
		cols.Enum("sex", []string{"l", "p"}, &EnumColumnProps{Default: "p"})
	})
	defer db.Close()

	buff := new(bytes.Buffer)
	err = parseTableTemplate(buff, table)

	if err != nil {
		t.Error(err)
	}

	_, _ = db.Exec("DROP TABLE IF EXISTS users")
	_, err = db.Exec(buff.String())

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

func TestCreateTableFunc(t *testing.T) {
	table := CreateTable("users", func(cols *TableColumns) {
		cols.Varchar("first_name", 50, &TextColumnProps{})
		cols.Varchar("last_name", 50, &TextColumnProps{Nullable: true})
	})

	tableLength := len(table.Columns)

	if tableLength != 2 {
		t.Errorf("Expected 2 columns, but got %d", tableLength)

	}
}

func TestCreateTableFuncRun(t *testing.T) {
	db, err := NewConnection("root:root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		t.Error(err)
	}

	table := CreateTable("items", func(cols *TableColumns) {
		cols.Varchar("name", 50, nil)
		cols.Varchar("sku", 50, &TextColumnProps{Nullable: false, Unique: true})
		cols.Float("mark", nil)
		cols.Double("price", nil)
	})

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}
