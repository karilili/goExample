package main

import (
	"fmt"
	"time"
)

func main() {

	// //用goc创建承载一个形参为空，返回值为空的一个函数
	// go func() {
	// 	defer fmt.Println("A.defer")
	// 	func() {
	// 		defer fmt.Println("B.defer")
	// 		//return //退出这个匿名函数
	// 		//退出当前goroutine
	// 		runtime.Goexit()

	// 		fmt.Println("B")
	// 	}() //如果没有()则这就是一个函数定义，加上()后，就会调用该函数

	// }()
	// fmt.Println("A")

	go func(a int, b int) bool {
		fmt.Println("a= ", a, "b= ", b)
		return true //不支持在go前加一个 flag:= 来接受函数结果，需要用到channel管道机制
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
