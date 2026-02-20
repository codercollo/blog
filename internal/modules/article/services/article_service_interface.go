package services

import (
	"github.com/codercollo/blog/internal/modules/article/requests/articles"
	ArticleResponse "github.com/codercollo/blog/internal/modules/article/responses"
	UserResponse "github.com/codercollo/blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}
