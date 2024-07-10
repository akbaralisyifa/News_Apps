package services_test

import (
	"errors"
	"newsapps/internal/features/articles"
	"newsapps/internal/features/articles/services"
	"newsapps/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateArticle(t *testing.T) {
	qry := mocks.NewQuery(t)
	vldt := mocks.NewAccountUtilityInterface(t)
	srv := services.NewArticlesServices(qry, vldt);
	input := articles.Article{
		Title: "barre",UserID: 1, Content: "desc content", Image: "http://123.img",
	}

	t.Run("Success Create Articles", func(t *testing.T) {
		inputQry :=  articles.Article{Title: "barre",UserID: 1, Content: "desc content", Image: "http://123.img"}

		// test validasi
		vldt.On("CreateArticlesValidator", input.Title, input.Content).Return(nil).Once();
		// test query
		qry.On("CreateArticles", inputQry).Return(nil).Once();
		// func crate article ini klo sukses akan mengasilkan error nil
		err := srv.CreateArticles(input);

		assert.Nil(t, err);
	})

	t.Run("Validator Error", func(t *testing.T) {
		vldt.On("CreateArticlesValidator", input.Title, input.Content).Return(errors.New("validator error")).Once();

		err := srv.CreateArticles(input);

		assert.NotNil(t, err);
		assert.Equal(t, "validate is not empty", err.Error())
	})

	t.Run("Query error", func(t *testing.T) {
		inputQry :=  articles.Article{Title: "barre",UserID: 1, Content: "desc content", Image: "http://123.img"};

		vldt.On("CreateArticlesValidator", input.Title, input.Content).Return(nil).Once();

		qry.On("CreateArticles", inputQry).Return(errors.New("query error")).Once();

		err := srv.CreateArticles(input);

		assert.NotNil(t, err)
		assert.Equal(t, "an error occurred on the server while processing data", err.Error())
	})

}

func TestGetArticles(t *testing.T) {
	qry := mocks.NewQuery(t);
	vldt := mocks.NewAccountUtilityInterface(t)
	srv := services.NewArticlesServices(qry, vldt);

	t.Run("Success Get Data", func(t *testing.T) {
		articleList := []articles.Article{
			{ID: 1, Title: "indonesia", Content: "content indo", Image: "http://img.jpg"},
		}

		// test query 
		qry.On("GetArticles").Return(articleList, nil).Once();

		result, err := srv.GetArticles();

		assert.Nil(t, err);
		assert.Equal(t, articleList, result);
	})

	t.Run("Query Get All Error", func(t *testing.T) {
		// test qerry 
		qry.On("GetArticles").Return([]articles.Article{}, errors.New("query error")).Once()

		result, err := srv.GetArticles();

		assert.NotNil(t, err);
		assert.Equal(t, "an error occurred on the server while processing data", err.Error());
		assert.Empty(t, result);
	})
};

func TestGetArticlesById(t *testing.T) {
	qry := mocks.NewQuery(t);
	vldt := mocks.NewAccountUtilityInterface(t);
	srv := services.NewArticlesServices(qry, vldt);


	t.Run("Success Get Article By Id", func(t *testing.T) {
		article := articles.Article{ID: 1, Title: "indonesia", Content: "desc indoesia", Image: "http://img.png"};

		qry.On("GetArticlesByID", uint(1)).Return(article, nil).Once();

		result, err := srv.GetArticlesByID(1);

		assert.Nil(t, err);
		assert.Equal(t, article, result);
	});

	t.Run("Query Get By Id error", func(t *testing.T) {
		// query gagal
		qry.On("GetArticlesByID", uint(1)).Return(articles.Article{}, errors.New("query error")).Once();

		result, err := srv.GetArticlesByID(1);

		assert.NotNil(t, err)
		assert.Equal(t, "an error occurred on the server while processing data", err.Error())
		assert.Equal(t, articles.Article{}, result)
	});
};

func TestUpdateArticles(t *testing.T) {
	qry := mocks.NewQuery(t);
	vldt := mocks.NewAccountUtilityInterface(t)
	srv := services.NewArticlesServices(qry, vldt);

	t.Run("Success Update Article", func(t *testing.T) {
		updateData := articles.Article{Title: "update Indo", Content: "konoha negaraku", Image: "http://img.png"};

		qry.On("UpdateArticles", uint(1), updateData).Return(nil).Once();

		err := srv.UpdateArticles(1, updateData);

		assert.Nil(t, err);
	});

	t.Run("Query update error", func(t *testing.T) {
		updateData := articles.Article{Title: "update Indo", Content: "konoha negaraku", Image: "http://img.png"};

		qry.On("UpdateArticles", uint(1), updateData).Return(errors.New("query error")).Once();

		err := srv.UpdateArticles(1, updateData);

		assert.NotNil(t, err)
		assert.Equal(t, "an error occurred on the server while processing data", err.Error())
	})
};

func TestDeleteArticles(t *testing.T) {
	qry := mocks.NewQuery(t);
	vldt:= mocks.NewAccountUtilityInterface(t);
	srv := services.NewArticlesServices(qry, vldt);

	t.Run("Success Delete", func(t *testing.T) {
		// delete qry
		qry.On("DeleteArticles", uint(1), uint(1)).Return(nil).Once();

		err := srv.DeleteArticles(1, 1);

		assert.Nil(t, err);
	});

	t.Run("Qerry Delete Error", func(t *testing.T) {
		// error query 
		qry.On("DeleteArticles", uint(1), uint(1)).Return(errors.New("query error")).Once();

		err := srv.DeleteArticles(1, 1);

		assert.NotNil(t, err)
		assert.Equal(t, "an error occurred on the server while processing data", err.Error())
	})
}

