package main

// 3818. Minimum Prefix Removal to Make Array Strictly Increasing
// You are given an integer array nums.

// You need to remove exactly one prefix (possibly empty) from nums.

// Return an integer denoting the minimum length of the removed prefix such that the remaining array is strictly increasing.

// Example 1:
// Input: nums = [1,-1,2,3,3,4,5]
// Output: 4
// Explanation:
// Removing the prefix = [1, -1, 2, 3] leaves the remaining array [3, 4, 5] which is strictly increasing.

// Example 2:
// Input: nums = [4,3,-2,-5]
// Output: 3
// Explanation:
// Removing the prefix = [4, 3, -2] leaves the remaining array [-5] which is strictly increasing.

// Example 3:
// Input: nums = [1,2,3,4]
// Output: 0
// Explanation:
// The array nums = [1, 2, 3, 4] is already strictly increasing so removing an empty prefix is sufficient.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func minimumPrefixLength(nums []int) int {
    // 倒序遍历 nums，如果发现 nums[i−1] ≥ nums[i]，那么前缀 [0,i−1] 必须删除，长度为 i。
    for i := len(nums) - 1; i > 0; i-- {
        if nums[i-1] >= nums[i] {
            return i // 移除前缀 [0, i-1]，长度为 i
        }
    }
    return 0 // 如果 nums 已经是严格递增的，返回 0。
}

func main() {
    // Example 1:
    // Input: nums = [1,-1,2,3,3,4,5]
    // Output: 4
    // Explanation:
    // Removing the prefix = [1, -1, 2, 3] leaves the remaining array [3, 4, 5] which is strictly increasing.
    fmt.Println(minimumPrefixLength([]int{1,-1,2,3,3,4,5})) // 4
    // Example 2:
    // Input: nums = [4,3,-2,-5]
    // Output: 3
    // Explanation:
    // Removing the prefix = [4, 3, -2] leaves the remaining array [-5] which is strictly increasing.
    fmt.Println(minimumPrefixLength([]int{4,3,-2,-5})) // 3
    // Example 3:
    // Input: nums = [1,2,3,4]
    // Output: 0
    // Explanation:
    // The array nums = [1, 2, 3, 4] is already strictly increasing so removing an empty prefix is sufficient.
    fmt.Println(minimumPrefixLength([]int{1,2,3,4})) // 0

    fmt.Println(minimumPrefixLength([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumPrefixLength([]int{9,8,7,6,5,4,3,2,1})) // 8
}