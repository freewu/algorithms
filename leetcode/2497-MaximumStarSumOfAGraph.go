package main

// 2497. Maximum Star Sum of a Graph
// There is an undirected graph consisting of n nodes numbered from 0 to n - 1. 
// You are given a 0-indexed integer array vals of length n where vals[i] denotes the value of the ith node.

// You are also given a 2D integer array edges where edges[i] = [ai, bi] denotes 
// that there exists an undirected edge connecting nodes ai and bi.

// A star graph is a subgraph of the given graph having a center node containing 0 or more neighbors. 
// In other words, it is a subset of edges of the given graph such that there exists a common node for all edges.

// The image below shows star graphs with 3 and 4 neighbors respectively, centered at the blue node.
// <img src="https://assets.leetcode.com/uploads/2022/11/07/max-star-sum-descdrawio.png" />

// The star sum is the sum of the values of all the nodes present in the star graph.

// Given an integer k, return the maximum star sum of a star graph containing at most k edges.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/11/07/max-star-sum-example1drawio.png" />
// Input: vals = [1,2,3,4,10,-10,-20], edges = [[0,1],[1,2],[1,3],[3,4],[3,5],[3,6]], k = 2
// Output: 16
// Explanation: The above diagram represents the input graph.
// The star graph with the maximum star sum is denoted by blue. It is centered at 3 and includes its neighbors 1 and 4.
// It can be shown it is not possible to get a star graph with a sum greater than 16.

// Example 2:
// Input: vals = [-5], edges = [], k = 0
// Output: -5
// Explanation: There is only one possible star graph, which is node 0 itself.
// Hence, we return -5.

// Constraints:
//     n == vals.length
//     1 <= n <= 10^5
//     -10^4 <= vals[i] <= 10^4
//     0 <= edges.length <= min(n * (n - 1) / 2, 10^5)
//     edges[i].length == 2
//     0 <= ai, bi <= n - 1
//     ai != bi
//     0 <= k <= n - 1

import "fmt"
import "sort"
import "container/heap"

func maxStarSum(vals []int, edges [][]int, k int) int {
    adjList := make([][]int, len(vals))
    for _, edge := range edges { // 创建邻接表
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
        adjList[edge[1]] = append(adjList[edge[1]], edge[0])
    }
    res := vals[0]
    for index, list := range adjList {
        currentSum := vals[index]
        neighbourValues := make([]int, 0)
        for _, neighbours := range list {
            neighbourValues = append(neighbourValues, vals[neighbours])
        }
        sort.Ints(neighbourValues)
        for i := len(neighbourValues) - 1; i >= 0; i-- {
            if ((len(neighbourValues) - 1) - i) == k {
                break
            }
            if (currentSum + neighbourValues[i]) > currentSum {
                currentSum += neighbourValues[i]
            }
        }
        if currentSum > res {
            res = currentSum
        }
    }
    return res
}

type MinHeap []int16

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int16)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func maxStarSum1(vals []int, edges [][]int, k int) int {
    if k == 0 {
        sort.Ints(vals)
        return vals[len(vals)-1]
    }
    connections := make([]MinHeap, len(vals))
    for _, edge := range edges {
        for i := 0; i <= 1; i++ {
            if v := int16(vals[edge[1-i]]); v > 0 {
                if len(connections[edge[i]]) < k {
                    heap.Push(&connections[edge[i]], v)
                } else if v > connections[edge[i]][0] {
                    connections[edge[i]][0] = v
                    heap.Fix(&connections[edge[i]], 0)
                }
            }
        }
    }
    res := -10000
    for starIdx, sum := range vals {
        for i := 0; i < len(connections[starIdx]); i++ {
            sum += int(connections[starIdx][i])
        }
        if sum > res {
            res = sum
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/11/07/max-star-sum-example1drawio.png" />
    // Input: vals = [1,2,3,4,10,-10,-20], edges = [[0,1],[1,2],[1,3],[3,4],[3,5],[3,6]], k = 2
    // Output: 16
    // Explanation: The above diagram represents the input graph.
    // The star graph with the maximum star sum is denoted by blue. It is centered at 3 and includes its neighbors 1 and 4.
    // It can be shown it is not possible to get a star graph with a sum greater than 16.
    fmt.Println(maxStarSum([]int{1,2,3,4,10,-10,-20}, [][]int{{0,1},{1,2},{1,3},{3,4},{3,5},{3,6}}, 2)) // 16
    // Example 2:
    // Input: vals = [-5], edges = [], k = 0
    // Output: -5
    // Explanation: There is only one possible star graph, which is node 0 itself.
    // Hence, we return -5.
    fmt.Println(maxStarSum([]int{-5}, [][]int{}, 0)) // -5

    fmt.Println(maxStarSum1([]int{1,2,3,4,10,-10,-20}, [][]int{{0,1},{1,2},{1,3},{3,4},{3,5},{3,6}}, 2)) // 16
    fmt.Println(maxStarSum1([]int{-5}, [][]int{}, 0)) // -5
}