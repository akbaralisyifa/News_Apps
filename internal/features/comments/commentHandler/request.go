package commentHandler

import (
	"newsapps/internal/features/comments"
)

type CommentsRequest struct {
	UserID    uint   `json:"user_id"`
	ArticleID uint   `json:"article"`
	Comment   string `json:"comment"`
}

func ToRequeteComments(cr CommentsRequest) comments.Comment {
	return comments.Comment{
		UserID:    cr.UserID,
		ArticleID: cr.ArticleID,
		Comments:  cr.Comment,
	}
}
