package main

import "fmt"

func main() {
	result, boolean := sum(49, 2)

	fmt.Println(result)
	fmt.Println(boolean)
}

func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}
