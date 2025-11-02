package main

// 3727. Maximum Alternating Sum of Squares
// You are given an integer array nums. You may rearrange the elements in any order.

// The alternating score of an array arr is defined as:
//     score = arr[0]2 - arr[1]2 + arr[2]2 - arr[3]2 + ...

// Return an integer denoting the maximum possible alternating score of nums after rearranging its elements.

// Example 1:
// Input: nums = [1,2,3]
// Output: 12
// Explanation:
// A possible rearrangement for nums is [2,1,3], which gives the maximum alternating score among all possible rearrangements.
// The alternating score is calculated as:
// score = 22 - 12 + 32 = 4 - 1 + 9 = 12

// Example 2:
// Input: nums = [1,-1,2,-2,3,-3]
// Output: 16
// Explanation:
// A possible rearrangement for nums is [-3,-1,-2,1,3,2], which gives the maximum alternating score among all possible rearrangements.
// The alternating score is calculated as:
// score = (-3)2 - (-1)2 + (-2)2 - (1)2 + (3)2 - (2)2 = 9 - 1 + 4 - 1 + 9 - 4 = 16

// Constraints:
//     1 <= nums.length <= 10^5
//     -4 * 10^4 <= nums[i] <= 4 * 10^4

import "fmt"
import "slices"

func maxAlternatingSum(nums []int) int64 {
    for i, v := range nums {
        nums[i] *= v
    }
    slices.Sort(nums)
    res, m := 0, len(nums) / 2
    // 交替和：减去小的，加上大的
    for _, v := range nums[:m] {
        res -= v
    }
    for _, v := range nums[m:] {
        res += v
    }
    return int64(res)
}

func maxAlternatingSum1(nums []int) int64 {
    res, n := 0, len(nums)
    for i := range nums {
        if nums[i] < 0 {
            nums[i] = -nums[i]
        }
    }
    slices.Sort(nums)
    i, j := 0, n - 1
    for i < j {
        res += nums[j]*nums[j] - nums[i]*nums[i]
        i++
        j--
    }
    if i == j {
        res += nums[i] * nums[i]
    }
    return int64(res)   
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 12
    // Explanation:
    // A possible rearrangement for nums is [2,1,3], which gives the maximum alternating score among all possible rearrangements.
    // The alternating score is calculated as:
    // score = 22 - 12 + 32 = 4 - 1 + 9 = 12
    fmt.Println(maxAlternatingSum([]int{1,2,3})) // 12
    // Example 2:
    // Input: nums = [1,-1,2,-2,3,-3]
    // Output: 16
    // Explanation:
    // A possible rearrangement for nums is [-3,-1,-2,1,3,2], which gives the maximum alternating score among all possible rearrangements.
    // The alternating score is calculated as:  
    // score = (-3)2 - (-1)2 + (-2)2 - (1)2 + (3)2 - (2)2 = 9 - 1 + 4 - 1 + 9 - 4 = 16
    fmt.Println(maxAlternatingSum([]int{1,-1,2,-2,3,-3})) // 16

    fmt.Println(maxAlternatingSum([]int{1,2,3,4,5,6,7,8,9})) // 225
    fmt.Println(maxAlternatingSum([]int{9,8,7,6,5,4,3,2,1})) // 225

    fmt.Println(maxAlternatingSum1([]int{1,2,3})) // 12
    fmt.Println(maxAlternatingSum1([]int{1,-1,2,-2,3,-3})) // 16
    fmt.Println(maxAlternatingSum1([]int{1,2,3,4,5,6,7,8,9})) // 225
    fmt.Println(maxAlternatingSum1([]int{9,8,7,6,5,4,3,2,1})) // 225
}   
