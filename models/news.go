package models

type News struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ShortDesc string `json:"short_desc" db:"short_desc"`
	ExpireAt string `json:"expire_at" db:"expire_at"`
	Img string `json:"img"`
	FullDesc string `json:"full_desc" db:"full_desc"`
	Status bool `json:"status"`
}

type MiniNews struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ShortDesc string `json:"short_desc" db:"short_desc"`
	Img string `json:"img"`
}
