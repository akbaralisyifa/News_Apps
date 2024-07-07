package commentHandler

import (
	"newsapps/internal/features/comments"
	"newsapps/internal/helper"
	"newsapps/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CommentssController struct {
	srv comments.Services
}

// GetComments implements comments.Handler.
func (hc *CommentssController) GetComments() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := hc.srv.GetComments()

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get data", result))
	}
}

// DeleteComments implements comments.Handler.
func (hc *CommentssController) DeleteComments() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))

		var userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token))

		err := hc.srv.DeleteComments(uint(id), uint(userID))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success delete data", nil))
	}
}

// CreateComments implements comments.Handler.
func (hc *CommentssController) CreateComments() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input CommentsRequest
		err := c.Bind(&input)

		if err != nil {
			c.Logger().Error("create comments error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		var userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token))
		// (("user").(*jwt.Token))

		if input.UserID != uint(userID) {
			return c.JSON(401, helper.ResponseFormat(400, "Unauthorized", nil))
		}

		err = hc.srv.CreateComments(ToRequeteComments(input))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(201, helper.ResponseFormat(201, "success create articles", nil))
	}
}

func NewCommentsController(s comments.Services) comments.Handler {
	return &CommentssController{
		srv: s,
	}
}
