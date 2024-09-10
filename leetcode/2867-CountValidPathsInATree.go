package main

// 2867. Count Valid Paths in a Tree
// There is an undirected tree with n nodes labeled from 1 to n. 
// You are given the integer n and a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates that there is an edge between nodes ui and vi in the tree.
// Return the number of valid paths in the tree.
// A path (a, b) is valid if there exists exactly one prime number among the node labels in the path from a to b.
// Note that:
//         The path (a, b) is a sequence of distinct nodes starting with node a and ending with node b such that every two adjacent nodes in the sequence share an edge in the tree.
//         Path (a, b) and path (b, a) are considered the same and counted only once.
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/08/27/example1.png"/>
// Input: n = 5, edges = [[1,2],[1,3],[2,4],[2,5]]
// Output: 4
// Explanation: The pairs with exactly one prime number on the path between them are: 
// - (1, 2) since the path from 1 to 2 contains prime number 2. 
// - (1, 3) since the path from 1 to 3 contains prime number 3.
// - (1, 4) since the path from 1 to 4 contains prime number 2.
// - (2, 4) since the path from 2 to 4 contains prime number 2.
// It can be shown that there are only 4 valid paths.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/08/27/example2.png"/>
// Input: n = 6, edges = [[1,2],[1,3],[2,4],[3,5],[3,6]]
// Output: 6
// Explanation: The pairs with exactly one prime number on the path between them are: 
// - (1, 2) since the path from 1 to 2 contains prime number 2.
// - (1, 3) since the path from 1 to 3 contains prime number 3.
// - (1, 4) since the path from 1 to 4 contains prime number 2.
// - (1, 6) since the path from 1 to 6 contains prime number 3.
// - (2, 4) since the path from 2 to 4 contains prime number 2.
// - (3, 6) since the path from 3 to 6 contains prime number 3.
// It can be shown that there are only 6 valid paths.
 
// Constraints:
//         1 <= n <= 10^5
//         edges.length == n - 1
//         edges[i].length == 2
//         1 <= ui, vi <= n
//         The input is generated such that edges represent a valid tree.

import "fmt"

// 埃氏筛
const N = 100001
var is_prime [N]bool
func init() {
    for i := 0; i < N; i++ {
        is_prime[i] = true
    }
    is_prime[1] = false
    for i := 2; i*i < N; i++ {
        if is_prime[i] {
            for j := i * i; j < N; j += i {
                is_prime[j] = false
            }
        }
    }
}

func countPaths(n int, edges [][]int) int64 {
    graph := make([][]int, n + 1)
    for _, edge := range edges {
        i, j := edge[0], edge[1]
        graph[i] = append(graph[i], j)
        graph[j] = append(graph[j], i)
    }
    var dfs func(int, int)
    var seen []int
    dfs = func(i, pre int) {
        seen = append(seen, i)
        for _, j := range graph[i] {
            if j != pre && !is_prime[j] {
                dfs(j, i)
            }
        }
    }
    res := int64(0)
    count := make([]int64, n+1)
    for i := 1; i <= n; i++ {
        if !is_prime[i] {
            continue
        }
        cur := int64(0)
        for _, j := range graph[i] {
            if is_prime[j] {
                continue
            }
            if count[j] == 0 {
                seen = []int{}
                dfs(j, 0)
                cnt := int64(len(seen))
                for _, k := range seen {
                    count[k] = cnt
                }
            }
            res += count[j] * cur
            cur += count[j]
        }
        res += cur
    }
    return res
}

func main() {
    fmt.Println(countPaths(
        5,
        [][]int{ []int{1,2}, []int{1,3}, []int{2,4}, []int{2,5} },
    )) // 4
    fmt.Println(countPaths(
        6,
        [][]int{ []int{1,2}, []int{1,3}, []int{2,4}, []int{3,5}, []int{3,6} },
    )) // 6
}