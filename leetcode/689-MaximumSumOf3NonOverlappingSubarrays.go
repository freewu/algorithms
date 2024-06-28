package main

// 689. Maximum Sum of 3 Non-Overlapping Subarrays
// Given an integer array nums and an integer k, 
// find three non-overlapping subarrays of length k with maximum sum and return them.

// Return the result as a list of indices representing the starting position of each interval (0-indexed). 
// If there are multiple answers, return the lexicographically smallest one.

// Example 1:
// Input: nums = [1,2,1,2,6,7,5,1], k = 2
// Output: [0,3,5]
// Explanation: Subarrays [1, 2], [2, 6], [7, 5] correspond to the starting indices [0, 3, 5].
// We could have also taken [2, 1], but an answer of [1, 3, 5] would be lexicographically larger.

// Example 2:
// Input: nums = [1,2,1,2,1,2,1,2,1], k = 2
// Output: [0,2,4]

// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     1 <= nums[i] < 2^16
//     1 <= k <= floor(nums.length / 3)

import "fmt"

// 思路：类似三数之和
func maxSumOfThreeSubarrays(nums []int, k int) []int {
    sum1, maxSum1,  maSumIdx1 := 0, 0, 0  // 第一个数的最大值，此时的和和最大和及最大和其实index
    sum2, maxSum12, maxSum12Idx1, maxSum12Idx2 := 0, 0, 0, 0 // 第一个和第二个数的最大值，第二个滑动窗口的此时的和
    sum3, maxTotal := 0, 0 // 三个滑动窗口和的最大值 & 第三个滑动窗口的此时的和
    res := []int{}
    for i := 2*k; i < len(nums); i++ {
        sum1 += nums[i-2*k]
        sum2 += nums[i-k]
        sum3 += nums[i] // 三个滑动窗口的此时和
        if i >= 3*k -1 {
            if sum1 > maxSum1 { // 记录第一个数的最大值和起始下标
                maSumIdx1 = i - 3 * k + 1
                maxSum1 = sum1
            } 
            if maxSum1 + sum2 > maxSum12 { // 记录前两个数的最大值和起始下标
                maxSum12 = maxSum1+sum2
                maxSum12Idx1 = maSumIdx1
                maxSum12Idx2 = i-2 * k + 1
            }
            if maxSum12 + sum3 > maxTotal {// 记录三个数的最大值和起始下标
                maxTotal = maxSum12+sum3
                res = []int{ maxSum12Idx1, maxSum12Idx2, i - k + 1 }
            } 
            sum1 -= nums[i-3*k+1]
            sum2 -= nums[i-2*k+1]
            sum3 -= nums[i-k+1]
        }
    }
    return res
}

func maxSumOfThreeSubarrays1(nums []int, k int) []int {
    n := len(nums)
    sum1, sum2, sum3 := 0, 0, 0
    maxSum1, maxSum12, maxSumTotal := 0, 0, 0
    index1, index2, index12 := -1, -1, -1
    res := []int{0, 0, 0}
    for i := 2 * k; i < n; i ++ {
        sum1 += nums[i - 2 * k]
        sum2 += nums[i - k]
        sum3 += nums[i]
        if i >= 3 * k - 1 { // 求和元素数量首次达到k个，记录最大值和起始索引，且为下次循环减去起始值
            if sum1 > maxSum1 {
                maxSum1 = sum1
                index1 = i - 3 * k + 1
            }
            if sum2 + maxSum1 > maxSum12 {
                maxSum12 = sum2 + maxSum1
                index12 = index1 // 只有第二个数取最大时记录第一个区间的起始索引，才能保证一二区间不重叠
                index2 = i - 2 * k + 1
            }
            if sum3 + maxSum12 > maxSumTotal {
                maxSumTotal = sum3 + maxSum12
                // 只有第三个数取最大时记录一二区间的起始索引，才能保证三个区间不重叠
                res = []int{index12, index2, i - k + 1 }
            }
            sum1 -= nums[i - 3 * k + 1]
            sum2 -= nums[i - 2 * k + 1]
            sum3 -= nums[i - k + 1]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,2,6,7,5,1], k = 2
    // Output: [0,3,5]
    // Explanation: Subarrays [1, 2], [2, 6], [7, 5] correspond to the starting indices [0, 3, 5].
    // We could have also taken [2, 1], but an answer of [1, 3, 5] would be lexicographically larger.
    fmt.Println(maxSumOfThreeSubarrays([]int{1,2,1,2,6,7,5,1}, 2)) // [0,3,5]
    // Example 2:
    // Input: nums = [1,2,1,2,1,2,1,2,1], k = 2
    // Output: [0,2,4]
    fmt.Println(maxSumOfThreeSubarrays([]int{1,2,1,2,1,2,1,2,1}, 2)) // [0,2,4]

    fmt.Println(maxSumOfThreeSubarrays1([]int{1,2,1,2,6,7,5,1}, 2)) // [0,3,5]
    fmt.Println(maxSumOfThreeSubarrays1([]int{1,2,1,2,1,2,1,2,1}, 2)) // [0,2,4]
}