package main

import "fmt"

// 1043. Partition Array for Maximum Sum
// Given an integer array arr, partition the array into (contiguous) subarrays of length at most k. 
// After partitioning, each subarray has their values changed to become the maximum value of that subarray.
// Return the largest sum of the given array after partitioning. 
// Test cases are generated so that the answer fits in a 32-bit integer.
 
// Example 1:
// Input: arr = [1,15,7,9,2,5,10], k = 3
// Output: 84
// Explanation: arr becomes [15,15,15,9,10,10,10]

// Example 2:
// Input: arr = [1,4,1,5,7,3,6,1,9,9,3], k = 4
// Output: 83

// Example 3:
// Input: arr = [1], k = 1
// Output: 1
 
// Constraints:
// 		1 <= arr.length <= 500
// 		0 <= arr[i] <= 10^9
// 		1 <= k <= arr.length

func maxSumAfterPartitioning(arr []int, k int) int {
	l := len(arr);
	dp := make([]int,l)
	dp[0] = arr[0];
	for i := 1; i < l; i = i + 1  {
		j := i
		maxInPart := arr[j]
		for ; j > i - k && j >= 0; j = j - 1 {
			maxInPart = max(maxInPart,arr[j]);
			prev := 0
			if j - 1 >= 0 {
				prev = dp[j - 1]
			} 
			dp[i] = max(dp[i], prev + maxInPart * (i - j + 1))

			fmt.Println(dp)
		}
	}
	return dp[l - 1]
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1,15,7,9,2,5,10}, 3)) // 84
	fmt.Println(maxSumAfterPartitioning([]int{1,4,1,5,7,3,6,1,9,9,3}, 4)) // 83
	fmt.Println(maxSumAfterPartitioning([]int{1}, 1)) // 1
}