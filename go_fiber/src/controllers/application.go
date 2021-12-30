package controllers

import (
	"github.com/PaHell/redesiigner-mono/models"
	"github.com/PaHell/redesiigner-mono/util"

	"github.com/gofiber/fiber/v2"
)

//ApplicationController ...
type ApplicationController struct{}

var applicationModel = new(models.ApplicationModel)

func (ctrl ApplicationController) ReadAll(c *fiber.Ctx) error {
	applications, err := applicationModel.All()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// return items
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    applications,
	})
}

func (ctrl ApplicationController) ReadOne(c *fiber.Ctx) error {
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// get item
	application, err := applicationModel.One(uint(id))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// return item
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    application,
	})
}

func (ctrl ApplicationController) Create(c *fiber.Ctx) error {
	// parse input
	parsed := new(models.Application)
	if err := c.BodyParser(parsed); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// perform creation
	created, err := applicationModel.Create(parsed)
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

func (ctrl ApplicationController) Update(c *fiber.Ctx) error {
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// parse body
	parsed := new(models.Application)
	if err := c.BodyParser(parsed); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// perform changes
	updated, err := applicationModel.Update(uint(id), parsed)
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

func (ctrl ApplicationController) Delete(c *fiber.Ctx) error {
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// perform deletion
	deleted, err := applicationModel.Delete(id)
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
