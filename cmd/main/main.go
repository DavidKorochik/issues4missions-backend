package main

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/DavidKorochik/issues4missions-backend/pkg/db"
	"github.com/DavidKorochik/issues4missions-backend/pkg/server"
)

func main() {
	config := config.LoadEnvVariables(".")

	db := db.ConnectToDB(config)
	server := server.NewServer(db)

	server.StartServer(config.Port)
}
