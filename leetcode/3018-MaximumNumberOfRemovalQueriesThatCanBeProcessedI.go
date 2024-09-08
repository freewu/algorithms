package main

// 3018. Maximum Number of Removal Queries That Can Be Processed I
// You are given a 0-indexed array nums and a 0-indexed array queries.
// You can do the following operation at the beginning at most once:
//     Replace nums with a subsequence of nums.

// We start processing queries in the given order; for each query, we do the following:
//     If the first and the last element of nums is less than queries[i], the processing of queries ends.
//     Otherwise, we choose either the first or the last element of nums if it is greater than or equal to queries[i], and we remove the chosen element from nums.

// Return the maximum number of queries that can be processed by doing the operation optimally.

// Example 1:
// Input: nums = [1,2,3,4,5], queries = [1,2,3,4,6]
// Output: 4
// Explanation: We don't do any operation and process the queries as follows:
// 1- We choose and remove nums[0] since 1 <= 1, then nums becomes [2,3,4,5].
// 2- We choose and remove nums[0] since 2 <= 2, then nums becomes [3,4,5].
// 3- We choose and remove nums[0] since 3 <= 3, then nums becomes [4,5].
// 4- We choose and remove nums[0] since 4 <= 4, then nums becomes [5].
// 5- We can not choose any elements from nums since they are not greater than or equal to 5.
// Hence, the answer is 4.
// It can be shown that we can't process more than 4 queries.

// Example 2:
// Input: nums = [2,3,2], queries = [2,2,3]
// Output: 3
// Explanation: We don't do any operation and process the queries as follows:
// 1- We choose and remove nums[0] since 2 <= 2, then nums becomes [3,2].
// 2- We choose and remove nums[1] since 2 <= 2, then nums becomes [3].
// 3- We choose and remove nums[0] since 3 <= 3, then nums becomes [].
// Hence, the answer is 3.
// It can be shown that we can't process more than 3 queries.

// Example 3:
// Input: nums = [3,4,3], queries = [4,3,2]
// Output: 2
// Explanation: First we replace nums with the subsequence of nums [4,3].
// Then we can process the queries as follows:
// 1- We choose and remove nums[0] since 4 <= 4, then nums becomes [3].
// 2- We choose and remove nums[0] since 3 <= 3, then nums becomes [].
// 3- We can not process any more queries since nums is empty.
// Hence, the answer is 2.
// It can be shown that we can't process more than 2 queries.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= queries.length <= 1000
//     1 <= nums[i], queries[i] <= 10^9

import "fmt"

func maximumProcessableQueries(nums []int, queries []int) int {
    res, n, m := 0, len(nums), len(queries)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for l := 0; l < n; l++ {
        for r := n - 1; r >= l; r-- {
            mx := 0
            if l > 0 {
                index := dp[l - 1][r]
                if (nums[l-1] >= queries[index]) {
                    mx = max(mx, index + 1)
                } else {
                    mx = max(mx, index)
                }
            }
            if r < n - 1 {
                index := dp[l][r+1]
                if nums[r+1] >= queries[index] {
                    mx = max(mx, index + 1)
                } else {
                    mx = max(mx, index)
                }
            }
            if mx == m { return m }
            dp[l][r] = mx
        }
    }
    for i := 0; i < n; i++ {
        index := dp[i][i]
        if nums[i] >= queries[index] {
            res = max(res, index + 1)
        } else {
            res = max(res, index)
        }
    }
    return res 
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], queries = [1,2,3,4,6]
    // Output: 4
    // Explanation: We don't do any operation and process the queries as follows:
    // 1- We choose and remove nums[0] since 1 <= 1, then nums becomes [2,3,4,5].
    // 2- We choose and remove nums[0] since 2 <= 2, then nums becomes [3,4,5].
    // 3- We choose and remove nums[0] since 3 <= 3, then nums becomes [4,5].
    // 4- We choose and remove nums[0] since 4 <= 4, then nums becomes [5].
    // 5- We can not choose any elements from nums since they are not greater than or equal to 5.
    // Hence, the answer is 4.
    // It can be shown that we can't process more than 4 queries.
    fmt.Println(maximumProcessableQueries([]int{1,2,3,4,5}, []int{1,2,3,4,6})) // 4
    // Example 2:
    // Input: nums = [2,3,2], queries = [2,2,3]
    // Output: 3
    // Explanation: We don't do any operation and process the queries as follows:
    // 1- We choose and remove nums[0] since 2 <= 2, then nums becomes [3,2].
    // 2- We choose and remove nums[1] since 2 <= 2, then nums becomes [3].
    // 3- We choose and remove nums[0] since 3 <= 3, then nums becomes [].
    // Hence, the answer is 3.
    // It can be shown that we can't process more than 3 queries.
    fmt.Println(maximumProcessableQueries([]int{2,3,2}, []int{2,2,3})) // 3
    // Example 3:
    // Input: nums = [3,4,3], queries = [4,3,2]
    // Output: 2
    // Explanation: First we replace nums with the subsequence of nums [4,3].
    // Then we can process the queries as follows:
    // 1- We choose and remove nums[0] since 4 <= 4, then nums becomes [3].
    // 2- We choose and remove nums[0] since 3 <= 3, then nums becomes [].
    // 3- We can not process any more queries since nums is empty.
    // Hence, the answer is 2.
    // It can be shown that we can't process more than 2 queries.
    fmt.Println(maximumProcessableQueries([]int{3,4,3}, []int{4,3,2})) // 2
}