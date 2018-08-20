package Route

import (
	"net/http"
	"strings"

	"../Controller/Cart"
	"../Controller/Home"
	"github.com/kataras/iris"
)

func RouteInit(ctx *iris.Application) {
	//首页，购物车，订单，评价，商品，秒杀，用户，通知，活动，积分，广告
	ctx.Get("/", Home.Index) //首页
	ctx.Get("/Home", Home.Index)
	ctx.Get("/Home/GetItemList", Home.GetItemList) //购物车
	ctx.Get("/Cart", Cart.Index)                   //购物车
	ctx.Get("/Cart/AddCart", Cart.AddCart)         //购物车
	ctx.Get("/Order", Cart.Index)                  //订单
	ctx.Get("/Comment", Cart.Index)                //评价
	ctx.Get("/Goods", Cart.Index)                  //商品
	ctx.Get("/User", Cart.Index)                   //用户
	ctx.Get("/Notify", Cart.Index)                 //通知
	ctx.Get("/Active", Cart.Index)                 //活动
	ctx.Get("/BP", Cart.Index)                     //积分
	ctx.Get("/AD", Cart.Index)                     //广告

	// Serve a controller based on the root Router, "/".
	//Home.HomeControlle值拷贝 是不是得改
	// mvc.New(ctx).Handle(new(Base.BaseController))

	// use of the .StaticHandler
	// which is the same as StaticWeb but it doesn't
	// registers the route, it just returns the handler.
	fileServer := ctx.StaticHandler("./Content", false, false)
	fileServerHF := ctx.StaticHandler("./", false, false)
	// fileServerHF := ctx.StaticHandler("./Content/Home_files", false, false)

	// wrap the router with a native net/http handler.
	// if url does not contain any "." (i.e: .css, .js...)
	// (depends on the app , you may need to add more file-server exceptions),
	// then the handler will execute the router that is responsible for the
	// registered routes (look "/" and "/profile/{username}")
	// if not then it will serve the files based on the root "/" path.
	ctx.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path
		// Note that if path has suffix of "index.html" it will auto-permant redirect to the "/",
		// so our first handler will be executed instead.

		if !strings.Contains(path, ".") {
			// if it's not a resource then continue to the router as normally. <-- IMPORTANT
			router(w, r)
			return
		}
		// acquire and release a context in order to use it to execute
		// our file server
		// remember: we use net/http.Handler because here we are in the "low-level", before the router itself.
		tctx := ctx.ContextPool.Acquire(w, r)
		fileServer(tctx)
		fileServerHF(tctx)
		ctx.ContextPool.Release(tctx)
	})
}
func BeforeActivation() {

}
