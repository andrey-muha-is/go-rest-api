package main

import (
	"fmt"
	"net/http"
	"github.com/jmoiron/sqlx"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	
	"./config"
	"./repositories"
	"./handlers"
	"./utils"
)

func main() {
	appConfig := config.GetAppConfig()

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	
	dataSourceName := utils.GetDataSourceName(&appConfig)

	zap.S().Debug("Connecting to DB...")
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		zap.S().Errorf("Failed to connect to DB: %s.", err)
	}
	zap.S().Debug("Successfully connected to DB")
	defer db.Close()
	
	channelsRepository := repositories.NewChannelsRepository(db, appConfig.DbChannelsTable)
	channelsHandler := handlers.NewChannelsHandler(channelsRepository)
	programsRepository := repositories.NewProgramsRepository(db, appConfig.DbProgramsTable)
	programsHandler := handlers.NewProgramsHandler(programsRepository)

	r := chi.NewRouter()

	r.Get("/channels", channelsHandler.FindAll)
	r.Get("/channels/{channelID}", channelsHandler.FindById)
	r.Get("/channels/{channelID}/programs", programsHandler.FindByChannelId)
	r.Get("/programs", programsHandler.FindAll)

	zap.S().Infof("Server was started on http://localhost:%s", appConfig.AppPort)

	http.ListenAndServe(fmt.Sprintf(":%s", appConfig.AppPort), r)
}
