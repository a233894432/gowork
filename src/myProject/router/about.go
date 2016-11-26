/*  project: router
    author:diogo
    time: 2016年11月26日15:58:55
*/
package router

import (
	"github.com/kataras/iris"
)

// About the handler func, we don't use struct here
// this is the simplest method to declare a route
func About(ctx *iris.Context) {
	// MustRender, same as Render but sends status 500 internal server error if rendering failed
	ctx.MustRender("about.html", nil)
}
