package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = b + a
	fmt.Printf("c : %d\n", c)
	c = a - b
	fmt.Printf("c : %d\n", c)
	c = a * b
	fmt.Printf("c : %d\n", c)
	c = a / b
	fmt.Printf("c : %d\n", c)
	c = a % b
	fmt.Printf("c : %d\n", c)
	a++
	fmt.Printf("c : %d\n", a)
	a = 21
	a--
	fmt.Printf("c : %d\n", a)
}
