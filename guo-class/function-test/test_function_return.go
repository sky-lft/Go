package main

import "fmt"

func test1() {
	fmt.Println("这个就是一个简单的函数，没有参数，也没有返回值")
}

func sum(a, b int) int {
	return a + b
}

func test3() (name string, age int) {
	name = "sky"
	age = 34
	return "sky", 78
	return //这种都行
}

func main() {
	fmt.Println("没有参数；", test1)
	fmt.Println("带有参数")

}
