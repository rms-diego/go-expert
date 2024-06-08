package main

import "fmt"

type customer struct {
	name     string
	taxId    string
	age      int
	isActive bool

	address
}

type address struct {
	zipCode string
	number  int
	street  string
}

func main() {
	customer := customer{
		name:     "Diego",
		age:      23,
		taxId:    "000.000.000-00",
		isActive: true,
		address:  address{zipCode: "00000000", number: 6, street: "groover street"},
	}

	fmt.Println(customer)
}
