package services

import (
	ArticleModel "github.com/codercollo/blog/internal/modules/article/models"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() []ArticleModel.Article
	GetStoriesArticles() []ArticleModel.Article
}
