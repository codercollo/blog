package view

import (
	"net/url"

	"github.com/codercollo/blog/internal/modules/user/helpers"
	"github.com/codercollo/blog/pkg/converters"
	"github.com/codercollo/blog/pkg/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_DATA"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["OLD"] = flattenURLValues(converters.StringToUrlValue(sessions.Flash(c, "old")))
	data["AUTH"] = helpers.Auth(c)
	return data
}

func flattenURLValues(vals url.Values) map[string]string {
	flat := make(map[string]string)
	for key, slice := range vals {
		if len(slice) > 0 {
			flat[key] = slice[0]
		}
	}
	return flat
}
