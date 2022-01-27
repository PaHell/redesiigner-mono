package main

import (
	"log"

	"github.com/PaHell/redesiigner-mono/database"
	"github.com/PaHell/redesiigner-mono/middleware"
	"github.com/PaHell/redesiigner-mono/models"
	"github.com/PaHell/redesiigner-mono/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// FRAMEWORK
	app := fiber.New()

	// DATABASE
	/* other clients:
	   mongoClient, ctx, cancel, err := database.ConnectMongo("MONGODB_CONNECTION_STRING");
	   defer cancel()
	   redisClient, ctx, cancel, err := database.ConnectRedis("REDIS_CONNECTION_STRING");
	   defer cancel()
	*/
	err := database.ConnectSQLite("app.db")
	if err != nil {
		log.Fatal("Could not connect to database")
		return
	}
	if database.InstanceGorm == nil {
		panic("No database instance")
	}

	// MIGRATIONS
	database.InstanceGorm.AutoMigrate(&models.User{})
	database.InstanceGorm.AutoMigrate(&models.Token{})
	database.InstanceGorm.AutoMigrate(&models.Application{})
	log.Print("Database migrated")

	// MIDDLEWARE
	middleware.UseLogger(app)
	//middleware.UseCors(app)
	//middleware.UseCSRF(app)
	// middleware.UseLimiter(app)
	log.Print("Middleware configured")

	// ROUTES
	routes.Setup(app)
	log.Print("Routes configured")

	// START
	log.Fatal(app.Listen(":3000"))

}
