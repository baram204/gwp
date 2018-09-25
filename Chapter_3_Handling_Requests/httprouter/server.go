package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// 세번째 인자로 httprouter.Params 를 넘김
func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
	fmt.Fprintf(w, "hello, %s!\n", p)
}

func main() {
	// mux 는 핸들러
	mux := httprouter.New()
	// handleFunc 대신, 함수를 사용
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		// 핸들러 넘김
		Handler: mux,
	}
	server.ListenAndServe()
}
