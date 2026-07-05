package main

// 3978. Unique Middle Element
// You are given an integer array nums of odd length n.

// Return true if the middle element of nums appears exactly once in the array. Otherwise return false.

// Example 1:
// Input: nums = [1,2,3]
// Output: true
// Explanation:
// The middle element of nums is 2, which appears exactly once.
// Thus, the answer is true.

// Example 2:
// Input: nums = [1,2,2]
// Output: false
// Explanation:
// The middle element of nums is 2, which appears twice.
// Thus, the answer is false.

// Constraints:
//     1 <= n == nums.length <= 100
//     n is odd.
//     1 <= nums[i] <= 100

import "fmt"

func isMiddleElementUnique(nums []int) bool {
    n := len(nums)
    mid := n / 2
    count := map[int]int{}
    for _, v := range nums {
        count[v]++
    }
    return count[nums[mid]] == 1    
}

func isMiddleElementUnique1(nums []int) bool {
    count, t := 0, nums[len(nums) >> 1]
    for _, v := range nums {
        if v == t {
            count++
            if count > 1 { return false } // 直接跳出
        }
    }
    return count == 1
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: true
    // Explanation:
    // The middle element of nums is 2, which appears exactly once.
    // Thus, the answer is true.
    fmt.Println(isMiddleElementUnique([]int{1,2,3})) // true 
    // Example 2:
    // Input: nums = [1,2,2]
    // Output: false
    // Explanation:
    // The middle element of nums is 2, which appears twice.
    // Thus, the answer is false.
    fmt.Println(isMiddleElementUnique([]int{1,2,2})) // false

    fmt.Println(isMiddleElementUnique([]int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(isMiddleElementUnique([]int{9,8,7,6,5,4,3,2,1})) // true

    fmt.Println(isMiddleElementUnique1([]int{1,2,3})) // true 
    fmt.Println(isMiddleElementUnique1([]int{1,2,2})) // false
    fmt.Println(isMiddleElementUnique1([]int{1,2,3,4,5,6,7,8,9})) // true
    fmt.Println(isMiddleElementUnique1([]int{9,8,7,6,5,4,3,2,1})) // true
}