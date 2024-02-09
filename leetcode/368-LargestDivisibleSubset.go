package main

// 368. Largest Divisible Subset
// Given a set of distinct positive integers nums, 
// return the largest subset answer such that every pair (answer[i], answer[j]) of elements in this subset satisfies:
// 		answer[i] % answer[j] == 0, or
// 		answer[j] % answer[i] == 0

// If there are multiple solutions, return any of them.

// Example 1:
// Input: nums = [1,2,3]
// Output: [1,2]
// Explanation: [1,3] is also accepted.

// Example 2:
// Input: nums = [1,2,4,8]
// Output: [1,2,4,8]

// Constraints:
// 		1 <= nums.length <= 1000
// 		1 <= nums[i] <= 2 * 10^9
// 		All the integers in nums are unique.

import "fmt"
import "sort"

func largestDivisibleSubset(nums []int) []int {
	// 先将集合排序
	sort.Ints(nums)
	dp, res := make([]int, len(nums)), []int{}
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	// 以某一个小的数作为基准，不断的选择能整除的数加入集合
	for i := 1; i < len(nums); i++ {
		for j, v := range nums[:i] {
			// 能整除 则 + 1
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	// 通过得到的 最大元素(maxVal) 反推出最大集合
	for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1,2,3})) // [1,2]
	fmt.Println(largestDivisibleSubset([]int{1,2,4,8})) // [1,2,4,8]
}