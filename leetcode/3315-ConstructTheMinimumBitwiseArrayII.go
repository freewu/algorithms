package main

// 3315. Construct the Minimum Bitwise Array II
// You are given an array nums consisting of n prime integers.

// You need to construct an array ans of length n, 
// such that, for each index i, the bitwise OR of ans[i] and ans[i] + 1 is equal to nums[i], 
// i.e. ans[i] OR (ans[i] + 1) == nums[i].

// Additionally, you must minimize each value of ans[i] in the resulting array.

// If it is not possible to find such a value for ans[i] that satisfies the condition, then set ans[i] = -1.

// Example 1:
// Input: nums = [2,3,5,7]
// Output: [-1,1,4,3]
// Explanation:
// For i = 0, as there is no value for ans[0] that satisfies ans[0] OR (ans[0] + 1) = 2, so ans[0] = -1.
// For i = 1, the smallest ans[1] that satisfies ans[1] OR (ans[1] + 1) = 3 is 1, because 1 OR (1 + 1) = 3.
// For i = 2, the smallest ans[2] that satisfies ans[2] OR (ans[2] + 1) = 5 is 4, because 4 OR (4 + 1) = 5.
// For i = 3, the smallest ans[3] that satisfies ans[3] OR (ans[3] + 1) = 7 is 3, because 3 OR (3 + 1) = 7.

// Example 2:
// Input: nums = [11,13,31]
// Output: [9,12,15]
// Explanation:
// For i = 0, the smallest ans[0] that satisfies ans[0] OR (ans[0] + 1) = 11 is 9, because 9 OR (9 + 1) = 11.
// For i = 1, the smallest ans[1] that satisfies ans[1] OR (ans[1] + 1) = 13 is 12, because 12 OR (12 + 1) = 13.
// For i = 2, the smallest ans[2] that satisfies ans[2] OR (ans[2] + 1) = 31 is 15, because 15 OR (15 + 1) = 31.

// Constraints:
//     1 <= nums.length <= 100
//     2 <= nums[i] <= 10^9
//     nums[i] is a prime number.

import "fmt"

func minBitwiseArray(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    for i := 0; i < n; i++ {
        v := nums[i]
        if v == 2 { // 没法找到符合 条件 的 ans[i] ，那么 ans[i] = -1
            res[i] = -1
            continue
        }
        for j := 0; j < 32; j++ {
            if ((v >> j) & 1) == 0 {
                res[i] = v ^ (1 << (j - 1))
                break
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,5,7]
    // Output: [-1,1,4,3]
    // Explanation:
    // For i = 0, as there is no value for ans[0] that satisfies ans[0] OR (ans[0] + 1) = 2, so ans[0] = -1.
    // For i = 1, the smallest ans[1] that satisfies ans[1] OR (ans[1] + 1) = 3 is 1, because 1 OR (1 + 1) = 3.
    // For i = 2, the smallest ans[2] that satisfies ans[2] OR (ans[2] + 1) = 5 is 4, because 4 OR (4 + 1) = 5.
    // For i = 3, the smallest ans[3] that satisfies ans[3] OR (ans[3] + 1) = 7 is 3, because 3 OR (3 + 1) = 7.
    fmt.Println(minBitwiseArray([]int{2,3,5,7})) // [-1,1,4,3]
    // Example 2:
    // Input: nums = [11,13,31]
    // Output: [9,12,15]
    // Explanation:
    // For i = 0, the smallest ans[0] that satisfies ans[0] OR (ans[0] + 1) = 11 is 9, because 9 OR (9 + 1) = 11.
    // For i = 1, the smallest ans[1] that satisfies ans[1] OR (ans[1] + 1) = 13 is 12, because 12 OR (12 + 1) = 13.
    // For i = 2, the smallest ans[2] that satisfies ans[2] OR (ans[2] + 1) = 31 is 15, because 15 OR (15 + 1) = 31.
    fmt.Println(minBitwiseArray([]int{11,13,31})) // [9,12,15]

    fmt.Println(minBitwiseArray([]int{2,3,5,7,11,13,17,19,23})) // [-1 1 4 3 9 12 16 17 19]
    fmt.Println(minBitwiseArray([]int{23,19,17,13,11,7,5,3,2})) // [19 17 16 12 9 3 4 1 -1]
}