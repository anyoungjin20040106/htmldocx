package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func Warning(w http.ResponseWriter, msg string) {
	warning, _ := template.ParseFiles("warning.html")
	data := struct {
		MSG string
	}{
		msg,
	}
	warning.Execute(w, data)
}
func main() {
	port := fmt.Sprintf(":%d", 10000)
	fileServer := http.FileServer(http.Dir("./chapter"))
	http.Handle("/view/", http.StripPrefix("/view/", fileServer))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/doc", func(w http.ResponseWriter, r *http.Request) {
		pw := r.FormValue("pw")
		if r.Method != "POST" {
			Warning(w, "암호를 입력해주세요")
		} else if pw != os.Getenv("pw") {
			Warning(w, "암호가 틀렸습니다")
		} else {
			http.ServeFile(w, r, "doc.html")
		}
	})
	log.Println("Server started on " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
