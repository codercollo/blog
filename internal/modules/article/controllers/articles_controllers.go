package controllers

import (
	"net/http"
	"strconv"

	ArticleService "github.com/codercollo/blog/internal/modules/article/services"
	"github.com/codercollo/blog/pkg/html"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "internal/templates/errors/html/500", gin.H{
			"title":   "Server error",
			"message": "error converting the id",
		})
		return
	}

	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "internal/templates/errors/html/404", gin.H{
			"title":   "Page not found",
			"message": err.Error(),
		})
		return
	}

	html.Render(c, http.StatusOK, "internal/modules/article/html/show", gin.H{
		"title":   "Show article",
		"article": article,
	})
}
