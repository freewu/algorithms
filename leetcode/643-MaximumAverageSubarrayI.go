package main

// 643. Maximum Average Subarray I
// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. 
// Any answer with a calculation error less than 10^-5 will be accepted.

// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000
 
// Constraints:
//     n == nums.length
//     1 <= k <= n <= 10^5
//     -10^4 <= nums[i] <= 10^4

import "fmt"
import "math"

func findMaxAverage(nums []int, k int) float64 {
    res, divider, sum,l := float64(math.MinInt64), float64(k), float64(0), 0
    for i := 0; i < len(nums); i++ {
        sum += float64(nums[i]) // 先累加到 k个数
        if i < k - 1 {
            continue
        }
        for (i - l) + 1 > k {// 滑动窗口
            sum -= float64(nums[l]) // 逐步把开头的数提出
            l++
        }
        res = math.Max(res, sum / divider)
    }
    return res
}

func findMaxAverage1(nums []int, k int) float64 {
    if len(nums) == 1 {
        return float64(nums[0])
    }
    sum := 0 // 计算最初窗口 sum 累加 k 个值 
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    res, avg := float64(sum) / float64(k), float64(0)
    for i := 1; i < len(nums); i++ {
        if i + k > len(nums) {
            break
        }
        sum = sum - nums[i-1] + nums[i-1 + k]
        avg = float64(sum) / float64(k)
        if avg > res {
            res = avg
        }
    }
    return res
}

func findMaxAverage2(nums []int, k int) float64 {
    sum := 0
    for i := 0; i < k; i++ { // 计算最初窗口 sum 累加 k - 1 个值 
        sum += nums[i]
    }
    t := sum
    for i := k; i < len(nums); i++ {
        t += nums[i] - nums[i-k]
        if t > sum {
            sum = t
        }
    }
    return float64(sum) / float64(k)
}

func main() {
    // Example 1:
    // Input: nums = [1,12,-5,-6,50,3], k = 4
    // Output: 12.75000
    // Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
    fmt.Println(findMaxAverage([]int{1,12,-5,-6,50,3}, 4)) // 12.75000
    // Example 2:
    // Input: nums = [5], k = 1
    // Output: 5.00000
    fmt.Println(findMaxAverage([]int{5}, 1)) // 5.00000

    fmt.Println(findMaxAverage1([]int{1,12,-5,-6,50,3}, 4)) // 12.75000
    fmt.Println(findMaxAverage1([]int{5}, 1)) // 5.00000

    fmt.Println(findMaxAverage2([]int{1,12,-5,-6,50,3}, 4)) // 12.75000
    fmt.Println(findMaxAverage2([]int{5}, 1)) // 5.00000
}