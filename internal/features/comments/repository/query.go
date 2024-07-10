package repository

import (
	"newsapps/internal/features/comments"

	"gorm.io/gorm"
)

type CommentModel struct {
	db *gorm.DB
}

// GetComments implements comments.Query.
func (qc *CommentModel) GetComments() ([]comments.Comment, error) {
	var result []comments.Comment

	err := qc.db.Find(&result).Error

	if err != nil {
		return []comments.Comment{}, err
	}

	return result, nil
}

// DeleteComments implements comments.Query.
func (qc *CommentModel) DeleteComments(id, userid uint) error {
	qry := qc.db.Where("id = ? AND user_id = ?", id, userid).Delete(&Comments{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// CreateComments implements comments.Query.
func (qc *CommentModel) CreateComments(newComments comments.Comment) error {
	resultData := ToArticlesQuery(newComments)
	err := qc.db.Create(&resultData).Error

	if err != nil {
		return err
	}

	return nil
}

func NewCommentModel(connection *gorm.DB) comments.Query {
	return &CommentModel{
		db: connection,
	}
}
