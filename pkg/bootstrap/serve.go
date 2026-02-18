package bootstrap

import (
	"github.com/codercollo/blog/pkg/config"
	"github.com/codercollo/blog/pkg/html"
	"github.com/codercollo/blog/pkg/routing"
	"github.com/codercollo/blog/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
