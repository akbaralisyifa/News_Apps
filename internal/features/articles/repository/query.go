package repository

import (
	"fmt"
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
	qry := `SELECT * FROM "newsapps"."articles"`
	err := am.db.Debug().Raw(qry).Scan(&result).Error;

	if err != nil {
		return []articles.Article{}, err;
	}

	return result, nil
}



// Get Article by ID
func (am *ArticleModel) GetArticlesByID(id uint) (articles.Article, error) {
	var result Articles;

	// Muat artikel beserta komentarnya
	err := am.db.Preload("Comments").First(&result, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return articles.Article{}, fmt.Errorf("article with ID %d not found", id)
		}
		return articles.Article{}, err
	}

		return result.ToArticlesEntityComments(), nil
}



// Create Articles
func (am *ArticleModel) CreateArticles(newArticles articles.Article) error {
	resultData := ToArticlesQuery(newArticles)
	err := am.db.Create(&resultData).Error;

	if err != nil {
		return err;
	}

	return nil;
}

// Update Aritcles
func (am *ArticleModel) UpdateArticles(id uint, updateArticles articles.Article) error {
	cnvData := ToArticlesQuery(updateArticles)

	// Jika salah satu saja yang di update
	if updateArticles.Title != "" {
		cnvData.Title = updateArticles.Title
	}

	if updateArticles.Content != "" {
		cnvData.Content = updateArticles.Content
	}

	if updateArticles.Image != "" {
		cnvData.Image = updateArticles.Image
	}

	qry := am.db.Model(Articles{}).Where("id = ?", id).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound;
	}
	
	return nil;
}

// Delete Articles
func (am *ArticleModel) DeleteArticles(id uint, userID uint) error {
	qry := am.db.Where("id = ? AND user_id = ?", id, userID).Delete(&Articles{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil;
}
