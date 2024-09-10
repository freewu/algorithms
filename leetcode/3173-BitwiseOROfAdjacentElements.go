package main

// 3173. Bitwise OR of Adjacent Elements
// Given an array nums of length n, return an array answer of length n - 1 
// such that answer[i] = nums[i] | nums[i + 1] where | is the bitwise OR operation.

// Example 1:
// Input: nums = [1,3,7,15]
// Output: [3,7,15]

// Example 2:
// Input: nums = [8,4,2]
// Output: [12,6]

// Example 3:
// Input: nums = [5,4,9,11]
// Output: [5,13,11] 

// Constraints:
//     2 <= nums.length <= 100
//     0 <= nums[i] <= 100

import "fmt"

func orArray(nums []int) []int {
    n := len(nums)
    res := make([]int, n - 1)
    for i := 0; i < n - 1; i++ {
        res[i] = nums[i] | nums[i + 1] // answer[i] = nums[i] | nums[i + 1] 
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,7,15]
    // Output: [3,7,15]
    fmt.Println(orArray([]int{1,3,7,15})) // [3,7,15]
    // Example 2:
    // Input: nums = [8,4,2]
    // Output: [12,6]
    fmt.Println(orArray([]int{8,4,2})) // [12,6]
    // Example 3:
    // Input: nums = [5,4,9,11]
    // Output: [5,13,11]
    fmt.Println(orArray([]int{5,4,9,11})) // [5,13,11]
}