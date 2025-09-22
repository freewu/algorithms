package main

// 3005. Count Elements With Maximum Frequency
// You are given an array nums consisting of positive integers.
// Return the total frequencies of elements in nums such that those elements all have the maximum frequency.
// The frequency of an element is the number of occurrences of that element in the array.

// Example 1:
// Input: nums = [1,2,2,3,1,4]
// Output: 4
// Explanation: The elements 1 and 2 have a frequency of 2 which is the maximum frequency in the array.
// So the number of elements in the array with maximum frequency is 4.

// Example 2:
// Input: nums = [1,2,3,4,5]
// Output: 5
// Explanation: All elements of the array have a frequency of 1 which is the maximum.
// So the number of elements in the array with maximum frequency is 5.

// Constraints:
//         1 <= nums.length <= 100
//         1 <= nums[i] <= 100

import "fmt"

func maxFrequencyElements(nums []int) int {
    // 先统计每个元素的频率
    mx, m := 1, make(map[int]int) 
    for _,v := range nums {
        m[v]++
    }
    // 找到最大的频率
    for _,v := range m {
        if v > mx {
            mx = v
        }
    }
    // 计算频率和
    res := 0
    for _,v := range m {
        if v == mx {
            res += v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,3,1,4]
    // Output: 4
    // Explanation: The elements 1 and 2 have a frequency of 2 which is the maximum frequency in the array.
    // So the number of elements in the array with maximum frequency is 4.
    fmt.Println(maxFrequencyElements([]int{1,2,2,3,1,4})) // 4
    // Example 2:
    // Input: nums = [1,2,3,4,5]
    // Output: 5
    // Explanation: All elements of the array have a frequency of 1 which is the maximum.
    // So the number of elements in the array with maximum frequency is 5.
    fmt.Println(maxFrequencyElements([]int{1,2,3,4,5})) // 5

    fmt.Println(maxFrequencyElements([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(maxFrequencyElements([]int{9,8,7,6,5,4,3,2,1})) // 9
}