package tests

import (
	"testing"

	gomigrator "github.com/suryaherdiyanto/go-migrator"
)

func TestCreateTablePostgres(t *testing.T) {
	db, err := gomigrator.NewConnection("postgres", "postgres://postgres:postgres@localhost/go-migrator?sslmode=disable")

	if err != nil {
		t.Error(err)
	}

	table := gomigrator.CreateTable("items", func(t *gomigrator.Table) {
		t.Increment("id")
		t.Varchar("name", 50, nil)
		t.Varchar("sku", 50, &gomigrator.TextColumnProps{Nullable: false, Unique: true})
		t.Real("mark", nil)
		t.DoublePrecision("price", nil)
		t.Enum("status", []string{"active", "inactive"}, &gomigrator.EnumColumnProps{Default: "inactive"})
		t.Text("description", nil)
	}, gomigrator.POSTGRES)

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateTableWithUUID(t *testing.T) {
	db, err := gomigrator.NewConnection("postgres", "postgres://postgres:root@localhost/testdb?sslmode=disable")

	if err != nil {
		t.Error(err)
	}

	table := gomigrator.CreateTable("items", func(t *gomigrator.Table) {
		t.Uuid("id", &gomigrator.TextColumnProps{PrimaryKey: true})
		t.Varchar("name", 50, nil)
		t.Int("grade", &gomigrator.NumericColumnProps{Default: 1})
	}, gomigrator.POSTGRES)

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}
