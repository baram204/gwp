package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	// 지정된 길이만큼의 바이트를 생성한다
	body := make([]byte, len)

	// 요청 본문을 읽어서 바이트 변수에 저장한다
	r.Body.Read(body)
	// 요청받은 내용을 그대로 두번 뿌려준다.
	fmt.Fprintln(w, string(body))
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
