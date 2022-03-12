package main

/*
26. Remove Duplicates from Sorted Array

Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

	Given nums = [1,1,2],
	Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
	It doesn't matter what you leave beyond the returned length.

Example 2:

	Given nums = [0,0,1,1,1,2,2,3,3,4],
	Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.
	It doesn't matter what values are set beyond the returned length.


*/

import (
	"fmt"
)


//allocate extra space for another array
func removeDuplicates1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var a = []int{nums[0]} // 声明一个数组来保存 去重后的数据
	var l = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != a[l - 1] { //
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

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

// best solution
func removeDuplicatesBest(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ { // 注意从 1 开始
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i] // 用已使用的地址来保存去重的数据
			// nums: 1 1 2 3 3
			// r1  nums: 1 1 2 3 3 j:0 i:1 因为 nums[0] == nums[1] 所以不进入
			// r2  nums: 1 2 2 3 3 j:1 i:2 因为 nums[0] ！+ nums[2] 所有 j++ (1) nums[1] = nums[2
			// r3  nums: 1 2 3 3 3 j:2 i:3
			// r3  nums: 1 2 3 3 3 j:2 i:4
		}
		fmt.Printf("\nround %v\n",i)
		fmt.Printf("nums = %v\n",nums)
		fmt.Printf("j = %v\n",j)
	}
	return j + 1
}

func main() {
	fmt.Printf("removeDuplicates([]int{1,1,5,6,7,8,9,9,10,11,23}) = %v\n",removeDuplicates([]int{1,1,5,6,7,8,9,9,10,11,23}))
	fmt.Printf("removeDuplicates1([]int{1,1,5,6,7,8,9,9,10,11,23}) = %v\n",removeDuplicates1([]int{1,1,5,6,7,8,9,9,10,11,23}))
	fmt.Printf("removeDuplicates2([]int{1,1,5,6,7,8,9,9,10,11,23}) = %v\n",removeDuplicates2([]int{1,1,5,6,7,8,9,9,10,11,23}))
	fmt.Printf("removeDuplicatesBest([]int{1,1,5,6,7,8,9,9,10,11,23}) = %v\n",removeDuplicatesBest([]int{1,1,5,6,7,8,9,9,10,11,23}))
	fmt.Printf("removeDuplicatesBest([]int{1,1,2,3,3}) = %v\n",removeDuplicatesBest([]int{1,1,2,3,3}))
}