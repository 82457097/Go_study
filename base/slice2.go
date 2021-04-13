package main

import "fmt"

func main() {
	var num []int
	printSlice(num)

	num = append(num, 0)
	printSlice(num)

	num = append(num, 1)
	printSlice(num)

	num = append(num, 2, 3, 4)
	printSlice(num)

	num = append(num, 5, 6)
	printSlice(num)

	num1 := make([]int, len(num), (cap(num))*2)
	copy(num1, num)
	printSlice(num1)
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}