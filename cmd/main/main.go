package main

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/DavidKorochik/issues4missions-backend/pkg/db"
)

func main() {
	config := config.LoadEnvVariables(".")

	db.ConnectToDB(config)
}
