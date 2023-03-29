package main

//以后还可以使用go mod
//不查询GOPATH 可能是因为版本比较新 go env -w GO111MODULE=off
import (
	_ "GolangStudy/5-init/lib1" //别名 _表示匿名别名 （满足一种特殊需求，使用init函数但不使用接口）
	//不使用包中接口，则不能导入包，不会使用其中的init函数。但通过匿名别名可以在不使用接口的导入包
	mylib2 "GolangStudy/5-init/lib2"

	//. "GolangStudy/5-init/lib2" //前面加. 则不必使用包名.的方式来使用接口，直接用就行了 //不要轻易使用，避免同名歧义
	"fmt"
)

//函数名首字母大写，代表这是一个对外开放的函数
//如果首字母小写，则只能在当前包内使用

func main() {
	fmt.Println("---main---")
	//lib1.Lib1Test() //如果不使用lib1的接口，则编译器报错，要求不导入lib1包
	//lib2.Lib2Test() //起了别名，包原名就用不了了
	mylib2.Lib2Test()
	//Lib2Test()
}
