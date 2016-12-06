package main

import (
	"github.com/kataras/go-template/django"
	"github.com/kataras/iris"
)

func main() {
	iris.Config.IsDevelopment = true
	iris.Config.Gzip = true

	iris.UseTemplate(django.New()).Directory("./templates", ".html")

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("mypage.html", map[string]interface{}{"username": "iris", "is_admin": true})
	})

	iris.Listen(":8080")
}
