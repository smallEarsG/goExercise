package main

import (
	"net/http"
	"reptile/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler("frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
