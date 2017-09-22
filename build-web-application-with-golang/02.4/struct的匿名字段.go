package main

import (
	. "fmt"
)

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，struct
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
	weight     int // 重载Human中的weight属性
}

func main() {
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology", weight: 101}
	Println(jane.name) // 可以直接访问匿名字段的属性
	Println(jane.age)
	Println(jane.weight)       // 会访问重载后的属性
	Println(jane.Human.weight) // 通过Human可以访问原属性
	Println(jane.speciality)
	// 修改Skills值
	jane.Skills = []string{"anatomy"}
	Println(jane.Skills)
	jane.Skills = append(jane.Skills, "physics", "golang")
	Println(jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	Println(jane.int)
}
