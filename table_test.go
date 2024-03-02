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
