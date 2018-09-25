package main

import (
	"fmt"
	"net/http"
)

// 헤더는 키-값, Allow 가 키면 GET, POST ... 값이 있음..
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	y := r.Header["Accept-Encoding"] // 문자열 얻기
	z := r.Header.Get("Accept-Encoding") // 콤마로 분리된 값 목록 얻기
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, y)
	fmt.Fprintln(w, z)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
