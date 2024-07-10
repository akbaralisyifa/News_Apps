package comments

import "github.com/labstack/echo/v4"

type Comment struct {
	ID        uint
	UserID    uint
	ArticleID uint
	Comments  string
}

type Handler interface {
	CreateComments() echo.HandlerFunc
	GetComments() echo.HandlerFunc
	DeleteComments() echo.HandlerFunc
}

type Services interface {
	CreateComments(newComments Comment) error
	GetComments() ([]Comment, error)
	DeleteComments(id, userid uint) error
}

type Query interface {
	GetComments() ([]Comment, error)
	CreateComments(newComments Comment) error
	DeleteComments(id, userid uint) error
}

// initial validator
type CommentsValidate struct {
	Comments string `validate:"required"`
	// Content string `validate:"required"`
}
