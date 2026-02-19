package routes

import (
	articleRoutes "github.com/codercollo/blog/internal/modules/article/routes"
	homeRoutes "github.com/codercollo/blog/internal/modules/home/routes"
	userRoutes "github.com/codercollo/blog/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
