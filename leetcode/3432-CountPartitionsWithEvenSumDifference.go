package main

// 3432. Count Partitions with Even Sum Difference
// You are given an integer array nums of length n.

// A partition is defined as an index i where 0 <= i < n - 1, splitting the array into two non-empty subarrays such that:
//     1. Left subarray contains indices [0, i].
//     2. Right subarray contains indices [i + 1, n - 1].

// Return the number of partitions where the difference between the sum of the left and right subarrays is even.

// Example 1:
// Input: nums = [10,10,3,7,6]
// Output: 4
// Explanation:
// The 4 partitions are:
// [10], [10, 3, 7, 6] with a sum difference of 10 - 26 = -16, which is even.
// [10, 10], [3, 7, 6] with a sum difference of 20 - 16 = 4, which is even.
// [10, 10, 3], [7, 6] with a sum difference of 23 - 13 = 10, which is even.
// [10, 10, 3, 7], [6] with a sum difference of 30 - 6 = 24, which is even.

// Example 2:
// Input: nums = [1,2,2]
// Output: 0
// Explanation:
// No partition results in an even sum difference.

// Example 3:
// Input: nums = [2,4,6,8]
// Output: 3
// Explanation:
// All partitions result in an even sum difference. 

// Constraints:
//     2 <= n == nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func countPartitions(nums []int) int {
    res, sum, left, n := 0, 0, 0, len(nums)
    for _, v := range nums {
        sum += v
    }
    for i := 0; i < n - 1; i++ {
        left += nums[i]
        if (left - (sum-left)) % 2 == 0 {
            res++
        }
    }
    return res
}

func countPartitions1(nums []int) int {
    res, n := 0, len(nums)
    prefix := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        prefix[i] = prefix[i - 1] + nums[i-1]
    }
    for i := 1; i < n; i++ {
        if (prefix[n] - prefix[i] - prefix[i]) % 2 == 0 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,10,3,7,6]
    // Output: 4
    // Explanation:
    // The 4 partitions are:
    // [10], [10, 3, 7, 6] with a sum difference of 10 - 26 = -16, which is even.
    // [10, 10], [3, 7, 6] with a sum difference of 20 - 16 = 4, which is even.
    // [10, 10, 3], [7, 6] with a sum difference of 23 - 13 = 10, which is even.
    // [10, 10, 3, 7], [6] with a sum difference of 30 - 6 = 24, which is even.
    fmt.Println(countPartitions([]int{10,10,3,7,6})) // 4
    // Example 2:
    // Input: nums = [1,2,2]
    // Output: 0
    // Explanation:
    // No partition results in an even sum difference.
    fmt.Println(countPartitions([]int{1,2,2})) // 0
    // Example 3:
    // Input: nums = [2,4,6,8]
    // Output: 3
    // Explanation:
    // All partitions result in an even sum difference. 
    fmt.Println(countPartitions([]int{2,4,6,8})) // 3

    fmt.Println(countPartitions([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countPartitions([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(countPartitions1([]int{10,10,3,7,6})) // 4
    fmt.Println(countPartitions1([]int{1,2,2})) // 0
    fmt.Println(countPartitions1([]int{2,4,6,8})) // 3
    fmt.Println(countPartitions1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countPartitions1([]int{9,8,7,6,5,4,3,2,1})) // 0
}