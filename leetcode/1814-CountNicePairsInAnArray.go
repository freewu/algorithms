package main

// 1814. Count Nice Pairs in an Array
// You are given an array nums that consists of non-negative integers. 
// Let us define rev(x) as the reverse of the non-negative integer x. 
// For example, rev(123) = 321, and rev(120) = 21. 
// A pair of indices (i, j) is nice if it satisfies all of the following conditions:
//     0 <= i < j < nums.length
//     nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])

// Return the number of nice pairs of indices. 
// Since that number can be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [42,11,1,97]
// Output: 2
// Explanation: The two pairs are:
//  - (0,3) : 42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121.
//  - (1,2) : 11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12.

// Example 2:
// Input: nums = [13,10,35,24,76]
// Output: 4

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"

func countNicePairs(nums []int) int {
    rev := func(n int) int {
        res := 0
        for n > 0 {
            remainder := n % 10
            res *= 10
            res += remainder
            n /= 10
        }
        return res
    }
    res, mp := 0, make(map[int]int)
    for i := 0; i < len(nums); i++ {
        rev := rev(nums[i])
        diff := nums[i] - rev
        mp[diff]++
    }
    for _, v := range mp {
        if v == 1 { continue } // 单个的不能成对
        res = (res + v * (v - 1) / 2) % 1_000_000_007 //  v * (v - 1) / 2  v = 4  6
    }
    return res
}

func countNicePairs1(nums []int) int {
    res, mp := 0, make(map[int]int)
    rev := func(num int) int {
        res := 0
        for i := num; i > 0; i /= 10 {
            res = res * 10 + i % 10
        }
        return res
    }
    for _, v := range nums {
        diff := v - rev(v)
        if w, ok := mp[diff]; ok {
            res += w
        }
        mp[diff]++
    }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: nums = [42,11,1,97]
    // Output: 2
    // Explanation: The two pairs are:
    //  - (0,3) : 42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121.
    //  - (1,2) : 11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12.
    fmt.Println(countNicePairs([]int{42,11,1,97})) // 2
    // Example 2:
    // Input: nums = [13,10,35,24,76]
    // Output: 4
    fmt.Println(countNicePairs([]int{13,10,35,24,76})) // 4

    fmt.Println(countNicePairs1([]int{42,11,1,97})) // 2
    fmt.Println(countNicePairs1([]int{13,10,35,24,76})) // 4
}