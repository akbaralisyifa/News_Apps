package routes

import (
	"newsapps/internal/features/users"

	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ur users.Handler) {

	c.POST("/users", ur.Register()) // register -> umum (boleh diakses semua orang)
	c.POST("/login", ur.Login())

}
