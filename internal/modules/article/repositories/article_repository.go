package repositories

import (
	ArticleModel "github.com/codercollo/blog/internal/modules/article/models"
	"github.com/codercollo/blog/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article
	articleRepository.DB.Limit(limit).Order("rand()").Preload("User").Find(&articles)
	return articles
}
