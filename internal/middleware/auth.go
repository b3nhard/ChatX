package middleware

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Auth(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		next := c.Path()
		s, err := store.Get(c)
		if err != nil {
			return c.Redirect("/login")
		}
		if len(s.Keys()) > 0 {

			user := s.Get("user")
			if user == nil {

				return c.Redirect(fmt.Sprintf("/login?next=%s", next))
			}
			c.Locals("user", user.(string))
			log.Println("Middleware User: ", user)
		} else {
			return c.Redirect(fmt.Sprintf("/login?next=%s", next))
		}
		return c.Next()
	}
}
