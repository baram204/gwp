package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

// json 마샬/언마샬 용 Post 구조체
type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

// main handler function
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	// 요청 메소드에 따라서 분기를 탄다
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Retrieve a post
// GET /post/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// /post/2  -> 2, /post -> post
	var urlPath = r.URL.Path
	var basePath = path.Base(urlPath)
	id, err := strconv.Atoi(basePath) // 문자열2 에서 숫자 2로 변환.
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	// 구조체를 전송하기 위해 바이너리로 마샬한다
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	// 해더 컨텐츠 타입을 json 으로 한 다음에 응답에 써버린다.
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// Create a post
// POST /post/
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	// 요청 본문을 읽어서 저장한다.
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	// 구조체, 이건 왜 또 값 타입이야?
	var post Post
	// 오잉? 줄 때는 주소를 주네.. 아 헷갈려
	json.Unmarshal(body, &post)

	// 내부 생성
	err = post.create()
	if err != nil {
		return
	}
	// 성공했다고 해더를 작성하기
	w.WriteHeader(200)
	return
}

// Update a post
// PUT /post/1
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {

	// 베이스 경로 구해서 문자열로 컨버팅
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 해당 아이디 내용 가져오기
	post, err := retrieve(id)
	if err != nil {
		return
	}

	// 요청 본문 읽기
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	// 언마셜해서 바이트배열의 본문 내용을 구조체화 하기
	json.Unmarshal(body, &post)

	// 갱신
	err = post.update()
	if err != nil {
		return
	}

	// 완료 헤더보내기
	w.WriteHeader(200)
	return
}

// Delete a post
// DELETE /post/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {

	// 베이스 경로 문자에서 숫자로
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	// 일단 가져오기
	post, err := retrieve(id)
	if err != nil {
		return
	}

	// 이걸 구지 post 에서 해야하나? 음..
	err = post.delete()
	if err != nil {
		return
	}

	// 완료 때리기
	w.WriteHeader(200)
	return
}
