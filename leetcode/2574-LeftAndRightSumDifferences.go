package main

// 2574. Left and Right Sum Differences
// Given a 0-indexed integer array nums, find a 0-indexed integer array answer where:
//     1. answer.length == nums.length.
//     2. answer[i] = |leftSum[i] - rightSum[i]|.

// Where:
//     1. leftSum[i] is the sum of elements to the left of the index i in the array nums. 
//        If there is no such element, leftSum[i] = 0.
//     2. rightSum[i] is the sum of elements to the right of the index i in the array nums. 
//        If there is no such element, rightSum[i] = 0.

// Return the array answer.

// Example 1:
// Input: nums = [10,4,8,3]
// Output: [15,1,11,22]
// Explanation: The array leftSum is [0,10,14,22] and the array rightSum is [15,11,3,0].
// The array answer is [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22].

// Example 2:
// Input: nums = [1]
// Output: [0]
// Explanation: The array leftSum is [0] and the array rightSum is [0].
// The array answer is [|0 - 0|] = [0].

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^5

import "fmt"

func leftRightDifference(nums []int) []int {
    res, left, right := []int{}, 0, 0
    for _, v := range nums {
        right += v
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, v := range nums {
        right -= v
        res = append(res, abs(left - right))
        left += v
    }
    return res
}

func leftRightDifference1(nums []int) []int {
    n := len(nums)
    res, left, right := make([]int, n), make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        if i > 0 {
            left[i] = left[i-1] + nums[i-1]
        }
        if n - i < n {
            right[n - i - 1] = right[n - i] + nums[n - i]
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n; i++ {
        res[i] = abs(left[i] - right[i])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,4,8,3]
    // Output: [15,1,11,22]
    // Explanation: The array leftSum is [0,10,14,22] and the array rightSum is [15,11,3,0].
    // The array answer is [|0 - 15|,|10 - 11|,|14 - 3|,|22 - 0|] = [15,1,11,22].
    fmt.Println(leftRightDifference([]int{10,4,8,3})) // [15,1,11,22]
    // Example 2:
    // Input: nums = [1]
    // Output: [0]
    // Explanation: The array leftSum is [0] and the array rightSum is [0].
    // The array answer is [|0 - 0|] = [0].
    fmt.Println(leftRightDifference([]int{1})) // [0]

    fmt.Println(leftRightDifference([]int{1,2,3,4,5,6,7,8,9})) // [44 41 36 29 20 9 4 19 36]
    fmt.Println(leftRightDifference([]int{9,8,7,6,5,4,3,2,1})) // [36 19 4 9 20 29 36 41 44]

    fmt.Println(leftRightDifference1([]int{10,4,8,3})) // [15,1,11,22]
    fmt.Println(leftRightDifference1([]int{1})) // [0]
    fmt.Println(leftRightDifference1([]int{1,2,3,4,5,6,7,8,9})) // [44 41 36 29 20 9 4 19 36]
    fmt.Println(leftRightDifference1([]int{9,8,7,6,5,4,3,2,1})) // [36 19 4 9 20 29 36 41 44]
}