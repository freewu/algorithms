package main

// 1425. Constrained Subsequence Sum
// Given an integer array nums and an integer k, 
// return the maximum sum of a non-empty subsequence of 
// that array such that for every two consecutive integers in the subsequence, 
// nums[i] and nums[j], where i < j, the condition j - i <= k is satisfied.

// A subsequence of an array is obtained by deleting some number of elements (can be zero) from the array, 
// leaving the remaining elements in their original order.

// Example 1:
// Input: nums = [10,2,-10,5,20], k = 2
// Output: 37
// Explanation: The subsequence is [10, 2, 5, 20].

// Example 2:
// Input: nums = [-1,-2,-3], k = 1
// Output: -1
// Explanation: The subsequence must be non-empty, so we choose the largest number.

// Example 3:
// Input: nums = [10,-2,-10,-5,20], k = 2
// Output: 23
// Explanation: The subsequence is [10, -2, -5, 20].

// Constraints:
//     1 <= k <= nums.length <= 10^5
//     -10^4 <= nums[i] <= 10^4

import "fmt"

func constrainedSubsetSum(nums []int, k int) int {
    n := len(nums)
    dp, queue := make([]int, n), []int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        dp[i] = nums[i];
        if len(queue) > 0 {
            dp[i] += max(queue[0], 0)
        }
        for len(queue) > 0 && queue[len(queue) - 1] < dp[i] {
            queue = queue[:len(queue) - 1] // pop
        }
        queue = append(queue, dp[i])
        if i >= k && dp[i-k] == queue[0] {
            queue = queue[1:]
        }
    }
    res := -1 << 31
    for _, v := range dp {
        res = max(res, v)
    }
    return res
}

func constrainedSubsetSum1(nums []int, k int) int {
    n, res := len(nums), nums[0]
    queue := []int{0} // 存储idx,并且nums用做dp数组
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        // 缩窗口,维持所有元素在窗口内,因为k>=1,且i-1的在上一轮必定在队列末尾,所以q不会为空
        if queue[0] < i-k { // 一次移动一格,if即可
            queue = queue[1:]
        }
        nums[i] = max(nums[i], nums[i] + nums[queue[0]]) // 选与不选(前缀)
        res = max(res, nums[i])                    // 以[i]结束的子序列
        for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
            queue = queue[:len(queue)-1] // pop 出
        }
        queue = append(queue, i) // push
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [10,2,-10,5,20], k = 2
    // Output: 37
    // Explanation: The subsequence is [10, 2, 5, 20].
    fmt.Println(constrainedSubsetSum([]int{10,2,-10,5,20}, 2)) // 37
    // Example 2:
    // Input: nums = [-1,-2,-3], k = 1
    // Output: -1
    // Explanation: The subsequence must be non-empty, so we choose the largest number.
    fmt.Println(constrainedSubsetSum([]int{-1,-2,-3}, 1)) // -1
    // Example 3:
    // Input: nums = [10,-2,-10,-5,20], k = 2
    // Output: 23
    // Explanation: The subsequence is [10, -2, -5, 20].
    fmt.Println(constrainedSubsetSum([]int{10,-2,-10,-5,20}, 2)) // 23

    fmt.Println(constrainedSubsetSum1([]int{10,2,-10,5,20}, 2)) // 37
    fmt.Println(constrainedSubsetSum1([]int{-1,-2,-3}, 1)) // -1
    fmt.Println(constrainedSubsetSum1([]int{10,-2,-10,-5,20}, 2)) // 23
}