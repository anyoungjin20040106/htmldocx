package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, msg string) {
	fail, err := template.ParseFiles("error.html")
	if err != nil {
		http.Error(w, "템플릿 에러", 505)
	}
	result := struct {
		MSG string
	}{
		MSG: msg,
	}
	fail.Execute(w, result)
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/doc", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			pw := r.FormValue("pw")
			fmt.Print(pw)
			if pw == "0106" {
				http.ServeFile(w, r, "doc.html")
			} else {
				err := "비밀번호를 확인해주세요\n당신이 입력한 비밀번호 : " + pw
				Error(w, err)
			}
		} else {
			Error(w, "비밀번호 입력해주세요")
		}
	})
	http.ListenAndServe(":1234", nil)
}
