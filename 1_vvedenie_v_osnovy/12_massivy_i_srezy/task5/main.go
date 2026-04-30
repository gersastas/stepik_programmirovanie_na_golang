package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]int, n)
	count := 0
	for i := 0; i < len(s); i++ {
		var x int
		fmt.Scan(&x)

		s[i] = x
		if s[i] > 0 {
			count++
		}
	}
	fmt.Println(count)
}
