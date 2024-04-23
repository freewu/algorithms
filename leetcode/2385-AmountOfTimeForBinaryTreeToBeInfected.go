package main

// 2385. Amount of Time for Binary Tree to Be Infected
// You are given the root of a binary tree with unique values, and an integer start.
// At minute 0, an infection starts from the node with value start.

// Each minute, a node becomes infected if:
//     The node is currently uninfected.
//     The node is adjacent to an infected node.

// Return the number of minutes needed for the entire tree to be infected.

// Example 1:
//                      1
//                    /   \
//                   5    [3]
//                    \   / \
//                     4 10  6
//                    / \
//                   9   2
// <img src="https://assets.leetcode.com/uploads/2022/06/25/image-20220625231744-1.png" />
// Input: root = [1,5,3,null,4,10,6,9,2], start = 3
// Output: 4
// Explanation: The following nodes are infected during:
// - Minute 0: Node 3
// - Minute 1: Nodes 1, 10 and 6
// - Minute 2: Node 5
// - Minute 3: Node 4
// - Minute 4: Nodes 9 and 2
// It takes 4 minutes for the whole tree to be infected so we return 4.

// Example 2:
// Input: root = [1], start = 1
// Output: 0
// Explanation: At minute 0, the only node in the tree is infected so we return 0.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 10^5
//     Each node has a unique value.
//     A node with a value of start exists in the tree.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// dfs
func amountOfTime(root *TreeNode, start int) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node *TreeNode, start int) (int, int)
    dfs = func(node *TreeNode, start int) (int, int) {
        if node == nil { return 0, 0 }
        rDay, rTotal := dfs(node.Right, start)
        lDay, lTotal := dfs(node.Left, start)
        switch {
            case node.Val == start:
                return 1, max(rTotal, lTotal)
            case rDay > 0:
                return rDay + 1, max(rDay + lTotal, rTotal)
            case lDay > 0:
                return lDay + 1, max(lDay + rTotal, lTotal)
            default:
                return -1, max(lTotal, rTotal) + 1
        }
    }
    _, res := dfs(root, start)
    return res
}

// bfs
func amountOfTime1(root *TreeNode, start int) int {
    graph := make(map[int][]int)
    var buildGraph func (node *TreeNode, parent int, graph map[int][]int) 
    buildGraph = func (node *TreeNode, parent int, graph map[int][]int) {
        if node == nil { return }
        if parent != -1 {
            graph[node.Val] = append(graph[node.Val], parent)
            graph[parent] = append(graph[parent], node.Val)
        }
        buildGraph(node.Left, node.Val, graph)
        buildGraph(node.Right, node.Val, graph)
    }
    bfs := func (graph map[int][]int, start int) int {
        visited := make(map[int]bool)
        queue := []int{ start }
        mins := -1
        for len(queue) > 0 {
            n := len(queue)
            for i := 0; i < n; i++ {
                node := queue[0]
                queue = queue[1:]
                if visited[node] {
                    continue
                }
                visited[node] = true
                for _, adjacent := range graph[node] {
                    if !visited[adjacent] {
                        queue = append(queue, adjacent)
                    }
                }
            }
            mins++
        }
        return mins
    }
    buildGraph(root, -1, graph)
    return bfs(graph, start)
}


func main() {
    // Explanation: The following nodes are infected during:
    // - Minute 0: Node 3
    // - Minute 1: Nodes 1, 10 and 6
    // - Minute 2: Node 5
    // - Minute 3: Node 4
    // - Minute 4: Nodes 9 and 2
    // It takes 4 minutes for the whole tree to be infected so we return 4.
    tree1 := &TreeNode {
        1,
        &TreeNode {
            5,
            nil,
            &TreeNode { 4 ,&TreeNode { 9 ,nil, nil }, &TreeNode { 2 ,nil, nil } },
        },
        &TreeNode { 3 ,&TreeNode { 10 ,nil, nil }, &TreeNode { 6 ,nil, nil } },
    }
    fmt.Println(amountOfTime(tree1, 3)) // 4
    // Explanation: At minute 0, the only node in the tree is infected so we return 0.
    tree2 := &TreeNode { 1 ,nil, nil }
    fmt.Println(amountOfTime(tree2, 1)) // 0

    fmt.Println(amountOfTime1(tree1, 3)) // 4
    fmt.Println(amountOfTime1(tree2, 1)) // 0
}