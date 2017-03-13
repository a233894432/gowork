/*
* 	Author:diogo
	time:2017年3月6日14:32:56
*/
package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"
)

func main() {
	app := iris.New()
	// Adapt the "httprouter", faster,
	// but it has limits on named path parameters' validation,
	// you can adapt "gorillamux" if you need regexp path validation!
	app.Adapt(httprouter.New())
	app.Adapt(iris.DevLogger())

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	app.Use(customLogger)

	app.HandleFunc("GET", "/", func(ctx *iris.Context) {
		ctx.Writef("hello world\n")
	})
	// log http errors
	errorLogger := logger.New()

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Writef("My Custom 404 error page ")
	})

	app.Listen(":8080")

}
