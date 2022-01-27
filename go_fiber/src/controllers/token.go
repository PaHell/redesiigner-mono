package controllers

import (
	"time"

	"github.com/PaHell/redesiigner-mono/forms"
	"github.com/PaHell/redesiigner-mono/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TokenController struct{}

var tokenModel = new(models.TokenModel)

func (ctrl TokenController) ReadAll(c *fiber.Ctx) error {
	// get all
	tokens, err := tokenModel.All()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// return list
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    tokens,
	})
}

func (ctrl TokenController) ReadAllForUser(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*models.User)
	// get item
	tokens, err := tokenModel.AllForUser(user.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	// return item
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    tokens,
	})
}

func (ctrl TokenController) Create(c *fiber.Ctx) error {
	var errors []string
	// parse body
	parsed := new(forms.TokenCreateForm)
	err := c.BodyParser(parsed)
	if err != nil {
		errors = append(errors, err.Error())
	}
	// validate
	v := validator.New()
	err = v.Struct(parsed)
	if err != nil {
		errors = append(errors, err.Error())
	}
	// return collected errors
	if len(errors) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  errors,
		})
	}
	// get user
	user, err := userModel.OneByEmailUsername(parsed.EmailUsername)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// compare password
	err = userModel.CheckPasswordHash(parsed.Password, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// create
	created, err := tokenModel.Create(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// return created
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    created,
	})
}

func (ctrl TokenController) Refresh(c *fiber.Ctx) error {
	// parse body
	parsed := new(forms.TokenUpdateForm)
	err := c.BodyParser(parsed)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	// get item
	old, err := tokenModel.OneByRefresh(parsed.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// check expired
	if old.RefreshExpires.Before(time.Now()) {
		// delete expired token (optionally)
		_, _ = tokenModel.DeleteByAccess(old.AccessToken)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"data":    []string{"Token expired"},
		})
	}
	// compare refresh token
	if old.RefreshToken != parsed.RefreshToken {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"data":    []string{"Token mismatch"},
		})
	}
	// create new
	new, err := tokenModel.Create(old.UserID)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"data":    []string{err.Error()},
		})
	}
	// delete old
	_, _ = tokenModel.DeleteByAccess(old.AccessToken)
	// return new
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": false,
		"data":    new,
	})
}

func (ctrl TokenController) DeleteCurrent(c *fiber.Ctx) error {
	tokenStr := c.Get("Authorization", "")
	deleted, err := tokenModel.DeleteByAccess(tokenStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// return deleted
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    deleted,
	})
}

func (ctrl TokenController) DeleteAll(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*models.User)
	deleted, err := tokenModel.DeleteAllForUser(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// return deleted
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    deleted,
	})
}
