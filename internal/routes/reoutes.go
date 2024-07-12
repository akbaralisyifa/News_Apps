package routes

import (
	"newsapps/configs"
	"newsapps/internal/features/articles"
	"newsapps/internal/features/comments"
	"newsapps/internal/features/users"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ur users.Handler, ac articles.Handler, cc comments.Handler) {

	secrateJwt := configs.ImportSetting().JWTSECRET

	c.POST("/register", ur.Register()) // register -> umum (boleh diakses semua orang)
	c.POST("/login", ur.Login())
	c.PUT("/users", ur.UpdateUserAccount(), echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	c.DELETE("/users", ur.DeleteUserAccount(), echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	c.GET("/articles", ac.GetArticles())
	c.GET("/comments", cc.GetComments())

	c.POST("/uploads", ac.UploadImage())

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

	b := c.Group("/comments")
	b.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(secrateJwt),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))

	b.POST("", cc.CreateComments())
	b.DELETE("/:id", cc.DeleteComments())
	a.GET("/:id", ac.GetArticlesByID())

}
