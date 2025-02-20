package main

// 2049. Count Nodes With the Highest Score
// There is a binary tree rooted at 0 consisting of n nodes. 
// The nodes are labeled from 0 to n - 1. 
// You are given a 0-indexed integer array parents representing the tree, where parents[i] is the parent of node i. 
// Since node 0 is the root, parents[0] == -1.

// Each node has a score. To find the score of a node, consider if the node and the edges connected to it were removed. 
// The tree would become one or more non-empty subtrees. 
// The size of a subtree is the number of the nodes in it. 
// The score of the node is the product of the sizes of all those subtrees.

// Return the number of nodes that have the highest score.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/03/example-1.png" />
// Input: parents = [-1,2,0,2,0]
// Output: 3
// Explanation:
// - The score of node 0 is: 3 * 1 = 3
// - The score of node 1 is: 4 = 4
// - The score of node 2 is: 1 * 1 * 2 = 2
// - The score of node 3 is: 4 = 4
// - The score of node 4 is: 4 = 4
// The highest score is 4, and three nodes (node 1, node 3, and node 4) have the highest score.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/03/example-2.png" />
// Input: parents = [-1,2,0]
// Output: 2
// Explanation:
// - The score of node 0 is: 2 = 2
// - The score of node 1 is: 2 = 2
// - The score of node 2 is: 1 * 1 = 1
// The highest score is 2, and two nodes (node 0 and node 1) have the highest score.

// Constraints:
//     n == parents.length
//     2 <= n <= 10^5
//     parents[0] == -1
//     0 <= parents[i] <= n - 1 for i != 0
//     parents represents a valid binary tree.

import "fmt"

func countHighestScoreNodes(parents []int) int {
    score, count, n := 0, 0, len(parents)
    tree := make([][]int, n)
    for i, p := range parents {
        if p < 0 { continue }
        tree[p] = append(tree[p], i)
    }
    var dfs func(node int) int 
    dfs = func(node int)  int { // dfs return tree size of node
        sum, product := 1, 1
        for _, child := range tree[node] {
            v := dfs(child)
            sum += v
            product *= v
        }
        if n - sum > 0 {
            product *= (n - sum)
        }
        if product == score {
            count++
        }
        if product > score {
            score, count = product, 1
        }
        return sum
    }
    dfs(0)
    return count
}

func countHighestScoreNodes1(parents []int) int {
    res, mx, n := 0, -1_000_000_000, len(parents)
    adjList := make([][]int, n)
    for i := 1; i < n; i++ {
        adjList[parents[i]] = append(adjList[parents[i]], i)
    }
    var dfs func(curr, n int) int
    dfs = func(curr, n int) int {
        l, r, score := 0, 0, 1
        if len(adjList[curr]) > 0 { l = dfs(adjList[curr][0], n) }
        if len(adjList[curr]) > 1 { r = dfs(adjList[curr][1], n) }
        if l > 0 { score *= l }
        if r > 0 { score *= r }
        if n - l - r - 1 > 0 { score *= (n - l - r - 1) }
        if score == mx {
            res++
        } else if score > mx {
            mx, res = score, 1
        }
        return 1 + l + r
    }
    dfs(0, n)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/03/example-1.png" />
    // Input: parents = [-1,2,0,2,0]
    // Output: 3
    // Explanation:
    // - The score of node 0 is: 3 * 1 = 3
    // - The score of node 1 is: 4 = 4
    // - The score of node 2 is: 1 * 1 * 2 = 2
    // - The score of node 3 is: 4 = 4
    // - The score of node 4 is: 4 = 4
    // The highest score is 4, and three nodes (node 1, node 3, and node 4) have the highest score.
    fmt.Println(countHighestScoreNodes([]int{-1,2,0,2,0})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/03/example-2.png" />
    // Input: parents = [-1,2,0]
    // Output: 2
    // Explanation:
    // - The score of node 0 is: 2 = 2
    // - The score of node 1 is: 2 = 2
    // - The score of node 2 is: 1 * 1 = 1
    // The highest score is 2, and two nodes (node 0 and node 1) have the highest score.
    fmt.Println(countHighestScoreNodes([]int{-1,2,0})) // 2

    fmt.Println(countHighestScoreNodes1([]int{-1,2,0,2,0})) // 3
    fmt.Println(countHighestScoreNodes1([]int{-1,2,0})) // 2
}