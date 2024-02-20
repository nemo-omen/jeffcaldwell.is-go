package custommiddleware

import "github.com/labstack/echo/v4"

type Mode struct {
	IsDev bool
}

func NewMode(isDev bool) *Mode {
	return &Mode{IsDev: isDev}
}

func (m *Mode) SetMode(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("isDev", m.IsDev)
		return next(c)
	}
}
