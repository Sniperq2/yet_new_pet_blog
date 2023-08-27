package controllers

import (
	"errors"
	"net/http"
	"text/template"
	"yet_new_pet_blog/internals/repositories"
)

var (
	indexTemplate      = template.Must(template.ParseFiles("./web/views/index.html"))
	ErrMethodMustBeGet = errors.New("method must be get")
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrMethodMustBeGet.Error(), http.StatusInternalServerError)
		return
	}

	posts := repositories.NewPostsRepository()

	err := indexTemplate.Execute(w, posts.GetPosts())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
