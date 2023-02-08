package server

import (
	"crypto/subtle"
	"os"

	"github.com/labstack/echo"
)

func requireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")

		if subtle.ConstantTimeCompare([]byte(auth), []byte(os.Getenv("SNIPPY_AUTH"))) != 1 {
			return c.String(401, "Unauthorized")
		}

		return next(c)
	}
}
