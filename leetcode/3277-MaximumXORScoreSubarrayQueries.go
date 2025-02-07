package main

// 3277. Maximum XOR Score Subarray Queries
// You are given an array nums of n integers, and a 2D integer array queries of size q, where queries[i] = [li, ri].

// For each query, you must find the maximum XOR score of any subarray of nums[li..ri].

// The XOR score of an array a is found by repeatedly applying the following operations on a so that only one element remains, 
// that is the score:
//     1. Simultaneously replace a[i] with a[i] XOR a[i + 1] for all indices i except the last one.
//     2. Remove the last element of a.

// Return an array answer of size q where answer[i] is the answer to query i.

// Example 1:
// Input: nums = [2,8,4,32,16,1], queries = [[0,2],[1,4],[0,5]]
// Output: [12,60,60]
// Explanation:
// In the first query, nums[0..2] has 6 subarrays [2], [8], [4], [2, 8], [8, 4], and [2, 8, 4] each with a respective XOR score of 2, 8, 4, 10, 12, and 6. The answer for the query is 12, the largest of all XOR scores.
// In the second query, the subarray of nums[1..4] with the largest XOR score is nums[1..4] with a score of 60.
// In the third query, the subarray of nums[0..5] with the largest XOR score is nums[1..4] with a score of 60.

// Example 2:
// Input: nums = [0,7,3,2,8,5,1], queries = [[0,3],[1,5],[2,4],[2,6],[5,6]]
// Output: [7,14,11,14,5]
// Explanation:
// Index	nums[li..ri]	Maximum XOR Score Subarray	Maximum Subarray XOR Score
// 0	[0, 7, 3, 2]	[7]	7
// 1	[7, 3, 2, 8, 5]	[7, 3, 2, 8]	14
// 2	[3, 2, 8]	[3, 2, 8]	11
// 3	[3, 2, 8, 5, 1]	[2, 8, 5, 1]	14
// 4	[5, 1]	[5]	5

// Constraints:
//     1 <= n == nums.length <= 2000
//     0 <= nums[i] <= 2^31 - 1
//     1 <= q == queries.length <= 10^5
//     queries[i].length == 2 
//     queries[i] = [li, ri]
//     0 <= li <= ri <= n - 1

import "fmt"

func maximumSubarrayXor(nums []int, queries [][]int) []int {
    res, n := []int{}, len(nums)
    precompute, best := make([][]int, n), make([][]int, n)
    for i := range precompute {
        precompute[i], best[i] = make([]int, n), make([]int, n)
        precompute[i][i], best[i][i] = nums[i], nums[i]
    }
    // fill from bottom row 
    for r := n - 1 ; r >= 0 ; r -- {
        for c := r + 1 ; c < n ; c ++ {
            precompute[r][c] =  precompute[r][c - 1] ^ precompute[r + 1][c]
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // fill from bottom row 
    for r := n - 1 ; r >= 0 ; r -- {
        for c := r + 1 ; c < n ; c ++ {
            best[r][c] = max(max(best[r][c - 1], best[r + 1][c]), precompute[r][c])
        }
    }
    for _, v := range queries {
        res = append(res,  best[v[0]][v[1]])
    }
    return res
}

func maximumSubarrayXor1(nums []int, queries [][]int) []int {
    n := len(nums)
    xorNums := make([][]int, n) // xorNums[i][j] i到j的数组异或值
    for i := range xorNums {
        xorNums[i] = make([]int, n)
        xorNums[i][i] = nums[i]
    }
    for i := n - 1 - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            xorNums[i][j] = xorNums[i][j-1] ^ xorNums[i+1][j]
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // i到j的所有子数组 最大值
    dp := xorNums
    for i := n - 1 - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            dp[i][j] = max(dp[i][j-1], max(dp[i+1][j], dp[i][j]))
        }
    }
    res := make([]int, len(queries))
    for i, v := range queries {
        res[i] = dp[v[0]][v[1]]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,8,4,32,16,1], queries = [[0,2],[1,4],[0,5]]
    // Output: [12,60,60]
    // Explanation:
    // In the first query, nums[0..2] has 6 subarrays [2], [8], [4], [2, 8], [8, 4], and [2, 8, 4] each with a respective XOR score of 2, 8, 4, 10, 12, and 6. The answer for the query is 12, the largest of all XOR scores.
    // In the second query, the subarray of nums[1..4] with the largest XOR score is nums[1..4] with a score of 60.
    // In the third query, the subarray of nums[0..5] with the largest XOR score is nums[1..4] with a score of 60.
    fmt.Println(maximumSubarrayXor([]int{2,8,4,32,16,1}, [][]int{{0,2},{1,4},{0,5}})) // [12,60,60]
    // Example 2:
    // Input: nums = [0,7,3,2,8,5,1], queries = [[0,3],[1,5],[2,4],[2,6],[5,6]]
    // Output: [7,14,11,14,5]
    // Explanation:
    // Index	nums[li..ri]	Maximum XOR Score Subarray	Maximum Subarray XOR Score
    // 0	[0, 7, 3, 2]	[7]	7
    // 1	[7, 3, 2, 8, 5]	[7, 3, 2, 8]	14
    // 2	[3, 2, 8]	[3, 2, 8]	11
    // 3	[3, 2, 8, 5, 1]	[2, 8, 5, 1]	14
    // 4	[5, 1]	[5]	5
    fmt.Println(maximumSubarrayXor([]int{0,7,3,2,8,5,1}, [][]int{{0,3},{1,5},{2,4},{2,6},{5,6}})) // [7,14,11,14,5]

    fmt.Println(maximumSubarrayXor([]int{1,2,3,4,5,6,7,8,9}, [][]int{{0,3},{1,5},{2,4},{2,6},{5,6}})) // [7 7 7 7 7]
    fmt.Println(maximumSubarrayXor([]int{9,8,7,6,5,4,3,2,1}, [][]int{{0,3},{1,5},{2,4},{2,6},{5,6}})) // [15 15 7 7 7]

    fmt.Println(maximumSubarrayXor1([]int{2,8,4,32,16,1}, [][]int{{0,2},{1,4},{0,5}})) // [12,60,60]
    fmt.Println(maximumSubarrayXor1([]int{0,7,3,2,8,5,1}, [][]int{{0,3},{1,5},{2,4},{2,6},{5,6}})) // [7,14,11,14,5]
    fmt.Println(maximumSubarrayXor1([]int{1,2,3,4,5,6,7,8,9}, [][]int{{0,3},{1,5},{2,4},{2,6},{5,6}})) // [7 7 7 7 7]
    fmt.Println(maximumSubarrayXor1([]int{9,8,7,6,5,4,3,2,1}, [][]int{{0,3},{1,5},{2,4},{2,6},{5,6}})) // [15 15 7 7 7]
}