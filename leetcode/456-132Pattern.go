package main

// 456. 132 Pattern
// Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] 
// such that i < j < k and nums[i] < nums[k] < nums[j].
// Return true if there is a 132 pattern in nums, otherwise, return false.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: false
// Explanation: There is no 132 pattern in the sequence.

// Example 2:
// Input: nums = [3,1,4,2]
// Output: true
// Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

// Example 3:
// Input: nums = [-1,3,2,0]
// Output: true
// Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].
 
// Constraints:
//     n == nums.length
//     1 <= n <= 2 * 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

// 单调栈
func find132pattern(nums []int) bool {
    if len(nums) < 3 {
        return false
    }
    num3, stack := -1 << 63, []int{}
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] < num3 {
            return true
        }
        for len(stack) != 0 && nums[i] > stack[len(stack)-1] {
            num3 = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, nums[i])
    }
    return false
}

func find132pattern1(nums []int) bool {
    k := -1 << 63  // 保存 第二大的数，也是出栈的数
    stack := []int{}
    for i := len(nums)-1; i >= 0; i-- {
        if nums[i] < k { // 说明此时k有值了。说明 stack里面有个最大值， 说明此事 nums[i] 是最小的
            return true
        }
        for len(stack)>0 && nums[stack[len(stack)-1]]<nums[i] {
            k = nums[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: false
    // Explanation: There is no 132 pattern in the sequence.
    fmt.Println(find132pattern([]int{1,2,3,4})) // false
    // Example 2:
    // Input: nums = [3,1,4,2]
    // Output: true
    // Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
    fmt.Println(find132pattern([]int{3,1,4,2})) // true
    // Example 3:
    // Input: nums = [-1,3,2,0]
    // Output: true
    // Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].
    fmt.Println(find132pattern([]int{-1,3,2,0})) // true

    fmt.Println(find132pattern1([]int{1,2,3,4})) // false
    fmt.Println(find132pattern1([]int{3,1,4,2})) // true
    fmt.Println(find132pattern1([]int{-1,3,2,0})) // true
}