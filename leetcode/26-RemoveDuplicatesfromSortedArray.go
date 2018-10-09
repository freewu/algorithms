package main

/*
Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
It doesn't matter what you leave beyond the new length.
*/

import (
	"fmt"
)


//allocate extra space for another array
func removeDuplicates1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var a = []int{nums[0]}
	var l = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != a[l - 1] {
			a = append(a,nums[i])
			l++
		} 
	}
    return l
}

func removeDuplicates(nums []int) int {
	var le = len(nums)
	if le < 2 {
		return le
	}
	var l = 1
	var t = nums[0]

	for i := 1; i < le; i++ {
		if nums[i] != t {
			t = nums[i]
			l++
		}
		nums[l - 1] = nums[i] // it is the point
	}
    return l
}

func main() {
	//var nums = []int{1,1,2}
	//fmt.Println(removeDuplicates(nums))

	var nums1 = []int{1,1,5,6,7,8,9,9,10,11,23}
	fmt.Println(removeDuplicates(nums1))

	//var nums2 = []int{1,1}
	//fmt.Println(removeDuplicates(nums2))
}