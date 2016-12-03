/*  project: router
    author:diogo
    time: 2016/11/22-15:41
*/
package router

import "github.com/kataras/iris"

func Profile(ctx *iris.Context) {
	username := ctx.Param("username")
	ctx.MustRender("profile.html", iris.Map{"Username": username})
}
