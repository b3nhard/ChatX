package utils

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
)

func NewSession() *session.Store {
	return session.New(session.Config{
		Storage:    redis.New(),
		Expiration: time.Minute * 30,
	})
}
