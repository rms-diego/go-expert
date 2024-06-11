package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// criando arquivo
	f, err := os.Create("1-manipulacao-de-arquivos/file.txt")
	if err != nil {
		panic(err)
	}

	// Escrevendo no arquivo
	length, err := f.Write([]byte("hello, world"))
	if err != nil {
		panic(err)
	}

	// imprimindo o tamanho do arquivo
	fmt.Printf("File length: %v bytes\n", length)
	f.Close()

	// lendo o arquivo
	d, err := os.ReadFile("1-manipulacao-de-arquivos/file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(d))

	// leitura em chunks
	file, err := os.Open("1-manipulacao-de-arquivos/file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

}
