package database

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestRestDB(t *testing.T) {
	ResetDB()
}
