package main

// 3688. Bitwise OR of Even Numbers in an Array
// You are given an integer array nums.

// Return the bitwise OR of all even numbers in the array.

// If there are no even numbers in nums, return 0.

// Example 1:
// Input: nums = [1,2,3,4,5,6]
// Output: 6
// Explanation:
// The even numbers are 2, 4, and 6. Their bitwise OR equals 6.

// Example 2:
// Input: nums = [7,9,11]
// Output: 0
// Explanation:
// There are no even numbers, so the result is 0.

// Example 3:
// Input: nums = [1,8,16]
// Output: 24
// Explanation:
// The even numbers are 8 and 16. Their bitwise OR equals 24.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func evenNumberBitwiseORs(nums []int) int {
    if len(nums) == 0 { return 0 }
    res := 0
    for _, v := range nums {
        if v % 2 == 0 { // even numbers
            res |= v // bitwise OR
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5,6]
    // Output: 6
    // Explanation:
    // The even numbers are 2, 4, and 6. Their bitwise OR equals 6.
    fmt.Println(evenNumberBitwiseORs([]int{1,2,3,4,5,6})) // 6
    // Example 2:
    // Input: nums = [7,9,11]
    // Output: 0
    // Explanation:
    // There are no even numbers, so the result is 0.
    fmt.Println(evenNumberBitwiseORs([]int{7,9,11})) // 0
    // Example 3:
    // Input: nums = [1,8,16]
    // Output: 24
    // Explanation:
    // The even numbers are 8 and 16. Their bitwise OR equals 24.
    fmt.Println(evenNumberBitwiseORs([]int{1,8,16})) // 24

    fmt.Println(evenNumberBitwiseORs([]int{1,2,3,4,5,6,7,8,9})) // 14
    fmt.Println(evenNumberBitwiseORs([]int{9,8,7,6,5,4,3,2,1})) // 14
}