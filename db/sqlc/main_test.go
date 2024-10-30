package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver  = "postgres"
	dbSourcer = "postgresql://root:secret@localhost:5432/hunt?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSourcer)
	if err != nil {
		log.Fatal("Não conectado ao banco de dados:", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
