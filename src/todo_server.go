package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

func main() {
	app := fiber.New()

	db := NewInMemoryStore()

	todoGroup := app.Group("/todos")

	todoGroup.Get("/", func(c *fiber.Ctx) error {
		todos := db.GetAll()
		if len(todos) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No todos found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(db.GetAll())
	})

	todoGroup.Post("/", func(c *fiber.Ctx) error {
		todo := new(Todo)
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		db.Set(todo)
		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	todoGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}
		todo, ok := db.Get(id)
		if !ok {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(todo)

	})

	todoGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}
		_, ok := db.Get(id)
		if !ok {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}
		db.Delete(id)
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Todo deleted",
		})
	})

	todoGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}
		todo, ok := db.Get(id)
		if !ok {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}
		todoUpdate := new(Todo)
		if err := c.BodyParser(todoUpdate); err != nil {
			return err
		}

		todo, _ = db.UpdateTodo(
			id,
			todoUpdate.Title,
			todoUpdate.Description,
			todoUpdate.Status,
		)

		return c.Status(fiber.StatusOK).JSON(todo)
	})

	todoGroup.Delete("/", func(c *fiber.Ctx) error {
		db.Clear()
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Todos deleted",
		})
	})

	err := app.Listen(":3000")
	if err != nil {
		os.Exit(1)
	}

}
