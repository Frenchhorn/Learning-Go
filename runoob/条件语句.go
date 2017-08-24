package main

import "fmt"

func main() {
	var a int = 10

	if a < 20 {
		fmt.Println("a 小于 20")
	} else {
		fmt.Println("a 大于等于 20")
	}

	var b int = 3

	switch b {
	case 1:
		fmt.Println("b 为 1")
	case 2:
		fmt.Println("b 为 2")
	case 3:
		fmt.Println("b 为 3")
	default:
		fmt.Println("b 不是 1,2,3")
	}

	switch {
	case b == 1:
		fmt.Println("b 为 1")
	case b == 2:
		fmt.Println("b 为 2")
	case b == 3:
		fmt.Println("b 为 3")
	default:
		fmt.Println("b 不是 1,2,3")
	}

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	fmt.Println("未知型")

	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
