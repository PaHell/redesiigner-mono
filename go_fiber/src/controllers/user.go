package controllers

import (
	"github.com/PaHell/redesiigner-mono/models"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

/*
func (ctrl UserController) ReadAll(c *fiber.Ctx) error {
	users, err := userModel.All()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// return items
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

func (ctrl UserController) ReadOne(c *fiber.Ctx) error {
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// get item
	user, err := userModel.One(uint(id))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// return item
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func (ctrl UserController) Register(c *fiber.Ctx) error {
	// parse input
	parsed := new(models.User)
	if err := c.BodyParser(parsed); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// perform creation
	created, err := userModel.Create(parsed)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// return item
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    created,
	})
}

func (ctrl UserController) Update(c *fiber.Ctx) error {
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// parse body
	parsed := new(models.User)
	if err := c.BodyParser(parsed); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// perform changes
	updated, err := userModel.Update(uint(id), parsed)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// return item
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    updated,
	})
}

func (ctrl UserController) Delete(c *fiber.Ctx) error {
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// perform deletion
	deleted, err := userModel.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// return item
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": false,
		"data":    deleted,
	})
}
*/

//#################################################################

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
