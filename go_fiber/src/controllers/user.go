package controllers

import (
	"github.com/PaHell/redesiigner-mono/forms"
	"github.com/PaHell/redesiigner-mono/models"

	"github.com/gofiber/fiber/v2"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

func (ctrl UserController) ReadAll(c *fiber.Ctx) error {
	// get all
	users, err := userModel.All()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	// return list
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

func (ctrl UserController) Read(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*models.User)
	// return item
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func (ctrl UserController) Create(c *fiber.Ctx) error {
	// parse input
	parsed := new(models.User)
	if err := c.BodyParser(parsed); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// create
	created, errors := userModel.Create(parsed)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  errors,
		})
	}
	// return created
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    created,
	})
}

func (ctrl UserController) UpdateInformation(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*models.User)
	// parse body
	parsed := new(models.User)
	if err := c.BodyParser(parsed); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	// perform changes
	updated, errors := userModel.UpdatePublic(user.ID, parsed)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  errors,
		})
	}
	// return updated
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    updated,
	})
}

func (ctrl UserController) UpdatePassword(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*models.User)
	// parse body
	parsed := new(forms.UserPasswordForm)
	if err := c.BodyParser(parsed); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	// compare password
	err := userModel.CheckPasswordHash(parsed.OldPassword, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	// perform changes
	updated, errors := userModel.UpdateSecret(user.ID, &models.User{
		Password: parsed.NewPassword,
	})
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  errors,
		})
	}
	// return updated
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    updated,
	})
}

func (ctrl UserController) UpdateEmail(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*models.User)
	// parse body
	parsed := new(forms.UserEmailForm)
	if err := c.BodyParser(parsed); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  err.Error(),
		})
	}
	// perform changes
	updated, errors := userModel.UpdateSecret(user.ID, &models.User{
		Email: parsed.NewEmail,
	})
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  errors,
		})
	}
	// return updated
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    updated,
	})
}

func (ctrl UserController) Delete(c *fiber.Ctx) error {
	user := c.UserContext().Value("user").(*models.User)
	// delete
	deleted, err := userModel.Delete(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	// return deleted
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": false,
		"data":    deleted,
	})
}

/*

//Login ...
func (ctrl UserController) Login(c *fiber.Ctx) {
	var loginForm forms.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, token, err := userModel.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

//Register ...
func (ctrl UserController) Register(c *gin.Context) {
	var registerForm forms.RegisterForm

	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, err := userModel.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}

//Logout ...
func (ctrl UserController) Logout(c *gin.Context) {

	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		return
	}

	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

*/
