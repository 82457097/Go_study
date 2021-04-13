package main

import "fmt"

func getSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNum := getSeq()
	i := 0
	for i < 3 {
		fmt.Println(nextNum())
		i++
	}

	nextNum2 :=getSeq()
	i = 0
	for i < 2 {
		fmt.Println(nextNum2())
		i++
	}
}
