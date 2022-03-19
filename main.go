package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{Id: 1, Name: "skyrun", Completed: true},
	{Id: 2, Name: "pringles", Completed: false},
}

func GetTodos(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(todos)

}

func CreateTodos(ctx *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}
	var body request

	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errer": "can not parse Json",
		})
	}
	todo := Todo{
		Id:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}

	todos = append(todos, &todo)

	return ctx.Status(fiber.StatusOK).JSON(todos)

}

//get Todo By ID

func GetTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errer": "cannot parse id",
		})

	}

	for _, todo := range todos {

		if todo.Id == id {
			return ctx.Status(fiber.StatusOK).JSON(todo)

		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON(todos)

}

func DeleteTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errer": "cannot parse id",
		})
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[0:i], todos[i+1:]...)

			// return ctx.Status(fiber.StatusNoContent)
		}
	}
	return ctx.Status(fiber.StatusNotFound).JSON(todos)
}

//Update Todos Function

func UpdateTodo(ctx *fiber.Ctx) error {

	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errer": "cannot parse id",
		})

	}
	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errer": "cannot parse id",
		})

	}
	var todo *Todo
  
	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}
	if todo.Id == 0 {
		ctx.Status(fiber.StatusNotFound)

	}

	if body.Name != nil {
		todo.Name = *body.Name
	}
	if body.Completed != nil {
		todo.Completed = *body.Completed
	}
	return ctx.Status(fiber.StatusOK).JSON(todo)
}

func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())


	todoApp := app.Group("/todos")
	todoApp.Get("/", GetTodos)
	todoApp.Post("/", CreateTodos)
	todoApp.Get("/:id", GetTodo)
	todoApp.Delete("/:id", DeleteTodo)
	todoApp.Patch("/:id", UpdateTodo)

	app.Listen(":8080")

}
