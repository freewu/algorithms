package main

// 3715. Sum of Perfect Square Ancestors
// You are given an integer n and an undirected tree rooted at node 0 with n nodes numbered from 0 to n - 1. 
// This is represented by a 2D array edges of length n - 1, where edges[i] = [ui, vi] indicates an undirected edge between nodes ui and vi.

// You are also given an integer array nums, where nums[i] is the positive integer assigned to node i.

// Define a value ti as the number of ancestors of node i such that the product nums[i] * nums[ancestor] is a perfect square.

// Return the sum of all ti values for all nodes i in range [1, n - 1].

// Note:
// In a rooted tree, the ancestors of node i are all nodes on the path from node i to the root node 0, excluding i itself.

// Example 1:
// Input: n = 3, edges = [[0,1],[1,2]], nums = [2,8,2]
// Output: 3
// Explanation:
// i	Ancestors	nums[i] * nums[ancestor]	Square Check	ti
// 1	[0]	nums[1] * nums[0] = 8 * 2 = 16	16 is a perfect square	1
// 2	[1, 0]	nums[2] * nums[1] = 2 * 8 = 16
// nums[2] * nums[0] = 2 * 2 = 4	Both 4 and 16 are perfect squares	2
// Thus, the total number of valid ancestor pairs across all non-root nodes is 1 + 2 = 3.

// Example 2:
// Input: n = 3, edges = [[0,1],[0,2]], nums = [1,2,4]
// Output: 1
// Explanation:
// i	Ancestors	nums[i] * nums[ancestor]	Square Check	ti
// 1	[0]	nums[1] * nums[0] = 2 * 1 = 2	2 is not a perfect square	0
// 2	[0]	nums[2] * nums[0] = 4 * 1 = 4	4 is a perfect square	1
// Thus, the total number of valid ancestor pairs across all non-root nodes is 1.

// Example 3:
// Input: n = 4, edges = [[0,1],[0,2],[1,3]], nums = [1,2,9,4]
// Output: 2
// Explanation:
// i	Ancestors	nums[i] * nums[ancestor]	Square Check	ti
// 1	[0]	nums[1] * nums[0] = 2 * 1 = 2	2 is not a perfect square	0
// 2	[0]	nums[2] * nums[0] = 9 * 1 = 9	9 is a perfect square	1
// 3	[1, 0]	nums[3] * nums[1] = 4 * 2 = 8
// nums[3] * nums[0] = 4 * 1 = 4	Only 4 is a perfect square	1
// Thus, the total number of valid ancestor pairs across all non-root nodes is 0 + 1 + 1 = 2.

// Constraints:
//     1 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] = [ui, vi]
//     0 <= ui, vi <= n - 1
//     nums.length == n
//     1 <= nums[i] <= 10^5
//     The input is generated such that edges represents a valid tree.

import "fmt"

// 预处理平方剩余核
const mx = 100_001
var core = [mx]int{}

func init() {
    for i := 1; i < mx; i++ {
        if core[i] == 0 {
            for j := 1; i*j*j < mx; j++ {
                core[i*j*j] = i
            }
        }
    }
}

func sumOfAncestors(n int, edges [][]int, nums []int) int64{
    res, count, graph := 0, make(map[int]int), make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        graph[x] = append(graph[x], y)
        graph[y] = append(graph[y], x)
    }
    var dfs func(x, fa int)
    dfs = func(x, fa int) {
        c := core[nums[x]]
        res += count[c] // 祖先不包括 x 自己
        count[c]++
        for _, y := range graph[x] {
            if y != fa {
                dfs(y, x)
            }
        }
        count[c]-- // 恢复现场
    }
    dfs(0, -1)
    return int64(res)
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1],[1,2]], nums = [2,8,2]
    // Output: 3
    // Explanation:
    // i	Ancestors	nums[i] * nums[ancestor]	Square Check	ti
    // 1	[0]	nums[1] * nums[0] = 8 * 2 = 16	16 is a perfect square	1
    // 2	[1, 0]	nums[2] * nums[1] = 2 * 8 = 16
    // nums[2] * nums[0] = 2 * 2 = 4	Both 4 and 16 are perfect squares	2
    // Thus, the total number of valid ancestor pairs across all non-root nodes is 1 + 2 = 3.
    fmt.Println(sumOfAncestors(3, [][]int{{0,1},{1,2}}, []int{2,8,2})) // 3
    // Example 2:
    // Input: n = 3, edges = [[0,1],[0,2]], nums = [1,2,4]
    // Output: 1
    // Explanation:
    // i	Ancestors	nums[i] * nums[ancestor]	Square Check	ti
    // 1	[0]	nums[1] * nums[0] = 2 * 1 = 2	2 is not a perfect square	0
    // 2	[0]	nums[2] * nums[0] = 4 * 1 = 4	4 is a perfect square	1
    // Thus, the total number of valid ancestor pairs across all non-root nodes is 1.
    fmt.Println(sumOfAncestors(3, [][]int{{0,1},{0,2}}, []int{1,2,4})) // 1
    // Example 3:
    // Input: n = 4, edges = [[0,1],[0,2],[1,3]], nums = [1,2,9,4]
    // Output: 2
    // Explanation:
    // i	Ancestors	nums[i] * nums[ancestor]	Square Check	ti
    // 1	[0]	nums[1] * nums[0] = 2 * 1 = 2	2 is not a perfect square	0
    // 2	[0]	nums[2] * nums[0] = 9 * 1 = 9	9 is a perfect square	1
    // 3	[1, 0]	nums[3] * nums[1] = 4 * 2 = 8
    // nums[3] * nums[0] = 4 * 1 = 4	Only 4 is a perfect square	1
    // Thus, the total number of valid ancestor pairs across all non-root nodes is 0 + 1 + 1 = 2.
    fmt.Println(sumOfAncestors(4, [][]int{{0,1},{0,2},{1,3}}, []int{1,2,9,4})) // 2
}