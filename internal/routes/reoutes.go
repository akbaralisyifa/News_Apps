package routes

import (
	"newsapps/configs"
	"newsapps/internal/features/articles"
	"newsapps/internal/features/users"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ur users.Handler, ac articles.Handler) {

	secrateJwt := configs.ImportSetting().JWTSECRET

	c.POST("/register", ur.Register()) // register -> umum (boleh diakses semua orang)
	c.POST("/login", ur.Login())

	c.GET("/articles", ac.GetArticles())

	a := c.Group("/articles")
	a.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	a.POST("", ac.CreateArticles())
	a.PUT("/:id", ac.UpdateArticles())
	a.DELETE("/:id", ac.DeleteArticles())
	a.GET("/:id", ac.GetArticlesByID())
}
