package main

import (
	"fmt"
)

/*
Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.

Example 1
Input:[1,2,3,4,8,9,0,3,4]
Output:[1,2,3,4,8,9,0,3,5]

Example 2
Input:[1,2,3,4,8,9,0,3,9]
Output:[1,2,3,4,8,9,0,4,0]

Example 3
Input:[9,9,9]
Output:[1,0,0,0]
*/

func plusOne(digits []int) []int {
	var l = len(digits)
	for i := l - 1; i >= 0; i-- {
		if (digits[i] + 1) >= 10 {
			digits[i] = ((digits[i] + 1) % 10)

			// 如果是9999这样子的情况
			if 0 == i {
				//var slice = make([]int, l+1)
				//copy(slice, []int{1})
				//return slice

				var slice = make([]int, l+1)
				slice[0] = 1
				return slice
			}

		} else {
			digits[i] = digits[i] + 1
			return digits
		}
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3, 4, 8, 9, 0, 3, 4}))
	fmt.Println(plusOne([]int{1, 2, 3, 4, 8, 9, 0, 3, 9}))
	fmt.Println(plusOne([]int{9, 9}))
}
