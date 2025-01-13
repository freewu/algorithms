package main

// 2587. Rearrange Array to Maximize Prefix Score
// You are given a 0-indexed integer array nums. 
// You can rearrange the elements of nums to any order (including the given order).

// Let prefix be the array containing the prefix sums of nums after rearranging it. 
// In other words, prefix[i] is the sum of the elements from 0 to i in nums after rearranging it. 
// The score of nums is the number of positive integers in the array prefix.

// Return the maximum score you can achieve.

// Example 1:
// Input: nums = [2,-1,0,1,-3,3,-3]
// Output: 6
// Explanation: We can rearrange the array into nums = [2,3,1,-1,-3,0,-3].
// prefix = [2,5,6,5,2,2,-1], so the score is 6.
// It can be shown that 6 is the maximum score we can obtain.

// Example 2:
// Input: nums = [-2,-3,0]
// Output: 0
// Explanation: Any rearrangement of the array will result in a score of 0.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^6 <= nums[i] <= 10^6

import "fmt"
import "sort"

func maxScore(nums []int) int {
    res, sum := 0, 0
    sort.SliceStable(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
    for _, v := range nums {
        sum += v
        if sum > 0 {
            res++
        } else {
            break
        }
    }
    return res
}

func maxScore1(nums []int) int {
    sort.Ints(nums)
    res := 0
    for i, sum := len(nums) - 1, 0; i >= 0; i-- {
        sum += nums[i]
        if sum <= 0 { break }
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,-1,0,1,-3,3,-3]
    // Output: 6
    // Explanation: We can rearrange the array into nums = [2,3,1,-1,-3,0,-3].
    // prefix = [2,5,6,5,2,2,-1], so the score is 6.
    // It can be shown that 6 is the maximum score we can obtain.
    fmt.Println(maxScore([]int{2,-1,0,1,-3,3,-3})) // 6
    // Example 2:
    // Input: nums = [-2,-3,0]
    // Output: 0
    // Explanation: Any rearrangement of the array will result in a score of 0.
    fmt.Println(maxScore([]int{-2,-3,0})) // 0

    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(maxScore1([]int{2,-1,0,1,-3,3,-3})) // 6
    fmt.Println(maxScore1([]int{-2,-3,0})) // 0
    fmt.Println(maxScore1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(maxScore1([]int{9,8,7,6,5,4,3,2,1})) // 9
}