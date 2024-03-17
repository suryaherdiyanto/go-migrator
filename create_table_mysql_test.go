package gomigrator

import (
	"testing"
)

func TestCreateTableMysql(t *testing.T) {
	db, err := NewConnection("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		t.Error(err)
	}

	table := CreateTable("items", func(t *Table) {
		t.Increment("id")
		t.Varchar("name", 50, nil)
		t.Varchar("sku", 50, &TextColumnProps{Nullable: false, Unique: true})
		t.Float("mark", nil)
		t.Double("price", nil)
		t.Enum("status", []string{"active", "inactive"}, &EnumColumnProps{Default: "inactive"})
		t.Text("description", nil)
	}, MYSQL)

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateTableWithUUIDMysql(t *testing.T) {
	db, err := NewConnection("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		t.Error(err)
	}

	table := CreateTable("items", func(t *Table) {
		t.Uuid("id", &TextColumnProps{PrimaryKey: true})
		t.Varchar("name", 50, nil)
		t.Int("grade", &NumericColumnProps{Default: 1})
	}, MYSQL)

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}
