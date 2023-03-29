package main

import "fmt"

func foo1(a string, b int) int { //没有返回值不用写类型
	fmt.Println("----foo1----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

func foo2(a string, b int) (int, int) { //匿名的返回值
	fmt.Println("----foo2----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

func foo3(a string, b int) (r1 int, r2 int) { //有名称的返回值，有作用域
	fmt.Println("----foo3----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//返回值一开始作为局部变量而存在
	//r1 r2 属于foo3的形参 初始化默认的值是0，为了避免野指针的出现
	//r1 r2 作用域 是foo3整个函数体的{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

func foo4(a string, b int) (r1, r2 int) { //返回值一样可以简写
	fmt.Println("----foo4----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c=", c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

	ret1, ret2 = foo4("foo4", 444)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

}
