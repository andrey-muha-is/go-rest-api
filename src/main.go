package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/jmoiron/sqlx"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	
	"./repositories"
	"./handlers"
)

const (
	dbUserName = "root"
	dbPassword = "root"
	dbHostName = "localhost"
	dbPort = "3306"
	dbName = "default"
	channelsTableName = "channel"
)

func main() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
		dbUserName,
		dbPassword,
		dbHostName,
		dbPort,
		dbName)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Panicf("Failed to connect to DB: %s.\n", err)
	}
	defer db.Close()
	
	channelsRepository := repositories.NewChannelsRepository(db, channelsTableName)
	channelsHandler := handlers.NewChannelsHandler(channelsRepository)

	r := chi.NewRouter()

	r.Get("/channels", channelsHandler.FindAll)
	r.Get("/channels/{channelID}", channelsHandler.FindById)

	http.ListenAndServe(":3000", r)
}
