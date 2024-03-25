package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "welcome to the homepage")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	data := "some data from the api"
	fmt.Println(w, data)
}
