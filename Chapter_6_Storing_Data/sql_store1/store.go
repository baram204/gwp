package main

import (
	"database/sql"
	"fmt"
	//	임포트 후 바로 사용하지 않고, init() 시점에 사용을 해야할 경우 일단 언더바로 오류를 스킵
	_ "github.com/lib/pq"

)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

// connect to the Db
func init() {
	var err error
	// 패키지-> 매인 순으로 init() 이 실행되는데 여튼, 처음 컴파일 때는 사용이 안 되고, 패키지를 읽거나 하는 나중 시점에
	// 사용이 되므로 ... 언더바로 임포트 처리

	// DB 접속
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// get all posts
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// 게시글 DB 저장 후 ID 를 가져와서 구조체에 넣는다.
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update a post
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// Delete a post
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

// Delete all posts
func DeleteAll() (err error) {
	_, err = Db.Exec("delete from posts")
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	// Create a post
	fmt.Println(post) // {0 Hello World! Sau Sheong}
	post.Create()
	fmt.Println(post) // {1 Hello World! Sau Sheong}

	// Get one post
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost) // {1 Hello World! Sau Sheong}

	// Update the post
	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	// Get all posts
	posts, _ := Posts(10)
	fmt.Println(posts) // [{1 Bonjour Monde! Pierre}]

	// Delete the post
	readPost.Delete()

	// Get all posts
	posts, _ = Posts(10)
	fmt.Println(posts) // []

	// Delete all posts
  // DeleteAll()
}
