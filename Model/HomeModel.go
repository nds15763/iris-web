package Model

//Home类
type HomeModel struct {
	B     BaseModel //继承
	ID    int64     `json:"ID"`
	PImg  string    `json:"pImg"`
	SImg  string    `json:"sImg"`
	PName string    `json:"pName"`
	Price string    `json:"price"`
}
