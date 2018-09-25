package main

import (
	"fmt"
	"net/http"
)

// struct 를 만들고 ServerHTTP 메소드를 선언해서 핸들러를 만들지 않아도
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// HandlerFunc 는 넘어온 함수를 핸들러로 변환 후 등록해준다.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
