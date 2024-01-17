package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	manage "github.com/b3nhard/chat-x/internal/manager"
	"github.com/b3nhard/chat-x/internal/models"
	"github.com/b3nhard/chat-x/internal/utils"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	username string
	password string
}
type Message struct {
	Message string
}

func GetLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{"Title": "ChatX | Login"})
}

func Login(store *session.Store, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		next := c.Params("next", "/")
		var cred Credentials
		var user models.User
		cred.username = c.FormValue("username")
		cred.password = c.FormValue("password")
		err := db.QueryRow("SELECT * FROM users WHERE username=?", cred.username).Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return c.Redirect("/login?error=invalid credentials 404")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.password)); err != nil {
			return c.Redirect("/login?error=invalid credentials 500")
		}
		sess, err := store.Get(c)
		if err != nil {
			return c.Redirect("/login?error=failed to fetch session")
		}

		if sess.Fresh() {
			sess.Set("user", user.Username)

			if err := sess.Save(); err != nil {
				return c.Redirect("/login?error=failed to save session")
			}
			return c.Redirect("/")

		}

		return c.Redirect(next)

	}
}

func GetSignup(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{"Title": "ChatX | Register"})
}

func Signup(store *session.Store, db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var cred Credentials
		var user models.User
		cred.username = c.FormValue("username")
		cred.password = c.FormValue("password")
		err := db.QueryRow("SELECT * FROM users WHERE username=?", cred.username).Scan(&user.Id, &user.Username, &user.Password)
		if err == nil {
			return c.Redirect("/register?error=user already exists")
		}
		if err != sql.ErrNoRows {
			return c.Redirect("/register?error=sql error")
		}
		hash, err := utils.HashPassword(cred.password)
		if err != nil {
			return c.Redirect("/register?error=failed to hash password")
		}
		_, err = db.Exec("INSERT INTO users(username,password) VALUES(?,?)", cred.username, hash)
		if err != nil {
			return c.Redirect("/register?error=failed to save user to db")
		}

		return c.Redirect("/login?success=User registered")
	}
}

func Index(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		log.Println("USER: ", user)
		return c.Render("index", fiber.Map{"user": user, "Title": "ChatX | Home"})
	}
}

func WebsocketHandler(db *sql.DB, store *session.Store, m *manage.Manager) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		c.Locals("user")
		var (
			mt  int
			msg []byte
			err error
		)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			var m Message
			if err := json.Unmarshal(msg, &m); err != nil {
				log.Println(err)
			}
			log.Println("Decoded MSG: ", m)

			text := fmt.Sprintf(
				`<div hx-swap-oob="beforeend:#messages">
				<div  class="flex justify-end mt-1 mb-4">
				<div class="bg-gray-300 p-2 rounded-lg rounded-tl-none  w-max max-w-xs">
				%s
				<br/>
				<small class="text-sm text-green-500">%s</small>
				</div>
				
			</div>
			</div>`, m.Message, c.Locals("user").(string))

			if err = c.WriteMessage(mt, []byte(text)); err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
}
