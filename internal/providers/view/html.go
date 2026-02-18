package view

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(data gin.H) gin.H {
	data["APP_DATA"] = viper.Get("App.Name")
	return data
}
