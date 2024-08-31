package main

// 962. Maximum Width Ramp
// A ramp in an integer array nums is a pair (i, j) for which i < j and nums[i] <= nums[j]. 
// The width of such a ramp is j - i.

// Given an integer array nums, return the maximum width of a ramp in nums. 
// If there is no ramp in nums, return 0.

// Example 1:
// Input: nums = [6,0,8,2,1,5]
// Output: 4
// Explanation: The maximum width ramp is achieved at (i, j) = (1, 5): nums[1] = 0 and nums[5] = 5.

// Example 2:
// Input: nums = [9,8,1,0,1,9,4,0,4,1]
// Output: 7
// Explanation: The maximum width ramp is achieved at (i, j) = (2, 9): nums[2] = 1 and nums[9] = 1.

// Constraints:
//     2 <= nums.length <= 5 * 10^4
//     0 <= nums[i] <= 5 * 10^4

import "fmt"

// stack
func maxWidthRamp(nums []int) int {
    stack := []int{}
    for i,v := range nums {
        if len(stack) == 0 || v < nums[stack[len(stack) - 1]] {
            stack = append(stack, i)
        }
    }
    res, j := 0, len(nums) - 1
    for ; len(stack) > 0 && j >= stack[len(stack) - 1]; j-- {
        for len(stack) > 0 && nums[j] >= nums[stack[len(stack) - 1]] {
            cur := j - stack[len(stack) - 1] // pop 
            stack = stack[:len(stack) - 1]
            if cur > res {
                res = cur
            }
        }
    }
    return res
}

func maxWidthRamp1(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    res, stack, top := 0, make([]int, n), 0
    for i := 0; i < n; i++ {
        if top > 0 && nums[stack[top - 1]] <= nums[i] { continue }
        stack[top] = i
        top++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        for top > 0 && nums[i] >= nums[stack[top - 1]] {
            res = max(res, i - stack[top - 1])
            top--
        }
        if top == 0 { break }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [6,0,8,2,1,5]
    // Output: 4
    // Explanation: The maximum width ramp is achieved at (i, j) = (1, 5): nums[1] = 0 and nums[5] = 5.
    fmt.Println(maxWidthRamp([]int{6,0,8,2,1,5})) // 4
    // Example 2:
    // Input: nums = [9,8,1,0,1,9,4,0,4,1]
    // Output: 7
    // Explanation: The maximum width ramp is achieved at (i, j) = (2, 9): nums[2] = 1 and nums[9] = 1.
    fmt.Println(maxWidthRamp([]int{9,8,1,0,1,9,4,0,4,1})) // 7

    fmt.Println(maxWidthRamp1([]int{6,0,8,2,1,5})) // 4
    fmt.Println(maxWidthRamp1([]int{9,8,1,0,1,9,4,0,4,1})) // 7
}