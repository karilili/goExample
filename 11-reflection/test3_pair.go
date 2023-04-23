package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体的类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	//b:pair<type:Book,value:Book{}地址>
	b := &Book{}

	//r:pair<type:,value:>
	var r Reader
	//r:pair<type:Book,value:Book{}地址>
	r = b

	r.ReadBook()

	var w Writer
	//r:pair<type:Book,value:Book{}地址>
	w = r.(Writer) //此处断言为什么会成功， 因为 w r 具体的type是一样的
	//很显然，pair的赋值传递不变性，都是建立在实现了多态的基础上的

	w.WriteBook()
}
