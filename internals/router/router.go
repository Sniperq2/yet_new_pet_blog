package router

import (
	"net/http"
	"yet_new_pet_blog/internals/controllers"
)

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	static := http.FileServer(http.Dir("./web/static"))
	r.Handle("/static/", http.StripPrefix("/static/", static))
	r.HandleFunc("/", controllers.Index)
	return r
}
