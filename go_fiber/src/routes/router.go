package routes

import (
	"github.com/PaHell/redesiigner-mono/controllers"
	"github.com/PaHell/redesiigner-mono/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// BASE
	groupApiV1 := app.Group("/api/v1")
	// User
	var userController controllers.UserController
	groupUser := groupApiV1.Group("/user", middleware.HandleAuth)
	/* Rx    */ groupUser.Get("/", userController.Read)
	/* Ux  */ groupUser.Patch("/info", userController.UpdateInformation)
	/* Ux  */ groupUser.Patch("/password", userController.UpdatePassword)
	/* Ux  */ groupUser.Patch("/email", userController.UpdateEmail) // TODO: Email confirmation
	/* D  */ groupUser.Delete("/", userController.Delete)
	// Token
	var tokenController controllers.TokenController
	groupToken := groupApiV1.Group("/token", middleware.HandleAuth)
	/* Rx    */ groupToken.Get("/", tokenController.ReadAllForUser)
	/* D  */ groupToken.Delete("/all", tokenController.DeleteAll)
	// Auth
	groupAuth := groupApiV1.Group("/auth")
	/* Register */ groupAuth.Post("/register", userController.Create)
	/* Login    */ groupAuth.Post("/login", tokenController.Create)
	/* Refresh */ groupAuth.Patch("/refresh", tokenController.Refresh)
	/* Logout */ groupAuth.Delete("/logout", middleware.HandleAuth, tokenController.DeleteCurrent)

	// MODELS
	// Application
	var appController controllers.ApplicationController
	groupApps := groupApiV1.Group("/app")
	/* C     */ groupApps.Put("/", appController.Create)
	/* R*    */ groupApps.Get("/", appController.ReadAll)
	/* Rx    */ groupApps.Get("/:ID", appController.ReadOne)
	/* U   */ groupApps.Patch("/:ID", appController.Update)
	/* D  */ groupApps.Delete("/:ID", appController.Delete)

	// ADMIN
	groupAdmin := groupApiV1.Group("/admin")
	/* R */ groupAdmin.Get("/users", userController.ReadAll)
	/* R */ groupAdmin.Get("/tokens", tokenController.ReadAll)

}
