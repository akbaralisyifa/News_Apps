package handler

import (
	"newsapps/internal/features/articles"
)

type ArticlesResponse struct {
	ID       uint     `json:"id"`
	UserID   uint     `json:"user_id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Image    string   `json:"image"`
	Comments []CommentResponse `json:"comments"` // Change to slice of CommentResponse
}

type CommentResponse struct {
	UserID  uint   `json:"user_id"`
	Comment string `json:"comment"`
}

// Converts a slice of articles.Article to a slice of ArticlesResponse.
func ToArticlesResponse(input []articles.Article) []ArticlesResponse {
	var responses []ArticlesResponse

	for _, val := range input {
		responses = append(responses, ToArticlesResponseById(val)) // Reuse the single article conversion
	}

	return responses
}

// Converts a single articles.Article to an ArticlesResponse.
func ToArticlesResponseById(input articles.Article) ArticlesResponse {
	comments := make([]CommentResponse, len(input.Comments))
	for i, comment := range input.Comments {
		comments[i] = CommentResponse{
			UserID:  comment.UserID,
			Comment: comment.Comment,
		}
	}

	return ArticlesResponse{
		ID:       input.ID,
		UserID:   input.UserID,
		Title:    input.Title,
		Content:  input.Content,
		Image:    input.Image,
		Comments: comments,
	}
}
