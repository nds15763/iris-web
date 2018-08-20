package ViewModel

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
