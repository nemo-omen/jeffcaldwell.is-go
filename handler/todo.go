package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
	"jeffcaldwell.is/view/todo"
)

type TodoHandler struct{}

func (h TodoHandler) HandleGetTodoIndex(c echo.Context) error {
	current := c.Request().URL.Path
	remoteAddr := c.Get("remoteAddr").(string)

	todoService := service.NewTodoService("./content/todo")

	pageData, err := todoService.GetTodos()

	if err != nil {
		return echo.NewHTTPError(501, fmt.Sprintf("Error getting todo markup: %v", err))
	}

	return render(c, todo.Index(current, remoteAddr, pageData))
}
