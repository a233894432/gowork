/*  project: router
    author:diogo
    time: 2016/11/22-15:41
*/
package router

import (
	"github.com/kataras/iris"
)

type index struct {
	Title   string
	Message string
	Body 	string
}

// Index index is the page for GET: / route
func Index() *index {
	return &index{
		Title:   "iris sample - index",
		Message: "This is just a sample index, it's empty because Iris doesnt wants influences!",
		Body: "name",
	}
}

func (i *index) Serve(ctx *iris.Context) {
	if err := ctx.Render("index.html", i); err != nil {
		// ctx.EmitError(iris.StatusInternalServerError) =>
		ctx.Panic()
	}
}
