package main

import "fmt"

type Person interface {
	disable()
}

type customer struct {
	Name     string
	lastName string
	taxId    string
	age      int
	isActive bool
}

func (c *customer) disable() {
	c.isActive = false
}

func main() {
	var customer Person = &customer{
		Name:     "Diego",
		lastName: "Ramos",
		age:      23,
		taxId:    "000.000.000-00",
		isActive: true,
	}

	fmt.Println(customer)
}
