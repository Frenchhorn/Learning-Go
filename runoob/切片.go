// 切片
package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3}
	b := a[:]           // 通过切片a初始化切片b
	c := []int{1, 2, 3} // 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	d := make([]int, 1) // 通过内置函数make()初始化切片d,[]int 标识为其元素类型为int的切片
	d = a[:]
	fmt.Println(a, b, c, d)
	fmt.Println("a长度为:", len(a))  // len
	fmt.Println("a最长可为:", cap(a)) // cap

	// 空切片
	var numbers []int
	printSlice(numbers)
	if numbers == nil {
		fmt.Println("切片是空的")
	}

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
