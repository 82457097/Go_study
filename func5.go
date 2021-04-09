package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.0
	fmt.Printf("c1的面积为:%f\n", c1.getArea())
}

//该方法属于Circle类的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
