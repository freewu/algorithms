package main

import "fmt"

/**
80. Remove Duplicates from Sorted Array II
Given an integer array nums sorted in non-decreasing order,
remove some duplicates in-place such that each unique element appears at most twice.
The relative order of the elements should be kept the same.
Since it is impossible to change the length of the array in some languages,
you must instead have the result be placed in the first part of the array nums.
More formally, if there are k elements after removing the duplicates,
then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
Return k after placing the final result in the first k slots of nums.
Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Constraints:

	1 <= nums.length <= 3 * 10^4
	-10^4 <= nums[i] <= 10^4
	nums is sorted in non-decreasing order.

Example 1:

	Input: nums = [1,1,1,2,2,3]
	Output: 5, nums = [1,1,2,2,3,_]
	Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:

	Input: nums = [0,0,1,1,1,1,2,3,3]
	Output: 7, nums = [0,0,1,1,2,3,3,_,_]
	Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).

解题思路:
	给定一个有序数组 nums，对数组中的元素进行去重，
	使得原数组中的每个元素最多暴露 2 个。
	最后返回去重以后数组的长度值。
	使用双指针的解法，双指针的关键点：移动两个指针的条件
 */

func removeDuplicatesII(nums []int) int {
	slow := 0
	for fast, v := range nums {
		// fast < 2 前两个数值 都进入
		// slow = 2 & nums[0] != nums[3] 说明  slow -2 是重点
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

// best solution
func removeDuplicatesIIBest(nums []int) int {
	idx := 1
	for i := 1; i < len(nums); i++ {
		if idx < 2 || nums[i] > nums[idx-2] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

func main() {
	fmt.Printf("removeDuplicatesII([]int{1,1,1,2,2,3}) = %v\n",removeDuplicatesII([]int{1,1,1,2,2,3})) // 5 [1,1,2,2,3]
	fmt.Printf("removeDuplicatesII([]int{0,0,1,1,1,1,2,3,3}) = %v\n",removeDuplicatesII([]int{0,0,1,1,1,1,2,3,3})) // 7 [0,0,1,1,2,3,3]

	fmt.Printf("removeDuplicatesIIBest([]int{1,1,1,2,2,3}) = %v\n",removeDuplicatesIIBest([]int{1,1,1,2,2,3})) // 5 [1,1,2,2,3]
	fmt.Printf("removeDuplicatesIIBest([]int{0,0,1,1,1,1,2,3,3}) = %v\n",removeDuplicatesIIBest([]int{0,0,1,1,1,1,2,3,3})) // 7 [0,0,1,1,2,3,3]
}
