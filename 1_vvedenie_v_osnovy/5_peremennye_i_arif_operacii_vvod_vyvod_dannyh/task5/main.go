package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	tensDigit := (num / 10) % 10
	fmt.Println(tensDigit)
}
