package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	nums := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	actual, _ := Fibonacci(11)

	if len(actual) != 11 {
		t.Fatalf("Expected nums length %d but got %v", 11, len(actual))
	}

	for idx, num := range actual {
		if num != nums[idx] {
			t.Errorf("Expected %v but got %v", nums, actual)
			break
		}
	}

	// return ArgumentError
	_, err := Fibonacci(0)
	if err == nil || err.Error() != "Bad argument: 0" {
		t.Error("Should raise ArgumentError when received 0")
	}

	_, err2 := Fibonacci(-10)
	if err2 == nil || err2.Error() != "Bad argument: -10" {
		t.Error("Should raise ArgumentError when received -10")
	}

	// limit length of array until maxNum
	actual2, _ := Fibonacci(100)
	if len(actual2) >= maxNum {
		t.Errorf("Should limit %d", maxNum)
	}
}
