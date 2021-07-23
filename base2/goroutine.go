package main

import "fmt"

var count int

func Count(ch chan int, i int) {
	ch <- 1
	fmt.Println("Counting ", count, i)
}

func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}

	for _, ch := range chs {
		count++
		fmt.Print(count)
		<-ch
	}
}
