package main

func main() {
	var a = 1
	for a < 5 {
		println(a)
		a++
	}

	var b = 1
	for b < 5 {
		println(b)
		b++
		if b == 3 {
			break
		}
	}

	var c = 1
	for c < 5 {
		println(c)
		c++
		if c == 2 {
			c++
			continue
		}
	}
}
