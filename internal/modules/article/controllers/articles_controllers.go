package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/codercollo/blog/internal/modules/article/requests/articles" // ✅ correct request
	ArticleService "github.com/codercollo/blog/internal/modules/article/services"
	"github.com/codercollo/blog/internal/modules/user/helpers"
	"github.com/codercollo/blog/pkg/converters"
	"github.com/codercollo/blog/pkg/errors"
	"github.com/codercollo/blog/pkg/html"
	"github.com/codercollo/blog/pkg/old"
	"github.com/codercollo/blog/pkg/sessions"
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

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "internal/modules/article/html/create", gin.H{
		"title": "Create article",
	})
}

func (controller *Controller) Store(c *gin.Context) {
	var request articles.StoreRequest // ✅ was auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))
		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(c) // returns responses.User

	article, err := controller.articleService.StoreAsUser(request, user)
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID)) // ✅ fixed /article → /articles
}
