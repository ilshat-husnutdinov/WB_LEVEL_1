package main

import "fmt"

type Human struct {
	name string
	age  uint8
}

func (h *Human) printName() {
	fmt.Printf("Name is: %v\n", h.name)
}

func (h *Human) printAge() {
	fmt.Printf("Age is: %v\n", h.age)
}

type Action struct {
	Human
}

func main() {
	testAction := Action{
		Human{
			name: "John",
			age:  30,
		},
	}
	testAction.printAge()        // вызов метода напрямую
	testAction.Human.printName() // вызов метода через структуру Human
}
