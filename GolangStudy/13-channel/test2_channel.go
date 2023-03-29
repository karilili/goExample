package main

import (
	"fmt"
	// "time"
)

func main() {
	var CLen int = 200
	c := make(chan int, CLen) //带有缓冲的channel
	fmt.Println("len(c) = ", len(c), ",cap(c)", cap(c))
	go func() {
		defer fmt.Println("子go结束")

		for i := 0; i < CLen+5; i++ {
			c <- i
			fmt.Println("子go程正在运行，发送的元素=", i, " len(c)=", len(c), " cap(c)=", cap(c))
		}
		//当channel已经满，再向里面写数据，就会阻塞
		//当channel为空，从里面取数据也会阻塞
		//阻塞与非阻塞状态的切换似乎没有那么快
		//当channel已经非满时，仍要过一会儿，才会恢复过来，以至于maingo已经取了更多的数据了
	}()

	// time.Sleep(2 * time.Second)
	fmt.Println("main begin get channel num")
	//channel或多或少会卡顿程序，当这一句执行后，子go可能放了1百多个数据了
	//main go才开始取数据

	for i := 0; i < CLen; i++ {
		num := <-c
		fmt.Println("num=", num, " len(c)=", len(c), " cap(c)=", cap(c))
	}
	fmt.Println("main go结束")
}
