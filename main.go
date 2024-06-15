package main

import (
	"net/http"
)

func main() {
	c := http.FileServer(http.Dir("chapter"))
	http.Handle("/c/", http.StripPrefix("/c/", c))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.ListenAndServe(":1234", nil)
}
