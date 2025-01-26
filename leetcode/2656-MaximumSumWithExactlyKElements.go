package main

// 2656. Maximum Sum With Exactly K Elements
// You are given a 0-indexed integer array nums and an integer k. 
// Your task is to perform the following operation exactly k times in order to maximize your score:
//     1. Select an element m from nums.
//     2. Remove the selected element m from the array.
//     3. Add a new element with a value of m + 1 to the array.
//     4. Increase your score by m.

// Return the maximum score you can achieve after performing the operation exactly k times.

// Example 1:
// Input: nums = [1,2,3,4,5], k = 3
// Output: 18
// Explanation: We need to choose exactly 3 elements from nums to maximize the sum.
// For the first iteration, we choose 5. Then sum is 5 and nums = [1,2,3,4,6]
// For the second iteration, we choose 6. Then sum is 5 + 6 and nums = [1,2,3,4,7]
// For the third iteration, we choose 7. Then sum is 5 + 6 + 7 = 18 and nums = [1,2,3,4,8]
// So, we will return 18.
// It can be proven, that 18 is the maximum answer that we can achieve.

// Example 2:
// Input: nums = [5,5,5], k = 2
// Output: 11
// Explanation: We need to choose exactly 2 elements from nums to maximize the sum.
// For the first iteration, we choose 5. Then sum is 5 and nums = [5,5,6]
// For the second iteration, we choose 6. Then sum is 5 + 6 = 11 and nums = [5,5,7]
// So, we will return 11.
// It can be proven, that 11 is the maximum answer that we can achieve.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100
//     1 <= k <= 100

import "fmt"
import "sort"
import "slices"

func maximizeSum(nums []int, k int) int {
    sort.Ints(nums)
    res := nums[len(nums)-1]
    start, last := res + 1, res + k
    for start < last{
        res += start
        start++
    }
    return res
}

func maximizeSum1(nums []int, k int) int {
    mx := slices.Max(nums)
    return mx * k + (k - 1) * k / 2
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], k = 3
    // Output: 18
    // Explanation: We need to choose exactly 3 elements from nums to maximize the sum.
    // For the first iteration, we choose 5. Then sum is 5 and nums = [1,2,3,4,6]
    // For the second iteration, we choose 6. Then sum is 5 + 6 and nums = [1,2,3,4,7]
    // For the third iteration, we choose 7. Then sum is 5 + 6 + 7 = 18 and nums = [1,2,3,4,8]
    // So, we will return 18.
    // It can be proven, that 18 is the maximum answer that we can achieve.
    fmt.Println(maximizeSum([]int{1,2,3,4,5}, 3)) // 18
    // Example 2:
    // Input: nums = [5,5,5], k = 2
    // Output: 11
    // Explanation: We need to choose exactly 2 elements from nums to maximize the sum.
    // For the first iteration, we choose 5. Then sum is 5 and nums = [5,5,6]
    // For the second iteration, we choose 6. Then sum is 5 + 6 = 11 and nums = [5,5,7]
    // So, we will return 11.
    // It can be proven, that 11 is the maximum answer that we can achieve.
    fmt.Println(maximizeSum([]int{5,5,5}, 2)) // 11

    fmt.Println(maximizeSum([]int{1,2,3,4,5,6,7,8,9}, 2)) // 19
    fmt.Println(maximizeSum([]int{9,8,7,6,5,4,3,2,1}, 2)) // 19

    fmt.Println(maximizeSum1([]int{1,2,3,4,5}, 3)) // 18
    fmt.Println(maximizeSum1([]int{5,5,5}, 2)) // 11
    fmt.Println(maximizeSum1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 19
    fmt.Println(maximizeSum1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 19
}