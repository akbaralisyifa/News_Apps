package services

import (
	"errors"
	"log"
	"newsapps/internal/features/articles"

	"github.com/go-playground/validator/v10"
)

type ArticlesServices struct {
	qry 	 articles.Query
	validate *validator.Validate
}

func NewArticlesServices( q articles.Query) articles.Services {
	return &ArticlesServices{
		qry: q,
		validate : validator.New(),
	}
}

func (as *ArticlesServices) CreateArticles(newArticles articles.Article) error{
	// validate articles
	err:= as.validate.Struct(
		&articles.ARticlesValidate{
			Title: newArticles.Title, 
			Content: newArticles.Content,
		});

	if err != nil {
		return errors.New("validate is not empty")
	}

	// query articles
	err = as.qry.CreateArticles(newArticles);

	if err != nil {
		log.Print("Create Articles Sql Error", err.Error())
		return errors.New("an error occurred on the server while processing data")
	}

	return nil;
}

func(as *ArticlesServices) GetArticles()( []articles.Article, error){
	result, err := as.qry.GetArticles();

	if err != nil {
		log.Print("Get Articles Sql Error", err.Error())
		return []articles.Article{}, errors.New("an error occurred on the server while processing data")
	}
	return result, nil;
}

func(as *ArticlesServices) GetArticlesByID(id uint)(articles.Article, error){
	result, err := as.qry.GetArticlesByID(id);

	if err != nil {
		log.Print("Get Article By ID Sql Error", err.Error())
		return articles.Article{}, errors.New("an error occurred on the server while processing data")
	}

	return result, nil;

}

func(as *ArticlesServices) UpdateArticles(id uint, updateArticles articles.Article) error {

	err := as.qry.UpdateArticles(id, updateArticles);

	if err != nil {
		log.Print("Update Articles Sql Error", err.Error())
		return errors.New("an error occurred on the server while processing data")
	}

	return nil;
}

func(as *ArticlesServices) DeleteArticles(id uint) error {
	err := as.qry.DeleteArticles(id);

	if err != nil {
		log.Print("Delete Articles Sql Error", err.Error())
		return errors.New("an error occurred on the server while processing data")
	}

	return nil
}

