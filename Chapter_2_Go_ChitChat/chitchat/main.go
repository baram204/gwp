package main

import (
	"net/http"
	"time"
)

func main() {

	p("ChitChat", version(), "started at", config.Address)

	// 정적 파일 다루기
	mux := http.NewServeMux() // 멀티 플렉서 1:n 매핑
	files := http.FileServer(http.Dir(config.Static)) // config 구조체 Static 필드에 담긴 값을 복사해서
	// 파일 서버가 그곳을 가리키게 한다. "Static"         : "public"
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	// /static/ 들어오면, /static/ 제거한 경로를 파일서버에 넘겨준다.
	// http://xxx.com/static/hello.css -> /public/hello.css

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
