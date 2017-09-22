package main

import "fmt"

func main() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)
	sum2 := 1
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Println("sum2 is equal to ", sum2)
}
