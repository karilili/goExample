package main

import "fmt"

func main() {
	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束") //这行代码执行永远在num:=<-c这之后
		fmt.Println("goroutine 正在运行....")
		c <- 666 //将666发送到c
		//此处也会等待。直到管道另一头得到数据
		//管道，没有缓冲
	}()

	//channel具备同步功能。此时会阻塞直到管道得到数据
	num := <-c //从c中接受数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}
