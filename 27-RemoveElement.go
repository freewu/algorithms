package main

/*
Given an array and a value, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example:

Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2.
*/

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	var t = 0
	
	for i := 0; i < len(nums); i++ {
		nums[t] = nums[i]
		if nums[i] != val {
			t++
		}
	}
    return t
}

func main() {
	var nums1 = []int{1,1,5,6,7,8,9,9,10,11,23}

	fmt.Println(removeElement(nums1,9))
	fmt.Println(removeElement(nums1,2))
}