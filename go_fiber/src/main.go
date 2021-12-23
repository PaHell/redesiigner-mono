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

    // establish db connection
    client, ctx, cancel, err := database.Connect("MONGODB_CONNECTION_STRING");
    if (*err != nil) {
        log.Fatal(*err);
    }
    log.Print(client, ctx, cancel, *err);

    // routes
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // entrypoint
    log.Fatal(app.Listen(":3000"))
}