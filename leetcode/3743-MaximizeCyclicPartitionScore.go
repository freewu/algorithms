package main

// 3743. maximize Cyclic Partition Score
// You are given a cyclic array nums and an integer k.

// Partition nums into at most k subarrays. As nums is cyclic, these subarrays may wrap around from the end of the array back to the beginning.

// The range of a subarray is the difference between its maximum and minimum values. The score of a partition is the sum of subarray ranges.

// Return the maximum possible score among all cyclic partitions.

// Example 1:
// Input: nums = [1,2,3,3], k = 2
// Output: 3
// Explanation:
// Partition nums into [2, 3] and [3, 1] (wrapped around).
// The range of [2, 3] is max(2, 3) - min(2, 3) = 3 - 2 = 1.
// The range of [3, 1] is max(3, 1) - min(3, 1) = 3 - 1 = 2.
// The score is 1 + 2 = 3.

// Example 2:
// Input: nums = [1,2,3,3], k = 1
// Output: 2
// Explanation:
// Partition nums into [1, 2, 3, 3].
// The range of [1, 2, 3, 3] is max(1, 2, 3, 3) - min(1, 2, 3, 3) = 3 - 1 = 2.
// The score is 2.

// Example 3:
// Input: nums = [1,2,3,3], k = 4
// Output: 3
// Explanation:
// Identical to Example 1, we partition nums into [2, 3] and [3, 1]. Note that nums may be partitioned into fewer than k subarrays.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^9
//     1 <= k <= nums.length

import "fmt"
import "math"

func maximumScore(nums []int, k int) int64 {
    n := len(nums)
    if n == 0 { return 0 }
    // 找到最小值的索引j
    j := 0
    for i := 0; i < n; i++ {
        if nums[i] < nums[j] {
            j = i
        }
    }
    // 构建数组 arr1
    arr1 := make([]int, n)
    for i := 0; i < n; i++ {
        arr1[i] = nums[(j+i) % n]
    }
    // 构建数组 arr2（旋转后反转）
    arr2 := make([]int, n)
    for i := 0; i < n; i++ {
        arr2[i] = nums[(j + 1 + i) % n]
    }
    reverse := func(arr []int) { // 反转数组
        left, right := 0, len(arr)-1
        for left < right {
            arr[left], arr[right] = arr[right], arr[left]
            left++
            right--
        }
    }
    f := func(arr []int, k int) int64 {
        n := len(arr)
        if n == 0 { return 0 }
        // 限制k的最大值为n，因为最多只能分成n个单个元素的子数组
        actualK := k
        if actualK > n {
            actualK = n
        }
        // dp[i][j]表示将前j个元素分成i个子数组的最大得分
        dp := make([][]int64, actualK + 1)
        for i := range dp {
            dp[i] = make([]int64, n + 1)
        }
        // 初始化i=1的情况：前j个元素作为一个子数组的得分
        mn, mx  := int64(arr[0]), int64(arr[0])
        dp[1][1] = mx - mn
        for j := 1; j < n; j++ {
            curr := int64(arr[j])
            if curr < mn {
                mn = curr
            }
            if curr > mx {
                mx = curr
            }
            dp[1][j+1] = mx - mn
        }
        res := dp[1][n]
        // 计算i从2到actualK的情况
        for i := 2; i <= actualK; i++ {
            x, y := int64(math.MinInt64), int64(math.MinInt64) // 用于跟踪dp[i-1][j] - a[j]的最大值, 用于跟踪dp[i-1][j] + a[j]的最大值
            // j从i-1开始，因为至少需要i个元素才能分成i个子数组
            for j := i - 1; j < n; j++ {
                // 更新x和y，考虑前j个元素分成i-1个子数组的情况
                x = max(x, dp[i-1][j]-int64(arr[j]))    
                y = max(y, dp[i-1][j]+int64(arr[j]))
                // 当前j+1个元素分成i个子数组的最大得分，要么是前j个元素的得分（不分割），要么是新分割的得分
                dp[i][j+1] = max(dp[i][j], max(x + int64(arr[j]), y - int64(arr[j])))
            }
            // 更新全局最大值
            if dp[i][n] > res {
                res = dp[i][n]
            }
        }
        return res
    }
    reverse(arr2)
    // 计算并返回最大值
    return max(f(arr1, k), f(arr2, k))      
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,3], k = 2
    // Output: 3
    // Explanation:
    // Partition nums into [2, 3] and [3, 1] (wrapped around).
    // The range of [2, 3] is max(2, 3) - min(2, 3) = 3 - 2 = 1.
    // The range of [3, 1] is max(3, 1) - min(3, 1) = 3 - 1 = 2.
    // The score is 1 + 2 = 3.
    fmt.Println(maximumScore([]int{1,2,3,3}, 2)) // 3
    // Example 2:
    // Input: nums = [1,2,3,3], k = 1
    // Output: 2
    // Explanation:
    // Partition nums into [1, 2, 3, 3].
    // The range of [1, 2, 3, 3] is max(1, 2, 3, 3) - min(1, 2, 3, 3) = 3 - 1 = 2.
    // The score is 2.
    fmt.Println(maximumScore([]int{1,2,3,3}, 1)) // 2
    // Example 3:
    // Input: nums = [1,2,3,3], k = 4
    // Output: 3
    // Explanation:
    // Identical to Example 1, we partition nums into [2, 3] and [3, 1]. Note that nums may be partitioned into fewer than k subarrays.
    fmt.Println(maximumScore([]int{1,2,3,3}, 2)) // 3

    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, 2)) // 14
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, 2)) // 14
}