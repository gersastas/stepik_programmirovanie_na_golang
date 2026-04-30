package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]int, n)
	for i := 0; i < len(s); i++ {
		var x int
		fmt.Scan(&x)
		s[i] = x

		if i%2 == 0 {
			fmt.Print(s[i], " ")
		}
	}
}
