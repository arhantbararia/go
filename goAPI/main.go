package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting REST API in go! ")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "return all comments")
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "posted a new comment")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintln(w, fmt.Sprintf("comment from id : %s", id))

	})

	//if no METHOD is written before path, default is GET

	err := http.ListenAndServe("localhost:8080", mux)

	if err != nil {
		fmt.Println(err.Error())
	}
}
