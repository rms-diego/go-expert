package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemon struct {
	Name string `json:"name"`
}

func main() {
	request, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	var dataInJson pokemon

	err = json.Unmarshal(data, &dataInJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(dataInJson.Name)
	request.Body.Close()
}
