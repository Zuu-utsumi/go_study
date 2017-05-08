package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	greet   string
}

func main() {
	fmt.Println("Hello, World!")

	var str string = ""
	for i := 0; i < 100; i++ {
		str += "a"
	}

	// var tom Person{"Top", 15, "Tokyo", "Hi"}
}
