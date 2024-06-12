package main

// 477. Total Hamming Distance
// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
// Given an integer array nums, return the sum of Hamming distances between all the pairs of the integers in nums.

// Example 1:
// Input: nums = [4,14,2]
// Output: 6
// Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
// showing the four bits relevant in this case).
// The answer will be:
// HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

// Example 2:
// Input: nums = [4,14,4]
// Output: 4

// Constraints:
//     1 <= nums.length <= 10^4
//     0 <= nums[i] <= 10^9
//     The answer for the given input will fit in a 32-bit integer.

import "fmt"

func totalHammingDistance(nums []int) int {
    res, n :=0, len(nums)
    for i:= 30; i >=0; i-- {
        setBits :=0
        multiplier := 1 << i
        for j := 0; j < n; j++ {
            if nums[j] & multiplier > 0 {
                setBits++
            }
        }
        res += (setBits * ( n - setBits))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,14,2]
    // Output: 6
    // Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
    // showing the four bits relevant in this case).
    // The answer will be:
    // HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.
    fmt.Println(totalHammingDistance([]int{4,14,2})) // 6
    // Example 2:
    // Input: nums = [4,14,4]
    // Output: 4
    fmt.Println(totalHammingDistance([]int{4,14,4})) // 4
}