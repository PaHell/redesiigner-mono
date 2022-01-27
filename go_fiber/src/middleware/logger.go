package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func UseLogger(app *fiber.App) {
	// Define file to logs
	file, err := os.OpenFile("./middleware.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file, // add file to save output
	}))
}
