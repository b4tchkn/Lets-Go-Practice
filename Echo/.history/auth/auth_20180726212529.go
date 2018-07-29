package interceptor

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password) bool {
		return username == "validUser" && password == "validPassword"
	})
}
