package main

import (
	"github.com/b3nhard/chat-x/internal/app"
	"github.com/b3nhard/chat-x/internal/database"
	"github.com/b3nhard/chat-x/internal/router"
)

func main() {
	db := database.InitDb()

	newApp := app.NewApp(db)

	router.SetupRutes(newApp.App, db, newApp.Store, newApp.Manager)
	newApp.Start()

}
