package main

import (
	"database/sql"
	"log"

	"github.com/SirBrunoTheWise/hunt/api"
	db "github.com/SirBrunoTheWise/hunt/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSourcer    = "postgresql://root:secret@localhost:5432/hunt?sslmode=disable"
	serverAdress = "0.0.0.0:8081"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSourcer)
	if err != nil {
		log.Fatal("Não conectado ao banco de dados:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)

	if err != nil {
		log.Fatal("Servidor não criado:", err)
	}
}
