package gomigrator

import "testing"

func TestCreateTablePostgres(t *testing.T) {
	db, err := NewConnection("postgres", "postgres://postgres:root@127.0.0.1/testdb?sslmode=disable")

	if err != nil {
		t.Error(err)
	}

	table := CreateTable("items", func(t *Table) {
		t.Increment("id")
		t.Varchar("name", 50, nil)
		t.Varchar("sku", 50, &TextColumnProps{Nullable: false, Unique: true})
		t.Real("mark", nil)
		t.DoublePrecision("price", nil)
		t.Enum("status", []string{"active", "inactive"}, &EnumColumnProps{Default: "inactive"})
		t.Text("description", nil)
	}, POSTGRES)

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}
