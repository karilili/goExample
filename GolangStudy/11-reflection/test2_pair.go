package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//tty:pair<type:*os.File,"dev/tty"文件描述符>  //不管赋值给谁，这个pair是不变的
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	//r:pair<type: ,value:>
	var r io.Reader
	//r:pair<type:*os.File,"dev/tty"文件描述符>
	r = tty

	//w:pair<type: ,value:>
	var w io.Writer
	//w:pair<type:*os.File,"dev/tty"文件描述符>
	w = r.(io.Writer) //强制转换为io.writer并赋值给w //也是类型断言，可以起到强转的作用

	w.Write([]byte("HELLO THIS is A TEST!!!\n"))

	//我比较疑惑的是没有隐式转换吗
	//我估计这种pair赋值不变性，是只针对接口的。需要变量类型先实现接口，然后用接口就可指向该变量
	//要明白类型断言的本质是利用pair
}
