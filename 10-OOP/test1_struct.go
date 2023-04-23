package main

import "fmt"

//go也是面向对象的语言，先熟悉一下struct的概念

// 声明一种行的数据类型 myint，是int的一个别名
type myint int

// 定义一个结构体  复杂数据类型
type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) {
	//传递一个book的副本  这是值传递
	book.auth = "666"
}

func changeBook2(book *Book) {
	//指针传递
	book.auth = "777"
}

func main() {
	// var a myint = 10
	// fmt.Println("a=", a)
	// fmt.Printf("type of a=%T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1) //打印任何一种格式数据

	changeBook1(book1)
	fmt.Println("-------")
	fmt.Printf("%v\n", book1) //打印任何一种格式数据

	changeBook2(&book1)
	fmt.Println("-------")
	fmt.Printf("%v\n", book1) //打印任何一种格式数据

	//接下来看看如何把一个结构体变成一个类
}
