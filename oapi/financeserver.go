package oapi

import (
	"net"
	"time"

	"github.com/labstack/echo/v4"
)

type TodoServer struct {
}

func (t TodoServer) GetTodos(ctx echo.Context, params GetTodosParams) error {
	// our logic to retrieve all todos from a persistent layer

	te := time.Now()
	todo := Todo{
		CompletedAt: &te,
		CreatedAt:   &te,
		Status:      TodoStatusDone,
		Task:        "task",
		User:        params.User,
	}
	err := ctx.JSON(200, todo)
	return err
}

func (t TodoServer) CreateTodo(ctx echo.Context) error {
	// our logic to store the todo into a persistent layer
	return nil
}

func (t TodoServer) DeleteTodo(ctx echo.Context, todoId int32) error {
	// our logic to delete a todo from the persistent layer
	return nil
}

func (t TodoServer) UpdateTodo(ctx echo.Context, todoId int32) error {
	// our logic to update the todo.
	return nil
}

func NewServer() {
	s := TodoServer{}
	e := echo.New()

	RegisterHandlers(e, s)
	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", "8080")))
}
