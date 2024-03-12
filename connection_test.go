package gomigrator

import "testing"

func TestConnection(t *testing.T) {
	_, err := NewConnection("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		t.Error(err)
	}
}
