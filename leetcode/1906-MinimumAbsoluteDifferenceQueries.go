package main

// 1906. Minimum Absolute Difference Queries
// The minimum absolute difference of an array a is defined as the minimum value of |a[i] - a[j]|, where 0 <= i < j < a.length and a[i] != a[j]. 
// If all elements of a are the same, the minimum absolute difference is -1.
//     For example, the minimum absolute difference of the array [5,2,3,7,2] is |2 - 3| = 1. 
//     Note that it is not 0 because a[i] and a[j] must be different.

// You are given an integer array nums and the array queries where queries[i] = [li, ri]. 
// For each query i, compute the minimum absolute difference of the subarray nums[li...ri] containing the elements of nums between the 0-based indices li and ri (inclusive).

// Return an array ans where ans[i] is the answer to the ith query.

// A subarray is a contiguous sequence of elements in an array.

// The value of |x| is defined as:
//     x if x >= 0.
//     -x if x < 0.

// Example 1:
// Input: nums = [1,3,4,8], queries = [[0,1],[1,2],[2,3],[0,3]]
// Output: [2,1,4,1]
// Explanation: The queries are processed as follows:
// - queries[0] = [0,1]: The subarray is [1,3] and the minimum absolute difference is |1-3| = 2.
// - queries[1] = [1,2]: The subarray is [3,4] and the minimum absolute difference is |3-4| = 1.
// - queries[2] = [2,3]: The subarray is [4,8] and the minimum absolute difference is |4-8| = 4.
// - queries[3] = [0,3]: The subarray is [1,3,4,8] and the minimum absolute difference is |3-4| = 1.

// Example 2:
// Input: nums = [4,5,2,2,7,10], queries = [[2,3],[0,2],[0,5],[3,5]]
// Output: [-1,1,1,3]
// Explanation: The queries are processed as follows:
// - queries[0] = [2,3]: The subarray is [2,2] and the minimum absolute difference is -1 because all the
//   elements are the same.
// - queries[1] = [0,2]: The subarray is [4,5,2] and the minimum absolute difference is |4-5| = 1.
// - queries[2] = [0,5]: The subarray is [4,5,2,2,7,10] and the minimum absolute difference is |4-5| = 1.
// - queries[3] = [3,5]: The subarray is [2,7,10] and the minimum absolute difference is |7-10| = 3.

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 100
//     1 <= queries.length <= 2 * 10^4
//     0 <= li < ri < nums.length

import "fmt"
import "sort"

func minDifference(nums []int, queries [][]int) []int {
    res, arr := make([]int, len(queries)), make([][]int, 100)
    for i, v := range nums {
        v-- 
        arr[v] = append(arr[v], i)
    }
    search := func(arr []int, from, to int) bool {
        n := len(arr)
        if n == 0 { return false }
        low := sort.Search(n, func(i int) bool { return from <= arr[i] })
        if low == n { return false }
        high := sort.Search(n - low, func(i int) bool { return arr[i + low] > to })
        return high != 0
    }
    for i, query := range queries {
        from, to, last, mn := query[0], query[1], -1, 100
        for j := range arr {
            if search(arr[j], from, to) {
                if last != -1 && j - last < mn {
                    mn = j - last
                }
                last = j
            }
        }
        if mn == 100 {
            res[i] = -1
        } else {
            res[i] = mn
        }
    } 
    return res
}

func minDifference1(nums []int, queries [][]int) []int {
    calc := func(arr1, arr2 [101]int) int {
        prev, res := 0, 200
        for i := 1; i <= 100; i++ {
            if arr2[i] - arr1[i] == 0 { continue }
            if prev != 0 {
                res = min(res, i - prev)
            }
            prev = i
        }
        if res == 200 { return -1 }
        return res
    }
    n := len(nums)
    prefix := make([][101]int, n + 1)
    for i, v := range nums {
        prefix[i + 1] = prefix[i]
        prefix[i + 1][v]++
    }   
    res := make([]int, len(queries))
    for i, v := range queries {
        res[i] = calc(prefix[v[0]], prefix[v[1] + 1])
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,4,8], queries = [[0,1],[1,2],[2,3],[0,3]]
    // Output: [2,1,4,1]
    // Explanation: The queries are processed as follows:
    // - queries[0] = [0,1]: The subarray is [1,3] and the minimum absolute difference is |1-3| = 2.
    // - queries[1] = [1,2]: The subarray is [3,4] and the minimum absolute difference is |3-4| = 1.
    // - queries[2] = [2,3]: The subarray is [4,8] and the minimum absolute difference is |4-8| = 4.
    // - queries[3] = [0,3]: The subarray is [1,3,4,8] and the minimum absolute difference is |3-4| = 1.
    fmt.Println(minDifference([]int{1,3,4,8}, [][]int{{0,1},{1,2},{2,3},{0,3}})) // [2,1,4,1]
    // Example 2:
    // Input: nums = [4,5,2,2,7,10], queries = [[2,3],[0,2],[0,5],[3,5]]
    // Output: [-1,1,1,3]
    // Explanation: The queries are processed as follows:
    // - queries[0] = [2,3]: The subarray is [2,2] and the minimum absolute difference is -1 because all the
    //   elements are the same.
    // - queries[1] = [0,2]: The subarray is [4,5,2] and the minimum absolute difference is |4-5| = 1.
    // - queries[2] = [0,5]: The subarray is [4,5,2,2,7,10] and the minimum absolute difference is |4-5| = 1.
    // - queries[3] = [3,5]: The subarray is [2,7,10] and the minimum absolute difference is |7-10| = 3.
    fmt.Println(minDifference([]int{4,5,2,2,7,10}, [][]int{{2,3},{0,2},{0,5},{3,5}})) // [-1,1,1,3]

    fmt.Println(minDifference1([]int{1,3,4,8}, [][]int{{0,1},{1,2},{2,3},{0,3}})) // [2,1,4,1]
    fmt.Println(minDifference1([]int{4,5,2,2,7,10}, [][]int{{2,3},{0,2},{0,5},{3,5}})) // [-1,1,1,3]
}