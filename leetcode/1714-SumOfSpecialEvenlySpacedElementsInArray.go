package main

// 1714. Sum Of Special Evenly-Spaced Elements In Array
// You are given a 0-indexed integer array nums consisting of n non-negative integers.

// You are also given an array queries, where queries[i] = [xi, yi]. 
// The answer to the ith query is the sum of all nums[j] where xi <= j < n and (j - xi) is divisible by yi.

// Return an array answer where answer.length == queries.length and answer[i] is the answer to the ith query modulo 10^9 + 7.

// Example 1:
// Input: nums = [0,1,2,3,4,5,6,7], queries = [[0,3],[5,1],[4,2]]
// Output: [9,18,10]
// Explanation: The answers of the queries are as follows:
// 1) The j indices that satisfy this query are 0, 3, and 6. nums[0] + nums[3] + nums[6] = 9
// 2) The j indices that satisfy this query are 5, 6, and 7. nums[5] + nums[6] + nums[7] = 18
// 3) The j indices that satisfy this query are 4 and 6. nums[4] + nums[6] = 10

// Example 2:
// Input: nums = [100,200,101,201,102,202,103,203], queries = [[0,7]]
// Output: [303]

// Constraints:
//     n == nums.length
//     1 <= n <= 5 * 10^4
//     0 <= nums[i] <= 10^9
//     1 <= queries.length <= 1.5 * 10^5
//     0 <= xi < n
//     1 <= yi <= 5 * 10^4

import "fmt"

func solve(nums []int, queries [][]int) []int {
    n, b, mod := len(nums), 0, 1_000_000_007
    for ; b * b < n; b++ {}
    cache := make([][]int, b)
    for i := 0; i < b; i++ {
        cache[i] = make([]int, n)
    }
    // 预计算每个下标的每隔i个数的前缀和
    for i := 1; i < b; i++ {
        for j := n - 1; j >= 0; j-- {
            if n-1 < i+j {
                cache[i][j] = nums[j]
            } else {
                cache[i][j] = nums[j] + cache[i][j+i]
                cache[i][j] %= mod
            }
        }
    }
    calc := func(nums []int, x,y int) int {
        res := 0
        for i := x; i < len(nums); i += y {
            res += nums[i]
            res %= mod
        }
        return res
    }
    res := make([]int, len(queries))
    for i, query := range queries {
        x, y := query[0], query[1]
        if y >= b { // 当 y 大于 根号n  时我们采用暴力算法
            res[i] = calc(nums,x,y)
        } else { // 小于根号 n 时，我们预计算好的
            res[i] = cache[y][x]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,2,3,4,5,6,7], queries = [[0,3],[5,1],[4,2]]
    // Output: [9,18,10]
    // Explanation: The answers of the queries are as follows:
    // 1) The j indices that satisfy this query are 0, 3, and 6. nums[0] + nums[3] + nums[6] = 9
    // 2) The j indices that satisfy this query are 5, 6, and 7. nums[5] + nums[6] + nums[7] = 18
    // 3) The j indices that satisfy this query are 4 and 6. nums[4] + nums[6] = 10
    fmt.Println(solve([]int{0,1,2,3,4,5,6,7}, [][]int{{0,3},{5,1},{4,2}})) // [9,18,10]
    // Example 2:
    // Input: nums = [100,200,101,201,102,202,103,203], queries = [[0,7]]
    // Output: [303]
    fmt.Println(solve([]int{100,200,101,201,102,202,103,203}, [][]int{{0,7}})) // [303]
}