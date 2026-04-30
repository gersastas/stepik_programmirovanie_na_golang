package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)

	slice := make([]int, size)
	for i := 0; i < len(slice); i++ {
		var n int
		fmt.Scan(&n)
		slice[i] = n
	}
	output := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == 0 {
			output++
		}
	}
	fmt.Println(output)
}
