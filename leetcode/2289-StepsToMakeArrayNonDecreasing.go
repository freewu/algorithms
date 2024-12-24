package main

// 2289. Steps to Make Array Non-decreasing
// You are given a 0-indexed integer array nums. 
// In one step, remove all elements nums[i] where nums[i - 1] > nums[i] for all 0 < i < nums.length.

// Return the number of steps performed until nums becomes a non-decreasing array.

// Example 1:
// Input: nums = [5,3,4,4,7,3,6,11,8,5,11]
// Output: 3
// Explanation: The following are the steps performed:
// - Step 1: [5,3,4,4,7,3,6,11,8,5,11] becomes [5,4,4,7,6,11,11]
// - Step 2: [5,4,4,7,6,11,11] becomes [5,4,7,11,11]
// - Step 3: [5,4,7,11,11] becomes [5,7,11,11]
// [5,7,11,11] is a non-decreasing array. Therefore, we return 3.

// Example 2:
// Input: nums = [4,5,7,7,13]
// Output: 0
// Explanation: nums is already a non-decreasing array. Therefore, we return 0.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func totalSteps(nums []int) int {
    res, n := 0, len(nums)
    stack, dp := []int{}, make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        for len(stack) > 0 && nums[i] > nums[stack[len(stack) - 1]] {
            dp[i] = max(dp[i] + 1, dp[stack[len(stack) - 1]])
            stack = stack[:len(stack) - 1]
            res = max(res, dp[i])
        }
        stack = append(stack, i)
    }
    return res
}

func totalSteps1(nums []int) int {
    stack := [][2]int{{nums[0], 0}}
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
        t := 0
        for len(stack) > 0 {
            top := stack[len(stack) - 1]
            if top[0] > nums[i] {
                if t == 0 {
                    t = 1
                }
                break
            } else {
                stack = stack[:len(stack) - 1]
                t = max(t, top[1] + 1)
            }
        }
        if len(stack) == 0 {
            t = 0
        }
        res = max(res, t)
        stack = append(stack, [2]int{ nums[i], t })
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,3,4,4,7,3,6,11,8,5,11]
    // Output: 3
    // Explanation: The following are the steps performed:
    // - Step 1: [5,3,4,4,7,3,6,11,8,5,11] becomes [5,4,4,7,6,11,11]
    // - Step 2: [5,4,4,7,6,11,11] becomes [5,4,7,11,11]
    // - Step 3: [5,4,7,11,11] becomes [5,7,11,11]
    // [5,7,11,11] is a non-decreasing array. Therefore, we return 3.
    fmt.Println(totalSteps([]int{5,3,4,4,7,3,6,11,8,5,11})) // 3
    // Example 2:
    // Input: nums = [4,5,7,7,13]
    // Output: 0
    // Explanation: nums is already a non-decreasing array. Therefore, we return 0.
    fmt.Println(totalSteps([]int{4,5,7,7,13})) // 0

    fmt.Println(totalSteps1([]int{5,3,4,4,7,3,6,11,8,5,11})) // 3
    fmt.Println(totalSteps1([]int{4,5,7,7,13})) // 0
}