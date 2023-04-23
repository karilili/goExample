package main

import "fmt"

func main() {
	//声明myMap1是一种map类型 key是string value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	//在使用map前，需要先用make给map分配数据空间 //如果不使用make分配空间就创建元素，则会报错
	myMap1 = make(map[string]string, 10) //如果用完了空间，再创建一个元素，则会动态开辟空间

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python" //顺序是乱的，是用的一种哈希的方式来存储数据

	fmt.Println(myMap1)

	//===>第二种声明方式
	myMap2 := make(map[int]string) //make其实不必要写元素个数，默认就会分配一些空间
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python" //顺序是乱的，是用的一种哈希的方式来存储数据
	fmt.Println(myMap2)

	//===>第三种声明方式   //如果需要初始化，这种用得较多
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "python", //最后一行依然要加上,
	}
	fmt.Println(myMap3)

}
