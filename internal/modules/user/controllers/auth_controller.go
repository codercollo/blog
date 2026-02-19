package controllers

import (
	"log"
	"net/http"

	"github.com/codercollo/blog/internal/modules/user/requests/auth"
	UserService "github.com/codercollo/blog/internal/modules/user/services"
	"github.com/codercollo/blog/pkg/converters"
	"github.com/codercollo/blog/pkg/errors"
	"github.com/codercollo/blog/pkg/html"
	"github.com/codercollo/blog/pkg/old"
	"github.com/codercollo/blog/pkg/sessions"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "internal/modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get() ))

		c.Redirect(http.StatusFound, "/register")
		return
	}
	user, err := controller.userService.Create(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}
	log.Printf("The user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}
