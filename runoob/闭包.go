package main

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()
	println(nextNumber())
	println(nextNumber())
	println(nextNumber())

	nextNumber2 := getSequence()
	println(nextNumber2())
	println(nextNumber2())
	println(nextNumber2())

}
