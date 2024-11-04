package main

// 1748. Sum of Unique Elements
// You are given an integer array nums. 
// The unique elements of an array are the elements that appear exactly once in the array.

// Return the sum of all the unique elements of nums.

// Example 1:
// Input: nums = [1,2,3,2]
// Output: 4
// Explanation: The unique elements are [1,3], and the sum is 4.

// Example 2:
// Input: nums = [1,1,1,1,1]
// Output: 0
// Explanation: There are no unique elements, and the sum is 0.

// Example 3:
// Input: nums = [1,2,3,4,5]
// Output: 15
// Explanation: The unique elements are [1,2,3,4,5], and the sum is 15.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func sumOfUnique(nums []int) int {
    res, mp := 0, make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    for k, v := range mp {
        if v == 1 { // unique element
            res += k
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,2]
    // Output: 4
    // Explanation: The unique elements are [1,3], and the sum is 4.
    fmt.Println(sumOfUnique([]int{1,2,3,2})) // 4
    // Example 2:
    // Input: nums = [1,1,1,1,1]
    // Output: 0
    // Explanation: There are no unique elements, and the sum is 0.
    fmt.Println(sumOfUnique([]int{1,1,1,1,1})) // 0
    // Example 3:
    // Input: nums = [1,2,3,4,5]
    // Output: 15
    // Explanation: The unique elements are [1,2,3,4,5], and the sum is 15.
    fmt.Println(sumOfUnique([]int{1,2,3,4,5})) // 15
}