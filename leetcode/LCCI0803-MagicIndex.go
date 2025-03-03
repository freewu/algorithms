package main

// 面试题 08.03. Magic Index LCCI
// A magic index in an array A[0...n-1] is defined to be an index such that A[i] = i. 
// Given a sorted array of integers, write a method to find a magic index, if one exists, in array A. 
// If not, return -1. If there are more than one magic index, return the smallest one.

// Example1:
// Input: nums = [0, 2, 3, 4, 5]
// Output: 0

// Example2:
// Input: nums = [1, 1, 1]
// Output: 1

// Note:
//     1 <= nums.length <= 1000000
//     This problem is the follow-up of the original problem in the book, i.e. the values are not distinct.

import "fmt"

func findMagicIndex(nums []int) int {
    for i, v := range nums {
        if i == v { return i }
    }
    return -1
}

func main() {
    // Example1:
    // Input: nums = [0, 2, 3, 4, 5]
    // Output: 0
    fmt.Println(findMagicIndex([]int{0, 2, 3, 4, 5})) // 0
    // Example2:
    // Input: nums = [1, 1, 1]
    // Output: 1
    fmt.Println(findMagicIndex([]int{1, 1, 1})) // 1

    fmt.Println(findMagicIndex([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(findMagicIndex([]int{9,8,7,6,5,4,3,2,1})) // -1
}