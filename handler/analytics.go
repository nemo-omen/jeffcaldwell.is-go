package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AnalyticsHandler struct{}

type PageView struct {
	IpHash []byte
	Path   string
	Time   time.Time
}

func (h AnalyticsHandler) PostPageview(c echo.Context) error {
	hash := md5.New()
	data := c.Request().Header.Get("remoteaddr")
	pv := PageView{
		IpHash: hash.Sum([]byte(data)),
		Path:   c.Request().URL.Path,
		Time:   time.Now(),
	}
	fmt.Printf("MD5: %x\n", pv.IpHash)
	return c.String(http.StatusOK, "Yep!")
}
