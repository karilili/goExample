package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close可以关闭一个channel
		close(c) //close也是go内置的关键字
		//如果不close，go能检测出死锁，因为main一直在读取，被卡死，但已经不再有写入了
	}()

	/*
		for {
			//ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
			if data, ok := <-c; ok { //if的一种句式，先执行data,ok:=<-c，再if判断ok
				fmt.Println(data)
			} else {
				break
			}
		}
	*/

	//channel与关键字的配合使用
	//1、channel与range
	//可以使用range来迭代不断操作channel
	//range在c有数据时，会返回数据
	//range在c没有数据但未关闭时，会阻塞
	for data := range c {
		fmt.Println(data)
	}

	//2、channel与select
	//单流程下一个go只能监控一个channel的状态，select可以完成监控多个channel的状态
	//类似于C的监听多个IO
	//select具备多路channel的监控状态功能
	//看下个文件

	fmt.Println("main finished")

	//1、channel不需要频繁关闭，只有当没有任何需要发送的数据了，或者想显示结束数据发送，才需要关闭
	//2、关闭channel后，无法向channel再发送数据（引发pannic错误后导致接受立即发挥零值）
	//3、关闭channel后，可以继续从channel接收数据
	//4、对于nil channel（声明一个chan数据类型，但是不通过make创建），无论收发都会被阻塞
}
