package main

// 3866. First Unique Even Element
// You are given an integer array nums.

// Return an integer denoting the first even integer (earliest by array index) that appears exactly once in nums. 
// If no such integer exists, return -1.

// An integer x is considered even if it is divisible by 2.

// Example 1:
// Input: nums = [3,4,2,5,4,6]
// Output: 2
// Explanation:
// Both 2 and 6 are even and they appear exactly once. Since 2 occurs first in the array, the answer is 2.

// Example 2:
// Input: nums = [4,4]
// Output: -1
// Explanation:
// No even integer appears exactly once, so return -1.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100

import "fmt"

func firstUniqueEven(nums []int) int {
    mp := make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    for _, v := range nums {
        if mp[v] == 1 && v % 2 == 0 {
            return v
        }
    }
    return -1
}

func main() {
// Example 1:
// Input: nums = [3,4,2,5,4,6]
// Output: 2
// Explanation:
// Both 2 and 6 are even and they appear exactly once. Since 2 occurs first in the array, the answer is 2.
fmt.Println(firstUniqueEven([]int{3,4,2,5,4,6})) // 2
// Example 2:
// Input: nums = [4,4]
// Output: -1
// Explanation:
// No even integer appears exactly once, so return -1.  
fmt.Println(firstUniqueEven([]int{4,4})) // -1

fmt.Println(firstUniqueEven([]int{1,2,3,4,5,6,7,8,9})) // 2
fmt.Println(firstUniqueEven([]int{9,8,7,6,5,4,3,2,1})) // 8
}