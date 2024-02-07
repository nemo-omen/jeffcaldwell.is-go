package middleware

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

type CurrentPathContext struct {
	echo.Context
	CurrentPath string
}

func (c *CurrentPathContext) GetCurrentPath() (string, error) {
	if c.CurrentPath == "" {
		return "", errors.New("current path is not defined")
	}
	fmt.Println("Current path: ", c.CurrentPath)
	return c.CurrentPath, nil
}
