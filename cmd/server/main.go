package main

import (
	"fmt"
	"github.com/joeymckenzie/brewdude/internal"
	"github.com/joeymckenzie/brewdude/pkg/logging"
	"github.com/joeymckenzie/brewdude/pkg/persistence"
	"github.com/spf13/viper"
	"net/http"
)

const environment = "local"
const configFilePath = "./config"
const loggerPrefix = "brewdude - "

func main() {
	logger := logging.New(loggerPrefix)

	viper.SetConfigName(environment)
	viper.AddConfigPath(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		logger.FatalWhile("reading configuration", err)
	}

	connectionString := viper.GetString("database.connection")
	port := viper.GetString("brewdude.port")

	db, err := persistence.InitializePostgres(logger, connectionString)

	if err != nil {
		logger.FatalWhile("initializing a postgres connection", err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: internal.BuildApiRouter(logger, db),
	}

	logger.Info(fmt.Sprintf("Server starting on port %s...", port))

	if err := server.ListenAndServe(); err != nil {
		logger.FatalWhile("while starting the server", err)
	}
}
