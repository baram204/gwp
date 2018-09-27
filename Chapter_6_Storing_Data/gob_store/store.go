package main

import (
	"bytes"
	// 스트립 관리, 인코더 디코더 간 변환, 
	// 직렬화 및 데이터 전송 지원
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// store data
func store(data interface{}, filename string) {

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	// 엔코더가 자료를 인코딩 한 결과를 버퍼에 저장한다
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	// 바이너리 자료가 담긴 버퍼를 바이트로 바꿔서 바이너리 파일에 쓴다
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// load the data
func load(data interface{}, filename string) {
	// 바이너리 파일을 읽는다
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	// 디코더는 바이너리 자료를 디코딩 해서 버퍼에 담는다
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	// 포스트 구조체 생성
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	// 인코딩 후 바이너리 파일로 저장
	store(post, "post1")
	var postRead Post
	// 파일에서 바이너리를 읽어서 디코딩 후 postRead 에 저장
	load(&postRead, "post1")
	fmt.Println(postRead)
}
