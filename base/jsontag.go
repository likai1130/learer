package main

import (
	"fmt"
	"reflect"
)

/**
	json tag使用

	对象属性的tag通过反射获取

	通过反射可以定义日志的格式化输出
 */

type Person struct {
	Name        string `label:"Name is: "`
	Age         int    `label:"Age is: "`
	Gender      string `label:"Gender is: " default:"unknown"`
}

func Print(obj interface{}) error {
	v := reflect.ValueOf(obj)
	//解析字段

	for i := 0; i < v.NumField(); i ++ {
		//取tag
		field := v.Type().Field(i)
		tag := field.Tag

		//解析label和default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == ""{
			//没有指定值，用默认值代替
			value = defaultValue
		}
		fmt.Println(label + value)
	}
	return nil
}

func main() {
	person := Person{
		Name: "KLI",
		Age: 20,
	}
	Print(person)
}
