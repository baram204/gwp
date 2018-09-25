package main

import (
	"net/http"
)

// 핸들러 매개변수가 nil 이면 defaultServerMux 가 사용됨
func main() {
	http.ListenAndServe("", nil)
}
