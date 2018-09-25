package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	// 템플릿의 함수 맵을 만들고
	funcMap := template.FuncMap{"fdate": formatDate}
	// 템플릿 생성시 해당 함수를 붙인다.
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl.html")
	// 그리고 함수 실행
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
