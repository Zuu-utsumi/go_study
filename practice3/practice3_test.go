package main

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	t.Log("Testing Sort() ascending")
	expect := []string{"a", "b", "z", "ほげ", "漢字"}
	actual := sort.StringSlice{"ほげ", "z", "a", "漢字", "b"}
	Sort(actual, "asc")

	for idx, s := range actual {
		if s != expect[idx] {
			t.Errorf("Expect %v but got %v", expect, actual)
		}
	}

	t.Log("Testing Sort() descending")
	expect2 := []string{"漢字", "ほげ", "z", "b", "a"}
	actual2 := sort.StringSlice{"ほげ", "z", "a", "漢字", "b"}
	Sort(actual2, "desc")

	for idx, s := range actual2 {
		if s != expect2[idx] {
			t.Errorf("Expect %v but got %v", expect2, actual2)
		}
	}
}
