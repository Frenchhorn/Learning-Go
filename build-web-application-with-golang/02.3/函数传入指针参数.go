package main

import "fmt"

// 插入地址
func add1(a *int) int {
	*a = *a + 1 // 给这个地址的值赋值
	return *a
}

func main() {
	x := 3
	fmt.Println(x)
	x1 := add1(&x)
	fmt.Println(x1)
	fmt.Println(x)
}
