package main

import "fmt"

func main() {
	values := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	fmt.Println("Row1")
	fmt.Println(values[0])
	fmt.Println("Row2")
	fmt.Println(values[1])

	fmt.Printf("第一个元素为:%d\n", values[0][0])
}
