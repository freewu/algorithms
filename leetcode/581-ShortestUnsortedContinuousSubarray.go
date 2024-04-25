package main

// 581. Shortest Unsorted Continuous Subarray
// Given an integer array nums, you need to find one continuous subarray 
// such that if you only sort this subarray in non-decreasing order, 
// then the whole array will be sorted in non-decreasing order.

// Return the shortest such subarray and output its length.

// Example 1:
// Input: nums = [2,6,4,8,10,9,15]
// Output: 5
// Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 0

// Example 3:
// Input: nums = [1]
// Output: 0
 
// Constraints:
//     1 <= nums.length <= 10^4
//     -10^5 <= nums[i] <= 10^5
 
// Follow up: Can you solve it in O(n) time complexity?

import "fmt"
import "sort"

func findUnsortedSubarray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    tmpNums := make([]int, len(nums))
    copy(tmpNums, nums)
    sort.Ints(tmpNums)
    
    min, max := -1, -1
    for i := 0; i < len(nums); i++ {
        if nums[i] != tmpNums[i] {
            min = i
            break
        }
    }
    for i := len(nums) - 1; i > -1; i-- {
        if nums[i] != tmpNums[i] {
            max = i
            break
        }
    }
    if min == -1 && max == - 1 {
        return 0
    }
    return max - min + 1
}

// 双指针
func findUnsortedSubarray1(nums []int) int {
    n := len(nums)
    minNum, maxNum := 1 << 32 - 1, -1 << 32 - 1
    left, right := -1, -1
    for i := 0; i < n; i++ {
        if maxNum > nums[i] { 
            right = i   //如果你找到了右边界，那right就不会再动了
        } else {
            maxNum = nums[i]
        }
    }
    if right == -1 {
        return 0
    }
    for i := n - 1; i >= 0; i-- {
        if minNum < nums[i] {
            left = i
        } else {
            minNum = nums[i]
        }
    }
    return right - left + 1
}

func main() {
    // Example 1:
    // Input: nums = [2,6,4,8,10,9,15]
    // Output: 5
    // Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
    fmt.Println(findUnsortedSubarray([]int{2,6,4,8,10,9,15})) // 5
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 0
    fmt.Println(findUnsortedSubarray([]int{1,2,3,4})) // 0
    // Example 3:
    // Input: nums = [1]
    // Output: 0
    fmt.Println(findUnsortedSubarray([]int{1})) // 0

    fmt.Println(findUnsortedSubarray1([]int{2,6,4,8,10,9,15})) // 5
    fmt.Println(findUnsortedSubarray1([]int{1,2,3,4})) // 0
    fmt.Println(findUnsortedSubarray1([]int{1})) // 0
}