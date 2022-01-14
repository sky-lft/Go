package main

func sum3(a, b int) int {
	return a + b
}

func sum1(a, b int) (result int) {
	result = a + b
	return 2
}
func sum2(a, b int) (result int) {
	result = a + b
	return result
}

func main() {
	println("sum", sum3(1, 2))
	println("sum1", sum1(1, 2))
	println("sum2", sum2(1, 2))
}
