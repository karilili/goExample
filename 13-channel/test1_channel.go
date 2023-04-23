package main

import (
	"fmt"
	"time"
)

func main() {
	//定义一个channel
	c := make(chan int, 1)
	//没有写cap的长度，则就是一个无缓冲chan
	//发送者必须等到接受者接收了数据后才取消阻塞
	//cap长度哪怕为1，也是一个有缓冲的chan

	go func() {
		defer fmt.Println("goroutine结束") //这行代码执行永远在num:=<-c这之后
		fmt.Println("goroutine 正在运行....1")
		c <- 666 //将666发送到c
		fmt.Println("goroutine 正在运行....2")
		//此处也会等待。直到管道另一头得到数据
		//管道，没有缓冲
	}()

	fmt.Println("before main get num")
	time.Sleep(100 * time.Second)
	//channel具备同步功能。此时会阻塞直到管道得到数据
	num := <-c //从c中接受数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}
