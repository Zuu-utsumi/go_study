/*
Fibonacci series, swapping two variables, finding maximum/minimum among a list of numbers.
フィボナッチ数列を表示する。二つの数字を入れ替える。数のリストの中から最大値と最小値を見つける。
*/
package main

import (
	"fmt"
	"strconv"
)

const maxNum = 500

type ArgumentError struct {
	message string
}

func (e ArgumentError) Error() string {
	return e.message
}

func Fibonacci(n int) (nums []int, e error) {
	if n <= 0 {
		return []int{}, ArgumentError{fmt.Sprintf("Bad argument: %d", n)}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			nums = append(nums, 0)
		} else if i == 1 {
			nums = append(nums, 1)
		} else {
			num := nums[i-2] + nums[i-1]

			if num > maxNum {
				break
			}

			nums = append(nums, num)
		}
	}

	return nums, e
}

func main() {
	nums, err := Fibonacci(100)

	if err != nil {
		fmt.Println(err.Error())
	}

	var str string
	for _, num := range nums {
		str += strconv.Itoa(num)
	}

	fmt.Println("[" + str + "]")
}
