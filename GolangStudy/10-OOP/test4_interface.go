package main

import "fmt"

// 本质是一个指针 内部指向了具体的类型，类型的函数列表
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

// 具体的类
type Cat struct {
	color string //猫的颜色
}

//go语言实现多态，不需要在类中写interface的类型，就像继承一样
//只需要实现接口定义的所有方法即可，随后可用接口指向该类
//但如果没有实现接口定义的方法，或者只实现了部分，则无法用该接口指向

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string //猫的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	var animal AnimalIF    //接口的数据类型，父类指针
	animal = &Cat{"Green"} //?这是否是一种匿名对象
	animal.Sleep()         //调用就是Cat.Sleep方法，多态的现象
	animal = &Dog{"Yellow"}
	animal.Sleep() //调用就是Dog.Sleep方法，多态的现象

	cat := Cat{"Green"}
	dog := Dog{"Red"}
	showAnimal(&cat)
	showAnimal(&dog)
	//使用(*Dog)这种函数绑定时，则用接口指向时，必须传入地址
	//不带地址，则提示错误:cannot use dog (variable of type Dog) as AnimalIF value in argument to showAnimal: Dog does not implement AnimalIF (method GetColor has pointer receiver)
}
