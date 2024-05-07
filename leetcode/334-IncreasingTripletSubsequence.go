package main

// 334. Increasing Triplet Subsequence
// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. 
// If no such indices exists, return false.

// Example 1:
// Input: nums = [1,2,3,4,5]
// Output: true
// Explanation: Any triplet where i < j < k is valid.

// Example 2:
// Input: nums = [5,4,3,2,1]
// Output: false
// Explanation: No triplet exists.

// Example 3:
// Input: nums = [2,1,5,0,4,6]
// Output: true
// Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.

// Constraints:
//     1 <= nums.length <= 5 * 10^5
//     -2^31 <= nums[i] <= 2^31 - 1
 
// Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?

import "fmt"

// func increasingTriplet(nums []int) bool {
//     for i := 2; i < len(nums); i++ {
//         if nums[i - 2] < nums[i - 1] && nums[i - 1] < nums[i] {
//             return true
//         }
//     }
//     return false
// }

func increasingTriplet(nums []int) bool {
    first, second := 1 << 32 - 1, 1 << 32 - 1 // Initialize two variables to track the smallest and second smallest elements
    for _, num := range nums { // Iterate through the elements in the input slice
        if num <= first { // Check if the current number is smaller than or equal to the first smallest
            first = num
        } else if num <= second { // Check if the current number is smaller than or equal to the second smallest
            second = num
        } else { // If neither condition is met, an increasing triplet subsequence is found
            return true
        }
    }
    return false // No increasing triplet subsequence was found
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5]
    // Output: true
    // Explanation: Any triplet where i < j < k is valid.
    fmt.Println(increasingTriplet([]int{1,2,3,4,5})) // true
    // Example 2:
    // Input: nums = [5,4,3,2,1]
    // Output: false
    // Explanation: No triplet exists.
    fmt.Println(increasingTriplet([]int{5,4,3,2,1})) // false
    // Example 3:
    // Input: nums = [2,1,5,0,4,6]
    // Output: true
    // Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
    fmt.Println(increasingTriplet([]int{2,1,5,0,4,6})) // true

    // [10,12,13]
    fmt.Println(increasingTriplet([]int{20,100,10,12,5,13})) // true
}