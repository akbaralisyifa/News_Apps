package repository

import (
	"newsapps/internal/features/articles"

	"gorm.io/gorm"
)

type ArticleModel struct {
	db *gorm.DB
};

func NewArticleModel(connection *gorm.DB) articles.Query {
	return &ArticleModel{
		db: connection,
	}
}

// Get All Articles
func (am *ArticleModel) GetArticles() ([]articles.Article, error){
	var result []articles.Article

	err := am.db.Find(&result).Error;

	if err != nil {
		return []articles.Article{}, err;
	}

	return result, nil
}

func (am *ArticleModel) CreateArticles(newArticles articles.Article) error {
	resultData := ToArticlesQuery(newArticles)
	err := am.db.Create(&resultData).Error;

	if err != nil {
		return err;
	}

	return nil;
}

func (am *ArticleModel) UpdateArticles(id uint, updateArticles articles.Article) error {
	qry := am.db.Model(articles.Article{}).Where("id = ?", id).Updates(updateArticles)

	if qry.Error != nil {
		return qry.Error
	}

	return nil;
}

func (am *ArticleModel) DeleteArticles(id uint) error {
	qry := am.db.Where("id = ?", id).Delete(Articles{});

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil;
}
