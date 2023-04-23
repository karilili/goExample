package main

import "fmt"

//了解slice的追加以及cap len的概念

func main() {
	var numbers = make([]int, 3, 5)

	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1，numbers len=4, [0,0,0,1] ,cap=5
	numbers = append(numbers, 1)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素2，numbers len=5, [0,0,0,1] ,cap=5
	numbers = append(numbers, 2)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	//向容量cap已满的slice追加元素，numbers len=6, [0,0,0,1] ,cap=10  //动态开辟空间，将cap翻倍
	numbers = append(numbers, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("======")
	var numbers2 = make([]int, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers2), cap(numbers2), numbers2)

}
