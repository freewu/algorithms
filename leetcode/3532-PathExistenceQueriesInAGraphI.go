package main

// 3532. Path Existence Queries in a Graph I
// You are given an integer n representing the number of nodes in a graph, labeled from 0 to n - 1.

// You are also given an integer array nums of length n sorted in non-decreasing order, and an integer maxDiff.

// An undirected edge exists between nodes i and j if the absolute difference between nums[i] and nums[j] is at most maxDiff (i.e., |nums[i] - nums[j]| <= maxDiff).

// You are also given a 2D integer array queries. 
// For each queries[i] = [ui, vi], determine whether there exists a path between nodes ui and vi.

// Return a boolean array answer, where answer[i] is true if there exists a path between ui and vi in the ith query and false otherwise.

// Example 1:
// Input: n = 2, nums = [1,3], maxDiff = 1, queries = [[0,0],[0,1]]
// Output: [true,false]
// Explanation:
// Query [0,0]: Node 0 has a trivial path to itself.
// Query [0,1]: There is no edge between Node 0 and Node 1 because |nums[0] - nums[1]| = |1 - 3| = 2, which is greater than maxDiff.
// Thus, the final answer after processing all the queries is [true, false].

// Example 2:
// Input: n = 4, nums = [2,5,6,8], maxDiff = 2, queries = [[0,1],[0,2],[1,3],[2,3]]
// Output: [false,false,true,true]
// Explanation:
// The resulting graph is:
// <img src="https://assets.leetcode.com/uploads/2025/03/25/screenshot-2025-03-26-at-122249.png" />
// Query [0,1]: There is no edge between Node 0 and Node 1 because |nums[0] - nums[1]| = |2 - 5| = 3, which is greater than maxDiff.
// Query [0,2]: There is no edge between Node 0 and Node 2 because |nums[0] - nums[2]| = |2 - 6| = 4, which is greater than maxDiff.
// Query [1,3]: There is a path between Node 1 and Node 3 through Node 2 since |nums[1] - nums[2]| = |5 - 6| = 1 and |nums[2] - nums[3]| = |6 - 8| = 2, both of which are within maxDiff.
// Query [2,3]: There is an edge between Node 2 and Node 3 because |nums[2] - nums[3]| = |6 - 8| = 2, which is equal to maxDiff.
// Thus, the final answer after processing all the queries is [false, false, true, true].

// Constraints:
//     1 <= n == nums.length <= 10^5
//     0 <= nums[i] <= 10^5
//     nums is sorted in non-decreasing order.
//     0 <= maxDiff <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i] == [ui, vi]
//     0 <= ui, vi < n

import "fmt"

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    var find func(u int, pa []int) int
    find = func(u int, pa []int) int {
        if pa[u] == u { return u }
        return find(pa[u], pa)
    }
    merge := func(u int, v int, pa []int) {
        u, v = find(u, pa), find(v, pa)
        if u > v {
            u, v = v, u
        }
        pa[v] = u
    }
    binarySearch := func(arr []int, target int) int {
        left, right := 0, len(arr) - 1
        for left <= right {
            mid := left + (right - left) / 2
            if arr[mid] <= target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        return right
    }
    pa := make([]int, n)
    for i := 0; i < n; i++ {
        pa[i] = i
    }
    for i := 0; i < n; i++ {
        index := binarySearch(nums, nums[i] + maxDiff)
        for j := i + 1; j <= index; j++ {
            merge(j, j - 1, pa)
        }
        if index > i {
            i = index - 1
        } else {
            i = index
        }
    }
    m := len(queries)
    res := make([]bool, m)
    for i := 0; i < m; i++ {
        u, v := queries[i][0], queries[i][1]
        if find(u, pa) == find(v, pa) {
            res[i] = true
        } else {
            res[i] = false
        }
    }
    return res
}

func pathExistenceQueries1(n int, nums []int, maxDiff int, queries [][]int) []bool {
    colors := make([]int, n)
    for i, color := 1, 0; i < n; i++ {
        if nums[i] - nums[i - 1] > maxDiff {
            color++
        }
        colors[i] = color
    }
    m := len(queries)
    res := make([]bool, m)
    for i, q := range queries {
        res[i] = (colors[q[0]] == colors[q[1]])
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, nums = [1,3], maxDiff = 1, queries = [[0,0],[0,1]]
    // Output: [true,false]
    // Explanation:
    // Query [0,0]: Node 0 has a trivial path to itself.
    // Query [0,1]: There is no edge between Node 0 and Node 1 because |nums[0] - nums[1]| = |1 - 3| = 2, which is greater than maxDiff.
    // Thus, the final answer after processing all the queries is [true, false].
    fmt.Println(pathExistenceQueries(2, []int{1,3}, 1, [][]int{{0,0},{0,1}})) // [true,false]
    // Example 2:
    // Input: n = 4, nums = [2,5,6,8], maxDiff = 2, queries = [[0,1],[0,2],[1,3],[2,3]]
    // Output: [false,false,true,true]
    // Explanation:
    // The resulting graph is:
    // <img src="https://assets.leetcode.com/uploads/2025/03/25/screenshot-2025-03-26-at-122249.png" />
    // Query [0,1]: There is no edge between Node 0 and Node 1 because |nums[0] - nums[1]| = |2 - 5| = 3, which is greater than maxDiff.
    // Query [0,2]: There is no edge between Node 0 and Node 2 because |nums[0] - nums[2]| = |2 - 6| = 4, which is greater than maxDiff.
    // Query [1,3]: There is a path between Node 1 and Node 3 through Node 2 since |nums[1] - nums[2]| = |5 - 6| = 1 and |nums[2] - nums[3]| = |6 - 8| = 2, both of which are within maxDiff.
    // Query [2,3]: There is an edge between Node 2 and Node 3 because |nums[2] - nums[3]| = |6 - 8| = 2, which is equal to maxDiff.
    // Thus, the final answer after processing all the queries is [false, false, true, true].
    fmt.Println(pathExistenceQueries(4, []int{2,5,6,8}, 2, [][]int{{0,1},{0,2},{1,3},{2,3}})) // [false,false,true,true]
    
    fmt.Println(pathExistenceQueries1(2, []int{1,3}, 1, [][]int{{0,0},{0,1}})) // [true,false]
    fmt.Println(pathExistenceQueries1(4, []int{2,5,6,8}, 2, [][]int{{0,1},{0,2},{1,3},{2,3}})) // [false,false,true,true]
}