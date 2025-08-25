package main

// 3660. Jump Game IX
// You are given an integer array nums.

// From any index i, you can jump to another index j under the following rules:
//     1. Jump to index j where j > i is allowed only if nums[j] < nums[i].
//     2. Jump to index j where j < i is allowed only if nums[j] > nums[i].

// For each index i, find the maximum value in nums that can be reached by following any sequence of valid jumps starting at i.

// Return an array ans where ans[i] is the maximum value reachable starting from index i.

// Example 1:
// Input: nums = [2,1,3]
// Output: [2,2,3]
// Explanation:
// For i = 0: No jump increases the value.
// For i = 1: Jump to j = 0 as nums[j] = 2 is greater than nums[i].
// For i = 2: Since nums[2] = 3 is the maximum value in nums, no jump increases the value.
// Thus, ans = [2, 2, 3].

// Example 2:
// Input: nums = [2,3,1]
// Output: [3,3,3]
// Explanation:
// For i = 0: Jump forward to j = 2 as nums[j] = 1 is less than nums[i] = 2, then from i = 2 jump to j = 1 as nums[j] = 3 is greater than nums[2].
// For i = 1: Since nums[1] = 3 is the maximum value in nums, no jump increases the value.
// For i = 2: Jump to j = 1 as nums[j] = 3 is greater than nums[2] = 1.
// Thus, ans = [3, 3, 3].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "sort"

func maxValue(nums []int) []int {
    n, pair := len(nums), make([][]int, 0)
    for i, v := range nums {
        pair = append(pair, []int{i, v})
    }
    sort.Slice(pair, func(i, j int) bool {
        return pair[i][1] > pair[j][1] || (pair[i][1] == pair[j][1] && pair[i][0] < pair[j][0])
    })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := make([]int, n)
    low := [2]int{ 1 << 31, 0 }
    for _, v := range pair {
        if low[0] < v[1] {
            v[1] = low[1]
        }
        for i := v[0]; i < n; i++ {
            res[i] = v[1]
            low[0] = min(low[0], nums[i])
        }
        n = min(v[0], n)
        low[1] = v[1]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3]
    // Output: [2,2,3]
    // Explanation:
    // For i = 0: No jump increases the value.
    // For i = 1: Jump to j = 0 as nums[j] = 2 is greater than nums[i].
    // For i = 2: Since nums[2] = 3 is the maximum value in nums, no jump increases the value.
    // Thus, ans = [2, 2, 3].
    fmt.Println(maxValue([]int{2,1,3})) // [2,2,3]
    // Example 2:
    // Input: nums = [2,3,1]
    // Output: [3,3,3]
    // Explanation:
    // For i = 0: Jump forward to j = 2 as nums[j] = 1 is less than nums[i] = 2, then from i = 2 jump to j = 1 as nums[j] = 3 is greater than nums[2].
    // For i = 1: Since nums[1] = 3 is the maximum value in nums, no jump increases the value.
    // For i = 2: Jump to j = 1 as nums[j] = 3 is greater than nums[2] = 1.
    // Thus, ans = [3, 3, 3].
    fmt.Println(maxValue([]int{2,3,1})) // [3,3,3]

    fmt.Println(maxValue([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(maxValue([]int{9,8,7,6,5,4,3,2,1})) // [9 9 9 9 9 9 9 9 9]
}