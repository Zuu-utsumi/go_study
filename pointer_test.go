package main

import (
	"fmt"
	"testing"
)

type List struct {
	items []Item
}

type Item struct {
	name     string
	children []int
}

func (list *List) Push(item Item) {
	list.items = append(list.items, item)
}

func BenchmarkPointer1(b *testing.B) {
	list := List{}

	for i := 0; i < b.N; i++ {
		item := Item{fmt.Sprintf("Item-%d", i+1), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
		list.Push(item)
	}
}

func BenchmarkPointer2(b *testing.B) {
	list := List{make([]Item, 0, b.N)}

	for i := 0; i < b.N; i++ {
		item := Item{fmt.Sprintf("Item-%d", i+1), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
		list.Push(item)
	}
}

func BenchmarkPointer3(b *testing.B) {
	list := List{make([]Item, 0, b.N*5)}

	for i := 0; i < b.N; i++ {
		item := Item{fmt.Sprintf("Item-%d", i+1), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
		list.Push(item)
	}
}

func (list *List) Push2(item *Item) {
	list.items = append(list.items, *item)
}

func BenchmarkPointer4(b *testing.B) {
	list := List{make([]Item, 0, b.N)}

	for i := 0; i < b.N; i++ {
		item := Item{fmt.Sprintf("Item-%d", i+1), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
		list.Push2(&item)
	}
}

func BenchmarkPointer5(b *testing.B) {
	list := List{make([]Item, 0, b.N)}
	ary := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		item := Item{fmt.Sprintf("Item-%d", i+1), ary}
		list.Push2(&item)
	}
}

func BenchmarkPointer6(b *testing.B) {
	list := List{make([]Item, 0, b.N)}
	ary := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		item := Item{fmt.Sprintf("Item-%d", i+1), ary[:]}
		list.Push2(&item)
	}
}
