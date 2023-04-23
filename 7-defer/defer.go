package main

import "fmt"

/*
知识点一：defer的执行顺序
*/
func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

/*
知识点二：defer和return谁先谁后
*/
func deferfunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnfunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	defer deferfunc() //...
	return returnfunc()
}

func main() {
	// //写入defer关键字
	// // 类似于析构函数，或者final关键字
	// defer fmt.Println("main end1") //类似于栈的概念，先入后出
	// defer fmt.Println("main end2")

	// fmt.Println("main::hello go 1")
	// fmt.Println("main::hello go 2")

	// func1()
	// func2()
	// func3()

	// defer func1() //类似于栈的概念，先入后出
	// defer func2()
	// defer func3()

	returnAndDefer() //先执行return。函数生命周期结束后，再执行defer
}
