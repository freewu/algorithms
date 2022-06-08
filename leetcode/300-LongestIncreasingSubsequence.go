package main

/**
300. Longest Increasing Subsequence
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some
or no elements without changing the order of the remaining elements.
For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].


Example 1:

	Input: nums = [10,9,2,5,3,7,101,18]
	Output: 4
	Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:

	Input: nums = [0,1,0,3,2,3]
	Output: 4

Example 3:

	Input: nums = [7,7,7,7,7,7,7]
	Output: 1


Constraints:

	1 <= nums.length <= 2500
	-10^4 <= nums[i] <= 10^4


Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?

给定一个无序的整数数组，找到其中最长上升子序列的长度
*/

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// O(n^2) DP
func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}
	return res
}

//  O(n log n) DP
func lengthOfLIS1(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}

func main() {
	fmt.Printf("lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}) =%v\n", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4 [2,3,7,101]
	fmt.Printf("lengthOfLIS([]int{0,1,0,3,2,3}) = %v\n", lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))                         // 4 [0,0,2,3]
	fmt.Printf("lengthOfLIS([]int{7,7,7,7,7,7,7}) = %v\n", lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))                    // 1

	fmt.Printf("lengthOfLIS1([]int{10, 9, 2, 5, 3, 7, 101, 18}) =%v\n", lengthOfLIS1([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4 [2,3,7,101]
	fmt.Printf("lengthOfLIS1([]int{0,1,0,3,2,3}) = %v\n", lengthOfLIS1([]int{0, 1, 0, 3, 2, 3}))                         // 4 [0,0,2,3]
	fmt.Printf("lengthOfLIS1([]int{7,7,7,7,7,7,7}) = %v\n", lengthOfLIS1([]int{7, 7, 7, 7, 7, 7, 7}))                    // 1
}
