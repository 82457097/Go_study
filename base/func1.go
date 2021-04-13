package main

import "fmt"

func max(num1, num2 int) int {
	var res int

	if(num1 > num2) {
		res = num1
	} else {
		res = num2
	}

	return res
}

func main() {
	a := 100
	b := 200
	var	ret int

	ret = max(a, b)
	fmt.Printf("最大值是%d\n", ret)
}
