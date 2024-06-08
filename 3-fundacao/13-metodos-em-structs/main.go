package main

import "fmt"

type customer struct {
	name     string
	taxId    string
	age      int
	isActive bool
}

func (c customer) greetings() string {
	return fmt.Sprintf(
		"My name is %v, i'm have a %v years old, my tax id is %v",
		c.name,
		c.age,
		c.taxId,
	)
}

func main() {
	customer := customer{
		name:     "Diego",
		age:      23,
		taxId:    "000.000.000-00",
		isActive: true,
	}

	fmt.Println(customer.greetings())
}
