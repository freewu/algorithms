package main

// 3929. Minimum Partition Score II
// You are given an integer array nums and an integer k.

// Your task is to partition nums into exactly k subarrays and return an integer denoting the minimum possible score among all valid partitions.

// The score of a partition is the sum of the values of all its subarrays.

// The value of a subarray is defined as sumArr * (sumArr + 1) / 2, where sumArr is the sum of its elements.

// Example 1:
// Input: nums = [5,1,2,1], k = 2
// Output: 25
// Explanation:
// We must partition the array into k = 2 subarrays. One optimal partition is [5] and [1, 2, 1].
// The first subarray has sum = 5 and value = 5 * 6 / 2 = 15.
// The second subarray has sum = 1 + 2 + 1 = 4 and value = 4 * 5 / 2 = 10.
// The score of this partition is 15 + 10 = 25, which is the minimum possible score.

// Example 2:
// Input: nums = [1,2,3,4], k = 1
// Output: 55
// Explanation:
// Since we must partition the array into k = 1 subarray, all elements belong to the same subarray: [1, 2, 3, 4].
// This subarray has sum = 1 + 2 + 3 + 4 = 10 and value = 10 * 11 / 2 = 55.​​​​​​​
// The score of this partition is 55, which is the minimum possible score.

// Example 3:
// Input: nums = [1,1,1], k = 3
// Output: 3
// Explanation:
// We must partition the array into k = 3 subarrays. The only valid partition is [1], [1], [1].
// Each subarray has sum = 1 and value = 1 * 2 / 2 = 1.
// The score of this partition is 1 + 1 + 1 = 3, which is the minimum possible score.

// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     1 <= nums[i] <= 10^3
//     1 <= k <= nums.length 

import "fmt"

// 超出时间限制 837 / 847 个通过的测试用例
func minPartitionScore(nums []int, k int) int64 {
    n := len(nums)
    // 前缀和 s[0] = 0, s[i] = nums[0]+...+nums[i-1]
    s := make([]int64, n+1)
    for i := 0; i < n; i++ {
        s[i+1] = s[i] + int64(nums[i])
    }
    // dp[i] = 前 i 个元素分成当前轮次段数的最小代价
    dp := make([]int64, n+1)
    for i := range dp {
        dp[i] = 1 << 61
    }
    dp[0] = 0
    // 分成 k 段，迭代 k 次
    for seg := 1; seg <= k; seg++ {
        nextDp := make([]int64, n+1)
        for i := range nextDp {
            nextDp[i] = 1 << 61
        }
        // 枚举右端点 i，左端点 j
        // 分成 seg 段 => j >= seg-1
        for i := seg; i <= n; i++ {
            for j := seg - 1; j < i; j++ {
                ds := s[i] - s[j]
                cost := ds * (ds + 1) / 2
                if dp[j]+cost < nextDp[i] {
                    nextDp[i] = dp[j] + cost
                }
            }
        }
        dp = nextDp
    }
    return dp[n]
}

func minPartitionScore1(nums []int, k int) int64 {
    n := len(nums)
    // 前缀和数组
    s := make([]int64, n+1)
    for i := 0; i < n; i++ {
        s[i+1] = s[i] + int64(nums[i])
    }
    // 前缀和的平方数组
    p2 := make([]int64, n+1)
    for i := 0; i <= n; i++ {
        p2[i] = s[i] * s[i]
    }
    // 特殊情况：只分1段
    if k == 1 {
        total := s[n]
        return total * (total + 1) / 2
    }
    // 核心计算函数：给定lambda，返回最小代价和分割段数
    calc := func(lamb int64) (int64, int) {
        f := make([]int64, n+1)  // 最小代价数组
        cnt := make([]int, n+1)   // 分割段数数组
        q := make([]int, n+1)     // 凸包优化队列
        head, tail := 0, 1
        for index := 1; index <= n; index++ {
            // 队列头部优化：移除非最优决策点
            for tail-1 > head {
                p0 := q[head]
                p1 := q[head+1]
                // 计算两个决策点的代价
                b0 := f[p0] + p2[p0] - s[p0] - 2*s[index]*s[p0]
                b1 := f[p1] + p2[p1] - s[p1] - 2*s[index]*s[p1]
                if b1 <= b0 {
                    // 代价相等时，保留段数更多的点
                    if b1 == b0 && cnt[p0] > cnt[p1] {
                        cnt[p1] = cnt[p0]
                    }
                    head++
                } else {
                    break
                }
            }
            // 取最优决策点更新状态
            p := q[head]
            d := s[index] - s[p]
            f[index] = f[p] + d*(d+1) + lamb
            cnt[index] = cnt[p] + 1
            // 队列尾部优化：维护凸包性质
            for tail-1 > head {
                xIdx, yIdx := s[index], f[index] + p2[index] - s[index]
                pTail1 := q[tail-1]
                xLast1 := s[pTail1]
                yLast1 := f[pTail1] + p2[pTail1] - s[pTail1]
                pTail2 := q[tail-2]
                xLast2 := s[pTail2]
                yLast2 := f[pTail2] + p2[pTail2] - s[pTail2]
                x1, y1 := xIdx - xLast2, yIdx - yLast2
                x2, y2 := xLast1 - xLast2, yLast1 - yLast2
                // 叉积判断，维护下凸包
                cross := x1*y2 - x2*y1
                if cross > 0 {
                    tail--
                } else {
                    break
                }
            }
            q[tail] = index
            tail++
        }
        return f[n], cnt[n]
    }
    // WQS 二分主体
    left, right := int64(0), int64(1e18)
    for left < right {
        mid := left + (right-left+1)/2 // 防止溢出
        _, count := calc(mid)
        if count < k {
            right = mid - 1
        } else {
            left = mid
        }
    }
    // 计算最终结果
    res, _ := calc(left)
    return (res - int64(k) * left) / 2
}

func main() {
    // Example 1:
    // Input: nums = [5,1,2,1], k = 2
    // Output: 25
    // Explanation:
    // We must partition the array into k = 2 subarrays. One optimal partition is [5] and [1, 2, 1].
    // The first subarray has sum = 5 and value = 5 * 6 / 2 = 15.
    // The second subarray has sum = 1 + 2 + 1 = 4 and value = 4 * 5 / 2 = 10.
    // The score of this partition is 15 + 10 = 25, which is the minimum possible score.
    fmt.Println(minPartitionScore([]int{5,1,2,1}, 2)) // 25 
    // Example 2:
    // Input: nums = [1,2,3,4], k = 1
    // Output: 55
    // Explanation:
    // Since we must partition the array into k = 1 subarray, all elements belong to the same subarray: [1, 2, 3, 4].
    // This subarray has sum = 1 + 2 + 3 + 4 = 10 and value = 10 * 11 / 2 = 55.​​​​​​​
    // The score of this partition is 55, which is the minimum possible score.
    fmt.Println(minPartitionScore([]int{1,2,3,4}, 1)) // 55 
    // Example 3:
    // Input: nums = [1,1,1], k = 3
    // Output: 3
    // Explanation:
    // We must partition the array into k = 3 subarrays. The only valid partition is [1], [1], [1].
    // Each subarray has sum = 1 and value = 1 * 2 / 2 = 1.
    // The score of this partition is 1 + 1 + 1 = 3, which is the minimum possible score.   
    fmt.Println(minPartitionScore([]int{1,1,1}, 3)) // 3 

    fmt.Println(minPartitionScore([]int{1,2,3,4,5,6,7,8,9}, 2)) // 531 
    fmt.Println(minPartitionScore([]int{9,8,7,6,5,4,3,2,1}, 2)) // 531 

    fmt.Println(minPartitionScore1([]int{1,2,3,4}, 1)) // 55    
    fmt.Println(minPartitionScore1([]int{1,1,1}, 3)) // 3 
    fmt.Println(minPartitionScore1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 531 
    fmt.Println(minPartitionScore1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 531 
}