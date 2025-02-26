package main

// 面试题 17.04. Missing Number LCCI
// An array contains all the integers from 0 to n, except for one number which is missing.  
// Write code to find the missing integer. Can you do it in O(n) time?

// Note: This problem is slightly different from the original one the book.

// Example 1:
// Input: [3,0,1]
// Output: 2

// Example 2:
// Input: [9,6,4,2,3,5,7,0,1]
// Output: 8

import "fmt"

func missingNumber(nums []int) int {
    n := len(nums)
    mp := make([]bool, n + 1)
    for _, v := range nums {
        mp[v] = true
    }
    for k, v := range mp {
        if !v { return k }
    }
    return n + 1
}

func missingNumber1(nums []int) int {
    sum := (1 + len(nums)) * len(nums) / 2
    for _, v := range nums {
        sum -= v
    }
    return sum
}

func main() {
    // Example 1:
    // Input: [3,0,1]
    // Output: 2
    fmt.Println(missingNumber([]int{3,0,1})) // 2
    // Example 2:
    // Input: [9,6,4,2,3,5,7,0,1]
    // Output: 8
    fmt.Println(missingNumber([]int{9,6,4,2,3,5,7,0,1})) // 8

    fmt.Println(missingNumber([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(missingNumber([]int{0,1,2,3,4,5,6,7,8})) // 9

    fmt.Println(missingNumber1([]int{3,0,1})) // 2
    fmt.Println(missingNumber1([]int{9,6,4,2,3,5,7,0,1})) // 8
    fmt.Println(missingNumber1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(missingNumber1([]int{0,1,2,3,4,5,6,7,8})) // 9
}