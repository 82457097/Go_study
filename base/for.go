package main 

import "fmt"

func main() {
	str := []string{"hello", "world"}
	for i, s := range str {
		fmt.Println(i, s)
	}

	num := [6]int{1, 2, 3, 4}
	for i, x := range num {
		fmt.Printf("第%d号元素的值为:%d\n", i, x)
	}
}
