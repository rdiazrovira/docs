package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
)

var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env: ENV,
		})

		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("version", "0.8.0")
				return next(c)
			}
		})
		app.GET("/", HomeHandler)
		app.GET("/docs/{name}", Docs)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
	}
	return app
}
