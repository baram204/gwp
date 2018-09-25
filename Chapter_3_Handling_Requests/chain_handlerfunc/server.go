package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// 핸들러 함수 선언
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// 핸들러 함수를 받아서 자기 할 일을 하고 핸들러 함수를 반환하는 함수 선언
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		// 인자로 받은 핸들러 함수에다가 리스폰스와 리퀘스르를 넘겨준다.
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// hello 핸들러함수가 처리되기 전에 현재 포인터의 이름을 reflect 로 가져와서 콘솔에 출력한다
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}
