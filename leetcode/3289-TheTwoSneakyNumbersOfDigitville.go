package main

// 3289. The Two Sneaky Numbers of Digitville
// In the town of Digitville, there was a list of numbers called nums containing integers from 0 to n - 1. 
// Each number was supposed to appear exactly once in the list, however, two mischievous numbers sneaked in an additional time, making the list longer than usual.

// As the town detective, your task is to find these two sneaky numbers. 
// Return an array of size two containing the two numbers (in any order), so peace can return to Digitville.

// Example 1:
// Input: nums = [0,1,1,0]
// Output: [0,1]
// Explanation:
// The numbers 0 and 1 each appear twice in the array.

// Example 2:
// Input: nums = [0,3,2,1,3,2]
// Output: [2,3]
// Explanation:
// The numbers 2 and 3 each appear twice in the array.

// Example 3:
// Input: nums = [7,1,5,4,3,4,6,0,9,5,8,2]
// Output: [4,5]
// Explanation:
// The numbers 4 and 5 each appear twice in the array.

// Constraints:
//     2 <= n <= 100
//     nums.length == n + 2
//     0 <= nums[i] < n
//     The input is generated such that nums contains exactly two repeated elements.

import "fmt"

func getSneakyNumbers(nums []int) []int {
    arr := make([]int, len(nums) - 2)
    for _, v := range nums {
        arr[v]++
    }
    res := []int{}
    for i, v := range arr {
        if v == 2 { 
            res = append(res, i)
        }
    }
    return res
}

func getSneakyNumbers1(nums []int) []int {
    arr := make([]int, len(nums) - 2)
    res := []int{}
    for _, v := range nums {
        arr[v]++
        if arr[v] == 2 { 
            res = append(res, v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,1,0]
    // Output: [0,1]
    // Explanation:
    // The numbers 0 and 1 each appear twice in the array.
    fmt.Println(getSneakyNumbers([]int{0,1,1,0})) // [0,1]
    // Example 2:
    // Input: nums = [0,3,2,1,3,2]
    // Output: [2,3]
    // Explanation:
    // The numbers 2 and 3 each appear twice in the array.
    fmt.Println(getSneakyNumbers([]int{0,3,2,1,3,2})) // [2,3]
    // Example 3:
    // Input: nums = [7,1,5,4,3,4,6,0,9,5,8,2]
    // Output: [4,5]
    // Explanation:
    // The numbers 4 and 5 each appear twice in the array.
    fmt.Println(getSneakyNumbers([]int{7,1,5,4,3,4,6,0,9,5,8,2})) // [4,5]

    fmt.Println(getSneakyNumbers([]int{0,1,2,3,4,5,6,7,8,9,1,2})) // [1,2]
    fmt.Println(getSneakyNumbers([]int{0,9,8,7,6,5,4,3,2,1,8,9})) // [8,9]

    fmt.Println(getSneakyNumbers1([]int{0,1,1,0})) // [0,1]
    fmt.Println(getSneakyNumbers1([]int{0,3,2,1,3,2})) // [2,3]
    fmt.Println(getSneakyNumbers1([]int{7,1,5,4,3,4,6,0,9,5,8,2})) // [4,5]
    fmt.Println(getSneakyNumbers1([]int{0,1,2,3,4,5,6,7,8,9,1,2})) // [1,2]
    fmt.Println(getSneakyNumbers1([]int{0,9,8,7,6,5,4,3,2,1,8,9})) // [8,9]
}