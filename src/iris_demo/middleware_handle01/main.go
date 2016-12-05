package main

import "github.com/kataras/iris"

func firstMiddleware(ctx *iris.Context) {
	ctx.Write("1. This is the first middleware, before any of route handlers \n")
	ctx.Next()
}

func secondMiddleware(ctx *iris.Context) {
	ctx.Write("2. This is the second middleware, before the '/' route handler \n")
	ctx.Next()
}

func thirdMiddleware(ctx *iris.Context) {
	ctx.Write("3. This is the 3rd middleware, after the '/' route handler \n")
	ctx.Next()
}

func lastAlwaysMiddleware(ctx *iris.Context) {
	ctx.Write("4. This is ALWAYS the last Handler \n")
}

func main() {
	// 初始化调用这个中间件
	iris.UseFunc(firstMiddleware)
	// 完成请求的时候 启动这个中间件
	iris.DoneFunc(lastAlwaysMiddleware)

	iris.Get("/", secondMiddleware, func(ctx *iris.Context) {
		ctx.Write("Hello from / \n")
		ctx.Next() // .Next because we 're using the third middleware after that, and lastAlwaysMiddleware also
	}, thirdMiddleware)

	iris.Listen(":8080")

}
