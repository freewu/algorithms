package main

// 3165. Maximum Sum of Subsequence With Non-adjacent Elements
// You are given an array nums consisting of integers. 
// You are also given a 2D array queries, where queries[i] = [posi, xi].

// For query i, we first set nums[posi] equal to xi, 
// then we calculate the answer to query i which is the maximum sum of a subsequence of nums 
// where no two adjacent elements are selected.

// Return the sum of the answers to all queries.

// Since the final answer may be very large, return it modulo 10^9 + 7.

// A subsequence is an array that can be derived from another array by deleting some 
// or no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [3,5,9], queries = [[1,-2],[0,-3]]
// Output: 21
// Explanation:
// After the 1st query, nums = [3,-2,9] and the maximum sum of a subsequence with non-adjacent elements is 3 + 9 = 12.
// After the 2nd query, nums = [-3,-2,9] and the maximum sum of a subsequence with non-adjacent elements is 9.

// Example 2:
// Input: nums = [0,-1], queries = [[0,-5]]
// Output: 0
// Explanation:
// After the 1st query, nums = [-5,-1] and the maximum sum of a subsequence with non-adjacent elements is 0 (choosing an empty subsequence).

// Constraints:
//     1 <= nums.length <= 5 * 10^4
//     -10^5 <= nums[i] <= 10^5
//     1 <= queries.length <= 5 * 10^4
//     queries[i] == [posi, xi]
//     0 <= posi <= nums.length - 1
//     -10^5 <= xi <= 10^5

import "fmt"

// // 超出时间限制 520 / 526 
// func maximumSumSubsequence(nums []int, queries [][]int) int {
//     res, n := 0, len(nums)
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     dp := make([]int, n)
//     dp[0] = max(0, nums[0])
//     if n > 1 {
//         dp[1] = max(dp[0], nums[1])
//     }
//     for i := 2; i < n; i++ {
//         dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
//     }
//     for _, v := range queries {
//         pos, value := v[0], v[1]
//         nums[pos] = value
//         if pos == 0 {
//             dp[0] = max(0, nums[0])
//             if n > 1 {
//                 dp[1] = max(dp[0], nums[1])
//             }
//         } else if pos == 1 {
//             dp[1] = max(dp[0], nums[1])
//         }
//         for i := max(2, pos); i < n; i++ {
//             dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
//         }
//         res = (res + dp[n - 1]) % 1_000_000_007
//     }
//     return res
// }

type data struct {
    f00 int // 第一个数一定不选，最后一个数一定不选
    f01 int // 第一个数一定不选，最后一个数可选可不选
    f10 int // 第一个数可选可不选，最后一个数一定不选
    f11 int // 第一个数可选可不选，最后一个数可选可不选，也就是没有任何限制
}

type seg []data

func (t seg) maintain(o int) {
    a, b := t[o<<1], t[o<<1|1]
    t[o] = data{
        max(a.f00+b.f10, a.f01+b.f00),
        max(a.f00+b.f11, a.f01+b.f01),
        max(a.f10+b.f10, a.f11+b.f00),
        max(a.f10+b.f11, a.f11+b.f01),
    }
}

func (t seg) build(a []int, o, l, r int) {
    if l == r {
        t[o].f11 = max(a[l], 0)
        return
    }
    m := (l + r) >> 1
    t.build(a, o<<1, l, m)
    t.build(a, o<<1|1, m+1, r)
    t.maintain(o)
}

func (t seg) update(o, l, r, i, val int) {
    if l == r {
        t[o].f11 = max(val, 0)
        return
    }
    m := (l + r) >> 1
    if i <= m {
        t.update(o<<1, l, m, i, val)
    } else {
        t.update(o<<1|1, m+1, r, i, val)
    }
    t.maintain(o)
}

// 线段树
func maximumSumSubsequence(nums []int, queries [][]int) int {
    res, n := 0, len(nums)
    t := make(seg, 2 << bits.Len(uint(n-1)))
    t.build(nums, 1, 0, n-1)
    for _, q := range queries {
        t.update(1, 0, n-1, q[0], q[1])
        res += t[1].f11 // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
    }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: nums = [3,5,9], queries = [[1,-2],[0,-3]]
    // Output: 21
    // Explanation:
    // After the 1st query, nums = [3,-2,9] and the maximum sum of a subsequence with non-adjacent elements is 3 + 9 = 12.
    // After the 2nd query, nums = [-3,-2,9] and the maximum sum of a subsequence with non-adjacent elements is 9.
    fmt.Println(maximumSumSubsequence([]int{3,5,9}, [][]int{{1,-2},{0,-3}})) // 21
    // Example 2:
    // Input: nums = [0,-1], queries = [[0,-5]]
    // Output: 0
    // Explanation:
    // After the 1st query, nums = [-5,-1] and the maximum sum of a subsequence with non-adjacent elements is 0 (choosing an empty subsequence).
    fmt.Println(maximumSumSubsequence([]int{0,-1}, [][]int{{0,-5}})) // 0
}