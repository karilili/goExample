package main

import "fmt"

func printArray(myArray [4]int) { //固定长度的数组在传参的时候，是严格匹配参数类型的；另外，传入时是值拷贝
	for index, value := range myArray {
		fmt.Println("index=", index, ",value=", value)
	}
	myArray[0] = 111
}

func main() {
	//固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4} //默认初始化值为0，前四个元素才有值
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 { //range是一个关键字，会根据遍历集合的不同，返回不同的值，对于数组/切片/动态数组，则是返回下标以及下标处的值
		fmt.Println("index=", index, ",value=", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3) //不同长度的数组类型不一样，myArray1是不能传入该参数的
	fmt.Println("------")
	for index, value := range myArray3 {
		fmt.Println("index=", index, ",value=", value)
	}

}
