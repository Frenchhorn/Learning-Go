package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

func main() {
	// 在Go 1.5将标识并发系统线程个数的runtime.GOMAXPROCS的初始值为了运行环境的CPU核数。
	runtime.GOMAXPROCS(1)
	go say("world") // 开一个新的Goroutines执行
	say("Hello")    // 当前Goroutines执行
}
