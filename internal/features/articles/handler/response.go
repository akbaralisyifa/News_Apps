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
	Comments []string `json:"comments"`
}

func ToArticlesResponse(input []articles.Article) []ArticlesResponse {
	var responses []ArticlesResponse

	for _, val := range input {
		responses = append(responses, ArticlesResponse{
			ID:       val.ID,
			UserID:   val.UserID,
			Title:    val.Title,
			Content:  val.Content,
			Image:    val.Image,
			Comments: val.Comments,
		})
	}

	return responses
}

func ToArticlesResponseById(input articles.Article) ArticlesResponse {
	return ArticlesResponse{
		ID:       input.ID,
		UserID:   input.UserID,
		Title:    input.Title,
		Content:  input.Content,
		Image:    input.Image,
		Comments: input.Comments,
	}
}
