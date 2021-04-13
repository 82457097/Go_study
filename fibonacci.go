package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}