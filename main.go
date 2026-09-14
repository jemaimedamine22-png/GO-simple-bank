package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // استيراد الموزع للتسجيل التلقائي
	"github/jemaimedamine22-png/simple_bank/api"
	db "github/jemaimedamine22-png/simple_bank/db/sqlc"
)

const(
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)
func main(){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatal("coonect to db failed", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err !=nil{
		log.Fatal("cannot start server:", err)
	}


}
