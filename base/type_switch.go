package main

import "fmt"

func main() {
	var x interface{}
	switch i := x.(type) {
		case nil:
			fmt.Printf("x 的类型是: %T\n", i)
		case int:
			fmt.Printf("x 的类型是: int\n")
		case float64:
			fmt.Printf("x 的类型是: float64\n")
		case func(int) float64:
			fmt.Printf("x 的类型是: func(int)\n")
		case bool, string:
			fmt.Printf("x 的类型是: bool或者string\n")
		default:
			fmt.Printf("未知类型!\n")
	}
}
