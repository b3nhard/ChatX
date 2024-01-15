package app

import (
	"database/sql"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql/v2"
)

type App struct {
	App   *fiber.App
	DB    *sql.DB
	Store *session.Store
}

func NewApp(_db *sql.DB) *App {
	storage := mysql.New(mysql.Config{
		Table: "sessions",
		Db:    _db,
	})
	store := session.New(session.Config{
		Storage:    storage,
		Expiration: time.Minute * 5,
	})
	engine := html.New("web/views", ".gotmpl")
	app := fiber.New(fiber.Config{
		AppName:     "ChatX",
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	return &App{
		App:   app,
		DB:    _db,
		Store: store,
	}
}

func (a *App) Start() {

	a.App.Static("/static", "web/static", fiber.Static{Compress: true})
	a.App.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	a.App.Listen(":8888")
}
