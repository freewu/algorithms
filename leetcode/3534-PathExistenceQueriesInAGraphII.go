package main

// 3534. Path Existence Queries in a Graph II
// You are given an integer n representing the number of nodes in a graph, labeled from 0 to n - 1.

// You are also given an integer array nums of length n and an integer maxDiff.

// An undirected edge exists between nodes i and j if the absolute difference between nums[i] and nums[j] is at most maxDiff (i.e., |nums[i] - nums[j]| <= maxDiff).

// You are also given a 2D integer array queries. 
// For each queries[i] = [ui, vi], find the minimum distance between nodes ui and vi. 
// If no path exists between the two nodes, return -1 for that query.

// Return an array answer, where answer[i] is the result of the ith query.

// Note: The edges between the nodes are unweighted.

// Example 1:
// Input: n = 5, nums = [1,8,3,4,2], maxDiff = 3, queries = [[0,3],[2,4]]
// Output: [1,1]
// Explanation:
// The resulting graph is:
// <img src="https://assets.leetcode.com/uploads/2025/03/25/4149example1drawio.png" />
// Query	Shortest Path	Minimum Distance
// [0, 3]	0 → 3	1
// [2, 4]	2 → 4	1
// Thus, the output is [1, 1].

// Example 2:
// Input: n = 5, nums = [5,3,1,9,10], maxDiff = 2, queries = [[0,1],[0,2],[2,3],[4,3]]
// Output: [1,2,-1,1]
// Explanation:
// The resulting graph is:
// <img src="https://assets.leetcode.com/uploads/2025/03/25/4149example2drawio.png" />
// Query	Shortest Path	Minimum Distance
// [0, 1]	0 → 1	1
// [0, 2]	0 → 1 → 2	2
// [2, 3]	None	-1
// [4, 3]	3 → 4	1
// Thus, the output is [1, 2, -1, 1].

// Example 3:
// Input: n = 3, nums = [3,6,1], maxDiff = 1, queries = [[0,0],[0,1],[1,2]]
// Output: [0,-1,-1]
// Explanation:
// There are no edges between any two nodes because:
// Nodes 0 and 1: |nums[0] - nums[1]| = |3 - 6| = 3 > 1
// Nodes 0 and 2: |nums[0] - nums[2]| = |3 - 1| = 2 > 1
// Nodes 1 and 2: |nums[1] - nums[2]| = |6 - 1| = 5 > 1
// Thus, no node can reach any other node, and the output is [0, -1, -1].

// Constraints:
//     1 <= n == nums.length <= 10^5
//     0 <= nums[i] <= 10^5
//     0 <= maxDiff <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i] == [ui, vi]
//     0 <= ui, vi < n

import "fmt"
// import "sort"

// Time Limit Exceeded 678 / 682 
// func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
//     type Elem struct { Value, Index int }
//     values := make([]Elem, n)
//     for i, v := range nums {
//         values[i] = Elem{ Value: v, Index: i}
//     }
//     sort.Slice(values, func(i, j int) bool {
//         return values[i].Value < values[j].Value
//     })
//     index := make(map[int]int)
//     for i, e := range values {
//         index[e.Index] = i
//     }
//     links := make([]int, n)
//     for i := 0; i < n; i++ {
//         target := values[i].Value + maxDiff
//         j := sort.Search(n, func(k int) bool {
//             return values[k].Value > target
//         }) - 1
//         if j < 0 {
//             j = 0
//         }
//         links[i] = j
//     }
//     res := make([]int, len(queries))
//     for i, q := range queries {
//         u, ok0 := index[q[0]]
//         v, ok1 := index[q[1]]
//         if !ok0 || !ok1 {
//             res[i] = -1
//             continue
//         }
//         if u > v {
//             u, v = v, u
//         }
//         steps, current := 0, u
//         for current < v {
//             next := links[current]
//             if next <= current {
//                 steps = -1
//                 break
//             }
//             steps++
//             current = next
//         }
//         res[i] = steps
//         if res[i] == -1 {
//             res[i] = -1
//         }
//     }
//     return res
// }

import "slices"
import "math/bits"

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
    index := make([]int, n)
    for i := range index {
        index[i] = i
    }
    slices.SortFunc(index, func(i, j int) int { 
        return nums[i] - nums[j] 
    })
    rank := make([]int, n) // rank[i] 表示 nums[i] 是 nums 中的第几小，或者说节点 i 在 idx 中的下标
    for i, j := range index {
        rank[j] = i
    }
    // 双指针，从第 i 小的数开始，向左一步，最远能跳到第 left 小的数
    pa := make([][]int, n)
    left, mx := 0, bits.Len(uint(n))
    for i, j := range index {
        for nums[j] - nums[index[left]] > maxDiff {
            left++
        }
        pa[i] = make([]int, mx)
        pa[i][0] = left
    }
    for i := 0; i < mx - 1; i++ { // 倍增
        for x := range pa {
            p := pa[x][i]
            pa[x][i + 1] = pa[p][i]
        }
    }
    res := make([]int, len(queries))
    for i, q := range queries {
        l, r := q[0], q[1]
        if l == r { continue } // 不用跳
        l, r = rank[l], rank[r]
        if l > r { // 保证 l 在 r 左边
            l, r = r, l
        }
        // 从 r 开始，向左跳到 l
        v := 0
        for k := mx - 1; k >= 0; k-- {
            if pa[r][k] > l {
                v |= 1 << k
                r = pa[r][k]
            }
        }
        if pa[r][0] > l { // 无法跳到 l
            res[i] = -1
        } else {
            res[i] = v + 1 // 再跳一步就能到 l
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5, nums = [1,8,3,4,2], maxDiff = 3, queries = [[0,3],[2,4]]
    // Output: [1,1]
    // Explanation:
    // The resulting graph is:
    // <img src="https://assets.leetcode.com/uploads/2025/03/25/4149example1drawio.png" />
    // Query	Shortest Path	Minimum Distance
    // [0, 3]	0 → 3	1
    // [2, 4]	2 → 4	1
    // Thus, the output is [1, 1].
    fmt.Println(pathExistenceQueries(5,[]int{1,8,3,4,2}, 3, [][]int{{0,3},{2,4}})) // [1,1]
    // Example 2:
    // Input: n = 5, nums = [5,3,1,9,10], maxDiff = 2, queries = [[0,1],[0,2],[2,3],[4,3]]
    // Output: [1,2,-1,1]
    // Explanation:
    // The resulting graph is:
    // <img src="https://assets.leetcode.com/uploads/2025/03/25/4149example2drawio.png" />
    // Query	Shortest Path	Minimum Distance
    // [0, 1]	0 → 1	1
    // [0, 2]	0 → 1 → 2	2
    // [2, 3]	None	-1
    // [4, 3]	3 → 4	1
    // Thus, the output is [1, 2, -1, 1].
    fmt.Println(pathExistenceQueries(5,[]int{5,3,1,9,10}, 2, [][]int{{0,1},{0,2},{2,3},{4,3}})) // [1,2,-1,1]
    // Example 3:
    // Input: n = 3, nums = [3,6,1], maxDiff = 1, queries = [[0,0],[0,1],[1,2]]
    // Output: [0,-1,-1]
    // Explanation:
    // There are no edges between any two nodes because:
    // Nodes 0 and 1: |nums[0] - nums[1]| = |3 - 6| = 3 > 1
    // Nodes 0 and 2: |nums[0] - nums[2]| = |3 - 1| = 2 > 1
    // Nodes 1 and 2: |nums[1] - nums[2]| = |6 - 1| = 5 > 1
    // Thus, no node can reach any other node, and the output is [0, -1, -1].
    fmt.Println(pathExistenceQueries(3,[]int{3,6,1}, 1, [][]int{{0,0},{0,1},{1,2}})) //  [0,-1,-1]
}