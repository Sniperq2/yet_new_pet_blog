package models

type Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Posts struct {
	PostsList []Post `json:"posts"`
}
