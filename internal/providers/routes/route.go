package routes

import (
	homeRoutes "github.com/codercollo/blog/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
}
