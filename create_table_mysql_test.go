package gomigrator

import (
	"testing"
)

func TestCreateTableMysql(t *testing.T) {
	db, err := NewConnection("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		t.Error(err)
	}

	table := CreateTable("items", func(cols *TableColumns) {
		cols.Varchar("name", 50, nil)
		cols.Varchar("sku", 50, &TextColumnProps{Nullable: false, Unique: true})
		cols.Float("mark", nil)
		cols.Double("price", nil)
		cols.Enum("status", []string{"active", "inactive"}, &EnumColumnProps{Default: "inactive"})
		cols.Text("description", nil)
	})

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}
