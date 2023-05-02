package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}
// 结构体是作为参数的值传递
// 如果想在函数里面改变结果体数据内容，需要传入指针
func changeBook(book *Books) {
	book.title = "book1_change"
}

func main() {
	var book1 Books
	book1.title = "book1"
	book1.author = "jiuyi"
	book1.bookId = 1
	changeBook(&book1) //传入指针
	fmt.Println(book1) //{book1_change jiuyi  1}


}
