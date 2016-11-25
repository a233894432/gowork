package main

import "github.com/kataras/iris"

func main() {

	iris.Get("/", func(c *iris.Context) {
		c.HTML(iris.StatusOK, "<h1> Hello World! </h1>")
	})

	iris.Listen(":8080")
}
