package main

// 1121. Divide Array Into Increasing Sequences
// Given an integer array nums sorted in non-decreasing order and an integer k, 
// return true if this array can be divided into one or more disjoint increasing subsequences of length at least k, 
// or false otherwise.

// Example 1:
// Input: nums = [1,2,2,3,3,4,4], k = 3
// Output: true
// Explanation: The array can be divided into two subsequences [1,2,3,4] and [2,3,4] with lengths at least 3 each.

// Example 2:
// Input: nums = [5,6,6,7,8], k = 3
// Output: false
// Explanation: There is no way to divide the array using the conditions required.

// Constraints:
//     1 <= k <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     nums is sorted in non-decreasing order.

import "fmt"

func canDivideIntoSubsequences(nums []int, k int) bool {
    freq, m := map[int]int{}, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 非递减数组中，递增子序列的概念：将一个数组中的数字去重，留下的元素就是递增子序列
    for _, v := range nums {
        freq[v]++
        m = max(freq[v], m)
    }
    return m * k <= len(nums) // 只要判断 数量 * 序列长度 <= 数组长度
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,3,3,4,4], k = 3
    // Output: true
    // Explanation: The array can be divided into two subsequences [1,2,3,4] and [2,3,4] with lengths at least 3 each.
    fmt.Println(canDivideIntoSubsequences([]int{1,2,2,3,3,4,4}, 3)) // true
    // Example 2:
    // Input: nums = [5,6,6,7,8], k = 3
    // Output: false
    // Explanation: There is no way to divide the array using the conditions required.
    fmt.Println(canDivideIntoSubsequences([]int{5,6,6,7,8}, 3)) // false
}