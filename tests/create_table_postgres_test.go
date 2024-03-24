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
	db, err := gomigrator.NewConnection("postgres", "postgres://postgres:postgres@localhost/go-migrator?sslmode=disable")

	if err != nil {
		t.Error(err)
	}

	table := gomigrator.CreateTable("items", func(t *gomigrator.Table) {
		t.Uuid("id", &gomigrator.UUIDColumnProps{PrimaryKey: true})
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

func TestCreateTableWithForeignKeyPostgres(t *testing.T) {
	db, err := gomigrator.NewConnection("postgres", "postgres://postgres:postgres@localhost/go-migrator?sslmode=disable")

	if err != nil {
		t.Error(err)
	}

	tableUser := gomigrator.CreateTable("users", func(t *gomigrator.Table) {
		t.Uuid("id", &gomigrator.UUIDColumnProps{PrimaryKey: true})
		t.Varchar("first_name", 50, nil)
		t.Varchar("last_name", 50, nil)
	}, gomigrator.POSTGRES)

	tableProfile := gomigrator.CreateTable("profiles", func(t *gomigrator.Table) {
		t.Uuid("id", &gomigrator.UUIDColumnProps{PrimaryKey: true})
		t.Uuid("user_id", nil)

		t.ForeignKey("user_id", &gomigrator.ForeignKeyOptions{ReferenceTable: "users", ReferenceColumn: "id"})
		t.Varchar("address", 100, nil)
	}, gomigrator.POSTGRES)

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

func TestCreateTableWithIndexPostres(t *testing.T) {
	db, err := gomigrator.NewConnection("postgres", "postgres://postgres:postgres@localhost/go-migrator?sslmode=disable")

	if err != nil {
		t.Error(err)
	}

	table := gomigrator.CreateTable("users", func(t *gomigrator.Table) {
		t.Increment("ID")
		t.Varchar("first_name", 50, nil)
		t.Varchar("last_name", 50, nil)
		t.Varchar("email", 50, &gomigrator.TextColumnProps{Unique: true})
		t.CreateIndex([]string{"first_name", "last_name"})
	}, gomigrator.POSTGRES)

	defer db.Close()

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}

}
