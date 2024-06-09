package main

func sum(a, b *int) {
	*a = 100
	*b = -10

	// return *a + *b
}

func main() {
	number1 := 10
	number2 := 10

	sum(&number1, &number2)
}
