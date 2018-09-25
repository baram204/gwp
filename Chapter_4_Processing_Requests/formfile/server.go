package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// multipart 어쩌구 하지 않아도 바로 파일을 가져와서 처리할 수 있다.
	// 첫번재 파일만 가져온다. 업로드할 파일이 하나인 경우에 더 빠르다
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		// 만약 원한다면 읽어온 파일의 내용을 검색할 수도 있을 것이다.
		if err == nil {
			// 또한 파일과 파일 헤더 모두를 동시에 반환한다
			// 그래서 다운로드 되는 대신, 브라우저에서 헤더에 맞춰 제대로 렌더링을 한다.
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
