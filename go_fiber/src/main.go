package main

import (
	"log"
    "github.com/PaHell/redesiigner-mono/database"

    "github.com/gofiber/fiber/v2"
	_ "github.com/Kamva/mgm/v2"
)

func main() {
    // create framework instance
    app := fiber.New()

    // establish db connections
    mongoClient, ctx, cancel, err := database.ConnectMongo("MONGODB_CONNECTION_STRING");
    defer cancel()
    log.Print(mongoClient, ctx, cancel, err)
    
    redisClient, ctx, cancel, err := database.ConnectRedis("REDIS_CONNECTION_STRING");
    defer cancel()
    log.Print(redisClient, ctx, cancel, err)
    
    // routes
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // entrypoint
    log.Fatal(app.Listen(":3000"))
}