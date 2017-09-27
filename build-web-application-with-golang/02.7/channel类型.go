package main

import "fmt"

// 必须使用make 创建channel：
// ci := make(chan int)
// cs := make(chan string)
// cf := make(chan interface{})

// channel通过操作符<-来接收和发送数据
// ch <- v    // 发送 v 到 channel ch
// v := <-ch  // 从 ch 中接收数据，并赋值给 v

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 默认情况下，channel接收和发送数据都是阻塞的
	fmt.Println(x, y, x+y)
}
