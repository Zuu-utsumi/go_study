package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Before
	v := m.Run()
	// After
	os.Exit(v)
}

func TestFizzBuzz(t *testing.T) {
	expect1 := "Fizz"
	actual1 := FizzBuzz(3)
	if actual1 != expect1 {
		t.Errorf("Expected '%s' but got '%s'", expect1, actual1)
	}

	expect2 := "Buzz"
	actual2 := FizzBuzz(5)
	if actual2 != expect2 {
		t.Errorf("Expected '%s' but got '%s'", expect2, actual2)
	}

	expect3 := "Fizz Buzz"
	actual3 := FizzBuzz(15)
	if actual3 != expect3 {
		t.Errorf("Expected '%s' but got '%s'", expect3, actual3)
	}

	expect4 := "-"
	actual4 := FizzBuzz(2)
	if actual4 != expect4 {
		t.Errorf("Expected '%s' but got '%s'", expect4, actual4)
	}
}
