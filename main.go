package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-master/api"
	db "github.com/go-master/db/sqlc"
	"github.com/go-master/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// コンフィグのデータを出力
	fmt.Println("===================================================")
	fmt.Println("DB Driver:", config.DBDriver)
	fmt.Println("DB Source:", config.DBSource)
	fmt.Println("Server Address:", config.ServerAddress)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		// panic(err)
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
