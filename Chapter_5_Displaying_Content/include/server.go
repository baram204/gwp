package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")

	// 템플릿 이름을 지정하지 않으면 첫번째 할당한 템플릿이 실행된다.
	t.Execute(w, "Hello World!")

	// 원하는 템플릿을 실행하려면 이름을 지정한다.
	// 	t.Execute(w, "t2.html", "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
