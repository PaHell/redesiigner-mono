package middleware

import (
	"context"

	"github.com/PaHell/redesiigner-mono/models"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/cors"
	// "github.com/gofiber/jwt"
)

func HandleAuth(c *fiber.Ctx) error {
	// get header
	tokenStr := c.Get("Authorization", "")
	if len(tokenStr) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized: Header not set",
		})
	}
	// get token
	tokenModel := new(models.TokenModel)
	token, err := tokenModel.OneByAccess(tokenStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized: Invalid Token",
		})
	}
	// get user
	userModel := new(models.UserModel)
	user, err := userModel.OneByID(token.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized: Could not get user",
		})
	}
	// save user context
	ctx := c.UserContext()
	// In just a couple of consecutive requests, this condition below will be triggered
	if ctx.Value("user") != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized: Could not set user on server",
		})
	}
	ctx = context.WithValue(ctx, "user", user)
	c.SetUserContext(ctx)
	// continue
	c.Next()
	return nil
}
