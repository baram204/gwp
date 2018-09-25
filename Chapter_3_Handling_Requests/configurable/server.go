package main

import (
	"net/http"
)

// 서버 Struct 를 사용한 설정
func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
