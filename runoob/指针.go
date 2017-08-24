// 指针
package main

import (
	"fmt"
)

func main() {
	var a int = 20 // 声明实际变量
	var ip *int    // 声明指针变量

	if ip == nil {
		fmt.Printf("空指针的值： %x\n", ip)
	}

	ip = &a // 指针变量的存储地址

	fmt.Printf("a 变量的地址： %x\n", &a)

	fmt.Printf("ip 变量存储的地址： %x\n", ip)

	fmt.Printf("ip 变量的访问值： %d\n", *ip)
}
