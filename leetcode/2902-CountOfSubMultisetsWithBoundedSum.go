package main

// 2902. Count of Sub-Multisets With Bounded Sum
// You are given a 0-indexed array nums of non-negative integers, and two integers l and r.

// Return the count of sub-multisets within nums where the sum of elements in each subset falls within the inclusive range of [l, r].

// Since the answer may be large, return it modulo 10^9 + 7.

// A sub-multiset is an unordered collection of elements of the array in which a given value x can occur 0, 1, ..., occ[x] times, where occ[x] is the number of occurrences of x in the array.

// Note that:
//     Two sub-multisets are the same if sorting both sub-multisets results in identical multisets.
//     The sum of an empty multiset is 0.

// Example 1:
// Input: nums = [1,2,2,3], l = 6, r = 6
// Output: 1
// Explanation: The only subset of nums that has a sum of 6 is {1, 2, 3}.

// Example 2:
// Input: nums = [2,1,4,2,7], l = 1, r = 5
// Output: 7
// Explanation: The subsets of nums that have a sum within the range [1, 5] are {1}, {2}, {4}, {2, 2}, {1, 2}, {1, 4}, and {1, 2, 2}.

// Example 3:
// Input: nums = [1,2,1,3,5,2], l = 3, r = 5
// Output: 9
// Explanation: The subsets of nums that have a sum within the range [3, 5] are {3}, {5}, {1, 2}, {1, 3}, {2, 2}, {2, 3}, {1, 1, 2}, {1, 1, 3}, and {1, 2, 2}.

// Constraints:
//     1 <= nums.length <= 2 * 10^4
//     0 <= nums[i] <= 2 * 10^4
//     Sum of nums does not exceed 2 * 10^4.
//     0 <= l <= r <= 2 * 10^4

import "fmt"

// Time Limit Exceeded 632 / 643 
func countSubMultisets(nums []int, l int, r int) int {
    res, mod := 0, 1_000_000_007
    count, memo := make(map[int]int), make([]int, r + 1)
    for _, v := range nums {
        count[v]++
    }
    memo[0] = 1
    for k, v := range count {
        for i := r; i >= 0; i-- {
            ways := 0
            for j := 0; j <= v && i - k * j >= 0; j++ {
                ways = (ways + memo[i - k * j]) % mod
            }
            memo[i] = ways
        }
    }
    for i := l; i <= r; i++ {
        res = (res + memo[i]) % mod
    }
    return res
}

// 式子变形 + 滚动数组
func countSubMultisets1(nums []int, l, r int) int {
    res, total, sum, mod := 0, 0, 0, 1_000_000_007
    count := make(map[int]int)
    for _, v := range nums {
        total += v
        count[v]++
    }
    if l > total { return res }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    r = min(r, total)
    dp := make([]int, r + 1)
    dp[0] = count[0] + 1
    delete(count, 0)
    for k, v := range count {
        arr := append([]int{}, dp...)
        sum = min(sum + k * v, r) // 到目前为止，能选的元素和至多为 sum
        for j := k; j <= sum; j++ { // 把循环上界从 r 改成 sum 可以快不少
            arr[j] += arr[j - k]
            if j >= (v + 1) * k {
                arr[j] -= dp[j-(v + 1) * k] // 注意这里有减法，可能产生负数
            }
            arr[j] %= mod
        }
        dp = arr
    }
    for _, v := range dp[l:] {
        res += v
    }
    return (res % mod + mod) % mod // 调整成非负数
}

// 同余前缀和优化
func countSubMultisets2(nums []int, l, r int) (ans int) {
    res, total, sum, mod := 0, 0, 0, 1_000_000_007
    count := make(map[int]int)
    for _, v := range nums {
        total += v
        count[v]++
    }
    if l > total { return res }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    r = min(r, total)
    dp := make([]int, r + 1)
    dp[0] = count[0] + 1
    delete(count, 0)
    for k, v := range count {
        sum = min(sum + k * v, r)
        for i := k; i <= sum; i++ {
            dp[i] = (dp[i] + dp[i - k]) % mod // 原地计算同余前缀和
        }
        for i := sum; i >= k * (v + 1); i-- {
            dp[i] = (dp[i] - dp[i - k * (v + 1)]) % mod // 两个同余前缀和的差
        }
    }
    for _, v := range dp[l:] {
        res += v
    }
    return (res % mod + mod) % mod // 调整成非负数
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,3], l = 6, r = 6
    // Output: 1
    // Explanation: The only subset of nums that has a sum of 6 is {1, 2, 3}.
    fmt.Println(countSubMultisets([]int{1,2,2,3}, 6, 6)) // 1
    // Example 2:
    // Input: nums = [2,1,4,2,7], l = 1, r = 5
    // Output: 7
    // Explanation: The subsets of nums that have a sum within the range [1, 5] are {1}, {2}, {4}, {2, 2}, {1, 2}, {1, 4}, and {1, 2, 2}.
    fmt.Println(countSubMultisets([]int{2,1,4,2,7}, 1, 5)) // 7
    // Example 3:
    // Input: nums = [1,2,1,3,5,2], l = 3, r = 5
    // Output: 9
    // Explanation: The subsets of nums that have a sum within the range [3, 5] are {3}, {5}, {1, 2}, {1, 3}, {2, 2}, {2, 3}, {1, 1, 2}, {1, 1, 3}, and {1, 2, 2}.
    fmt.Println(countSubMultisets([]int{1,2,1,3,5,2}, 3, 5)) // 9

    fmt.Println(countSubMultisets([]int{1,2,3,4,5,6,7,8,9}, 3, 5)) // 7
    fmt.Println(countSubMultisets([]int{9,8,7,6,5,4,3,2,1}, 3, 5)) // 7

    fmt.Println(countSubMultisets1([]int{1,2,2,3}, 6, 6)) // 1
    fmt.Println(countSubMultisets1([]int{2,1,4,2,7}, 1, 5)) // 7
    fmt.Println(countSubMultisets1([]int{1,2,1,3,5,2}, 3, 5)) // 9
    fmt.Println(countSubMultisets1([]int{1,2,3,4,5,6,7,8,9}, 3, 5)) // 7
    fmt.Println(countSubMultisets1([]int{9,8,7,6,5,4,3,2,1}, 3, 5)) // 7

    fmt.Println(countSubMultisets2([]int{1,2,2,3}, 6, 6)) // 1
    fmt.Println(countSubMultisets2([]int{2,1,4,2,7}, 1, 5)) // 7
    fmt.Println(countSubMultisets2([]int{1,2,1,3,5,2}, 3, 5)) // 9
    fmt.Println(countSubMultisets2([]int{1,2,3,4,5,6,7,8,9}, 3, 5)) // 7
    fmt.Println(countSubMultisets2([]int{9,8,7,6,5,4,3,2,1}, 3, 5)) // 7
}