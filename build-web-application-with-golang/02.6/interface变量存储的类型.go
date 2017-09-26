package main

import (
	"fmt"
	"strconv"
)

// 空interface,可以存储任意类型的数值
type Element interface{}

// 空interface的列表
type List []Element

type Person struct {
	name string
	age  int
}

// 定义String方法，实现fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

// Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)
// value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型
func main() {
	list := make(List, 3)
	list[0] = 1       // int
	list[1] = "Hello" //sring
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	fmt.Println("***********")

	// element.(type)语法不能在switch外的逻辑里面使用
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
