package repository

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	UserID 		uint	`json:"user_id"`
	ArticleID 	uint	`json:"article_id"`
	Comment		string	`json:"comment"`
}
