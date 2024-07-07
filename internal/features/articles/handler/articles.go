package handler

import (
	"newsapps/internal/features/articles"
	"newsapps/internal/helper"
	"newsapps/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ArticlesController struct {
	srv articles.Services
}

func NewArticlesController(s articles.Services) articles.Handler {
	return &ArticlesController{
		srv : s,
	}
}


func(ac *ArticlesController) CreateArticles() echo.HandlerFunc {
	return func(c echo.Context) error{

		var input ArticlesRequeste;
		err := c.Bind(&input);

		if err != nil {
			c.Logger().Error("create article error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}
		
		var userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token));

		// if input.UserID != uint(userID){
		// 	return c.JSON(401, helper.ResponseFormat(400, "Unauthorized", nil))
		// }

		input.UserID = uint(userID)

		err = ac.srv.CreateArticles(ToRequeteArticles(input));

		if err != nil{
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(201, helper.ResponseFormat(201, "success create articles", nil))
	}
}


func(ac *ArticlesController) GetArticles() echo.HandlerFunc{
	return func(c echo.Context) error {

		result, err := ac.srv.GetArticles();

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(404, helper.ResponseFormat(404, "article not found", nil))
			}
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return	c.JSON(200, helper.ResponseFormat(200, "success get data", ToArticlesResponse(result)))
	}
}

// get articles by id
func(ac *ArticlesController) GetArticlesByID() echo.HandlerFunc{
	return func (c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		result, err := ac.srv.GetArticlesByID(uint(id));

		if err != nil{
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get data", ToArticlesResponseById(result)))
	}
}


func(ac *ArticlesController) UpdateArticles() echo.HandlerFunc{
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		article, _ := ac.srv.GetArticlesByID(uint(id))

		var input ArticlesRequeste
		err := c.Bind(&input);

		if err != nil {
			c.Logger().Error("create article error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		var userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token));

		input.UserID = uint(userID);

		if input.UserID != uint(userID){
			return c.JSON(401, helper.ResponseFormat(400, "Unauthorized", nil))
		}

		err = ac.srv.UpdateArticles(article.ID, ToRequeteArticles(input))

		if err != nil{
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success update data", nil))
	}
}


func(ac *ArticlesController) DeleteArticles() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		var userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token));

		// // get article by Id
		// article, err := ac.srv.GetArticlesByID(uint(id));

		// if err != nil {
		// 	if err == gorm.ErrRecordNotFound {
		// 		return c.JSON(404, helper.ResponseFormat(404, "article not found", nil))
		// 	}
		// 	return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		// }

		// if  article.UserID != uint(userID) {
		// 	return c.JSON(401, helper.ResponseFormat(401, "Unauthorized", nil))
		// }

		// fungsi delete article
		err := ac.srv.DeleteArticles(uint(id), uint(userID));

		if err != nil{
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success delete data", nil))
	}
}