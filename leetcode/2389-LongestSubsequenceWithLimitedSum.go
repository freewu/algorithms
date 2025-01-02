package main

// 2389. Longest Subsequence With Limited Sum
// You are given an integer array nums of length n, and an integer array queries of length m.

// Return an array answer of length m where answer[i] is the maximum size of a subsequence 
// that you can take from nums such that the sum of its elements is less than or equal to queries[i].

// A subsequence is an array that can be derived from another array by deleting some 
// or no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [4,5,2,1], queries = [3,10,21]
// Output: [2,3,4]
// Explanation: We answer the queries as follows:
// - The subsequence [2,1] has a sum less than or equal to 3. It can be proven that 2 is the maximum size of such a subsequence, so answer[0] = 2.
// - The subsequence [4,5,1] has a sum less than or equal to 10. It can be proven that 3 is the maximum size of such a subsequence, so answer[1] = 3.
// - The subsequence [4,5,2,1] has a sum less than or equal to 21. It can be proven that 4 is the maximum size of such a subsequence, so answer[2] = 4.

// Example 2:
// Input: nums = [2,3,4,5], queries = [1]
// Output: [0]
// Explanation: The empty subsequence is the only subsequence that has a sum less than or equal to 1, so answer[0] = 0.

// Constraints:
//     n == nums.length
//     m == queries.length
//     1 <= n, m <= 1000
//     1 <= nums[i], queries[i] <= 10^6

import "fmt"
import "sort"

func answerQueries(nums []int, queries []int) []int {
    n := len(nums)
    res:= make([]int, len(queries))
    sort.Ints(nums)
    for i := 1; i < n; i++ {
        nums[i] += nums[i - 1]
    }
    for i, query := range queries {
        low, high := 0, n
        for low < high {
            mid := low + ((high - low) >> 1)
            if nums[mid] <= query {
                low = mid + 1
            } else {
                high = mid
            }
        }
        res[i] = low
    }
    return res
}

func answerQueries1(nums []int, queries []int) []int {
    sort.Ints(nums)
    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i-1]
    }
    LowerBound := func(nums []int, target int) int {
        l, r := 0, len(nums) - 1
        for l <= r { // 搜索结果[l,r]
            mid := l + (r-l)/2
            if nums[mid] <= target {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        return l // 最终的结果是 l > r   l = r + 1 所以返回 l 或者 r + 1都可
    }
    for i, query := range queries {
        queries[i] = LowerBound(nums, query)
    }
    return queries
}

func answerQueries2(nums []int, queries []int) []int {
    sort.Ints(nums)
    res := make([]int, 0)
    for i := 0; i < len(queries); i++ {
        total, index := 0, 0
        for j := 0; j < len(nums); j++ {
            total += nums[j]
            if total > queries[i] { break }
            index++
        }
        res = append(res, index)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,5,2,1], queries = [3,10,21]
    // Output: [2,3,4]
    // Explanation: We answer the queries as follows:
    // - The subsequence [2,1] has a sum less than or equal to 3. It can be proven that 2 is the maximum size of such a subsequence, so answer[0] = 2.
    // - The subsequence [4,5,1] has a sum less than or equal to 10. It can be proven that 3 is the maximum size of such a subsequence, so answer[1] = 3.
    // - The subsequence [4,5,2,1] has a sum less than or equal to 21. It can be proven that 4 is the maximum size of such a subsequence, so answer[2] = 4.
    fmt.Println(answerQueries([]int{4,5,2,1}, []int{3,10,21})) // [2,3,4]
    // Example 2:
    // Input: nums = [2,3,4,5], queries = [1]
    // Output: [0]
    // Explanation: The empty subsequence is the only subsequence that has a sum less than or equal to 1, so answer[0] = 0.
    fmt.Println(answerQueries([]int{2,3,4,5}, []int{1})) // [0]

    fmt.Println(answerQueries1([]int{4,5,2,1}, []int{3,10,21})) // [2,3,4]
    fmt.Println(answerQueries1([]int{2,3,4,5}, []int{1})) // [0]

    fmt.Println(answerQueries2([]int{4,5,2,1}, []int{3,10,21})) // [2,3,4]
    fmt.Println(answerQueries2([]int{2,3,4,5}, []int{1})) // [0]
}