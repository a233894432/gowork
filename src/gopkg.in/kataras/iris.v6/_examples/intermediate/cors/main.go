package main

import (
	"log"
	"os"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

var myLogFile *os.File

func init() {
	// open an output file
	f, err := os.Create("logs.txt")
	if err != nil {
		panic(err)
	}
	myLogFile = f
}
func myFileLogger() iris.LoggerPolicy {

	// you can use a *File or an io.Writer,
	// we want to log with timestamps so we use the log.New.
	myLogger := log.New(myLogFile, "", log.LstdFlags)

	// the logger is just a func,
	// will be used in runtime
	return func(mode iris.LogMode, message string) {
		// optionally, check for production or development log message mode
		// two modes: iris.ProdMode and iris.DevMode
		if mode == iris.DevMode {
			// log only production-mode log messages
			myLogger.Println(message)
		}
	}
}

func main() {

	defer func() {
		if err := myLogFile.Close(); err != nil {
			panic(err)
		}
	}()

	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(myFileLogger())
	app.Adapt(httprouter.New())

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	app.Adapt(crs) // this line should be added
	// adaptor supports cors allowed methods, middleware does not.

	// if you want per-route-only cors
	// then you should check https://github.com/iris-contrib/middleware/tree/master/cors

	v1 := app.Party("/api/v1")
	{
		v1.Post("/home", func(c *iris.Context) {
			app.Log(iris.DevMode, "home")
			c.WriteString("Hello from /home")
		})
		v1.Get("/g", func(c *iris.Context) {
			app.Log(iris.DevMode, "g")
			c.WriteString("Hello from /home")
		})
		v1.Post("/h", func(c *iris.Context) {
			app.Log(iris.DevMode, "h")
			c.WriteString("Hello from /home")
		})
	}

	app.Listen(":8080")
}
