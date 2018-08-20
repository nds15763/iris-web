package Model

type CLItem struct {
	CLId       int
	CLImg      string
	CLProduct  string
	CLPrice    int
	CLQuantity int
	CLTotal    int
}

//Home类
type CartModel struct {
	B         BaseModel //继承
	CList     []CLItem
	CSubTotal int
	CTotal    int
}
