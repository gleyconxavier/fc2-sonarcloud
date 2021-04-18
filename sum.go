package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(2, 2))

	return
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func findPercentage(num, percentage int) int {
	return (num / 100) * percentage
}
