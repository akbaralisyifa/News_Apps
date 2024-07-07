package services

import (
	"errors"
	"log"
	"newsapps/internal/features/comments"

	"github.com/go-playground/validator/v10"
)

type CommentsServices struct {
	qry      comments.Query
	validate *validator.Validate
}

// GetComments implements comments.Services.
func (sc *CommentsServices) GetComments() ([]comments.Comment, error) {
	result, err := sc.qry.GetComments()

	if err != nil {
		log.Print("Get Comments Sql Error", err.Error())
		return []comments.Comment{}, errors.New("an error occurred on the server while processing data")
	}
	return result, nil
}

// DeleteComments implements comments.Services.
func (sc *CommentsServices) DeleteComments(id, userid uint) error {
	err := sc.qry.DeleteComments(id, userid)

	if err != nil {
		log.Print("Delete Comment Sql Error", err.Error())
		return errors.New("an error occurred on the server while processing data")
	}

	return nil
}

// CreateComments implements comments.Services.
func (sc *CommentsServices) CreateComments(newComments comments.Comment) error {
	err := sc.validate.Struct(
		&comments.CommentsValidate{
			Comments: newComments.Comments,
		})

	if err != nil {
		return errors.New("validate is not empty")
	}

	// query articles
	err = sc.qry.CreateComments(newComments)
	if err != nil {
		log.Print("Create comments Sql Error", err.Error())
		return errors.New("an error occurred on the server while processing data")
	}

	return nil
}

func NewCommentServices(q comments.Query) comments.Services {
	return &CommentsServices{
		qry:      q,
		validate: validator.New(),
	}
}
