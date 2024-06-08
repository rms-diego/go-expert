package main

import "fmt"

func main() {
	salaries := map[string]int{"Diego": 100, "Tatiana": 100}

	/*
	*		Para acessar os valores do 'map' é possível percorrer ou acessando diretamente
	* 	Usando 'for'
	 */
	for index, values := range salaries {
		fmt.Printf("No chave: %v\nTem o valor: %v", index, values)
		fmt.Printf("\n\n")
	}

	// Acessando diretamente
	fmt.Println(salaries["Diego"])
	fmt.Println(salaries["Tatiana"])

	// Deletando valores do 'map'
	delete(salaries, "Diego")
}
