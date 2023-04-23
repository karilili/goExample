package main

import (
	"fmt"
	"reflect"
)

// 作用，其他包导入该包时，该包对该属性进行一个说明
type resume struct {
	Name string `info:"name" doc:"我的名字"` //给属性绑定一个标签，可以写多个
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() //获取所有的元素

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, "doc: ", tagdoc)
	}

}

func main() {
	var re resume
	findTag(&re)
}
