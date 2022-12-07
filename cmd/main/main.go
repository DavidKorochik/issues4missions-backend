package main

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/DavidKorochik/issues4missions-backend/pkg/db"
	"github.com/DavidKorochik/issues4missions-backend/pkg/handlers"
	"github.com/DavidKorochik/issues4missions-backend/pkg/routes"
	"github.com/DavidKorochik/issues4missions-backend/pkg/store"
)

func main() {
	config := config.LoadEnvVariables(".")

	db := db.ConnectToDB(config)

	is := store.NewIssueStore(db)
	us := store.NewUserStore(db)

	handler := handlers.NewHandler(us, is)
	routes.NewRouter(handler)
}
