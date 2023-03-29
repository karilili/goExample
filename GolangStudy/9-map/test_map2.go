package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap是一个引用传递，似乎map本身就是一个地址
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}
}

func changValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}

	//删除 delete是go内置的关键字，与make一样
	delete(cityMap, "China")

	//修改
	cityMap["USA"] = "DC"
	changValue(cityMap)
	fmt.Println("_______")

	//遍历
	for key, value := range cityMap {
		fmt.Println("key=", key)
		fmt.Println("value=", value)
	}

}
