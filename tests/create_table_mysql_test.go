package tests

import (
	"testing"

	gomigrator "github.com/suryaherdiyanto/go-migrator"
)

func TestCreateTableMysql(t *testing.T) {
	db, err := gomigrator.NewConnection("mysql", "root:mariadb@tcp(localhost:3306)/go-migrator")

	if err != nil {
		t.Error(err)
	}

	table := gomigrator.CreateTable("items", func(t *gomigrator.Table) {
		t.Increment("id")
		t.Varchar("name", 50, nil)
		t.Varchar("sku", 50, &gomigrator.TextColumnProps{Nullable: false, Unique: true})
		t.Float("mark", nil)
		t.Double("price", nil)
		t.Enum("status", []string{"active", "inactive"}, &gomigrator.EnumColumnProps{Default: "inactive"})
		t.Text("description", nil)
	}, gomigrator.MYSQL)

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateTableWithUUIDMysql(t *testing.T) {
	db, err := gomigrator.NewConnection("mysql", "root:mariadb@tcp(127.0.0.1:3306)/go-migrator")

	if err != nil {
		t.Error(err)
	}

	table := gomigrator.CreateTable("items", func(t *gomigrator.Table) {
		t.Uuid("id", &gomigrator.TextColumnProps{PrimaryKey: true})
		t.Varchar("name", 50, nil)
		t.Int("grade", &gomigrator.NumericColumnProps{Default: 1})
	}, gomigrator.MYSQL)

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}
