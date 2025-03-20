package main

// 3489. Zero Array Transformation IV
// You are given an integer array nums of length n and a 2D array queries, where queries[i] = [li, ri, vali].

// Each queries[i] represents the following action on nums:
//     1. Select a subset of indices in the range [li, ri] from nums.
//     2. Decrement the value at each selected index by exactly vali.

// A Zero Array is an array with all its elements equal to 0.

// Return the minimum possible non-negative value of k, such that after processing the first k queries in sequence, nums becomes a Zero Array. 
// If no such k exists, return -1.

// Example 1:
// Input: nums = [2,0,2], queries = [[0,2,1],[0,2,1],[1,1,3]]
// Output: 2
// Explanation:
// For query 0 (l = 0, r = 2, val = 1):
// Decrement the values at indices [0, 2] by 1.
// The array will become [1, 0, 1].
// For query 1 (l = 0, r = 2, val = 1):
// Decrement the values at indices [0, 2] by 1.
// The array will become [0, 0, 0], which is a Zero Array. Therefore, the minimum value of k is 2.

// Example 2:
// Input: nums = [4,3,2,1], queries = [[1,3,2],[0,2,1]]
// Output: -1
// Explanation:
// It is impossible to make nums a Zero Array even after all the queries.

// Example 3:
// Input: nums = [1,2,3,2,1], queries = [[0,1,1],[1,2,1],[2,3,2],[3,4,1],[4,4,1]]
// Output: 4
// Explanation:
// For query 0 (l = 0, r = 1, val = 1):
// Decrement the values at indices [0, 1] by 1.
// The array will become [0, 1, 3, 2, 1].
// For query 1 (l = 1, r = 2, val = 1):
// Decrement the values at indices [1, 2] by 1.
// The array will become [0, 0, 2, 2, 1].
// For query 2 (l = 2, r = 3, val = 2):
// Decrement the values at indices [2, 3] by 2.
// The array will become [0, 0, 0, 0, 1].
// For query 3 (l = 3, r = 4, val = 1):
// Decrement the value at index 4 by 1.
// The array will become [0, 0, 0, 0, 0]. Therefore, the minimum value of k is 4.

// Example 4:
// Input: nums = [1,2,3,2,6], queries = [[0,1,1],[0,2,1],[1,4,2],[4,4,4],[3,4,1],[4,4,5]]
// Output: 4

// Constraints:
//     1 <= nums.length <= 10
//     0 <= nums[i] <= 1000
//     1 <= queries.length <= 1000
//     queries[i] = [li, ri, vali]
//     0 <= li <= ri < nums.length
//     1 <= vali <= 10

import "fmt"
import "math/big"

func minZeroArray(nums []int, queries [][]int) int {
    res, left, right := -1, 0, len(queries)
    check := func(arr []int, target int) bool {
        n := len(arr)
        if target == 0 { return true } 
        if n == 0 { return false }
        // Subset Sum DP: Check if we can form `target`
        dp := make([]bool, target + 1);
        dp[0] = true;
        for _, v := range arr {
            for j := target; j >= v; j-- {
                dp[j] = dp[j] || dp[j - v]
            }
        }
        return dp[target]
    }
    canMakeZero := func(mid int) bool {
        n := len(nums)
        mp := make(map[int][]int)
        // Apply first mid queries
        for i := 0; i < mid; i++ {
            start, end, val := queries[i][0], queries[i][1], queries[i][2]
            for j := start; j <= end; j++ {
                mp[j] = append(mp[j], val)
            }
        }
        // Check if each index can be made zero using subset sum
        for i := 0; i < n; i++ {
            if !check(mp[i], nums[i]) {
                return false
            }
        }
        return true
    }
    for left <= right {
        mid := (left + right) / 2
        if canMakeZero(mid) {
            res, right = mid, mid - 1
        } else {
            left = mid + 1
        }
    }
    return res
}

func minZeroArray1(nums []int, queries [][]int) int {
    allZero := true
    for _, v := range nums {
        if v > 0 {
            allZero = false
            break
        }
    }
    if allZero { return 0 } // 全部是0 直接返回 0
    f := make([]*big.Int, len(nums))
    for i := range f {
        f[i] = big.NewInt(1)
    }
    p := new(big.Int)
    for k, q := range queries {
        flag, val := false, uint(q[2])
        for i := q[0]; i <= q[1]; i++ {
            if f[i].Bit(nums[i]) == 0 {
                f[i].Or(f[i], p.Lsh(f[i], val))
            }
        }
        for i, v := range nums {
            if f[i].Bit(v) == 0 {
                flag = true
                break
            }
        }
        if flag { continue } // 还没结束还需要走大外循环
        return k + 1
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [2,0,2], queries = [[0,2,1],[0,2,1],[1,1,3]]
    // Output: 2
    // Explanation:
    // For query 0 (l = 0, r = 2, val = 1):
    // Decrement the values at indices [0, 2] by 1.
    // The array will become [1, 0, 1].
    // For query 1 (l = 0, r = 2, val = 1):
    // Decrement the values at indices [0, 2] by 1.
    // The array will become [0, 0, 0], which is a Zero Array. Therefore, the minimum value of k is 2.
    fmt.Println(minZeroArray([]int{2,0,2}, [][]int{{0,2,1},{0,2,1},{1,1,3}})) // 2
    // Example 2:
    // Input: nums = [4,3,2,1], queries = [[1,3,2],[0,2,1]]
    // Output: -1
    // Explanation:
    // It is impossible to make nums a Zero Array even after all the queries.
    fmt.Println(minZeroArray([]int{4,3,2,1}, [][]int{{1,3,2},{0,2,1}})) // -1
    // Example 3:
    // Input: nums = [1,2,3,2,1], queries = [[0,1,1],[1,2,1],[2,3,2],[3,4,1],[4,4,1]]
    // Output: 4
    // Explanation:
    // For query 0 (l = 0, r = 1, val = 1):
    // Decrement the values at indices [0, 1] by 1.
    // The array will become [0, 1, 3, 2, 1].
    // For query 1 (l = 1, r = 2, val = 1):
    // Decrement the values at indices [1, 2] by 1.
    // The array will become [0, 0, 2, 2, 1].
    // For query 2 (l = 2, r = 3, val = 2):
    // Decrement the values at indices [2, 3] by 2.
    // The array will become [0, 0, 0, 0, 1].
    // For query 3 (l = 3, r = 4, val = 1):
    // Decrement the value at index 4 by 1.
    // The array will become [0, 0, 0, 0, 0]. Therefore, the minimum value of k is 4.
    fmt.Println(minZeroArray([]int{1,2,3,2,1}, [][]int{{0,1,1},{1,2,1},{2,3,2},{3,4,1},{4,4,1}})) // 4
    // Example 4:
    // Input: nums = [1,2,3,2,6], queries = [[0,1,1],[0,2,1],[1,4,2],[4,4,4],[3,4,1],[4,4,5]]
    // Output: 4
    fmt.Println(minZeroArray([]int{1,2,3,2,6}, [][]int{{0,1,1},{0,2,1},{1,4,2},{4,4,4},{3,4,1},{4,4,5}})) // 4

    fmt.Println(minZeroArray1([]int{2,0,2}, [][]int{{0,2,1},{0,2,1},{1,1,3}})) // 2
    fmt.Println(minZeroArray1([]int{4,3,2,1}, [][]int{{1,3,2},{0,2,1}})) // -1
    fmt.Println(minZeroArray1([]int{1,2,3,2,1}, [][]int{{0,1,1},{1,2,1},{2,3,2},{3,4,1},{4,4,1}})) // 4
    fmt.Println(minZeroArray1([]int{1,2,3,2,6}, [][]int{{0,1,1},{0,2,1},{1,4,2},{4,4,4},{3,4,1},{4,4,5}})) // 4
}