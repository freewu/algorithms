package main 

// 2538. Difference Between Maximum and Minimum Price Sum
// There exists an undirected and initially unrooted tree with n nodes indexed from 0 to n - 1. 
// You are given the integer n and a 2D integer array edges of length n - 1, 
// where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// Each node has an associated price. 
// You are given an integer array price, where price[i] is the price of the ith node.

// The price sum of a given path is the sum of the prices of all nodes lying on that path.

// The tree can be rooted at any node root of your choice. 
// The incurred cost after choosing root is the difference between the maximum and minimum price sum amongst all paths starting at root.

// Return the maximum possible cost amongst all possible root choices.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/12/01/example14.png" />
// Input: n = 6, edges = [[0,1],[1,2],[1,3],[3,4],[3,5]], price = [9,8,7,6,10,5]
// Output: 24
// Explanation: The diagram above denotes the tree after rooting it at node 2. The first part (colored in red) shows the path with the maximum price sum. The second part (colored in blue) shows the path with the minimum price sum.
// - The first path contains nodes [2,1,3,4]: the prices are [7,8,6,10], and the sum of the prices is 31.
// - The second path contains the node [2] with the price [7].
// The difference between the maximum and minimum price sum is 24. It can be proved that 24 is the maximum cost.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/11/24/p1_example2.png" />
// Input: n = 3, edges = [[0,1],[1,2]], price = [1,1,1]
// Output: 2
// Explanation: The diagram above denotes the tree after rooting it at node 0. The first part (colored in red) shows the path with the maximum price sum. The second part (colored in blue) shows the path with the minimum price sum.
// - The first path contains nodes [0,1,2]: the prices are [1,1,1], and the sum of the prices is 3.
// - The second path contains node [0] with a price [1].
// The difference between the maximum and minimum price sum is 2. It can be proved that 2 is the maximum cost.

// Constraints:
//     1 <= n <= 10^5
//     edges.length == n - 1
//     0 <= ai, bi <= n - 1
//     edges represents a valid tree.
//     price.length == n
//     1 <= price[i] <= 10^5

import "fmt"

func maxOutput(n int, edges [][]int, price []int) int64 {
    res, graph := 0, make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i, fa int) (int, int)
    dfs = func(i, fa int) (int, int) {
        a, b := price[i], 0
        for _, j := range graph[i] {
            if j != fa {
                c, d := dfs(j, i)
                res = max(res, max(a + d, b + c))
                a, b = max(a, price[i] + c), max(b, price[i] + d)
            }
        }
        return a, b
    }
    dfs(0, -1)
    return int64(res)
}

func maxOutput1(n int, edges [][]int, price []int) int64 {
    res, graph := 0, make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int, int) (int, int)
    dfs = func(x, fa int) (int, int) {
        l, r, p := 0, price[x],  price[x]
        for _, v := range graph[x] {
            if v == fa { continue }
            lm, rx := dfs(v, x)
            res = max(res, max(lm + r, l + rx))
            l, r = max(l, lm + p), max(r, rx + p)
        }
        return l, r
    }
    dfs(0,-1)
    return int64(res)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/12/01/example14.png" />
    // Input: n = 6, edges = [[0,1],[1,2],[1,3],[3,4],[3,5]], price = [9,8,7,6,10,5]
    // Output: 24
    // Explanation: The diagram above denotes the tree after rooting it at node 2. The first part (colored in red) shows the path with the maximum price sum. The second part (colored in blue) shows the path with the minimum price sum.
    // - The first path contains nodes [2,1,3,4]: the prices are [7,8,6,10], and the sum of the prices is 31.
    // - The second path contains the node [2] with the price [7].
    // The difference between the maximum and minimum price sum is 24. It can be proved that 24 is the maximum cost.
    fmt.Println(maxOutput(6, [][]int{{0,1},{1,2},{1,3},{3,4},{3,5}}, []int{9,8,7,6,10,5})) // 24
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/11/24/p1_example2.png" />
    // Input: n = 3, edges = [[0,1],[1,2]], price = [1,1,1]
    // Output: 2
    // Explanation: The diagram above denotes the tree after rooting it at node 0. The first part (colored in red) shows the path with the maximum price sum. The second part (colored in blue) shows the path with the minimum price sum.
    // - The first path contains nodes [0,1,2]: the prices are [1,1,1], and the sum of the prices is 3.
    // - The second path contains node [0] with a price [1].
    // The difference between the maximum and minimum price sum is 2. It can be proved that 2 is the maximum cost.
    fmt.Println(maxOutput(3, [][]int{{0,1},{1,2}}, []int{1,1,1})) // 2

    fmt.Println(maxOutput1(6, [][]int{{0,1},{1,2},{1,3},{3,4},{3,5}}, []int{9,8,7,6,10,5})) // 24
    fmt.Println(maxOutput1(3, [][]int{{0,1},{1,2}}, []int{1,1,1})) // 2
}