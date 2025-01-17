package main

// 2518. Number of Great Partitions
// You are given an array nums consisting of positive integers and an integer k.

// Partition the array into two ordered groups such that each element is in exactly one group. 
// A partition is called great if the sum of elements of each group is greater than or equal to k.

// Return the number of distinct great partitions. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Two partitions are considered distinct if some element nums[i] is in different groups in the two partitions.

// Example 1:
// Input: nums = [1,2,3,4], k = 4
// Output: 6
// Explanation: The great partitions are: ([1,2,3], [4]), ([1,3], [2,4]), ([1,4], [2,3]), ([2,3], [1,4]), ([2,4], [1,3]) and ([4], [1,2,3]).

// Example 2:
// Input: nums = [3,3,3], k = 4
// Output: 0
// Explanation: There are no great partitions for this array.

// Example 3:
// Input: nums = [6,6], k = 2
// Output: 2
// Explanation: We can either put nums[0] in the first partition or in the second partition.
// The great partitions will be ([6], [6]) and ([6], [6]).

// Constraints:
//     1 <= nums.length, k <= 1000
//     1 <= nums[i] <= 10^9

import "fmt"

func countPartitions(nums []int, k int) int {
    sum, n := 0, len(nums)
    for _, v := range nums {
        sum += v
    }
    if sum < k * 2 { return 0 }
    res, mod := 1, 1_000_000_007
    dp := make([]int, k)
    dp[0] = 1
    for _, v := range nums {
        for j := k - 1; j >= v; j-- {
            dp[j] = (dp[j] + dp[j - v]) % mod
        }
    }
    for i := 0; i < n; i++ {
        res = (res * 2) % mod
    }
    for _, v := range dp {
        res = (res - v * 2 % mod + mod) % mod
    }
    return res
}

func countPartitions1(nums []int, k int) int {
    res, sum, mod := 1, 0, 1_000_000_007
    for _ , v := range nums {
        sum += v
    }
    if sum < 2 * k { return 0 }
    dp := make([]int, k)
    dp[0] = 1
    for _ , v := range nums {
        res = (2 * res) % mod
        for j := k - 1; j - v >= 0; j-- {
            dp[j] = (dp[j] + dp[j - v]) % mod
        }
    }
    for _, v := range dp {
        res = (res - 2 * v) % mod
    }
    return (res + mod) % mod
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4], k = 4
    // Output: 6
    // Explanation: The great partitions are: ([1,2,3], [4]), ([1,3], [2,4]), ([1,4], [2,3]), ([2,3], [1,4]), ([2,4], [1,3]) and ([4], [1,2,3]).
    fmt.Println(countPartitions([]int{1,2,3,4}, 4)) // 6
    // Example 2:
    // Input: nums = [3,3,3], k = 4
    // Output: 0
    // Explanation: There are no great partitions for this array.
    fmt.Println(countPartitions([]int{3,3,3}, 4)) // 0
    // Example 3:
    // Input: nums = [6,6], k = 2
    // Output: 2
    // Explanation: We can either put nums[0] in the first partition or in the second partition.
    // The great partitions will be ([6], [6]) and ([6], [6]).
    fmt.Println(countPartitions([]int{6,6}, 2)) // 2

    fmt.Println(countPartitions([]int{1,2,3,4,5,6,7,8,9}, 2)) // 508
    fmt.Println(countPartitions([]int{9,8,7,6,5,4,3,2,1}, 2)) // 508

    fmt.Println(countPartitions1([]int{1,2,3,4}, 4)) // 6
    fmt.Println(countPartitions1([]int{3,3,3}, 4)) // 0
    fmt.Println(countPartitions1([]int{6,6}, 2)) // 2
    fmt.Println(countPartitions1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 508
    fmt.Println(countPartitions1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 508
}