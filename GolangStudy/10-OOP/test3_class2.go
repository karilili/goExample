package main

import "fmt"

// go语言简化省去了private public等访问权限，取而代之的是包内包外的访问权限
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human //superMan类继承了Human的方法，go没有c++访问权限那一套东西
	//把父类类名写在{}中
	level int
}

// 重定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("superMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("superMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"zhang3", "famale"}

	h.Eat()
	h.Walk()

	//定义一个子类对象
	var s SuperMan
	s.name = "lie4"
	s.sex = "male"
	s.level = 88

	//s := SuperMan{Human{"lie4", "famele"}, 88}

	s.Walk() //父类的方法
	s.Eat()  //子类的方法
	s.Fly()  //子类的方法

	s.Print()
}
