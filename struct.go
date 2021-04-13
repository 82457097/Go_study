package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	fmt.Println(Books{"Go 语言", "hello world", "Go 语言教程", 123456})
	fmt.Println(Books{title : "Go 语言", author : "hello world", subject : "Go 语言教程", book_id : 123456})
	fmt.Println(Books{title : "Go 语言", author : "hello world"})
}