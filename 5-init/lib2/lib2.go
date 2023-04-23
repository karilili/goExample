package lib2 //文件名与包名一样，方便导入

import "fmt"

//当前lib2包提供的API
func Lib2Test(){
	fmt.Println("lib2Test()...")
}

func init(){
	fmt.Println("lib1.init()")
}