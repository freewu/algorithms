package main

// 3004. Maximum Subtree of the Same Color
// You are given a 2D integer array edges representing a tree with n nodes, 
// numbered from 0 to n - 1, rooted at node 0, where edges[i] = [ui, vi] means there is an edge between the nodes vi and ui.

// You are also given a 0-indexed integer array colors of size n, where colors[i] is the color assigned to node i.

// We want to find a node v such that every node in the subtree  of v has the same color.

// Return the size of such subtree with the maximum number of nodes possible.
// <img src="https://assets.leetcode.com/static_assets/others/20231216-134026.png">

// Example 1:
// Input: edges = [[0,1],[0,2],[0,3]], colors = [1,1,2,3]
// Output: 1
// Explanation: 
//     Each color is represented as: 1 -> Red, 2 -> Green, 3 -> Blue. 
//     We can see that the subtree rooted at node 0 has children with different colors. 
//     Any other subtree is of the same color and has a size of 1. 
//     Hence, we return 1.

// Example 2:
// Input: edges = [[0,1],[0,2],[0,3]], colors = [1,1,1,1]
// Output: 4
// Explanation: 
//     The whole tree has the same color, and the subtree rooted at node 0 has the most number of nodes which is 4. 
//     Hence, we return 4.
//     <img src="https://assets.leetcode.com/static_assets/others/20231216-134017.png">

// Example 3:
// Input: edges = [[0,1],[0,2],[2,3],[2,4]], colors = [1,2,3,3,3]
// Output: 3
// Explanation: 
//     Each color is represented as: 1 -> Red, 2 -> Green, 3 -> Blue. 
//     We can see that the subtree rooted at node 0 has children with different colors. 
//     Any other subtree is of the same color, but the subtree rooted at node 2 has a size of 3 which is the maximum.
//     Hence, we return 3.

// Constraints:
//     n == edges.length + 1
//     1 <= n <= 5 * 10^4
//     edges[i] == [ui, vi]
//     0 <= ui, vi < n
//     colors.length == n
//     1 <= colors[i] <= 10^5
//     The input is generated such that the graph represented by edges is a tree.

import "fmt"

func maximumSubtreeSize(edges [][]int, colors []int) int {
    mx, n := 0, len(edges)
    graph := make([][]int, n + 1)
    for i := 0; i < n; i++ {
        graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
        graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var  reRoot func(x, parent int) int
    reRoot = func(x, parent int) int { // 以x为根节点，相同颜色的子树的大小 返回0，说明x为根节点的子树不是所有节点的颜色都一样
        color := colors[x] // 当前节点的颜色
        res := 1 // 以x为根节点，相同颜色的子树的大小
        flag := true //以x为根节点的子树，是否所有节点的颜色都一样 true：所有的颜色都一样 | false：不是所有的颜色都一样
        for _, next := range graph[x] {
            if next != parent {
                childSize := reRoot(next, x) // 换根dp
                if 0 == childSize || color != colors[next] {
                    // 以next为根节点的子树，不是所有节点的颜色都一样
                    // 或者x的颜色和next的颜色不一样
                    flag = false
                } else {
                    res += childSize
                }
            }
        }
        if !flag { return 0 }
        mx = max(mx, res)
        return res
    }
    reRoot(0, -1)
    return mx
}

func main() {
    // Example 1:
    // Input: edges = [[0,1],[0,2],[0,3]], colors = [1,1,2,3]
    // Output: 1
    // Explanation: 
    //     Each color is represented as: 1 -> Red, 2 -> Green, 3 -> Blue. 
    //     We can see that the subtree rooted at node 0 has children with different colors. 
    //     Any other subtree is of the same color and has a size of 1. 
    //     Hence, we return 1.
    fmt.Println(maximumSubtreeSize([][]int{{0,1},{0,2},{0,3}}, []int{1,1,2,3})) // 1
    // Example 2:
    // Input: edges = [[0,1],[0,2],[0,3]], colors = [1,1,1,1]
    // Output: 4
    // Explanation: 
    //     The whole tree has the same color, and the subtree rooted at node 0 has the most number of nodes which is 4. 
    //     Hence, we return 4.
    //     <img src="https://assets.leetcode.com/static_assets/others/20231216-134017.png">
    fmt.Println(maximumSubtreeSize([][]int{{0,1},{0,2},{0,3}}, []int{1,1,1,1})) // 4
    // Example 3:
    // Input: edges = [[0,1],[0,2],[2,3],[2,4]], colors = [1,2,3,3,3]
    // Output: 3
    // Explanation: 
    //     Each color is represented as: 1 -> Red, 2 -> Green, 3 -> Blue. 
    //     We can see that the subtree rooted at node 0 has children with different colors. 
    //     Any other subtree is of the same color, but the subtree rooted at node 2 has a size of 3 which is the maximum.
    //     Hence, we return 3.
    fmt.Println(maximumSubtreeSize([][]int{{0,1},{0,2},{2,3},{2,4}}, []int{1,2,3,3,3})) // 3
}