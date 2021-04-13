package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var stockCode = "000987"
	var enddate = "2021-4-13"
	var url = "Code = %s & endDate = %s"
	var target_url = fmt.Sprintf(url, stockCode, enddate)
	fmt.Println(target_url)

	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s)
}