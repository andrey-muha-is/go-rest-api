package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/jmoiron/sqlx"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	
	"./config"
	"./repositories"
	"./handlers"
	"./utils"
)

func main() {
	appConfig := config.GetAppConfig() 
	
	dataSourceName := utils.GetDataSourceName(&appConfig)

	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Panicf("Failed to connect to DB: %s.\n", err)
	}
	defer db.Close()
	
	channelsRepository := repositories.NewChannelsRepository(db, appConfig.DbChannelsTable)
	channelsHandler := handlers.NewChannelsHandler(channelsRepository)

	r := chi.NewRouter()

	r.Get("/channels", channelsHandler.FindAll)
	r.Get("/channels/{channelID}", channelsHandler.FindById)

	http.ListenAndServe(fmt.Sprintf(":%s", appConfig.AppPort), r)
}
