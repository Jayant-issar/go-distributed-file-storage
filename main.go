package main

import "fmt"

type Employee struct {
	name     string
	age      int
	isRomote bool
	address  *Address
}

type Address struct {
	street string
	city   string
}

func (e *Employee) updateName(newName string) {
	e.name = newName
}

func main() {
	jayantAddress := Address{
		street: "moti bagh",
		city:   "New Delhi",
	}
	jayant := Employee{
		name:     "jayant",
		age:      21,
		isRomote: false,
		address:  &jayantAddress,
	}

	jayant.address.street = "satya niketan"
	fmt.Println(jayant.address.street)
	fmt.Println(jayantAddress.street)
}
