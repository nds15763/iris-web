package About

import (
	"github.com/kataras/iris"
)

func Index(ctx iris.Context) {
	ctx.ViewData("message", "Hello world!")
	ctx.View("/About/Index.html")
}
