package main

import "fmt"

func main() {
	numbers := []int{00, 10, 20, 30, 40, 50, 60, 70, 80, 100}

	fmt.Printf("Capacidade do array: %v\nTamanho do array: %v\nValores: %v", cap(numbers), len(numbers), numbers)

	/**
	*		Essa sintaxe de [:index] significa que todos os indices após o definido
	*		Serão excluídos. A mesma logica pode ser usada inversamente
	* 	exemplo [index:]
	**/

	numbers = numbers[:5]
	fmt.Printf("\n\nCapacidade do array: %v\nTamanho do array: %v\nValores: %v", cap(numbers), len(numbers), numbers)
}
