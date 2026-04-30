package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)

	a *= a
	b *= b
	c = a + b
	fmt.Println(c)
}
