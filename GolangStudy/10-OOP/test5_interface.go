package main

import "fmt"

// int string float struct 都实现了interface{}接口，因此interface{}这种类型可以引用任意的数据类型
// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{} 该如何区分，此时引用的底层数据类型到底是什么

	//go语言给interface{}提供了”类型断言“机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type,value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{auth: "Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)

}
