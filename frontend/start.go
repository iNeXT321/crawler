package main

import (
	"net/http"
	"test/crawler/frontend/controller"
)
const templateName = "crawler/frontend/view/template.html"
func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		templateName))
	err := http.ListenAndServe(":8882", nil)
	if err != nil{
		panic(err)
	}
}
