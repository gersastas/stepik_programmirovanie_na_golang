package main

import "fmt"

func main() {
	a := 8
	const b int = 10
	a += b
	fmt.Println(a)
}
