package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	//  CSV 파일 생성
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	// 모든 작업 끝나고 닫기 지연
	defer csvFile.Close()

	// 모든 게시글 슬라이스 생성
	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	// csv 쓰기를 생성
	writer := csv.NewWriter(csvFile)

	// 모든 글에서 글을 하나씩 꺼내서
	for _, post := range allPosts {
		// 문자열 슬라이스를 만들어서
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		// 씀
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// 다 끝나면 날려버림
	writer.Flush()

	// 써놓은 CSV 파일을 열기
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 읽기 생성
	reader := csv.NewReader(file)
	// 필드당 레코드 -1 으로.. 이거 무슨 옵션인데 뭔지 모르겠네
	reader.FieldsPerRecord = -1
	// 모두 읽어재낌
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	// 그걸 하나씩 꺼내서
	for _, item := range record {
		// id 를 문자에서 숫자로 바꾸고
		id, _ := strconv.ParseInt(item[0], 0, 0)
		// Post 구조체에 담아서 생성하고는
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		// 다시 하나씩 차곡차곡 담음
		posts = append(posts, post)
	}

	// 첫번째 것의 내용을 출력
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
