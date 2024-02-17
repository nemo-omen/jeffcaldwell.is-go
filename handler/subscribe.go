package handler

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/subscribe"
)

type SubscribeHandler struct{}

func (h SubscribeHandler) HandleSubscribeIndex(c echo.Context) error {
	return render(c, subscribe.Index())
}
