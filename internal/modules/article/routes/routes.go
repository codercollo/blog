package routes

import (
	articleCtrl "github.com/codercollo/blog/internal/modules/article/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articlesController := articleCtrl.New()
	router.GET("/articles/:id", articlesController.Show)
}
