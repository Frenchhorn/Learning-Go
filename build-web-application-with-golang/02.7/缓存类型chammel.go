package main

import "fmt"

// 当 value=0 时，是无缓冲阻塞读写的
// 当 value>0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入
// ch := make(chan type, value)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 显式的关闭channel
	// 关闭channel后就无法再发送任何数据
	// 可以通过语法v, ok := <-ch测试channel是否被关闭
}

func main() {
	c := make(chan int, 2) // 修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("*************")
	// 可以通过range，像操作slice或者map一样操作缓存类型的channel
	c2 := make(chan int, 10)
	fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}
}
