package gomigrator

import "testing"

func TestConnection(t *testing.T) {
	err := NewConnection("root:root@tcp(127.0.0.1:3306)/")

	if err != nil {
		t.Error(err)
	}
}
