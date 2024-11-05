package main

// 1752. Check if Array Is Sorted and Rotated
// Given an array nums, return true if the array was originally sorted in non-decreasing order, 
// then rotated some number of positions (including zero). 
// Otherwise, return false.

// There may be duplicates in the original array.

// Note: An array A rotated by x positions results in an array B of the same length 
// such that A[i] == B[(i+x) % A.length], where % is the modulo operation.

// Example 1:
// Input: nums = [3,4,5,1,2]
// Output: true
// Explanation: [1,2,3,4,5] is the original sorted array.
// You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].

// Example 2:
// Input: nums = [2,1,3,4]
// Output: false
// Explanation: There is no sorted array once rotated that can make nums.

// Example 3:
// Input: nums = [1,2,3]
// Output: true
// Explanation: [1,2,3] is the original sorted array.
// You can rotate the array by x = 0 positions (i.e. no rotation) to make nums.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func check(nums []int) bool {
    rotation := 0
    for i, v := range nums{
        // Note: here we are taking mod, because we have to check last element with first 
        // element. If element is sorted rotated, then first element should always be greater  than the last element
        if v > nums[(i+1) % len(nums)] {
            rotation++
        }
        if rotation > 1 {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: nums = [3,4,5,1,2]
    // Output: true
    // Explanation: [1,2,3,4,5] is the original sorted array.
    // You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].
    fmt.Println(check([]int{3,4,5,1,2})) // true
    // Example 2:
    // Input: nums = [2,1,3,4]
    // Output: false
    // Explanation: There is no sorted array once rotated that can make nums.
    fmt.Println(check([]int{2,1,3,4})) // false
    // Example 3:
    // Input: nums = [1,2,3]
    // Output: true
    // Explanation: [1,2,3] is the original sorted array.
    // You can rotate the array by x = 0 positions (i.e. no rotation) to make nums.
    fmt.Println(check([]int{1,2,3})) // true
}