package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Use(&myMiddleware{})

	iris.Get("/home", func(c *iris.Context) {
		c.HTML(iris.StatusOK, "<h1>Hello from /home </h1>")
	})

	iris.Listen(":8080")
}

type myMiddleware struct{}

func (m myMiddleware) Serve(c *iris.Context) {
	shouldContinueToTheNextHandler := false

	if shouldContinueToTheNextHandler {

		c.Next()
	} else {
		c.Text(403, "Forbidden !!")
	}

}
