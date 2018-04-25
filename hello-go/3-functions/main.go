package main

import "fmt"

func main() {
	fmt.Println(sumTwo(1, 2))
	fmt.Println(sumN(1, 2, 3))
}

func sumTwo(x, y int) int {
	return sumN(x, y)
}

// Functions can have N arguments
func sumN(numbers ...int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}

	return sum
}
