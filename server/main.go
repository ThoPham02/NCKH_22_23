package main

import (
	"database/sql"
	"github/ThoPham02/research_management/api"
	"github/ThoPham02/research_management/api/db/store"
	"github/ThoPham02/research_management/config"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config::", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to DB:", err)
	}

	store := store.NewStore(conn)
	server := api.NewServer(*store)

	// go func ()  {

	// }
	err = server.Router.Run(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't run server:", err)
	}

	// os.Signal()
}
