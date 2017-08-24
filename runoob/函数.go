package main

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)

	println("最大值为:", ret)

	var x = "asdf"
	var y = "qwer"
	x, y = swap(x, y)
	println(x, y)
}

func max(a, b int) int {
	var result int
	if a >= b {
		result = a
	} else {
		result = b
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
