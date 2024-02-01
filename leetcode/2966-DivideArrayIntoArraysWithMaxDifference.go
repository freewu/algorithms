package main

import (
	"fmt"
	"sort"
	"slices"
)

// 2966. Divide Array Into Arrays With Max Difference
// You are given an integer array nums of size n and a positive integer k.
// Divide the array into one or more arrays of size 3 satisfying the following conditions:
// 		Each element of nums should be in exactly one array.
// 		The difference between any two elements in one array is less than or equal to k.

// Return a 2D array containing all the arrays. If it is impossible to satisfy the conditions, 
// return an empty array. And if there are multiple answers, return any of them.


// Example 1:
// Input: nums = [1,3,4,8,7,9,3,5,1], k = 2
// Output: [[1,1,3],[3,4,5],[7,8,9]]
// Explanation: We can divide the array into the following arrays: [1,1,3], [3,4,5] and [7,8,9].
// The difference between any two elements in each array is less than or equal to 2.
// Note that the order of elements is not important.

// Example 2:
// Input: nums = [1,3,3,2,7,3], k = 3
// Output: []
// Explanation: It is not possible to divide the array satisfying all the conditions.

// Constraints:

// 		n == nums.length
// 		1 <= n <= 10^5
// 		n is a multiple of 3.
// 		1 <= nums[i] <= 10^5
// 		1 <= k <= 10^5

func divideArray(nums []int, k int) [][]int {
	l := len(nums)
	if l % 3 != 0 {
		return [][]int{}
	}
	// 先将数组排序
	sort.Ints(nums)
	for i := 0; i < l; i= i + 3 {
		// 子数组中 任意 两个元素的差必须小于或等于 k [n1,n2,n3] 取 n1 n3 做判断即可
		if (nums[i + 2] - nums[i]) > k {
			return [][]int{}
		}
	}
	// result := [][]int{}
	result := make([][]int, 0, l/3)
	for i := 0; i < l; i = i+3 {
		//result = append(result,[]int{nums[i],nums[i+1],nums[i+2]})
		result = append(result, nums[i:i+3])
	}
	return result
}

// best solution
func divideArray1(nums []int, k int) (ans [][]int) {
	slices.Sort(nums)
	for i := 2; i < len(nums); i += 3 {
		if nums[i]-nums[i-2] > k {
			return nil
		}
		ans = append(ans, nums[i-2:i+1])
	}
	return
}

func main() {
	fmt.Println(divideArray([]int{1,3,4,8,7,9,3,5,1},2))
	fmt.Println(divideArray([]int{1,3,3,2,7,3},3))

	fmt.Println(divideArray1([]int{1,3,4,8,7,9,3,5,1},2))
	fmt.Println(divideArray1([]int{1,3,3,2,7,3},3))
}