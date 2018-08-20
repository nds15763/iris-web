package Home

import (
	"../../Model"
	"../../Service"
	"github.com/kataras/iris"
)

type HomeController struct {
	Service Service.MovieService
}

// type newModel map[int64]Model.HomeModel
// type repo Repositories.MovieRepository
// type movieService Service.MovieService

// func Init() {
// 	newModel := map[int64]Model.HomeModel{}
// 	// Create our movie repository with some (memory) data from the datasource.
// 	repo = Repositories.NewMovieRepository(newModel)

// 	// Create our movie service, we will bind it to the movie app's dependencies.
// 	movieService = Service.NewMovieService(repo)
// }

type itemList struct {
	ID     int
	Sort   int
	Name   string
	Status string
	Forums []forumsList
}

type forumsList struct {
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
type PageItem struct {
	ID           int    `json:"ID"`
	PrimaryImg   string `json:"pImg"`
	SecondaryImg string `json:"sImg"`
	ProductName  string `json:"pName"`
	Price        string `json:"price"`
}

type PageMap map[string]string

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *HomeController) GetHello() interface{} {

	return map[string]PageItem{
		"PItem1": PageItem{
			ID:           1,
			PrimaryImg:   "Home_files/product_pic_2.jpg",
			SecondaryImg: "Home_files/product_pic_1.jpg",
			ProductName:  "Donec non est at",
			Price:        "$111111111250.00",
		},
	}
}

func Index(ctx iris.Context) {
	res2D := &PageItem{
		ID:           1,
		PrimaryImg:   "Home_files/product_pic_2.jpg",
		SecondaryImg: "Home_files/product_pic_1.jpg",
		ProductName:  "啊",
		Price:        "$111111111250.00",
	}
	// res2B, _ := json.Marshal(res2D)
	// fmt.Println(string(res2B))

	ctx.ViewData("pItem", res2D)
	ctx.View("/Home/Index.html")
}

// Get returns list of the movies.
// Demo:
// curl -i http://localhost:8080/movies
//
// The correct way if you have sensitive data:
// func (c *MovieController) Get() (results []viewmodels.Movie) {
//     data := c.Service.GetAll()
//
//     for _, movie := range data {
//         results = append(results, viewmodels.Movie{movie})
//     }
//     return
// }
// otherwise just return the datamodels.
func (c *HomeController) Get() (results []Model.HomeModel) {
	// c.View("/Home/Index.html")
	return c.Service.GetAll()
}

// GetBy returns a movie.
// Demo:
// curl -i http://localhost:8080/movies/1
func (c *HomeController) GetBy(id int64) (movie Model.HomeModel, found bool) {
	return c.Service.GetByID(id) // it will throw 404 if not found.
}
func GetItemList(ctx iris.Context) {
	flist := []forumsList{
		forumsList{
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
		},
	}
	rtList := itemList{
		ID:     1,
		Sort:   1,
		Name:   "每日好货",
		Status: "n",
	}
	rtList.Forums = flist
	ctx.JSON(rtList)
}
