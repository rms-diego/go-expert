package main

import "fmt"

func main() {
	var numbers [10]int

	for i := range numbers {
		numbers[i] = i + 1

		fmt.Printf("value: %v\n", numbers[i])
	}

	fmt.Println(numbers)
}
