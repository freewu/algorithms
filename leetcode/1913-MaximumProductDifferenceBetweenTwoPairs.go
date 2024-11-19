package main

// 1913. Maximum Product Difference Between Two Pairs
// The product difference between two pairs (a, b) and (c, d) is defined as (a * b) - (c * d).
//     For example, the product difference between (5, 6) and (2, 7) is (5 * 6) - (2 * 7) = 16.
    
// Given an integer array nums, choose four distinct indices w, x, y, and z 
// such that the product difference between pairs (nums[w], nums[x]) and (nums[y], nums[z]) is maximized.

// Return the maximum such product difference.

// Example 1:
// Input: nums = [5,6,2,7,4]
// Output: 34
// Explanation: We can choose indices 1 and 3 for the first pair (6, 7) and indices 2 and 4 for the second pair (2, 4).
// The product difference is (6 * 7) - (2 * 4) = 34.

// Example 2:
// Input: nums = [4,2,5,9,7,4,8]
// Output: 64
// Explanation: We can choose indices 3 and 6 for the first pair (9, 8) and indices 1 and 5 for the second pair (2, 4).
// The product difference is (9 * 8) - (2 * 4) = 64.

// Constraints:
//     4 <= nums.length <= 10^4
//     1 <= nums[i] <= 10^4

import "fmt"
import "sort"

func maxProductDifference(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums) - 1] * nums[len(nums) - 2] - nums[0] * nums[1]
}

func maxProductDifference1(nums []int) int {
    if len(nums) < 4 { return 0 }
    arr :=  make([]int, 10001)
    for _, v := range nums {
        arr[v]++
    }
    res := []int{}
    for k, v := range arr {
        for i := 0; i < v; i++ {
            res = append(res, k)
        }
    }
    n := len(res)
    return res[n - 1] * res[n - 2] - res[0] * res[1]
}

func maxProductDifference2(nums []int) int {
    a, b, c, d := 0, 0, 0, 0
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[a] { a = i }
        if nums[i] > nums[d] { d = i }
    }
    b, c = d, a
    for i := 0; i < len(nums); i++ {
        if nums[i] <= nums[b] && i != a { b = i }
        if nums[i] >= nums[c] && i != d { c = i }
    }
    return nums[c] * nums[d] - nums[b] * nums[a]
}

func main() {
    // Example 1:
    // Input: nums = [5,6,2,7,4]
    // Output: 34
    // Explanation: We can choose indices 1 and 3 for the first pair (6, 7) and indices 2 and 4 for the second pair (2, 4).
    // The product difference is (6 * 7) - (2 * 4) = 34.
    fmt.Println(maxProductDifference([]int{5,6,2,7,4})) // 34
    // Example 2:
    // Input: nums = [4,2,5,9,7,4,8]
    // Output: 64
    // Explanation: We can choose indices 3 and 6 for the first pair (9, 8) and indices 1 and 5 for the second pair (2, 4).
    // The product difference is (9 * 8) - (2 * 4) = 64.
    fmt.Println(maxProductDifference([]int{4,2,5,9,7,4,8})) // 64

    fmt.Println(maxProductDifference1([]int{5,6,2,7,4})) // 34
    fmt.Println(maxProductDifference1([]int{4,2,5,9,7,4,8})) // 64

    fmt.Println(maxProductDifference2([]int{5,6,2,7,4})) // 34
    fmt.Println(maxProductDifference2([]int{4,2,5,9,7,4,8})) // 64
}