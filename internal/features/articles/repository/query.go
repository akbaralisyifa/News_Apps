package repository

import "gorm.io/gorm"

type ArticleModel struct {
	db *gorm.DB
}

// func NewUserModel(connection *gorm.DB) Article.Query {
// 	return &ArticleModel{
// 		db: connection,
// 	}
// }
