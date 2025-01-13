package main

// 2440. Create Components With Same Value
// There is an undirected tree with n nodes labeled from 0 to n - 1.

// You are given a 0-indexed integer array nums of length n where nums[i] represents the value of the ith node. 
// You are also given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// You are allowed to delete some edges, splitting the tree into multiple connected components. 
// Let the value of a component be the sum of all nums[i] for which node i is in the component.

// Return the maximum number of edges you can delete, 
// such that every connected component in the tree has the same value.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/08/26/diagramdrawio.png" />
// Input: nums = [6,2,2,2,6], edges = [[0,1],[1,2],[1,3],[3,4]] 
// Output: 2 
// Explanation: The above figure shows how we can delete the edges [0,1] and [3,4]. The created components are nodes [0], [1,2,3] and [4]. The sum of the values in each component equals 6. It can be proven that no better deletion exists, so the answer is 2.

// Example 2:
// Input: nums = [2], edges = []
// Output: 0
// Explanation: There are no edges to be deleted.

// Constraints:
//     1 <= n <= 2 * 10^4
//     nums.length == n
//     1 <= nums[i] <= 50
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= edges[i][0], edges[i][1] <= n - 1
//     edges represents a valid tree.

import "fmt"

func componentValue(nums []int, edges [][]int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    sum, mx := 0, -1
    for _, v := range nums {
        sum += v
        mx = max(mx, v)
    }
    n := len(nums)
    graph := make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    t := 0
    var dfs func(int, int) int
    dfs = func(i, fa int) int {
        x := nums[i]
        for _, j := range graph[i] {
            if j != fa {
                y := dfs(j, i)
                if y == -1 { return -1 }
                x += y
            }
        }
        if x > t { return -1 }
        if x < t { return x  }
        return 0
    }
    for k := min(n, sum / mx); k > 1; k-- {
        if sum % k == 0 {
            t = sum / k
            if dfs(0, -1) == 0 {
                return k - 1
            }
        }
    }
    return 0
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/08/26/diagramdrawio.png" />
    // Input: nums = [6,2,2,2,6], edges = [[0,1],[1,2],[1,3],[3,4]] 
    // Output: 2 
    // Explanation: The above figure shows how we can delete the edges [0,1] and [3,4]. The created components are nodes [0], [1,2,3] and [4]. The sum of the values in each component equals 6. It can be proven that no better deletion exists, so the answer is 2.
    fmt.Println(componentValue([]int{6,2,2,2,6}, [][]int{{0,1},{1,2},{1,3},{3,4}} )) // 2
    // Example 2:
    // Input: nums = [2], edges = []
    // Output: 0
    // Explanation: There are no edges to be deleted.
    fmt.Println(componentValue([]int{2}, [][]int{})) // 0
}