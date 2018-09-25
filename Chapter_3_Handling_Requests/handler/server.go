package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		// 핸들러는 인터페이스고 ServeHTTP(http.res..., * http.Request) 메소드 모양을 가지는 모든 것들을 수용한다.
		Handler: &handler,
	}
	server.ListenAndServe()
}
