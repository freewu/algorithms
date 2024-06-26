package main

// 666. Path Sum IV
// If the depth of a tree is smaller than 5, then this tree can be represented by an array of three-digit integers. 
// For each integer in this array:
//     1. The hundreds digit represents the depth d of this node where 1 <= d <= 4.
//     2. The tens digit represents the position p of this node in the level it belongs to where 1 <= p <= 8. 
//        The position is the same as that in a full binary tree.
//     3. The units digit represents the value v of this node where 0 <= v <= 9.

// Given an array of ascending three-digit integers nums representing a binary tree with a depth smaller than 5, 
// return the sum of all paths from the root towards the leaves.

// It is guaranteed that the given array represents a valid connected binary tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/30/pathsum4-1-tree.jpg">
// Input: nums = [113,215,221]
// Output: 12
// Explanation: The tree that the list represents is shown.
// The path sum is (3 + 5) + (3 + 1) = 12.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/30/pathsum4-2-tree.jpg">
// Input: nums = [113,221]
// Output: 4
// Explanation: The tree that the list represents is shown. 
// The path sum is (3 + 1) = 4.

// Constraints:
//     1 <= nums.length <= 15
//     110 <= nums[i] <= 489
//     nums represents a valid binary tree with depth less than 5.

import "fmt"
import "math"

func pathSum(nums []int) int {
    l  := len(nums)
    nowLevel := int(nums[l-1] / 100)
    nowLevelWidth := int(math.Pow(2, float64(nowLevel-1)))
    nowNodeNums := make([]int, nowLevelWidth)
    now, sum := 0, 0
    for i := l - 1; i >= 0; i-- {
        now = nums[i]
        level, pos := int(now / 100), int(now/10) % 10
        if level != nowLevel {
            nowLevelWidth = nowLevelWidth / 2
            tmp := make([]int, nowLevelWidth)
            for j := 0; j < nowLevelWidth; j++ {
                tmp[j] = nowNodeNums[j*2] + nowNodeNums[j*2+1]
            }
            nowNodeNums = tmp
            nowLevel = level
        }
        if nowNodeNums[pos-1] == 0 {
            nowNodeNums[pos-1] = 1
        }
        sum += ((now % 10) * nowNodeNums[pos-1])
    }
    return sum
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/30/pathsum4-1-tree.jpg">
    // Input: nums = [113,215,221]
    // Output: 12
    // Explanation: The tree that the list represents is shown.
    // The path sum is (3 + 5) + (3 + 1) = 12.
    fmt.Println(pathSum([]int{113,215,221})) // 12
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/30/pathsum4-2-tree.jpg">
    // Input: nums = [113,221]
    // Output: 4
    // Explanation: The tree that the list represents is shown. 
    // The path sum is (3 + 1) = 4.
    fmt.Println(pathSum([]int{113,221})) // 4
}