package routes

import (
	_ "log"

	"github.com/PaHell/redesiigner-mono/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	groupApiV1 := app.Group("/api/v1")

	//var userController controllers.UserController
	//groupUser := groupApiV1.Group("/users")
	/* C  */ // groupUser.Put("/", userController.Create) // Register
	/* R* */ // groupUser.Get("/", userController.ReadAll)
	/* Rx */ // groupUser.Get("/:ID", userController.ReadOne)
	/* U  */ // groupUser.Patch("/:ID", userController.Update)
	/* D  */ // groupUser.Delete("/:ID", userController.Delete)

	// Applications
	var appController controllers.ApplicationController
	groupApps := groupApiV1.Group("/apps")
	/* C  */ groupApps.Put("/", appController.Create)
	/* R* */ groupApps.Get("/", appController.ReadAll)
	/* Rx */ groupApps.Get("/:ID", appController.ReadOne)
	/* U  */ groupApps.Patch("/:ID", appController.Update)
	/* D  */ groupApps.Delete("/:ID", appController.Delete)

}
