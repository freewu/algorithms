package main

// 1636. Sort Array by Increasing Frequency
// Given an array of integers nums, sort the array in increasing order based on the frequency of the values. 
// If multiple values have the same frequency, sort them in decreasing order.

// Return the sorted array.

// Example 1:
// Input: nums = [1,1,2,2,2,3]
// Output: [3,1,1,2,2,2]
// Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.

// Example 2:
// Input: nums = [2,3,1,3,2]
// Output: [1,3,3,2,2]
// Explanation: '2' and '3' both have a frequency of 2, so they are sorted in decreasing order.

// Example 3:
// Input: nums = [-1,1,-6,4,5,-6,1,4,1]
// Output: [5,-1,4,4,-6,-6,1,1,1]

// Constraints:
//     1 <= nums.length <= 100
//     -100 <= nums[i] <= 100

import "fmt"
import "sort"

func frequencySort(nums []int) []int {
    mp := map[int]int{}
    for _, n := range nums { // 统计出现频次
        mp[n]++
    }
    sort.Slice(nums, func(i,j int) bool {
        if mp[nums[i]] == mp[nums[j]] { // 如果有多个值的频率相同
            return nums[i] > nums[j] // 请你按照数值本身将它们 降序 排序
        } 
        return mp[nums[i]] < mp[nums[j]] // 按照每个值的频率 升序 排序
    })
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [1,1,2,2,2,3]
    // Output: [3,1,1,2,2,2]
    // Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.
    fmt.Println(frequencySort([]int{1,1,2,2,2,3})) // [3,1,1,2,2,2]
    // Example 2:
    // Input: nums = [2,3,1,3,2]
    // Output: [1,3,3,2,2]
    // Explanation: '2' and '3' both have a frequency of 2, so they are sorted in decreasing order.
    fmt.Println(frequencySort([]int{2,3,1,3,2})) // [1,3,3,2,2]
    // Example 3:
    // Input: nums = [-1,1,-6,4,5,-6,1,4,1]
    // Output: [5,-1,4,4,-6,-6,1,1,1]
    fmt.Println(frequencySort([]int{-1,1,-6,4,5,-6,1,4,1})) // [5,-1,4,4,-6,-6,1,1,1]
}