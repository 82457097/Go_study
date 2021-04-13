package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Printf("1.case is false.\n")
		fallthrough
	case true:
		fmt.Printf("2.case is true.\n")
		fallthrough
	case false:
		fmt.Printf("3.case is false.\n")
		fallthrough
	case true:
		fmt.Printf("4.case is true.\n")
	case false:
		fmt.Printf("5.case is false.\n")
		fallthrough
	default:
		fmt.Printf("6.default case.\n")
	}
}
