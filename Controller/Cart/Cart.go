package Cart

import (
	"time"

	"github.com/kataras/iris"
)

type CartItem struct {
	ID        int
	Fgroup    int
	Sort      int
	Name      string
	ShowName  string
	Msg       string
	Interval  int
	CreatedAt string
	UpdateAt  string
	Status    string
}

func Index(ctx iris.Context) {

	ctx.ViewData("message", "Hello Cart!")
	ctx.View("/Cart/Index.html")
}

func Showf(ctx iris.Context) {

	// ctx.JSON()
}

func CartList(ctx iris.Context) {
	//购物车列表

	ctx.JSON(nil)
}
func ReadCookieWhenUserLogin() {
	//用户执行登陆操作时读取cookie信息
	//将cookie中的商品信息录入到数据库中
	cookieList := ReadCartCookie()
	if len(cookieList) > 0 {
		//是
		//将cookie中的商品信息录入到数据库中
	} else {
		//否
		//继续完成登录
	}
}
func ReadCartCookie() []CartItem {
	return []CartItem{}
}
func AddCart(ctx iris.Context) {
	//添加到购物车

	//是否登录
	isLogin := true

	if isLogin {
		//是
	} else {
		//否
		//将商品信息添加到cookie
		AddCartToCookie()
	}
	//查询已有的购物车信息

	//将商品信息添加到购物车中

	//持久化数据
	c := make(chan bool)
	go WriteIntoDb(c)

	ctx.JSON(<-c)
}

func AddCartToCookie() {

}

func WriteIntoDb(c chan bool) {
	println("Start Write Into Db")
	time.Sleep(1 * time.Second)
	println("Write Finished")
	c <- true
}
func Update() {

}
func ReadFromDb() CartItem {
	clist := CartItem{
		ID:        1,
		Fgroup:    4,
		Sort:      1,
		Name:      "每日好货",
		ShowName:  "",
		Msg:       "msg",
		Interval:  1,
		CreatedAt: "2018-07-20 15:49:28",
		UpdateAt:  "2018-07-20 15:49:28",
		Status:    "n",
	}
	return clist
}
