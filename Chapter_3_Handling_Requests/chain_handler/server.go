package main

import (
	"fmt"
	"net/http"
)

// 핸들러
type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// 핸들러를 받아서 핸들러를 반환하는 함수
// 중간에 핸들러 호출을 콘솔에 출력한다.
func log(h http.Handler) http.Handler {
	// HandlerFunc 는 Handler 를 반환한다.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		// 지 할일 다 하면 핸들러의 serveHTTP 메소드를 반환한다.
		h.ServeHTTP(w, r)
		// HandlerFunc 메소드의 인자로 들어가면, Handler 로 반환된다.
	})
}

// 핸들러를 받아서 핸들러를 반환하는 함수
// 중간에 사용자가 인증되었는지를 확인
func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// some code to make sure the user is authorized
		// 지 할일 다 하면 핸들러의 serveHTTP 에다가 w,r 넣어준다.
		h.ServeHTTP(w, r)
		// HandlerFunc 메소드의 인자로 들어가면, Handler 로 반환된다.
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	hello := HelloHandler{}
	// 핸들러를 log 에 넘겨주면 핸드러를 반환하고 그것이 protect 에 넘겨주면 또 handler 를 반환한다.
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
}
