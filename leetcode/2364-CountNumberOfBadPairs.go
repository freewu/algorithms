package main

// 2364. Count Number of Bad Pairs
// You are given a 0-indexed integer array nums. 
// A pair of indices (i, j) is a bad pair if i < j and j - i != nums[j] - nums[i].

// Return the total number of bad pairs in nums.

// Example 1:
// Input: nums = [4,1,3,3]
// Output: 5
// Explanation: The pair (0, 1) is a bad pair since 1 - 0 != 1 - 4.
// The pair (0, 2) is a bad pair since 2 - 0 != 3 - 4, 2 != -1.
// The pair (0, 3) is a bad pair since 3 - 0 != 3 - 4, 3 != -1.
// The pair (1, 2) is a bad pair since 2 - 1 != 3 - 1, 1 != 2.
// The pair (2, 3) is a bad pair since 3 - 2 != 3 - 3, 1 != 0.
// There are a total of 5 bad pairs, so we return 5.

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: 0
// Explanation: There are no bad pairs.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func countBadPairs(nums []int) int64 {
    n := len(nums)
    mp := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        mp[nums[i] - i]++
    }
    res := int64(n * (n - 1) / 2)
    for _, v := range mp {
        if v > 1 {
            res -= int64(v * (v - 1) / 2)
        }
    }
    return res
}

func countBadPairs1(nums []int) int64 {
    good, n := 0, len(nums)
    mp := make(map[int]int) 
    for i , v := range nums {
        mp[v - i] ++
    }
    for _ , v := range mp {
        good += (v * (v - 1) / 2)
    }
    return int64(n * (n-1) / 2 - good)
}

func main() {
    // Example 1:
    // Input: nums = [4,1,3,3]
    // Output: 5
    // Explanation: The pair (0, 1) is a bad pair since 1 - 0 != 1 - 4.
    // The pair (0, 2) is a bad pair since 2 - 0 != 3 - 4, 2 != -1.
    // The pair (0, 3) is a bad pair since 3 - 0 != 3 - 4, 3 != -1.
    // The pair (1, 2) is a bad pair since 2 - 1 != 3 - 1, 1 != 2.
    // The pair (2, 3) is a bad pair since 3 - 2 != 3 - 3, 1 != 0.
    // There are a total of 5 bad pairs, so we return 5.
    fmt.Println(countBadPairs([]int{4,1,3,3})) // 5
    // Example 2:
    // Input: nums = [1,2,3,4,5]
    // Output: 0
    // Explanation: There are no bad pairs.
    fmt.Println(countBadPairs([]int{1,2,3,4,5})) // 0

    fmt.Println(countBadPairs1([]int{4,1,3,3})) // 5
    fmt.Println(countBadPairs1([]int{1,2,3,4,5})) // 0
}