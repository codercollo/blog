package repositories

import (
	ArticleModel "github.com/codercollo/blog/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
	Create(article ArticleModel.Article) ArticleModel.Article
}
