package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1024)
	// <input type="file" name="uploaded"> 을 읽어서 파일 헤더에 넣는다.
	fileHeader := r.MultipartForm.File["uploaded"][0]
	// 파일 헤더를 열고
	file, err := fileHeader.Open()
	// 오류가 없으면
	if err == nil {
		// 다 읽어서 data 에 저장하고
		data, err := ioutil.ReadAll(file)
		// 오류 없으면
		if err == nil {
			// 응답 라이터에다가 쓴다
			fmt.Print(w,string(data))

			// 아래처럼 하면 파일로 다운된다.
			//fmt.Fprintln(w, string(data))
		}
	}


	// 멀티파트 폼에서 키-값 쌍만 얻으려면

	// url 의 key-value
	// <form action="http://localhost:8080/process?hello=world&thread=123" method="post" enctype="multipart/form-data">
	fmt.Fprintln(w,"(1)",r.FormValue("hello"))

	// 폼 내부의 key-value
	//<input type="text" name="hello" value="sau sheong"/>
	//<input type="text" name="post" value="456"/>
	fmt.Fprintln(w,"(2)",r.PostFormValue("hello"))
	fmt.Fprintln(w,"(3)",r.PostForm)

	// 자동으로 r.ParseMultipartForm 호출되어 보여진다.
	fmt.Fprintln(w,"(4)", r.MultipartForm)

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
