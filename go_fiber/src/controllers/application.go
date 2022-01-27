package controllers

import (
	"github.com/PaHell/redesiigner-mono/models"
	"github.com/PaHell/redesiigner-mono/util"

	"github.com/gofiber/fiber/v2"
)

type ApplicationController struct{}

var applicationModel = new(models.ApplicationModel)

func (ctrl ApplicationController) ReadAll(c *fiber.Ctx) error {
	// get all
	applications, err := applicationModel.All()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// return list
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    applications,
	})
}

func (ctrl ApplicationController) ReadOne(c *fiber.Ctx) error {
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// get item
	application, err := applicationModel.One(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// create
	created, errors := applicationModel.Create(parsed)
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

func (ctrl ApplicationController) Update(c *fiber.Ctx) error {
	errors := []string{}
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		errors = append(errors, err.Error())
	}
	// parse body
	parsed := new(models.Application)
	if err := c.BodyParser(parsed); err != nil {
		errors = append(errors, err.Error())
	}
	// return collected errors from id and body
	if len(errors) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  errors,
		})
	}
	// perform changes
	updated, errors := applicationModel.Update(uint(id), parsed)
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

func (ctrl ApplicationController) Delete(c *fiber.Ctx) error {
	// parse id
	id, err := util.ParseID(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// delete
	deleted, err := applicationModel.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  []string{err.Error()},
		})
	}
	// return deleted
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": false,
		"data":    deleted,
	})
}
