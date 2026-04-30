package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)

	numLast := num / 100
	numMid := (num / 10) % 10
	numStart := num % 10
	output := strconv.Itoa(numStart) + strconv.Itoa(numMid) + strconv.Itoa(numLast)
	fmt.Println(output)
}
