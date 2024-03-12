package gomigrator

import "testing"

func TestCreateTablePostgres(t *testing.T) {
	db, err := NewConnection("postgres", "postgres://postgres:root@127.0.0.1/testdb?sslmode=disable")

	if err != nil {
		t.Error(err)
	}

	table := CreateTable("items", func(cols *TableColumns) {
		cols.Serial("id")
		cols.Varchar("name", 50, nil)
		cols.Varchar("sku", 50, &TextColumnProps{Nullable: false, Unique: true})
		cols.Real("mark", nil)
		cols.DoublePrecision("price", nil)
		cols.Enum("status", []string{"active", "inactive"}, &EnumColumnProps{Default: "inactive", Dialect: POSTGRES})
		cols.Text("description", nil)
	})

	defer db.Close()

	_, _ = db.Exec("DROP TABLE IF EXISTS items")

	err = table.Run(db)

	if err != nil {
		t.Error(err)
	}
}
