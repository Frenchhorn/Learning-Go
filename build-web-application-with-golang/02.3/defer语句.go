// defer语句
package main

import (
	"fmt"
)

// defer后指定的函数会在函数退出前调用。
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 后进先出
	}
}
