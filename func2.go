package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var str1 string = "world"
	var str2 string = "hello"
	str1, str2 = swap(str1, str2)
	fmt.Printf("%s %s\n", str1, str2)
}
