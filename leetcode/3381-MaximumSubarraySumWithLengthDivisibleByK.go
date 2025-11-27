package main

// 3381. Maximum Subarray Sum With Length Divisible by K
// You are given an array of integers nums and an integer k.

// Return the maximum sum of a subarray of nums, such that the size of the subarray is divisible by k.

// Example 1:
// Input: nums = [1,2], k = 1
// Output: 3
// Explanation:
// The subarray [1, 2] with sum 3 has length equal to 2 which is divisible by 1.

// Example 2:
// Input: nums = [-1,-2,-3,-4,-5], k = 4
// Output: -10
// Explanation:
// The maximum sum subarray is [-1, -2, -3, -4] which has length equal to 4 which is divisible by 4.

// Example 3:
// Input: nums = [-5,1,2,-3,4], k = 2
// Output: 4
// Explanation:
// The maximum sum subarray is [1, 2, -3, 4] which has length equal to 4 which is divisible by 2.

// Constraints:
//     1 <= k <= nums.length <= 2 * 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"

// Kadane's Algorithm
func maxSubarraySum(nums []int, k int) int64 {
    res, sum, n := -1 << 61, 0, len(nums)
    arr := make([]int, n - k + 1)
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    arr[0] = sum
    for i := k; i < n; i++ {
        sum += (nums[i] - nums[i - k])
        arr[i - k + 1] = sum
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < k; i++ {
        curr := 0
        for j := i; j + k <= n; j = j + k {
            curr = max(arr[j], curr + arr[j])
            res = max(res, curr)
        }
    }
    return int64(res)
}

func maxSubarraySum1(nums []int, k int) int64 {
    sum := make([]int, len(nums) + 1)
    for i, v := range nums {
        sum[i+1] = sum[i] + v
    }
    mn := make([]int, k)
    for i := range mn {
        mn[i] = 1 << 61
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := -1 << 61
    for j, v := range sum {
        i := j % k
        res = max(res, v - mn[i])
        mn[i] = min(mn[i], v)
    }
    return int64(res)
}

func maxSubarraySum2(nums []int, k int) int64 {
    prefix, n := int64(0), len(nums)
    arr := make([]int64, k)
    for i := 0; i < k-1; i++ {
        prefix += int64(nums[i])
        arr[i+1] = prefix
    }
    res := prefix + int64(nums[k-1])
    for i := k-1; i < n; i++ {
        prefix += int64(nums[i])
        cur := prefix - arr[(i+1)%k]
        if cur > res {
            res = cur
        }
        if prefix < arr[(i + 1) % k] {
            arr[(i+1)%k] = prefix   
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2], k = 1
    // Output: 3
    // Explanation:
    // The subarray [1, 2] with sum 3 has length equal to 2 which is divisible by 1.
    fmt.Println(maxSubarraySum([]int{1,2}, 1)) // 3
    // Example 2:
    // Input: nums = [-1,-2,-3,-4,-5], k = 4
    // Output: -10
    // Explanation:
    // The maximum sum subarray is [-1, -2, -3, -4] which has length equal to 4 which is divisible by 4.
    fmt.Println(maxSubarraySum([]int{-1,-2,-3,-4,-5}, 4)) // -10
    // Example 3:
    // Input: nums = [-5,1,2,-3,4], k = 2
    // Output: 4
    // Explanation:
    // The maximum sum subarray is [1, 2, -3, 4] which has length equal to 4 which is divisible by 2.
    fmt.Println(maxSubarraySum([]int{-5,1,2,-3,4}, 2)) // 4

    fmt.Println(maxSubarraySum([]int{1,2,3,4,5,6,7,8,9}, 2)) // 44
    fmt.Println(maxSubarraySum([]int{9,8,7,6,5,4,3,2,1}, 2)) // 44

    fmt.Println(maxSubarraySum1([]int{1,2}, 1)) // 3
    fmt.Println(maxSubarraySum1([]int{-1,-2,-3,-4,-5}, 4)) // -10
    fmt.Println(maxSubarraySum1([]int{-5,1,2,-3,4}, 2)) // 4
    fmt.Println(maxSubarraySum1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 44
    fmt.Println(maxSubarraySum1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 44

    fmt.Println(maxSubarraySum2([]int{1,2}, 1)) // 3
    fmt.Println(maxSubarraySum2([]int{-1,-2,-3,-4,-5}, 4)) // -10
    fmt.Println(maxSubarraySum2([]int{-5,1,2,-3,4}, 2)) // 4
    fmt.Println(maxSubarraySum2([]int{1,2,3,4,5,6,7,8,9}, 2)) // 44
    fmt.Println(maxSubarraySum2([]int{9,8,7,6,5,4,3,2,1}, 2)) // 44
}