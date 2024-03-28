package main

// 704. Binary Search
// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. 
// If target exists, then return its index. Otherwise, return -1.
// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Example 2:
// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1
 
// Constraints:
//     1 <= nums.length <= 10^4
//     -10^4 < nums[i], target < 10^4
//     All the integers in nums are unique.
//     nums is sorted in ascending order.

import "fmt"

func search(nums []int, target int) int {
    l, r, mid := 0, len(nums) - 1, 0
    for l <= r {
        mid = l + (r - l) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target { // 向 <- 靠拢
            r = mid - 1
        } else { // 向 -> 靠拢
            l = mid + 1
        }
    }
    return -1
}

func search1(nums []int, target int) int {
    start, end := 0, len(nums)-1
    for start + 1 < end {
        mid := start + (end - start) >> 1
        if nums[mid] >= target {
            end = mid
        } else {
            start = mid
        }
    }
    if nums[start] == target {
        return start
    }
    if nums[end] == target {
        return end
    }
    return -1
}

func main() {
    // Explanation: 9 exists in nums and its index is 4
    fmt.Println(search([]int{1,0,3,5,9,12}, 9)) // 4
    // Explanation: 2 does not exist in nums so return -1
    fmt.Println(search([]int{-1,0,3,5,9,12}, 2)) // -1
    fmt.Println(search([]int{-1,0,3,5,9,12}, 13)) // -1
    fmt.Println(search([]int{-1,0,3,5,9,12}, -2)) // -1

    // Explanation: 9 exists in nums and its index is 4
    fmt.Println(search1([]int{1,0,3,5,9,12}, 9)) // 4
    // Explanation: 2 does not exist in nums so return -1
    fmt.Println(search1([]int{-1,0,3,5,9,12}, 2)) // -1
    fmt.Println(search1([]int{-1,0,3,5,9,12}, 13)) // -1
    fmt.Println(search1([]int{-1,0,3,5,9,12}, -2)) // -1
}