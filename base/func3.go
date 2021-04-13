package main

import "fmt"
import "math"

func main() {
	getSqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSqrt(9))
}
