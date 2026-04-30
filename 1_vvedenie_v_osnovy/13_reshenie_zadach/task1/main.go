package main

import "fmt"

const (
	ten      = 10
	hundered = 100
)

func main() {
	var num int
	fmt.Scan(&num)

	numFirst := num / hundered
	numTwo := (num / ten) % ten
	numThree := num % ten
	sum := numFirst + numTwo + numThree
	fmt.Println(sum)
}
