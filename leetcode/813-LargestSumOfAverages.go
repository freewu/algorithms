package main

// 813. Largest Sum of Averages
// You are given an integer array nums and an integer k. 
// You can partition the array into at most k non-empty adjacent subarrays. 
// The score of a partition is the sum of the averages of each subarray.

// Note that the partition must use every integer in nums, and that the score is not necessarily an integer.

// Return the maximum score you can achieve of all the possible partitions. 
// Answers within 10^-6 of the actual answer will be accepted.

// Example 1:
// Input: nums = [9,1,2,3,9], k = 3
// Output: 20.00000
// Explanation: 
// The best choice is to partition nums into [9], [1, 2, 3], [9]. The answer is 9 + (1 + 2 + 3) / 3 + 9 = 20.
// We could have also partitioned nums into [9, 1], [2], [3, 9], for example.
// That partition would lead to a score of 5 + 2 + 6 = 13, which is worse.

// Example 2:
// Input: nums = [1,2,3,4,5,6,7], k = 4
// Output: 20.50000

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 10^4
//     1 <= k <= nums.length

import "fmt"

// dp + memo
func largestSumOfAverages(nums []int, k int) float64 {
    sum := make([]float64, len(nums) + 1)
    memo := make([][]float64, k + 1)
    for i := range memo {
        memo[i] = make([]float64, len(nums) + 1)
    }
    for i, v := range nums {
        sum[i+1] = sum[i] + float64(v)
    }
    max := func (x, y float64) float64 { if x > y { return x; }; return y; }
    var dfs func(n, k int) float64 
    dfs = func(n, k int) float64 {
        if memo[k][n] > 0 { 
            return memo[k][n]
        }
        if k == 1 { 
            return sum[n] / float64(n) 
        }
        for i := k - 1; i < n; i++ {
            memo[k][n] = max(memo[k][n], dfs(i, k - 1) + (sum[n] - sum[i]) / float64(n - i))
        }
        return memo[k][n]
    }
    return dfs(len(nums), k)
}

func main() {
    // Example 1:
    // Input: nums = [9,1,2,3,9], k = 3
    // Output: 20.00000
    // Explanation: 
    // The best choice is to partition nums into [9], [1, 2, 3], [9]. The answer is 9 + (1 + 2 + 3) / 3 + 9 = 20.
    // We could have also partitioned nums into [9, 1], [2], [3, 9], for example.
    // That partition would lead to a score of 5 + 2 + 6 = 13, which is worse.
    fmt.Println(largestSumOfAverages([]int{9,1,2,3,9}, 3)) // 20.00000
    // Example 2:
    // Input: nums = [1,2,3,4,5,6,7], k = 4
    // Output: 20.50000
    fmt.Println(largestSumOfAverages([]int{1,2,3,4,5,6,7}, 4)) // 20.50000
}