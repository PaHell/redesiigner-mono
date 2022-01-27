package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func UseLimiter(app *fiber.App) {
	/*
		app.Use(limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.IP() == "127.0.0.1" // limit will apply to this IP
			},
			Max:        20,                // max count of connections
			Expiration: 30 * time.Second,  // expiration time of the limit
			Storage:    myCustomStorage{}, // used to store the state of the middleware
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.Get("x-forwarded-for") // allows you to generate custom keys
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.SendFile("./too-fast-page.html") // called when a request hits the limit
			},
		}))
	*/
}
