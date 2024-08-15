package main

// 3117. Minimum Sum of Values by Dividing Array
// You are given two arrays nums and andValues of length n and m respectively.
// The value of an array is equal to the last element of that array.

// You have to divide nums into m disjoint contiguous subarrays such that for the ith subarray [li, ri], 
// the bitwise AND of the subarray elements is equal to andValues[i], 
// in other words, nums[li] & nums[li + 1] & ... & nums[ri] == andValues[i] for all 1 <= i <= m, where & represents the bitwise AND operator.

// Return the minimum possible sum of the values of the m subarrays nums is divided into. 
// If it is not possible to divide nums into m subarrays satisfying these conditions, return -1.

// Example 1:
// Input: nums = [1,4,3,3,2], andValues = [0,3,3,2]
// Output: 12
// Explanation:
// The only possible way to divide nums is:
// [1,4] as 1 & 4 == 0.
// [3] as the bitwise AND of a single element subarray is that element itself.
// [3] as the bitwise AND of a single element subarray is that element itself.
// [2] as the bitwise AND of a single element subarray is that element itself.
// The sum of the values for these subarrays is 4 + 3 + 3 + 2 = 12.

// Example 2:
// Input: nums = [2,3,5,7,7,7,5], andValues = [0,7,5]
// Output: 17
// Explanation:
// There are three ways to divide nums:
// [[2,3,5],[7,7,7],[5]] with the sum of the values 5 + 7 + 5 == 17.
// [[2,3,5,7],[7,7],[5]] with the sum of the values 7 + 7 + 5 == 19.
// [[2,3,5,7,7],[7],[5]] with the sum of the values 7 + 7 + 5 == 19.
// The minimum possible sum of the values is 17.

// Example 3:
// Input: nums = [1,2,3,4], andValues = [2]
// Output: -1
// Explanation:
// The bitwise AND of the entire array nums is 0. As there is no possible way to divide nums into a single subarray to have the bitwise AND of elements 2, return -1.

// Constraints:
//     1 <= n == nums.length <= 10^4
//     1 <= m == andValues.length <= min(n, 10)
//     1 <= nums[i] < 10^5
//     0 <= andValues[j] < 10^5

import "fmt"

// DP Top Down
func minimumValueSum(nums []int, andValues []int) int {
    cache := make(map[string]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func( i, j, currentValue int) int 
    dfs = func( i, j, currentValue int) int {
        if i == len(nums) && j == len(andValues) { return 0  }
        if i == len(nums) || j == len(andValues) { return -1 }
        currentValue &= nums[i]
        if currentValue < andValues[j] { return -1 }
        key := fmt.Sprintf("%d-%d-%d", i, j, currentValue)
        if v,ok := cache[key]; ok {
            return v
        }
        continueTheArray := dfs(i + 1, j, currentValue)
        //can't break the subarray here, need to continue
        if currentValue != andValues[j] {
            cache[key] = continueTheArray
            return cache[key]
        }
        // break and evaluate
        next := dfs(i + 1, j + 1, ^0)  
        if next == -1 {
            cache[key] = continueTheArray
            return cache[key]
        }
        if continueTheArray == -1 {
            cache[key] = nums[i] + next
            return cache[key]
        }
        cache[key] = min(nums[i] + next, continueTheArray)
        return cache[key]
    }
    return dfs(0, 0, ^0)
}

// 单调队列
func minimumValueSum1(nums, andValues []int) int {
    n, inf := len(nums), 1 << 31
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = inf
    }
    nf := make([]int, n+1)
    for _, target := range andValues {
        type pair struct{ and, l int }
        a := []pair{} // logTrick 子数组 AND 和子数组左端点
        q := []int{}  // 单调队列，保存 dp 的下标
        qi := 0       // 单调队列目前处理到 f[qi]
        nf[0] = inf
        for i, x := range nums {
            for j := range a {
                a[j].and &= x
            }
            a = append(a, pair{x, i})
            // 原地去重
            last, j := -1, 0
            for _, p := range a {
                if p.and >= target && p.and != last {
                    a[j] = p
                    j++
                    last = p.and
                }
            }
            a = a[:j]
            // 去掉无用数据
            for len(a) > 0 && a[0].and < target {
                a = a[1:]
            }
            // 上面这一大段的目的是求出子数组右端点为 i 时，子数组左端点的最小值和最大值
            // 下面是单调队列的滑窗过程
            if len(a) > 0 && a[0].and == target {
                // 现在 a[0].l 和 a[1].l-1 分别是子数组左端点的最小值和最大值
                r := i
                if len(a) > 1 {
                    r = a[1].l - 1
                }
                // 单调队列：右边入
                for ; qi <= r; qi++ {
                    for len(q) > 0 && dp[qi] <= dp[q[len(q)-1]] {
                        q = q[:len(q)-1]
                    }
                    q = append(q, qi)
                }
                // 单调队列：左边出
                for q[0] < a[0].l {
                    q = q[1:]
                }
                // 单调队列：计算答案
                nf[i+1] = dp[q[0]] + x // 队首就是最小值
            } else {
                nf[i+1] = inf
            }
        }
        dp, nf = nf, dp
    }
    if dp[n] < inf {
        return dp[n]
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [1,4,3,3,2], andValues = [0,3,3,2]
    // Output: 12
    // Explanation:
    // The only possible way to divide nums is:
    // [1,4] as 1 & 4 == 0.
    // [3] as the bitwise AND of a single element subarray is that element itself.
    // [3] as the bitwise AND of a single element subarray is that element itself.
    // [2] as the bitwise AND of a single element subarray is that element itself.
    // The sum of the values for these subarrays is 4 + 3 + 3 + 2 = 12.
    fmt.Println(minimumValueSum([]int{1,4,3,3,2},[]int{0,3,3,2})) // 12
    // Example 2:
    // Input: nums = [2,3,5,7,7,7,5], andValues = [0,7,5]
    // Output: 17
    // Explanation:
    // There are three ways to divide nums:
    // [[2,3,5],[7,7,7],[5]] with the sum of the values 5 + 7 + 5 == 17.
    // [[2,3,5,7],[7,7],[5]] with the sum of the values 7 + 7 + 5 == 19.
    // [[2,3,5,7,7],[7],[5]] with the sum of the values 7 + 7 + 5 == 19.
    // The minimum possible sum of the values is 17.
    fmt.Println(minimumValueSum([]int{2,3,5,7,7,7,5},[]int{0,7,5})) // 17
    // Example 3:
    // Input: nums = [1,2,3,4], andValues = [2]
    // Output: -1
    // Explanation:
    // The bitwise AND of the entire array nums is 0. As there is no possible way to divide nums into a single subarray to have the bitwise AND of elements 2, return -1.
    fmt.Println(minimumValueSum([]int{1,2,3,4},[]int{2})) // -1

    fmt.Println(minimumValueSum1([]int{1,4,3,3,2},[]int{0,3,3,2})) // 12
    fmt.Println(minimumValueSum1([]int{2,3,5,7,7,7,5},[]int{0,7,5})) // 17
    fmt.Println(minimumValueSum1([]int{1,2,3,4},[]int{2})) // -1
}