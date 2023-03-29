package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("User is called...")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Aceld", 18}
	user.Call()
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType.Name())
	fmt.Println("inputType is:", inputType)

	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	//通过type 获取里面的字段
	//1、获取interface的reflect.Type，通过Type获取到NumField，进行遍历
	//2、得到每个field，数据类型
	//3、通过field有一个Interface()方法 得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s:%v = %v\n", field.Name, field.Type, value)
	}

	//通过type 获取里面的方法，调用
	fmt.Println("inputType.NumMethod() = ", inputType.NumMethod())
	fmt.Println("inputType.NumField() = ", inputType.NumField())
	for j := 0; j < inputType.NumMethod(); j++ {
		n := inputType.Method(j)
		fmt.Printf("%s:%v\n", n.Name, n.Type)
	}
	//根据弹幕（不是很了解为何如此）
	//go 1.17中，NumMethod()已经不可用，只能导出接口的不公共的方法
	//实测，如果方法中不使用指针传递，则反射可以获取到该方法
}
