package gomigrator

import (
	"testing"
)

func TestCreateColumn(t *testing.T) {
	column := CreateColumn("first_name", &MysqlDataType{
		Type: "varchar",
		Size: 50,
	})
	stmt := column.ParseColumn()
	expected := "first_name varchar(50)"

	if stmt != expected {
		t.Errorf("Expected: %s, and got %q", expected, stmt)
	}
}
