package main

// 2680. Maximum OR
// You are given a 0-indexed integer array nums of length n and an integer k. 
// In an operation, you can choose an element and multiply it by 2.

// Return the maximum possible value of nums[0] | nums[1] | ... | nums[n - 1] that can be obtained after applying the operation on nums at most k times.

// Note that a | b denotes the bitwise or between two integers a and b.

// Example 1:
// Input: nums = [12,9], k = 1
// Output: 30
// Explanation: If we apply the operation to index 1, our new array nums will be equal to [12,18]. Thus, we return the bitwise or of 12 and 18, which is 30.

// Example 2:
// Input: nums = [8,1,2], k = 2
// Output: 35
// Explanation: If we apply the operation twice on index 0, we yield a new array of [32,1,2]. Thus, we return 32|1|2 = 35.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 15

import "fmt"

func maximumOr(nums []int, k int) int64 {
    multiple, any := 0, 0  // O(n) - determine bits that appear in multiple numbers and those that appear in any number
    for _, v := range nums {
        multiple |= any & v
        any |= v
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := any
    for _, v := range nums {
        res = max(res, (multiple | (any & ^v) | v << k))
    }
    return int64(res)
}

// 贪心 + 前后缀分解
func maximumOr1(nums []int, k int) int64 {
    n := len(nums)
    suffix := make([]int, n + 1) // // suffix[i] 表示 nums[i+1] 到 nums[n-1] 的 OR
    for i := n - 1; i > 0; i-- {
        suffix[i] = suffix[i + 1] | nums[i]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, pre := 0, 0 // pre 表示 nums[0] 到 nums[i-1] 的 OR
    for i, v := range nums {
        res = max(res, pre| v << k| suffix[i + 1])
        pre |= v
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [12,9], k = 1
    // Output: 30
    // Explanation: If we apply the operation to index 1, our new array nums will be equal to [12,18]. Thus, we return the bitwise or of 12 and 18, which is 30.
    fmt.Println(maximumOr([]int{12, 9}, 1)) // 30
    // Example 2:
    // Input: nums = [8,1,2], k = 2
    // Output: 35
    // Explanation: If we apply the operation twice on index 0, we yield a new array of [32,1,2]. Thus, we return 32|1|2 = 35.
    fmt.Println(maximumOr([]int{8,1,2}, 2)) // 35

    fmt.Println(maximumOr([]int{1,2,3,4,5,6,7,8,9}, 2)) // 47
    fmt.Println(maximumOr([]int{9,8,7,6,5,4,3,2,1}, 2)) // 47

    fmt.Println(maximumOr1([]int{12, 9}, 1)) // 30
    fmt.Println(maximumOr1([]int{8,1,2}, 2)) // 35
    fmt.Println(maximumOr1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 47
    fmt.Println(maximumOr1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 47
}