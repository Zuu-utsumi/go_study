/*
Accepting series of numbers, strings from keyboard and sorting them ascending, descending order.
キーボードから数字や文字列のリストを入力し、降順、昇順でソートする。
*/
package main

import (
	"fmt"
	"sort"
	// "strings"
)

func Sort(ss sort.StringSlice, orderBy string) {
	// if strings.ToLower(orderBy) == "desc" {
	if orderBy == "desc" {
		rev := sort.Reverse(ss)
		sort.Sort(rev)
	} else {
		sort.Sort(ss)
	}
}

func main() {
	ss := sort.StringSlice{}
	orders := [2]string{"desc", "asc"}

	for {
		s := ""
		fmt.Print("Input: ")
		fmt.Scanln(&s)

		ss = append(ss, s)

		for _, order := range orders {
			fmt.Printf("Order by %s:\n", order)
			Sort(ss, order)

			for idx, str := range ss {
				fmt.Printf("  %d: %s\n", idx+1, str)
			}
		}

		fmt.Println("---")
	}
}
