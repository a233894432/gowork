package main

import (
	"myProject/router"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

func main() {

	// set the isDevelopment
	iris.Config.IsDevelopment = true

	// set the template engine
	iris.UseTemplate(html.New(html.Config{Layout: "layout.html"})).Directory("./templates", ".html")
	// set static folder(s)

	iris.StaticWeb("/public", "./static", 1)

	// set the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
	})

	// set the global middlewares
	iris.Use(logger.New())
	// register the routes & the public API
	registerRoutes()

	iris.Listen(":8080")
}

func registerRoutes() {
	// register index using a 'Handler'
	iris.Handle("GET", "/", router.Index())

	// this is other way to declare a route
	// using a 'HandlerFunc'
	iris.Get("/about", router.About)

	// Dynamic route

	iris.Get("/profile/:username", router.Profile)("user-profile")
	// user-profile is the custom,optional, route's Name: with this we can use the {{ url "user-profile" $username}} inside userlist.html

	// iris.Get("/all", routes.UserList)
}
