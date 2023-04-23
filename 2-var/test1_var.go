package main

/*
   四种变量的声明方式
*/

import "fmt"

// 声明全局变量  方法一二三是可以的
var gA int = 100
var gB = 200

//方法四不能用来声明全局变量，只能用于函数体中进行声明
// gc:=300

func main() {
	//方法一：声明一个变量 默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb = "abcd"
	fmt.Printf("bb = %s,type of bb = %T\n", bb, bb)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s,type of cc = %T\n", cc, cc)

	//方法四：（常用的方法）省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e=%T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f=%T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g=%T\n", g)

	//快捷键操作，按住shift选中多行 shift+alt+上箭头+下箭头也选中多行，按esc可取消
	//利用alt+鼠标左键点击，可以选择多行粘贴，每一需粘贴行的位置; 也可以用于一次选中多个单词
	//ubuntu下 ctrl+shift+i 代码对齐  home end跳转代码行末或行尾
	//shift + 左/右 每次移动，多选中一个字母  //ctrl+shift +  左/右 快速选完一个词
	//ctrl+shift+k 删除一行
	//ctrl+鼠标左键，跳转函数定义
	//ctrl+1回到主编辑区（一个编辑区可以有多个文件）
	// 使用ctrl+tab可以切换编辑的文件，也可起到切换回编辑区的作用
	//ctrl+`（esc旁边那个），回到终端、开关终端（不使用rime输入法）
	//（rime输入法下）ctrl+shift+3可以切换半角全角;ctrl+`可以控制rime输入模式
	//设置了快捷键，（rime输入法下）使用ctrl+escape切换回终端

	//=====
	fmt.Println("gA = ", gA, ",gB = ", gB)
	// fmt.Println("gC = ",gC)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx==", xx, "yy=", yy)

	var kk, ll = 100, "Aceld" //不同数据类型，只能自动匹配了
	fmt.Println("kk==", kk, "ll=", ll)

	//多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv=", vv, ",jj=", jj)

}
