package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			return
		}
		// 在select里面还有default语法，当监听的channel都没有准备好的时候，默认执行的
	}
}

// select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行
// 当多个channel都准备好的时候，select是随机的选择一个执行的
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
