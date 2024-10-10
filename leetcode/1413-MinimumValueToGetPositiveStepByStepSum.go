package main

// 1413. Minimum Value to Get Positive Step by Step Sum
// Given an array of integers nums, you start with an initial positive value startValue.

// In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).

// Return the minimum positive value of startValue such that the step by step sum is never less than 1.

// Example 1:
// Input: nums = [-3,2,-3,4,2]
// Output: 5
// Explanation: If you choose startValue = 4, in the third iteration your step by step sum is less than 1.
// step by step sum
// startValue = 4 | startValue = 5 | nums
//   (4 -3 ) = 1  | (5 -3 ) = 2    |  -3
//   (1 +2 ) = 3  | (2 +2 ) = 4    |   2
//   (3 -3 ) = 0  | (4 -3 ) = 1    |  -3
//   (0 +4 ) = 4  | (1 +4 ) = 5    |   4
//   (4 +2 ) = 6  | (5 +2 ) = 7    |   2

// Example 2:
// Input: nums = [1,2]
// Output: 1
// Explanation: Minimum start value should be positive. 

// Example 3:
// Input: nums = [1,-2,-3]
// Output: 5

// Constraints:
//     1 <= nums.length <= 100
//     -100 <= nums[i] <= 100

import "fmt"

func minStartValue(nums []int) int {
    res, sum := 1 << 31, 0
    for _, v := range nums {
        sum += v
        if sum < res {
            res = sum
        }
    }
    if res > 0 { return 1 }
    return -res + 1
}

func minStartValue1(nums []int) int {
    mn, sum := 0, 0
    for _, v := range nums {
        sum += v
        if sum < mn {
            mn = sum
        }
    }
    return 1 - mn
}

func main() {
    // Example 1:
    // Input: nums = [-3,2,-3,4,2]
    // Output: 5
    // Explanation: If you choose startValue = 4, in the third iteration your step by step sum is less than 1.
    // step by step sum
    // startValue = 4 | startValue = 5 | nums
    //   (4 -3 ) = 1  | (5 -3 ) = 2    |  -3
    //   (1 +2 ) = 3  | (2 +2 ) = 4    |   2
    //   (3 -3 ) = 0  | (4 -3 ) = 1    |  -3
    //   (0 +4 ) = 4  | (1 +4 ) = 5    |   4
    //   (4 +2 ) = 6  | (5 +2 ) = 7    |   2
    fmt.Println(minStartValue([]int{-3,2,-3,4,2})) // 5
    // Example 2:
    // Input: nums = [1,2]
    // Output: 1
    // Explanation: Minimum start value should be positive. 
    fmt.Println(minStartValue([]int{1,2})) // 1
    // Example 3:
    // Input: nums = [1,-2,-3]
    // Output: 5
    fmt.Println(minStartValue([]int{1,-2,-3})) // 5

    fmt.Println(minStartValue1([]int{-3,2,-3,4,2})) // 5
    fmt.Println(minStartValue1([]int{1,2})) // 1
    fmt.Println(minStartValue1([]int{1,-2,-3})) // 5
}