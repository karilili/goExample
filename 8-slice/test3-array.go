package main

//slice的四种声明方式
import "fmt"

func main() {
	//声明slice1是一个切片，并且初始化，默认值是1，2，3，长度len是3
	//slice1 := []int{1, 2, 3}

	//声明slice1是一个切片，但是并没有给slice1分配空间
	var slice1 []int
	//slice1 = make([]int, 3) //通过make来开辟空间

	//声明slice2是一个切片，同时给slice2分配3个空间，初始化值是0
	// var slice2 []int = make([]int, 3)

	//声明slice3是一个切片，同时给slice3分配3个空间，初始化值是0，通过:=推导出slice3是一个切片
	// slice3 := make([]int, 3)

	// slice1[0] = 100
	// fmt.Printf("len=%d,slice=%v\n", len(slice1), slice1) //%v代表打印出所有信息
	// fmt.Printf("len=%d,slice=%v\n", len(slice2), slice2)
	// fmt.Printf("len=%d,slice=%v\n", len(slice3), slice3)

	//判断一个slice是否为0
	if slice1 == nil {
		fmt.Println("slice1是一个空切片")
	} else {
		fmt.Println("slice1s是有空间的")
	}

}
