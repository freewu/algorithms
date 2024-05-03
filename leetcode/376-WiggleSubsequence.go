package main

// 376. Wiggle Subsequence
// A wiggle sequence is a sequence where the differences between successive numbers strictly alternate between positive and negative. 
// The first difference (if one exists) may be either positive or negative. 
// A sequence with one element and a sequence with two non-equal elements are trivially wiggle sequences.
//     For example, [1, 7, 4, 9, 2, 5] is a wiggle sequence because the differences (6, -3, 5, -7, 3) alternate between positive and negative.
//     In contrast, [1, 4, 7, 2, 5] and [1, 7, 4, 5, 5] are not wiggle sequences. 
//     The first is not because its first two differences are positive, and the second is not because its last difference is zero.

// A subsequence is obtained by deleting some elements (possibly zero) from the original sequence, leaving the remaining elements in their original order.
// Given an integer array nums, return the length of the longest wiggle subsequence of nums.

// Example 1:
// Input: nums = [1,7,4,9,2,5]
// Output: 6
// Explanation: The entire sequence is a wiggle sequence with differences (6, -3, 5, -7, 3).

// Example 2:
// Input: nums = [1,17,5,10,13,15,10,5,16,8]
// Output: 7
// Explanation: There are several subsequences that achieve this length.
// One is [1, 17, 10, 13, 10, 16, 8] with differences (16, -7, 3, -3, 6, -8).

// Example 3:
// Input: nums = [1,2,3,4,5,6,7,8,9]
// Output: 2
 
// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] <= 1000
 
// Follow up: Could you solve this in O(n) time?

import "fmt"

func wiggleMaxLength(nums []int) int {
    valley, peek := 1, 1 // 记录当前序列的上升和下降的趋势
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[i-1] {
            valley = peek + 1
        } else if nums[i] > nums[i-1] {
            peek = valley + 1
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(valley, peek)
}

func wiggleMaxLength1(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    res := 1
    prevDiff := nums[1] - nums[0]
    if prevDiff != 0 {
        res = 2
    }
    for i := 2; i < len(nums); i++ {
        diff := nums[i] - nums[i-1]
        if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
            res++
            prevDiff = diff
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,7,4,9,2,5]
    // Output: 6
    // Explanation: The entire sequence is a wiggle sequence with differences (6, -3, 5, -7, 3).
    fmt.Println(wiggleMaxLength([]int{1,7,4,9,2,5})) // 6
    // Example 2:
    // Input: nums = [1,17,5,10,13,15,10,5,16,8]
    // Output: 7
    // Explanation: There are several subsequences that achieve this length.
    // One is [1, 17, 10, 13, 10, 16, 8] with differences (16, -7, 3, -3, 6, -8).
    fmt.Println(wiggleMaxLength([]int{1,17,5,10,13,15,10,5,16,8})) // 7
    // Example 3:
    // Input: nums = [1,2,3,4,5,6,7,8,9]
    // Output: 2
    fmt.Println(wiggleMaxLength([]int{1,2,3,4,5,6,7,8,9})) // 2

    fmt.Println(wiggleMaxLength1([]int{1,7,4,9,2,5})) // 6
    fmt.Println(wiggleMaxLength1([]int{1,17,5,10,13,15,10,5,16,8})) // 7
    fmt.Println(wiggleMaxLength1([]int{1,2,3,4,5,6,7,8,9})) // 2
}