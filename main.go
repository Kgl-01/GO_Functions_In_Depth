package main

import "fmt"

func main() {
	fmt.Println("Do Something")
}

func variadicFunctions(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
