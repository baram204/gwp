package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// 해석한다.
	r.ParseForm()
	// 해석된 타입은 Map 이라 키-값 쌍은 맞는데 정렬이 안 됨
	fmt.Fprintln(w, r.Form)
	// map[post:[456] thread:[123] hello:[sau sheong world]]
	fmt.Fprintln(w, r.Form["hello"])
	// [sau sheong world]

	// PostForm 은 URL 의 키-값만 가져오고 폼의 키-값 쌍은 얻지 못한다.
	// form-urlencoded 만 지원하기 때문에
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
