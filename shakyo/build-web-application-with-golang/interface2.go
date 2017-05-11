package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "Person{name: " + p.name + ", age: " + strconv.Itoa(p.age) + "}"
}

func main() {
	list := make(List, 4)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}
	list[3] = []int{}

	for idx, elem := range list {
		// if value, ok := elem.(int); ok {
		// 	fmt.Printf("list[%d] is an int and its value is %d\n", idx, value)
		// } else if value, ok := elem.(string); ok {
		// 	fmt.Printf("list[%d] is an string and its value is '%s'\n", idx, value)
		// } else if value, ok := elem.(Person); ok {
		// 	fmt.Printf("list[%d] is an Person and its value is %s\n", idx, value)
		// } else {
		// 	fmt.Printf("list[%d] is of a different type\n", idx)
		// }

		switch value := elem.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", idx, value)
		case string:
			fmt.Printf("list[%d] is an string and its value is '%s'\n", idx, value)
		case Person:
			fmt.Printf("list[%d] is an Person and its value is %s\n", idx, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", idx)
		}
	}
}
