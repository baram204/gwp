package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

// 컨텐츠 타입이 없으니 512 바이트로 html 판단
func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	// 바이트 배열을 취한다.
	// write 호출까지 컨텐츠 타입이 없으면 첫번째 512 바이트 데이터에 대한 만큼 콘텐츠 타입으로 감지해 사용한다.
	w.Write([]byte(str))
}

// 501 에러코드를 찍음
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	// 호출 후에 헤더 수정 막음
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

// 302 찍고 로케이션 키에 주소 찍어서 리다이렉트 시키기
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	// WriteHdader 전에 헤더 수정을 해야함.
	w.WriteHeader(302)
}

// 컨텐츠 타입 json 으로 설정 후 구조체를 json 으로 마샬해서 쓰기
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
