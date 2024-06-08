package main

import "fmt"

type customer struct {
	name     string
	taxId    string
	age      int
	isActive bool
}

func main() {
	customer := customer{
		name:     "Diego",
		age:      23,
		taxId:    "000.000.000-00",
		isActive: true,
	}

	fmt.Println(customer)

	customer.name = "aoba"
	fmt.Println(customer)
}
