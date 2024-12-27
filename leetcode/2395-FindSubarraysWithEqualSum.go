package main

// 2395. Find Subarrays With Equal Sum
// Given a 0-indexed integer array nums, determine whether there exist two subarrays of length 2 with equal sum. 
// Note that the two subarrays must begin at different indices.

// Return true if these subarrays exist, and false otherwise.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [4,2,4]
// Output: true
// Explanation: The subarrays with elements [4,2] and [2,4] have the same sum of 6.

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: false
// Explanation: No two subarrays of size 2 have the same sum.

// Example 3:
// Input: nums = [0,0,0]
// Output: true
// Explanation: The subarrays [nums[0],nums[1]] and [nums[1],nums[2]] have the same sum of 0. 
// Note that even though the subarrays have the same content, the two subarrays are considered different because they are in different positions in the original array.

// Constraints:
//     2 <= nums.length <= 1000
//     -10^9 <= nums[i] <= 10^9

import "fmt"

func findSubarrays(nums []int) bool {
    mp := make(map[int]int)
    for i := 0; i < len(nums) - 1; i++ {
        sum := nums[i] + nums[i + 1]
        mp[sum]++
        if mp[sum] == 2 {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [4,2,4]
    // Output: true
    // Explanation: The subarrays with elements [4,2] and [2,4] have the same sum of 6.
    fmt.Println(findSubarrays([]int{4,2,4})) // true
    // Example 2:
    // Input: nums = [1,2,3,4,5]
    // Output: false
    // Explanation: No two subarrays of size 2 have the same sum.
    fmt.Println(findSubarrays([]int{1,2,3,4,5})) // false
    // Example 3:
    // Input: nums = [0,0,0]
    // Output: true
    // Explanation: The subarrays [nums[0],nums[1]] and [nums[1],nums[2]] have the same sum of 0. 
    // Note that even though the subarrays have the same content, the two subarrays are considered different because they are in different positions in the original array.
    fmt.Println(findSubarrays([]int{0,0,0})) // true

    fmt.Println(findSubarrays([]int{1,2,3,4,5,6,7,8,9})) // false
    fmt.Println(findSubarrays([]int{5,5,5,5,5,5,5,5,5})) // true
}