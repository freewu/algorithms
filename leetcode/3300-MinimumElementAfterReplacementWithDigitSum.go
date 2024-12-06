package main

// 3300. Minimum Element After Replacement With Digit Sum
// You are given an integer array nums.

// You replace each element in nums with the sum of its digits.

// Return the minimum element in nums after all replacements.

// Example 1:
// Input: nums = [10,12,13,14]
// Output: 1
// Explanation:
// nums becomes [1, 3, 4, 5] after all replacements, with minimum element 1.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 1
// Explanation:
// nums becomes [1, 2, 3, 4] after all replacements, with minimum element 1.

// Example 3:
// Input: nums = [999,19,199]
// Output: 10
// Explanation:
// nums becomes [27, 10, 19] after all replacements, with minimum element 10.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 10^4

import "fmt"

func minElement(nums []int) int {
    res := 1 << 31
    calc := func(n int) int { // 123 => 1 + 2 + 3
        res := 0
        for n > 0 {
            res += n % 10
            n /= 10
        }
        return res
    }
    for _, v := range nums {
        res = min(res, calc(v))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,12,13,14]
    // Output: 1
    // Explanation:
    // nums becomes [1, 3, 4, 5] after all replacements, with minimum element 1.
    fmt.Println(minElement([]int{10,12,13,14})) // 1
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 1
    // Explanation:
    // nums becomes [1, 2, 3, 4] after all replacements, with minimum element 1.
    fmt.Println(minElement([]int{1,2,3,4})) // 1
    // Example 3:
    // Input: nums = [999,19,199]
    // Output: 10
    // Explanation:
    // nums becomes [27, 10, 19] after all replacements, with minimum element 10.
    fmt.Println(minElement([]int{999,19,199})) // 10
}