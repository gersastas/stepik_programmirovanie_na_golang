package main

import "fmt"

func main() {
	var d, h, m int
	fmt.Scan(&d)

	h = d / 30
	m = 2 * (d % 30)
	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
