package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			//如果c可写，则该case就会进来
			x = y
			y = x + y
		case <-quit:
			//如果quit可读
			fmt.Println("quit")
			return

			//下述default会打印很多很多，看来不应该轻易地写default，或者说default最好也是可以阻塞的
			//default:
			//	fmt.Println("not c or quit, just test")
		}
	}
	fmt.Println("fibonacii end1")
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	//main go
	fibonacii(c, quit) //channel是引用传递
	fmt.Println("fibonacii end2")
}
