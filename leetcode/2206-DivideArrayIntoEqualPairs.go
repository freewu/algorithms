package main

// 2206. Divide Array Into Equal Pairs
// You are given an integer array nums consisting of 2 * n integers.

// You need to divide nums into n pairs such that:
//     Each element belongs to exactly one pair.
//     The elements present in a pair are equal.

// Return true if nums can be divided into n pairs, otherwise return false.

// Example 1:
// Input: nums = [3,2,3,2,2,2]
// Output: true
// Explanation: 
// There are 6 elements in nums, so they should be divided into 6 / 2 = 3 pairs.
// If nums is divided into the pairs (2, 2), (3, 3), and (2, 2), it will satisfy all the conditions.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: false
// Explanation: 
// There is no way to divide nums into 4 / 2 = 2 pairs such that the pairs satisfy every condition.

// Constraints:
//     nums.length == 2 * n
//     1 <= n <= 500
//     1 <= nums[i] <= 500

import "fmt"
import "sort"

func divideArray(nums []int) bool {
    if len(nums) % 2 != 0 { return false }
    count := make(map[int]int)
    for _, v := range nums {
        count[v]++
    }
    for _, v := range count {
        if v % 2 != 0 { return false }
    }
    return true
}

func divideArray1(nums []int) bool {
    if len(nums) % 2 != 0 { return false }
    sort.Ints(nums)
    sum := 0
    for i, v := range nums {
        if i % 2 == 0 {
            sum += v
        } else {
            sum -= v
        }
    }
    return sum == 0
}

func main() {
    // Example 1:
    // Input: nums = [3,2,3,2,2,2]
    // Output: true
    // Explanation: 
    // There are 6 elements in nums, so they should be divided into 6 / 2 = 3 pairs.
    // If nums is divided into the pairs (2, 2), (3, 3), and (2, 2), it will satisfy all the conditions.
    fmt.Println(divideArray([]int{3,2,3,2,2,2})) // true
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: false
    // Explanation: 
    // There is no way to divide nums into 4 / 2 = 2 pairs such that the pairs satisfy every condition.
    fmt.Println(divideArray([]int{1,2,3,4})) // false

    fmt.Println(divideArray1([]int{3,2,3,2,2,2})) // true
    fmt.Println(divideArray1([]int{1,2,3,4})) // false
}