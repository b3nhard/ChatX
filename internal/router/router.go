package router

import (
	"database/sql"

	"github.com/b3nhard/chat-x/internal/handlers"
	manage "github.com/b3nhard/chat-x/internal/manager"
	"github.com/b3nhard/chat-x/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SetupRutes(app *fiber.App, db *sql.DB, store *session.Store, m *manage.Manager) {
	app.Get("/", middleware.Auth(store), handlers.Index(store))
	app.Get("/login", handlers.GetLogin)
	app.Post("/login", handlers.Login(store, db))
	app.Get("/register", handlers.GetSignup)
	app.Post("/register", handlers.Signup(store, db))
	app.Get("/ws", middleware.Auth(store), handlers.WebsocketHandler(db, store, m))
}
