package main

// 665. Non-decreasing Array
// Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.
// We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).

// Example 1:
// Input: nums = [4,2,3]
// Output: true
// Explanation: You could modify the first 4 to 1 to get a non-decreasing array.

// Example 2:
// Input: nums = [4,2,1]
// Output: false
// Explanation: You cannot get a non-decreasing array by modifying at most one element.
 
// Constraints:
//     n == nums.length
//     1 <= n <= 10^4
//     -10^5 <= nums[i] <= 10^5

import "fmt"

func checkPossibility(nums []int) bool {
    count, index := 0, 0
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] { // 出现前面更大情况
            index = i
            count++
        }
    }
    if count == 0 { // 没有出现前面更大情况直接返回 true
        return true 
    }
    if count == 1 { // 只有前面更大的情况中出现1次
        if index == 1 || index == len(nums) - 1 { // 出现在第二位或倒数第二位置直接返回 true
            return true 
        }
        if nums[index +1] >= nums[index-1] || nums[index] >= nums[index-2] { 
            return true 
        }
    }
    return false
}

func checkPossibility1(nums []int) bool {
    count := 0
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            count++
            if count > 1 {
                return false
            }
            if i > 0 && nums[i+1] < nums[i-1] {
                nums[i+1] = nums[i]
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: nums = [4,2,3]
    // Output: true
    // Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
    fmt.Println(checkPossibility([]int{4,2,3})) // true
    // Example 2:
    // Input: nums = [4,2,1]
    // Output: false
    // Explanation: You cannot get a non-decreasing array by modifying at most one element.
    fmt.Println(checkPossibility([]int{4,2,1})) // false

    fmt.Println(checkPossibility1([]int{4,2,3})) // true
    fmt.Println(checkPossibility1([]int{4,2,1})) // false
}