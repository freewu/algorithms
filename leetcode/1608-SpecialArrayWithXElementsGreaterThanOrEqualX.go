package main

// 1608. Special Array With X Elements Greater Than or Equal X
// You are given an array nums of non-negative integers. 
// nums is considered special if there exists a number x such that there are exactly x numbers in nums that are greater than or equal to x.

// Notice that x does not have to be an element in nums.

// Return x if the array is special, otherwise, return -1. 
// It can be proven that if nums is special, the value for x is unique.

// Example 1:
// Input: nums = [3,5]
// Output: 2
// Explanation: There are 2 values (3 and 5) that are greater than or equal to 2.

// Example 2:
// Input: nums = [0,0]
// Output: -1
// Explanation: No numbers fit the criteria for x.
// If x = 0, there should be 0 numbers >= x, but there are 2.
// If x = 1, there should be 1 number >= x, but there are 0.
// If x = 2, there should be 2 numbers >= x, but there are 0.
// x cannot be greater since there are only 2 numbers in nums.

// Example 3:
// Input: nums = [0,4,3,0,4]
// Output: 3
// Explanation: There are 3 values that are greater than or equal to 3.

// Constraints:
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 1000

import "fmt"
import "sort"

func specialArray(nums []int) int {
    res, m := -1, map[int]int{};
    for _, v := range nums { // 统计每个字符出现次数
        m[v]++
    }
    for i := 1; i <= len(nums); i++ {
        count := 0
        for k, v := range m {
            if k >= i {
                count += v
            } 
        }
        if count == i { // 恰好有 x 个元素 大于或者等于 x
            res = i
        }
    }
    return res
}

func specialArray1(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    if nums[0] >= n {
        return n
    }
    for i := 1; i < n; i++ {
        if n - i <= nums[i] && n - i > nums[i - 1]{
            return n - i
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [3,5]
    // Output: 2
    // Explanation: There are 2 values (3 and 5) that are greater than or equal to 2.
    fmt.Println(specialArray([]int{3,5})) // 2
    // Example 2:
    // Input: nums = [0,0]
    // Output: -1
    // Explanation: No numbers fit the criteria for x.
    // If x = 0, there should be 0 numbers >= x, but there are 2.
    // If x = 1, there should be 1 number >= x, but there are 0.
    // If x = 2, there should be 2 numbers >= x, but there are 0.
    // x cannot be greater since there are only 2 numbers in nums.
    fmt.Println(specialArray([]int{0, 0})) // -1
    // Example 3:
    // Input: nums = [0,4,3,0,4]
    // Output: 3
    // Explanation: There are 3 values that are greater than or equal to 3.
    fmt.Println(specialArray([]int{0,4,3,0,4})) // 3

    fmt.Println(specialArray1([]int{3,5})) // 2
    fmt.Println(specialArray1([]int{0, 0})) // -1
    fmt.Println(specialArray1([]int{0,4,3,0,4})) // 3
}