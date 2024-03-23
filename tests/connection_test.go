package tests

import (
	"testing"

	gomigrator "github.com/suryaherdiyanto/go-migrator"
)

func TestConnection(t *testing.T) {
	_, err := gomigrator.NewConnection("mysql", "root:mariadb@tcp(localhost:3306)/go-migrator")

	if err != nil {
		t.Error(err)
	}
}
