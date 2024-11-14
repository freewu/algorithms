package main

// 1862. Sum of Floored Pairs
// Given an integer array nums, 
// return the sum of floor(nums[i] / nums[j]) for all pairs of indices 0 <= i, j < nums.length in the array. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// The floor() function returns the integer part of the division.

// Example 1:
// Input: nums = [2,5,9]
// Output: 10
// Explanation:
// floor(2 / 5) = floor(2 / 9) = floor(5 / 9) = 0
// floor(2 / 2) = floor(5 / 5) = floor(9 / 9) = 1
// floor(5 / 2) = 2
// floor(9 / 2) = 4
// floor(9 / 5) = 1
// We calculate the floor of the division for every pair of indices in the array then sum them up.

// Example 2:
// Input: nums = [7,7,7,7,7,7,7]
// Output: 49

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

func sumOfFlooredPairs(nums []int) int {
    res, mx, mod := 0, -1, 1_000_000_007
    for _, v := range nums { // 找到最大值
        if v > mx { mx = v }
    }
    freq := make([]int, mx + 1)
    for _, v := range nums {
        freq[v]++
    }
    prefix := make([]int, mx + 1)
    for i := 1; i <= mx; i++ {
        prefix[i] = prefix[i - 1] + freq[i]
    }
    for i := 0; i <= mx; i++ {
        if freq[i] == 0 { continue }
        sum := 0
        for j := 1; j * i <= mx; j++ {
            left, right := j * i, (j + 1) * i - 1
            if right > mx { right = mx }
            sum = (sum + (j * (prefix[right] - prefix[left - 1]))) % mod
        }
        res = (res + freq[i] * sum) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,5,9]
    // Output: 10
    // Explanation:
    // floor(2 / 5) = floor(2 / 9) = floor(5 / 9) = 0
    // floor(2 / 2) = floor(5 / 5) = floor(9 / 9) = 1
    // floor(5 / 2) = 2
    // floor(9 / 2) = 4
    // floor(9 / 5) = 1
    // We calculate the floor of the division for every pair of indices in the array then sum them up.
    fmt.Println(sumOfFlooredPairs([]int{2,5,9})) // 10
    // Example 2:
    // Input: nums = [7,7,7,7,7,7,7]
    // Output: 49
    fmt.Println(sumOfFlooredPairs([]int{7,7,7,7,7,7,7})) // 49
}