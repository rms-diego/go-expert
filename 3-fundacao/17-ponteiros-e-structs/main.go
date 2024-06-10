package main

import "fmt"

type customer struct {
	name string
}

func newCustomer(name string) *customer {
	return &customer{name: name}
}

func (c *customer) greetings() string {
	return "Hello, my name is " + c.name
}

func main() {

	person := newCustomer("Diego Ramos")

	fmt.Println(person.greetings())
}
