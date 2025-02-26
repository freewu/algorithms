package main

// 面试题 17.10. Find Majority Element LCCI
// A majority element is an element that makes up more than half of the items in an array. 
// Given a integers array, find the majority element. 
// If there is no majority element, return -1. Do this in O(N) time and O(1) space.

// Example 1:
// Input: [1,2,5,9,5,9,5,5,5]
// Output: 5

// Example 2:
// Input: [3,2]
// Output: -1

// Example 3:
// Input: [2,2,1,1,1,2,2]
// Output: 2

import "fmt"

func majorityElement(nums []int) int {
    mp, half := make(map[int]int), len(nums) / 2 + 1
    for _, v := range nums {
        mp[v]++
        if mp[v] >= half {
            return v
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: [1,2,5,9,5,9,5,5,5]
    // Output: 5
    fmt.Println(majorityElement([]int{1,2,5,9,5,9,5,5,5})) // 5
    // Example 2:
    // Input: [3,2]
    // Output: -1
    fmt.Println(majorityElement([]int{3,2})) // -1
    // Example 3:
    // Input: [2,2,1,1,1,2,2]
    // Output: 2
    fmt.Println(majorityElement([]int{2,2,1,1,1,2,2})) // 2
}