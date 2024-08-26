package main

// 2815. Max Pair Sum in an Array
// You are given an integer array nums. 
// You have to find the maximum sum of a pair of numbers from nums such that the largest digit in both numbers is equal.

// For example, 2373 is made up of three distinct digits: 2, 3, and 7, where 7 is the largest among them.

// Return the maximum sum or -1 if no such pair exists.

// Example 1:
// Input: nums = [112,131,411]
// Output: -1
// Explanation:
// Each numbers largest digit in order is [2,3,4].

// Example 2:
// Input: nums = [2536,1613,3366,162]
// Output: 5902
// Explanation:
// All the numbers have 6 as their largest digit, so the answer is 2536 + 3366 = 5902.

// Example 3:
// Input: nums = [51,71,17,24,42]
// Output: 88
// Explanation:
// Each number's largest digit in order is [5,7,7,4,4].
// So we have only two possible pairs, 71 + 17 = 88 and 24 + 42 = 66.

// Constraints:
//     2 <= nums.length <= 100
//     1 <= nums[i] <= 10^4

import "fmt"

func maxSum(nums []int) int {
    mp, res := [10]int{}, -1
    maxDigit := func(a int) int { // 请最大的数字 如: 1342  最大数字为 4
        res := 0
        for ; a > 0; a /= 10 {
            d := a % 10
            if d > res {
                res = d
            }
        }
        return res
    }
    for _, v := range nums {
        d := maxDigit(v)
        if mp[d] > 0 && mp[d] + v > res {
            res = mp[d] + v
        }
        if v > mp[d] { // mp 保存最大的 即可
            mp[d] = v
        }
    }
    return res
}

func maxSum1(nums []int) int {
    mp, res := [10]int{}, -1
    for i := range mp { // 初始化为-1，为-1则表示现在还没有数
        mp[i] = -1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    maxDigit := func(a int) int { // 请最大的数字 如: 1342  最大数字为 4
        res := 0
        for ; a > 0; a /= 10 {
            d := a % 10
            if d > res {
                res = d
            }
        }
        return res
    }
    for _, v := range nums {
        d := maxDigit(v)
        if mp[d] == -1 {
            mp[d] = v
            continue // 第一次出现，直接继续即可
        }
        res = max(mp[d] + v, res)
        mp[d] = max(mp[d], v)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [112,131,411]
    // Output: -1
    // Explanation:
    // Each numbers largest digit in order is [2,3,4].
    fmt.Println(maxSum([]int{112,131,411})) // -1
    // Example 2:
    // Input: nums = [2536,1613,3366,162]
    // Output: 5902
    // Explanation:
    // All the numbers have 6 as their largest digit, so the answer is 2536 + 3366 = 5902.
    fmt.Println(maxSum([]int{2536,1613,3366,162})) // 5902 (2536 + 3366 = 5902)
    // Example 3:
    // Input: nums = [51,71,17,24,42]
    // Output: 88
    // Explanation:
    // Each number's largest digit in order is [5,7,7,4,4].
    // So we have only two possible pairs, 71 + 17 = 88 and 24 + 42 = 66.
    fmt.Println(maxSum([]int{51,71,17,24,42})) // 71 + 17 = 88 and 24 + 42 = 66

    fmt.Println(maxSum1([]int{112,131,411})) // -1
    fmt.Println(maxSum1([]int{2536,1613,3366,162})) // 5902 (2536 + 3366 = 5902)
    fmt.Println(maxSum1([]int{51,71,17,24,42})) // 71 + 17 = 88 and 24 + 42 = 66
}