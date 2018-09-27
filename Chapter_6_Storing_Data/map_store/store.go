package main

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// Post 의 포인터를 매핑하는 이유는, 가져올 때마다 복사하지 않기 위함
var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	// 저자가 같으면 슬라이스 내부에 쌓인다. append 로 누적시키는 형태 생소하다 진짜...
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {

	// map 은 사용하기 전에 반드시 make 해줘야 한다.
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
