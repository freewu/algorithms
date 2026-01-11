package main

// 3804. Number of Centered Subarrays
// You are given an integer array nums.

// A subarray of nums is called centered if the sum of its elements is equal to at least one element within that same subarray.

// Return the number of centered subarrays of nums.

// Example 1:
// Input: nums = [-1,1,0]
// Output: 5
// Explanation:
// All single-element subarrays ([-1], [1], [0]) are centered.
// The subarray [1, 0] has a sum of 1, which is present in the subarray.
// The subarray [-1, 1, 0] has a sum of 0, which is present in the subarray.
// Thus, the answer is 5.

// Example 2:
// Input: nums = [2,-3]
// Output: 2
// Explanation:
// Only single-element subarrays ([2], [-3]) are centered.

// Constraints:
//     1 <= nums.length <= 500
//     -10^5 <= nums[i] <= 10^5

import "fmt"

func centeredSubarrays(nums []int) int {
    res, n := 0, len(nums)
    set := make(map[int]bool)
    for start := range n {
        clear(set)
        sum := 0
        for i := start; i < n; i++ {
            sum += nums[i]
            set[nums[i]] = true
            if _, ok := set[sum]; ok {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [-1,1,0]
    // Output: 5
    // Explanation:
    // All single-element subarrays ([-1], [1], [0]) are centered.
    // The subarray [1, 0] has a sum of 1, which is present in the subarray.
    // The subarray [-1, 1, 0] has a sum of 0, which is present in the subarray.
    // Thus, the answer is 5.
    fmt.Println(centeredSubarrays([]int{-1,1,0})) // 5
    // Example 2:
    // Input: nums = [2,-3]
    // Output: 2
    // Explanation:
    // Only single-element subarrays ([2], [-3]) are centered.
    fmt.Println(centeredSubarrays([]int{2,-3})) // 2

    fmt.Println(centeredSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(centeredSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 9
}