package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(getSum(257, 486))
}
