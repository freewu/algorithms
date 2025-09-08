package main

// 3676. Count Bowl Subarrays
// You are given an integer array nums with distinct elements.

// A subarray nums[l...r] of nums is called a bowl if:
//     1. The subarray has length at least 3. That is, r - l + 1 >= 3.
//     2. The minimum of its two ends is strictly greater than the maximum of all elements in between. 
//        That is, min(nums[l], nums[r]) > max(nums[l + 1], ..., nums[r - 1]).

// Return the number of bowl subarrays in nums.

// Example 1:
// Input: nums = [2,5,3,1,4]
// Output: 2
// Explanation:
// The bowl subarrays are [3, 1, 4] and [5, 3, 1, 4].
// [3, 1, 4] is a bowl because min(3, 4) = 3 > max(1) = 1.
// [5, 3, 1, 4] is a bowl because min(5, 4) = 4 > max(3, 1) = 3.

// Example 2:
// Input: nums = [5,1,2,3,4]
// Output: 3
// Explanation:
// The bowl subarrays are [5, 1, 2], [5, 1, 2, 3] and [5, 1, 2, 3, 4].

// Example 3:
// Input: nums = [1000000000,999999999,999999998]
// Output: 0
// Explanation:
// No subarray is a bowl.

// Constraints:
//     3 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     nums consists of distinct elements.

import "fmt"

// Monotonic stack
func bowlSubarrays(nums []int) int64 {
    res, stack := 0, []int{}
    for i := range nums {
        // remove elements from monotonic stack, that are less or equal to current element
        for len(stack) > 0 && nums[stack[len(stack) - 1]] <= nums[i] {
            if i - stack[len(stack) - 1] + 1 >= 3 {
                res++
            }
            stack = stack[:len(stack) - 1]
        }
        // add +1 if distance is atleast 3
        if len(stack) > 0 && i - stack[len(stack) - 1] + 1 >= 3 {
            res++
        }
        stack = append(stack, i)
    }
    return int64(res)
}

func bowlSubarrays1(nums []int) int64 {
    res, n := 0, len(nums)
    stack := make([]int, 0, n)
    for i, v := range nums {
        flag := false
        for len(stack) > 0 && nums[stack[len(stack) - 1]] < v {
            if i - stack[len(stack) - 1] >= 2 {
                res++
            }
            stack = stack[:len(stack) - 1]
            flag = true
        }
        if len(stack) > 0 && nums[stack[0]] >= v && i - stack[0] >= 2 && flag {
            res++
        }
        stack = append(stack, i)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [2,5,3,1,4]
    // Output: 2
    // Explanation:
    // The bowl subarrays are [3, 1, 4] and [5, 3, 1, 4].
    // [3, 1, 4] is a bowl because min(3, 4) = 3 > max(1) = 1.
    // [5, 3, 1, 4] is a bowl because min(5, 4) = 4 > max(3, 1) = 3.
    fmt.Println(bowlSubarrays([]int{2,5,3,1,4})) // 2
    // Example 2:
    // Input: nums = [5,1,2,3,4]
    // Output: 3
    // Explanation:
    // The bowl subarrays are [5, 1, 2], [5, 1, 2, 3] and [5, 1, 2, 3, 4].
    fmt.Println(bowlSubarrays([]int{5,1,2,3,4})) // 3
    // Example 3:
    // Input: nums = [1000000000,999999999,999999998]
    // Output: 0
    // Explanation:
    // No subarray is a bowl.
    fmt.Println(bowlSubarrays([]int{1000000000,999999999,999999998})) // 0

    fmt.Println(bowlSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(bowlSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(bowlSubarrays1([]int{2,5,3,1,4})) // 2
    fmt.Println(bowlSubarrays1([]int{5,1,2,3,4})) // 3
    fmt.Println(bowlSubarrays1([]int{1000000000,999999999,999999998})) // 0
    fmt.Println(bowlSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(bowlSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // 0
}