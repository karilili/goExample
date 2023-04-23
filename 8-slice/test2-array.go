package main

import "fmt"

func printArray(myArray []int) {
	//动态数组/切片是引用传递，实际上传递的是一个指针。可以理解为，动态数组本身就是一个指针
	//_表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组，切片 slice
	fmt.Printf("myArray type is %T\n", myArray)
	printArray(myArray)

	fmt.Println("=======")
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
