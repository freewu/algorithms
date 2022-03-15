package main

import "fmt"

/*
31. Next Permutation

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).
The replacement must be  in place and use only constant extra memory.

Example 1:

	Input: nums = [1,2,3]
	Output: [1,3,2]

Example 2:

	Input: nums = [3,2,1]
	Output: [1,2,3]

Example 3:

	Input: nums = [1,1,5]
	Output: [1,5,1]

Example 4:

	Input: nums = [1]
	Output: [1]

Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 100


*/

// 解法一
func nextPermutation(nums []int) {
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse(&nums, i+1, len(nums)-1)
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func main() {
	nums1 := []int{1,2,3}
	nextPermutation(nums1)
	fmt.Printf("nums1 = %v \n",nums1)

	nums2 := []int{3,2,1}
	nextPermutation(nums2)
	fmt.Printf("nums2 = %v \n",nums2)
}
