package main

// 3755. Find Maximum Balanced XOR Subarray Length
// Given an integer array nums, 
// return the length of the longest subarray that has a bitwise XOR of zero and contains an equal number of even and odd numbers. 
// If no such subarray exists, return 0.

// Example 1:
// Input: nums = [3,1,3,2,0]
// Output: 4
// Explanation:
// The subarray [1, 3, 2, 0] has bitwise XOR 1 XOR 3 XOR 2 XOR 0 = 0 and contains 2 even and 2 odd numbers.

// Example 2:
// Input: nums = [3,2,8,5,4,14,9,15]
// Output: 8
// Explanation:
// The whole array has bitwise XOR 0 and contains 4 even and 4 odd numbers.

// Example 3:
// Input: nums = [0]
// Output: 0
// Explanation:
// No non-empty subarray satisfies both conditions.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"

func maxBalancedSubarray(nums []int) int {
    mp := map[string]int{"0!0": -1}
    res, xor, eveodd := 0, 0, 0
    for i := 0; i < len(nums); i++ {
        xor ^= nums[i]
        if (nums[i] & 1) == 1 {
            eveodd++
        } else {
            eveodd--
        }
        curr := fmt.Sprintf("%d!%d", xor, eveodd)
        if index, ok := mp[curr]; ok {
            res = max(res, i - index)
        } else {
            mp[curr] = i
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,3,2,0]
    // Output: 4
    // Explanation:
    // The subarray [1, 3, 2, 0] has bitwise XOR 1 XOR 3 XOR 2 XOR 0 = 0 and contains 2 even and 2 odd numbers.
    fmt.Println(maxBalancedSubarray([]int{3,1,3,2,0})) // 4
    // Example 2:
    // Input: nums = [3,2,8,5,4,14,9,15]
    // Output: 8
    // Explanation:
    // The whole array has bitwise XOR 0 and contains 4 even and 4 odd numbers.
    fmt.Println(maxBalancedSubarray([]int{3,2,8,5,4,14,9,15})) // 8 
    // Example 3:
    // Input: nums = [0]
    // Output: 0
    // Explanation:
    // No non-empty subarray satisfies both conditions.
    fmt.Println(maxBalancedSubarray([]int{0})) // 0 

    fmt.Println(maxBalancedSubarray([]int{1,2,3,4,5,6,7,8,9})) // 8
    fmt.Println(maxBalancedSubarray([]int{9,8,7,6,5,4,3,2,1})) // 8
}