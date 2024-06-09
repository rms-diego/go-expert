package main

import "fmt"

func main() {
	number := 10

	// aponta para o endereço da variável "number"
	var pointer *int = &number
	fmt.Println(*pointer)

	// ao atualizar a variável "pointer", o "number" muda. porque são o mesmo endereço
	*pointer++
	fmt.Println(number)
}
