package main

import (
	"fmt"

	con "./config"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())
	// first parameter is the request path
	// second is the operating system directory
	app.StaticWeb("/", "../frontend")

	fmt.Println(con.Pv)
	app.Listen(":8080")
}
