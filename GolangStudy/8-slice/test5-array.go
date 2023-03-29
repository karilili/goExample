package main

import "fmt"

//了解切片的截取

func main() {
	s := []int{1, 2, 3} //len=3,cap=3

	//[0,2]

	s1 := s[0:2] //[1,2]   //这样写，s s1是指向同一处内存的

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	//copy，深拷贝，可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2=[0,0,0]  //先创建s2的空间

	//将s中的值，依次拷贝到s2中
	copy(s2, s)

	fmt.Println(s)
	fmt.Println(s2)

}
