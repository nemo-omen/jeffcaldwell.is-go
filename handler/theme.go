package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
)

type ThemeHandler struct{}

const _themeDir string = "./content/theme/"

func (h ThemeHandler) HandleGetTheme(c echo.Context) error {
	themeName := c.Param("themeName")
	fileService := service.FileService{}
	fullFilePath := _themeDir + themeName
	fileData, err := fileService.GetFileText(fullFilePath)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error reading theme file")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	return c.String(http.StatusOK, fileData)
}
