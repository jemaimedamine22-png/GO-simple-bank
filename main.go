package main

import (
	"database/sql"
	"log"

	"github/jemaimedamine22-png/simple_bank/api"
	db "github/jemaimedamine22-png/simple_bank/db/sqlc"
	"github/jemaimedamine22-png/simple_bank/util"

	_ "github.com/lib/pq" // استيراد الموزع للتسجيل التلقائي
)


func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBDriver)
	if err != nil{
		log.Fatal("coonect to db failed", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err !=nil{
		log.Fatal("cannot start server:", err)
	}


}
