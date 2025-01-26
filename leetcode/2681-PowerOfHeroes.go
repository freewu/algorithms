package main

// 2681. Power of Heroes
// You are given a 0-indexed integer array nums representing the strength of some heroes. 
// The power of a group of heroes is defined as follows:
//     1. Let i0, i1, ... ,ik be the indices of the heroes in a group. 
//        Then, the power of this group is max(nums[i0], nums[i1], ... ,nums[ik])2 * min(nums[i0], nums[i1], ... ,nums[ik]).

// Return the sum of the power of all non-empty groups of heroes possible. 
// Since the sum could be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [2,1,4]
// Output: 141
// Explanation: 
// 1st group: [2] has power = 22 * 2 = 8.
// 2nd group: [1] has power = 12 * 1 = 1. 
// 3rd group: [4] has power = 42 * 4 = 64. 
// 4th group: [2,1] has power = 22 * 1 = 4. 
// 5th group: [2,4] has power = 42 * 2 = 32. 
// 6th group: [1,4] has power = 42 * 1 = 16. 
// ​​​​​​​7th group: [2,1,4] has power = 42​​​​​​​ * 1 = 16. 
// The sum of powers of all groups is 8 + 1 + 64 + 4 + 32 + 16 + 16 = 141.

// Example 2:
// Input: nums = [1,1,1]
// Output: 7
// Explanation: A total of 7 groups are possible, and the power of each group will be 1. 
// Therefore, the sum of the powers of all groups is 7.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"

func sumOfPower(nums []int) int {
    res, sum, mod := 0, 0, 1_000_000_007
    sort.Ints(nums)
    for _, v := range nums {
        res = (res + (v * v % mod) * (v + sum)) % mod
        sum = (sum << 1 + v) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,4]
    // Output: 141
    // Explanation: 
    // 1st group: [2] has power = 22 * 2 = 8.
    // 2nd group: [1] has power = 12 * 1 = 1. 
    // 3rd group: [4] has power = 42 * 4 = 64. 
    // 4th group: [2,1] has power = 22 * 1 = 4. 
    // 5th group: [2,4] has power = 42 * 2 = 32. 
    // 6th group: [1,4] has power = 42 * 1 = 16. 
    // ​​​​​​​7th group: [2,1,4] has power = 42​​​​​​​ * 1 = 16. 
    // The sum of powers of all groups is 8 + 1 + 64 + 4 + 32 + 16 + 16 = 141.
    fmt.Println(sumOfPower([]int{2,1,4})) // 141
    // Example 2:
    // Input: nums = [1,1,1]
    // Output: 7
    // Explanation: A total of 7 groups are possible, and the power of each group will be 1. 
    // Therefore, the sum of the powers of all groups is 7.
    fmt.Println(sumOfPower([]int{1,1,1})) // 7

    fmt.Println(sumOfPower([]int{1,2,3,4,5,6,7,8,9})) // 67293
    fmt.Println(sumOfPower([]int{9,8,7,6,5,4,3,2,1})) // 67293
}