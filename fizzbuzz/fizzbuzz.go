package main

import "fmt"

// FizzBuzz ...
func FizzBuzz(i int) string {
	if i%3 == 0 {
		if i%5 == 0 {
			return "Fizz Buzz"
		}

		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return "-"
	}
}

func main() {
	const loopCount = 100
	for i := 1; i <= loopCount; i++ {
		fmt.Printf("%d: ", i)
		fmt.Println(FizzBuzz(i))
	}
}
