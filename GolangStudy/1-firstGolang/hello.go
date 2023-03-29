package main //程序的包名  java等语言存在包的概念，对于C++学习者，只需知道有main函数的就是main包
//表示这是main包，与文件名没有关系

/*
import (
	"fmt"
	"time"
) //导入fmt包
*/

//建议这种格式，整洁一些
import (
	"fmt"
	"time"
)

// main函数
func main() { //函数的 { 必须与函数名同行，否则编译错误
	//golang中的表达式，加“;”和不加 都可以 建议是不加
	fmt.Println("hello Go!")

	time.Sleep(1 * time.Second)
}

//go run 编译加运行，等同于go build 再./可执行程序
//但是go run似乎不会产生可执行程序
