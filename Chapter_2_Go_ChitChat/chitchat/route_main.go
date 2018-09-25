package main

import (
	// 로컬 패키지 임포트
	"./data"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

// 루트에 접근했을 때 실행될 handleFunc 인 index
func index(writer http.ResponseWriter, request *http.Request) {
	// 모든 쓰레드 자료를 가져오는데
	threads, err := data.Threads()
	// 오류 있으면
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	// 오류없으면
	} else {
		// 세션검사를 해서
		_, err := session(writer, request)
		// 통과를 못하면 로그인 한 사람이아니니 public-navbar 보여주고
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		// 통과하면 로그인 한 사용자니 private-navbar 를 보여준다.
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
