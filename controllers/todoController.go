package controllers

import (
	"goAuthTodo/database"
	"goAuthTodo/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllTodo(c *fiber.Ctx) error {
	var todos []models.Todo

	database.DB.Preload("User").Find(&todos)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todos": todos,
	})
}

func GetTodoByID(c *fiber.Ctx) error {
	todo := models.Todo{}
	todoId := c.Params("id")

	if err := database.DB.Preload("User").First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Todo tidak ada!",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo ditemukan",
		"data":    todo,
	})
}

func CreateTodo(c *fiber.Ctx) error {
	data := models.Todo{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	todos := models.Todo{
		Name:     data.Name,
		Desc:     data.Desc,
		Complete: data.Complete,
	}

	if err := database.DB.Create(&todos).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating todo",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo created successfully",
		"todos":   todos,
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	update := models.Todo{}

	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	todoId := c.Params("id")
	todoUp := models.Todo{}

	if err := database.DB.First(&todoUp, "id = ?", todoId).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Todo tidak ada!",
		})
	}

	todoUp.Name = update.Name
	todoUp.Desc = update.Desc
	todoUp.Complete = update.Complete

	if err := database.DB.Save(todoUp).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error save update",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo Berhasil diupdate",
		"todo":    todoUp,
	})

}

func DeleteTodo(c *fiber.Ctx) error {
	todoId := c.Params("id")
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Todo tidak ada!",
		})
	}

	if err := database.DB.Delete(&todo).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error Delete Todo",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"messaage": "Todo berhasil dihapus",
	})
}
