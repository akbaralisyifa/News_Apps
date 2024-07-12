package main

import (
	"fmt"
	"log"
	"newsapps/configs"
	"newsapps/internal/features/articles"
	"newsapps/internal/features/articles/handler"
	"newsapps/internal/features/articles/repository"
	articleRepository "newsapps/internal/features/articles/repository"
	articleService "newsapps/internal/features/articles/services"
	"newsapps/internal/features/comments"
	"newsapps/internal/features/comments/commentHandler"
	commentRepository "newsapps/internal/features/comments/repository"
	"newsapps/internal/features/comments/services"
	"newsapps/internal/features/users"
	userHandler "newsapps/internal/features/users/handler"
	userRepository "newsapps/internal/features/users/repository"
	userServices "newsapps/internal/features/users/services"
	"newsapps/internal/routes"
	"newsapps/internal/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	storage_go "github.com/supabase-community/storage-go"
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
	vldt := utils.NewAccountUtility(*validator.New())
	am := repository.NewArticleModel(db)
	as := articleService.NewArticlesServices(am, vldt)
	ac := handler.NewArticlesController(as)

	return ac
}

func InitialCommentRouter(db *gorm.DB) comments.Handler {
	cm := commentRepository.NewCommentModel(db)
	cs := services.NewCommentServices(cm)
	cc := commentHandler.NewCommentsController(cs)

	return cc
}

func main() {

	storageClient := storage_go.NewClient("https://ciotjxaqztwqcxlzhvis.supabase.co/storage/v1/s3", "eddefca34cea8d434883622a10b44150", nil)

	bucket, err := storageClient.GetBucket("newsAppsImage")

	if err != nil {
		log.Println("error xx", err.Error())
	}

	var builder strings.Builder

	builder.WriteString("[Bucket]\n")
	builder.WriteString(fmt.Sprintf("Id = %s\n", bucket.Id))
	builder.WriteString(fmt.Sprintf("Name = %s\n", bucket.Name))
	builder.WriteString(fmt.Sprintf("Owner = %s\n", bucket.Owner))
	builder.WriteString(fmt.Sprintf("Public = %t\n", bucket.Public))

	if bucket.FileSizeLimit != nil {
		builder.WriteString(fmt.Sprintf("FileSizeLimit = %d\n", *bucket.FileSizeLimit))
	} else {
		builder.WriteString("FileSizeLimit = \n")
	}

	if len(bucket.AllowedMimeTypes) > 0 {
		builder.WriteString(fmt.Sprintf("AllowedMimeTypes = %s\n", strings.Join(bucket.AllowedMimeTypes, ",")))
	} else {
		builder.WriteString("AllowedMimeTypes = \n")
	}

	builder.WriteString(fmt.Sprintf("CreatedAt = %s\n", bucket.CreatedAt))
	builder.WriteString(fmt.Sprintf("UpdatedAt = %s\n", bucket.UpdatedAt))
	return

	/*-----------------------------------------------------------------------------*/
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
	cc := InitialCommentRouter(connection)

	routes.InitRoute(e, ur, ac, cc)
	e.Logger.Fatal(e.Start(":8000"))
}
