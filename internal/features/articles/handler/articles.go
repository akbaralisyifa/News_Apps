package handler

import (
	"io"
	"log"
	"net/http"
	"newsapps/internal/features/articles"
	"newsapps/internal/helper"
	"newsapps/internal/utils"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ArticlesController struct {
	srv articles.Services
}

func NewArticlesController(s articles.Services) articles.Handler {
	return &ArticlesController{
		srv: s,
	}
}

func (ac *ArticlesController) CreateArticles() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input ArticlesRequeste
		err := c.Bind(&input)

		if err != nil {
			c.Logger().Error("create article error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		var userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token))

		// if input.UserID != uint(userID){
		// 	return c.JSON(401, helper.ResponseFormat(400, "Unauthorized", nil))
		// }

		input.UserID = uint(userID)

		err = ac.srv.CreateArticles(ToRequeteArticles(input))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(201, helper.ResponseFormat(201, "success create articles", nil))
	}
}

type Response struct {
	Message string
	Data    interface{}
}

func (ac *ArticlesController) UploadImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		isSuccess := true
		var response Response
		var fileType, fileName string
		var fileSize int64
		file, err := c.FormFile("file")
		if err != nil {
			isSuccess = false
			log.Println("cant find file", err.Error())
			return c.JSON(http.StatusFailedDependency, Response{})
		} else {
			src, err := file.Open()
			if err != nil {
				isSuccess = false
				log.Println("cant Open", err.Error())
				return c.JSON(http.StatusGone, Response{})

			} else {
				fileByte, _ := io.ReadAll(src)
				fileType = http.DetectContentType(fileByte)

				if fileType == "application/pdf" {
					fileName = "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
				} else {
					fileName = "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"

				}

				err = os.WriteFile(file.Filename, fileByte, 0777)
				if err != nil {
					isSuccess = false
					log.Println("cant write file", err.Error())
					return c.JSON(http.StatusBadGateway, Response{})
				} else {
					fileSize = file.Size
				}
			}
			defer src.Close()
		}
		if isSuccess {
			response = Response{
				Message: "succes upload file",
				Data: struct {
					Filename  string
					Filetype  string
					FilesSize int64
				}{
					Filename:  fileName,
					Filetype:  fileType,
					FilesSize: fileSize,
				},
			}
		} else {
			response = Response{
				Message: "failed upload file",
			}
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (ac *ArticlesController) GetArticles() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := ac.srv.GetArticles()

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(404, helper.ResponseFormat(404, "article not found", nil))
			}
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get data", ToArticlesResponse(result)))
	}
}

// get articles by id
func (ac *ArticlesController) GetArticlesByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		result, err := ac.srv.GetArticlesByID(uint(id))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get data", ToArticlesResponseById(result)))
	}
}

func (ac *ArticlesController) UpdateArticles() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		article, _ := ac.srv.GetArticlesByID(uint(id))

		var input ArticlesRequeste
		err := c.Bind(&input)

		if err != nil {
			c.Logger().Error("create article error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		var userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token))

		input.UserID = uint(userID)

		if input.UserID != uint(userID) {
			return c.JSON(401, helper.ResponseFormat(400, "Unauthorized", nil))
		}

		err = ac.srv.UpdateArticles(article.ID, ToRequeteArticles(input))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success update data", nil))
	}
}

func (ac *ArticlesController) DeleteArticles() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		var userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token))

		// fungsi delete article
		err := ac.srv.DeleteArticles(uint(id), uint(userID))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success delete data", nil))
	}
}
