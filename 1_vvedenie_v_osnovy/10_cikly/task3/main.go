package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x >= 10 && x <= 99 && x%8 == 0 {
			sum += x
		}
	}
	fmt.Println(sum)
}
