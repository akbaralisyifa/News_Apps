package main

import (
	"log"
	"newsapps/configs"
	"newsapps/internal/features/articles"
	"newsapps/internal/features/articles/handler"
	"newsapps/internal/features/articles/repository"
	articleRepository "newsapps/internal/features/articles/repository"
	"newsapps/internal/features/articles/services"
	commentRepository "newsapps/internal/features/comments/repository"
	"newsapps/internal/features/users"
	userHandler "newsapps/internal/features/users/handler"
	userRepository "newsapps/internal/features/users/repository"
	userServices "newsapps/internal/features/users/services"
	"newsapps/internal/routes"
	"newsapps/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitUserRoute(db *gorm.DB) users.Handler {
	um := userRepository.NewUserModel(db)
	pu := utils.NewPasswordUtility()
	jwt := utils.NewJwtUtility()
	vldt := utils.NewAccountUtility(*validator.New())
	us := userServices.NewUserService(um, vldt, pu, jwt)
	uc := userHandler.NewUserController(us)
	return uc
}

func InitialArticleRouter(db *gorm.DB) articles.Handler {
	am := repository.NewArticleModel(db)
	as := services.NewArticlesServices(am)
	ac := handler.NewArticlesController(as)

	return ac
}

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		log.Fatal("Stop program, masalah database", err.Error())
		return
	}

	err = connection.AutoMigrate(&userRepository.User{}, &articleRepository.Articles{}, &commentRepository.Comments{})

	if err != nil {
		log.Fatal("Stop program, masalah database ", err.Error())
		return
	}
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup

	ur := InitUserRoute(connection)
	ac := InitialArticleRouter(connection)

	routes.InitRoute(e, ur, ac)
	e.Logger.Fatal(e.Start(":8000"))
}
