package main

// 485. Max Consecutive Ones
// Given a binary array nums, return the maximum number of consecutive 1's in the array.

// Example 1:
// Input: nums = [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.

// Example 2:
// Input: nums = [1,0,1,1,0,1]
// Output: 2

// Constraints:
//     1 <= nums.length <= 10^5
//     nums[i] is either 0 or 1.

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
    res,count := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        if v == 1 {
            count++
            res = max(res, count)
        } else {
            count = 0
        }
    }
    return res
}

func findMaxConsecutiveOnes1(nums []int) int {
    res,count := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        if v == 1 {
            count++
        } else {
            res = max(res, count)
            count = 0
        }
    }
    res = max(res, count)
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,0,1,1,1]
    // Output: 3
    // Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
    fmt.Println(findMaxConsecutiveOnes([]int{1,1,0,1,1,1})) // 3
    // Example 2:
    // Input: nums = [1,0,1,1,0,1]
    // Output: 2
    fmt.Println(findMaxConsecutiveOnes([]int{1,0,1,1,0,1})) // 2

    fmt.Println(findMaxConsecutiveOnes1([]int{1,1,0,1,1,1})) // 3
    fmt.Println(findMaxConsecutiveOnes1([]int{1,0,1,1,0,1})) // 2
}