package main

// 1133. Largest Unique Number
// Given an integer array nums, return the largest integer that only occurs once. 
// If no integer occurs once, return -1.

// Example 1:
// Input: nums = [5,7,3,9,4,9,8,3,1]
// Output: 8
// Explanation: The maximum integer in the array is 9 but it is repeated. The number 8 occurs only once, so it is the answer.

// Example 2:
// Input: nums = [9,9,8,8]
// Output: -1
// Explanation: There is no number that occurs only once.
 
// Constraints:
//     1 <= nums.length <= 2000
//     0 <= nums[i] <= 1000

import "fmt"

func largestUniqueNumber(nums []int) int {
    res, m := -1, make(map[int]int)
    for _, v := range nums {
        m[v]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range m {
        if v == 1 { // 只判断出现一次的数字
            res = max(res, i)
        }
    } 
    return res
}

// 使用数组来代替 map
func largestUniqueNumber1(nums []int) int {
    res, arr := -1, make([]int, 1001 ) // 0 <= nums[i] <= 1000
    for _, v := range nums {
        arr[v]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range arr {
        if v == 1 { // 只判断出现一次的数字
            res = max(res, i)
        }
    } 
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,7,3,9,4,9,8,3,1]
    // Output: 8
    // Explanation: The maximum integer in the array is 9 but it is repeated. The number 8 occurs only once, so it is the answer.
    fmt.Println(largestUniqueNumber([]int{ 5,7,3,9,4,9,8,3,1 })) // 8
    // Example 2:
    // Input: nums = [9,9,8,8]
    // Output: -1
    // Explanation: There is no number that occurs only once.
    fmt.Println(largestUniqueNumber([]int{ 9,9,8,8 })) // -1

    fmt.Println(largestUniqueNumber1([]int{ 5,7,3,9,4,9,8,3,1 })) // 8
    fmt.Println(largestUniqueNumber1([]int{ 9,9,8,8 })) // -1
}