package main

import (
	"encoding/json"
	"os"
)

type account struct {
	Amount int
	Number int
}

func main() {
	a := account{Amount: 100, Number: 1}

	_, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(os.Stdout).Encode(a)
}
