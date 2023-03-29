package main

import "fmt"

//go语言中的main实际上就是通过struct绑定方法

// 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写，表示该属性是对外能够访问的，否则只能够包内访问
	Name  string
	Ad    int
	Level int
}

//对于方法，如果首字母大写，则其他包能够访问，如果首字母小写，则包内能够访问，其他包不能访问

//类名、方法、属性，首字母大写，代表其他包可以访问，小写，代表只有包内可以访问

// func (this Hero) Show() {
// 	fmt.Println("Name=", this.Name)
// 	fmt.Println("Ad=", this.Ad)
// 	fmt.Println("Level=", this.Level)
// }

// func (this Hero) GetName() string { //表示这个方法是绑定到Hero这个结构体上的
// 	fmt.Println("Name=", this.Name)
// 	return this.Name
// }

// func (this Hero) SetName(newName string) {
// 	//this 是调用该方法的对象的一个副本（拷贝）
// 	this.Name = newName
// }

func (this *Hero) Show() { //这里的this不是关键字，只是一个名
	fmt.Println("Name=", this.Name)
	fmt.Println("Ad=", this.Ad)
	fmt.Println("Level=", this.Level)
}

func (this *Hero) GetName() string {
	fmt.Println("Name=", this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//this 是调用该方法的对象的一个指针
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}
	hero.Show()
	hero.SetName("Li4")
	hero.Show()
}
