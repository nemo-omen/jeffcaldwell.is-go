package custommiddleware

import "github.com/labstack/echo/v4"

func SetRemoteAddr(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("remoteAddr", c.Request().RemoteAddr)
		return next(c)
	}
}
