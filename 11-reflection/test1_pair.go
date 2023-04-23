//golang中变量的构造
//type value
//static type  concrete type
//基本类型int string...    interface所指向的具体数据类型，系统runtime机制看得见的类型
//如果定义一个interface{}变量，则其有值，及一个具体类型
//type 要么是static type、要么是concrete type
//反射就是通过变量来找到其类型

package main

import "fmt"

func main() {
	var a string
	//pair<statictype:string,value:"aceld">
	a = "aceld"

	//一个value中一定有type指针，一个value指针，将a赋值给value，就是将a自身的两个指针不变地传递到value的两个指针上
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)

}
