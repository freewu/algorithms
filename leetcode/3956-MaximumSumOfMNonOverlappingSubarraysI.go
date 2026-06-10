package main

// 3956. Maximum Sum of M Non-Overlapping Subarrays I
// You are given an integer array nums of length n, and three integers m, l, and r.

// Your task is to select at least one and at most m non-overlapping subarrays from nums such that:
//     1. Each selected subarray has a length between [l, r] (inclusive).
//     2. The total sum of all selected subarrays is maximized.

// Return the maximum total sum you can achieve.

// Example 1:
// Input: nums = [4,1,-5,2], m = 2, l = 1, r = 3
// Output: 7
// Explanation:
// One optimal strategy is to:
// Select the subarray [4, 1] with sum 4 + 1 = 5 and the subarray [2] with sum 2. Both subarrays have length between [l, r].
// The total sum of these subarrays is 5 + 2 = 7, which is the maximum achievable sum with at most m = 2 subarrays.

// Example 2:
// Input: nums = [1,0,3,4], m = 2, l = 1, r = 2
// Output: 8
// Explanation:
// One optimal strategy is to:
// Select the subarray [1] with sum 1 and the subarray [3, 4] with sum 3 + 4 = 7. Both subarrays have length between [l, r].
// The total sum of these subarrays is 1 + 7 = 8, which is the maximum achievable sum with at most m = 2 subarrays.

// Example 3:
// Input: nums = [-1,7,-4], m = 1, l = 2, r = 3
// Output: 6
// Explanation:
// Select the subarray [-1, 7] from nums which has length between [l, r].
// The total sum of this subarray is -1 + 7 = 6, which is the maximum achievable sum with at most m = 1 subarray.

// Example 4:
// Input: nums = [-3,-4,-1], m = 2, l = 1, r = 2
// Output: -1
// Explanation:
// All subarrays of nums have negative sums. The optimal strategy is to select the subarray [-1], which has length between [l, r].
// The total sum of this subarray is -1, which is the maximum achievable sum with at most m = 2 subarrays.

// Constraints:
//     1 <= n == nums.length <= 1000
//     -10^9 <= nums[i] <= 10^9​​​​​​​
//     1 <= m <= n
//     1 <= l <= r <= n

import "fmt"
import "sort"

func maximumSum(nums []int, m, left, right int) int64 {
    res, n := -1 << 61, len(nums)
    prefix := make([]int, n + 1) // nums 的前缀和
    for i, v := range nums {
        prefix[i + 1] = prefix[i] + v
    }
    // f[i][j] 表示在前 j 个数（下标 0 到 j-1）中选出恰好 i 个子数组，所选元素之和的最大值
    f := make([]int, n + 1)
    for i := 1; i <= m; i++ {
        nf := make([]int, n + 1)
        for j := range nf {
            nf[j] = -1 << 61
        }
        q := []int{}
        // 前 i 个子数组至少占用了 i * left 个位置
        for j := i * left; j <= n; j++ {
            // 1. 入
            k := j - left
            v := f[k] - prefix[k]
            for len(q) > 0 && f[q[len(q)-1]] - prefix[q[len(q)-1]] <= v {
                q = q[:len(q)-1]
            }
            q = append(q, k)
            // 2. 更新
            // 不选 nums[j-1] vs 选一个以 j-1 结尾的子数组
            nf[j] = max(nf[j-1], f[q[0]] - prefix[q[0]] + prefix[j])
            // 3. 出，下一轮循环队首离开窗口
            if q[0] <= j-right {
                q = q[1:]
            }
        }
        // 枚举恰好选 i 个子数组
        f = nf
        res = max(res, f[n])
    }
    return int64(res)
}

func maximumSum1(nums []int, m, l, r int) int64 {
    type Pair struct{ f, count int } // DP 值, 子数组个数
    res, n := 0, len(nums)
    s := make([]int, n + 1) // nums 的前缀和
    posSum := 0 // nums 中的正数之和
    for i, v := range nums {
        s[i+1] = s[i] + v
        if v > 0 {
            posSum += v
        }
    }
    less := func (a, b Pair) bool {  // 相等的时候，子数组个数更大的劣
        return a.f < b.f || a.f == b.f && a.count > b.count
    }
    // 没有 m 约束，但每选一个子数组就要把元素和减少 k
    dpWithoutLimit := func(k int) Pair {
        f := make([]Pair, n+1)
        q := []int{}
        res := Pair{ -1 << 61, 0}
        for i := l; i <= n; i++ {
            // 1. 入
            j := i - l
            v := Pair{f[j].f - s[j], f[j].count}
            for len(q) > 0 && less(Pair{f[q[len(q)-1]].f - s[q[len(q)-1]], f[q[len(q)-1]].count}, v) {
                q = q[:len(q)-1]
            }
            q = append(q, j)
            // 2. 更新答案
            choose := Pair{f[q[0]].f - s[q[0]] + s[i] - k, f[q[0]].count + 1}
            if less(res, choose) {
                // choose 保证我们至少选了一个子数组
                res = choose
            }
            // 更新 DP
            if less(f[i-1], choose) {
                f[i] = choose
            } else { // 不选
                f[i] = f[i-1]
            }
            // 3. 出，下一轮循环队首离开窗口
            if q[0] <= i-r {
                q = q[1:]
            }
        }
        return res
    }
    pair := dpWithoutLimit(0)
    if pair.count <= m { // 直接满足题目要求
        return int64(pair.f)
    }
    // 现在专注于解决「选恰好 m 个子数组」的问题
    sort.Search(posSum, func(k int) bool {
        k++
        pair := dpWithoutLimit(k)
        if pair.count <= m {
            res = pair.f + m * k // 不需要取 max，二分最终会缩小到凸函数中的 x=m 所在的那条线段
            return true
        }
        return false
    })
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [4,1,-5,2], m = 2, l = 1, r = 3
    // Output: 7
    // Explanation:
    // One optimal strategy is to:
    // Select the subarray [4, 1] with sum 4 + 1 = 5 and the subarray [2] with sum 2. Both subarrays have length between [l, r].
    // The total sum of these subarrays is 5 + 2 = 7, which is the maximum achievable sum with at most m = 2 subarrays.
    fmt.Println(maximumSum([]int{4,1,-5,2}, 2, 1, 3)) // 7
    // Example 2:
    // Input: nums = [1,0,3,4], m = 2, l = 1, r = 2
    // Output: 8
    // Explanation:
    // One optimal strategy is to:
    // Select the subarray [1] with sum 1 and the subarray [3, 4] with sum 3 + 4 = 7. Both subarrays have length between [l, r].
    // The total sum of these subarrays is 1 + 7 = 8, which is the maximum achievable sum with at most m = 2 subarrays.
    fmt.Println(maximumSum([]int{1,0,3,4}, 2, 1, 2)) // 8
    // Example 3:
    // Input: nums = [-1,7,-4], m = 1, l = 2, r = 3
    // Output: 6
    // Explanation:
    // Select the subarray [-1, 7] from nums which has length between [l, r].
    // The total sum of this subarray is -1 + 7 = 6, which is the maximum achievable sum with at most m = 1 subarray.
    fmt.Println(maximumSum([]int{-1,7,-4}, 1, 2, 3)) // 6
    // Example 4:
    // Input: nums = [-3,-4,-1], m = 2, l = 1, r = 2
    // Output: -1
    // Explanation:
    // All subarrays of nums have negative sums. The optimal strategy is to select the subarray [-1], which has length between [l, r].
    // The total sum of this subarray is -1, which is the maximum achievable sum with at most m = 2 subarrays.
    fmt.Println(maximumSum([]int{-3,-4,-1}, 2, 1, 2)) // -1

    fmt.Println(maximumSum([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 2)) // 30
    fmt.Println(maximumSum([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 2)) // 30

    fmt.Println(maximumSum1([]int{4,1,-5,2}, 2, 1, 3)) // 7
    fmt.Println(maximumSum1([]int{1,0,3,4}, 2, 1, 2)) // 8
    fmt.Println(maximumSum1([]int{-1,7,-4}, 1, 2, 3)) // 6
    fmt.Println(maximumSum1([]int{-3,-4,-1}, 2, 1, 2)) // -1
    fmt.Println(maximumSum1([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 2)) // 30
    fmt.Println(maximumSum1([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 2)) // 30
}