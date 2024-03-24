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
	db, err := gomigrator.NewConnection("mysql", "root:mariadb@tcp(localhost:3306)/go-migrator")

	if err != nil {
		t.Error(err)
	}

	table := gomigrator.CreateTable("items", func(t *gomigrator.Table) {
		t.Uuid("id", &gomigrator.UUIDColumnProps{PrimaryKey: true})
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

func TestCreateTableWithForeignKeyMysql(t *testing.T) {
	db, err := gomigrator.NewConnection("mysql", "root:mariadb@tcp(localhost:3306)/go-migrator")

	if err != nil {
		t.Error(err)
	}

	tableUser := gomigrator.CreateTable("users", func(t *gomigrator.Table) {
		t.Uuid("id", &gomigrator.UUIDColumnProps{PrimaryKey: true})
		t.Varchar("first_name", 50, nil)
		t.Varchar("last_name", 50, nil)
	}, gomigrator.MYSQL)

	tableProfile := gomigrator.CreateTable("profiles", func(t *gomigrator.Table) {
		t.Uuid("id", &gomigrator.UUIDColumnProps{PrimaryKey: true})
		t.Uuid("user_id", nil)

		t.ForeignKey("user_id", &gomigrator.ForeignKeyOptions{ReferenceTable: "users", ReferenceColumn: "id"})
		t.Varchar("address", 100, nil)
	}, gomigrator.MYSQL)

	err = tableUser.Run(db)

	if err != nil {
		t.Error(err)
	}

	err = tableProfile.Run(db)

	if err != nil {
		t.Error(err)
	}

	defer db.Close()
}
