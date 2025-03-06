package main

// 3392. Count Subarrays of Length Three With a Condition
// Given an integer array nums, return the number of subarrays of length 3 such that the sum of the first and third numbers equals exactly half of the second number.

// Example 1:
// Input: nums = [1,2,1,4,1]
// Output: 1
// Explanation:
// Only the subarray [1,4,1] contains exactly 3 elements where the sum of the first and third numbers equals half the middle number.

// Example 2:
// Input: nums = [1,1,1]
// Output: 0
// Explanation:
// [1,1,1] is the only subarray of length 3. However, its first and third numbers do not add to half the middle number.

// Constraints:
//     3 <= nums.length <= 100
//     -100 <= nums[i] <= 100

import "fmt"

func countSubarrays(nums []int) int {
    res := 0
    for i := 1; i < len(nums) - 1; i++ {
        if nums[i] == (nums[i - 1] + nums[i + 1]) * 2 { // 第一个数和第三个数的和恰好为第二个数的一半
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,4,1]
    // Output: 1
    // Explanation:
    // Only the subarray [1,4,1] contains exactly 3 elements where the sum of the first and third numbers equals half the middle number.
    fmt.Println(countSubarrays([]int{1,2,1,4,1})) // 1
    // Example 2:
    // Input: nums = [1,1,1]
    // Output: 0
    // Explanation:
    // [1,1,1] is the only subarray of length 3. However, its first and third numbers do not add to half the middle number.
    fmt.Println(countSubarrays([]int{1,1,1})) // 0

    fmt.Println(countSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(countSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 0
}