package main

import "fmt"

func main() {
	var num = make([]int, 3, 5)

	printSlice(num)
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}