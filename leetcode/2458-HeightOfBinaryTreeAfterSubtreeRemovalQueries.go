package main

// 2458. Height of Binary Tree After Subtree Removal Queries
// You are given the root of a binary tree with n nodes. 
// Each node is assigned a unique value from 1 to n. 
// You are also given an array queries of size m.

// You have to perform m independent queries on the tree where in the ith query you do the following:
//     Remove the subtree rooted at the node with the value queries[i] from the tree. 
//     It is guaranteed that queries[i] will not be equal to the value of the root.

// Return an array answer of size m where answer[i] is the height of the tree after performing the ith query.

// Note:
//     The queries are independent, so the tree returns to its initial state after each query.
//     The height of a tree is the number of edges in the longest simple path from the root to some node in the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/07/binaryytreeedrawio-1.png" />
// Input: root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries = [4]
// Output: [2]
// Explanation: The diagram above shows the tree after removing the subtree rooted at node with value 4.
// The height of the tree is 2 (The path 1 -> 3 -> 2).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/09/07/binaryytreeedrawio-2.png" />
// Input: root = [5,8,9,2,1,3,7,4,6], queries = [3,2,4,8]
// Output: [3,2,3,2]
// Explanation: We have the following queries:
// - Removing the subtree rooted at node with value 3. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 4).
// - Removing the subtree rooted at node with value 2. The height of the tree becomes 2 (The path 5 -> 8 -> 1).
// - Removing the subtree rooted at node with value 4. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 6).
// - Removing the subtree rooted at node with value 8. The height of the tree becomes 2 (The path 5 -> 9 -> 3).

// Constraints:
//     The number of nodes in the tree is n.
//     2 <= n <= 10^5
//     1 <= Node.val <= n
//     All the values in the tree are unique.
//     m == queries.length
//     1 <= m <= min(n, 10^4)
//     1 <= queries[i] <= n
//     queries[i] != root.val

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
func treeQueries(root *TreeNode, queries []int) []int {
    depthMap, heightMap, h1Map, h2Map := make(map[int]int),make(map[int]int),make(map[int]int), make(map[int]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var calculateHeightDepth func(root *TreeNode, depth int) int 
    calculateHeightDepth = func(root *TreeNode, depth int) int {
        if root == nil {
            return 0
        }
        depthMap[root.Val] = depth
        leftHeight, rightHeight := calculateHeightDepth(root.Left, depth + 1), calculateHeightDepth(root.Right, depth + 1)
        heightMap[root.Val] = max(leftHeight, rightHeight) + 1
        if h1Map[depth] < heightMap[root.Val] {
            h2Map[depth] = h1Map[depth]
            h1Map[depth] = heightMap[root.Val]
        } else if h2Map[depth] < heightMap[root.Val] {
            h2Map[depth] = heightMap[root.Val]
        }
        return heightMap[root.Val]
    }
    calculateHeightDepth(root, 0)
    res := []int{}
    for _, query := range queries {
        depth := depthMap[query]
        if h1Map[depth] == heightMap[query] {
            res = append(res, depth + h2Map[depth]-1)
        } else {
            res = append(res, depth + h1Map[depth]-1)
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/09/07/binaryytreeedrawio-1.png" />
    // Input: root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries = [4]
    // Output: [2]
    // Explanation: The diagram above shows the tree after removing the subtree rooted at node with value 4.
    // The height of the tree is 2 (The path 1 -> 3 -> 2).
    tree1 := &TreeNode {
        1,
        &TreeNode { 3 , &TreeNode { 2 ,nil, nil }, nil },
        &TreeNode { 4 , &TreeNode { 6 ,nil, nil },  &TreeNode { 5 ,nil, &TreeNode { 7 ,nil, nil }, } },
    }
    fmt.Println(treeQueries(tree1,[]int{4})) // [2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/09/07/binaryytreeedrawio-2.png" />
    // Input: root = [5,8,9,2,1,3,7,4,6], queries = [3,2,4,8]
    // Output: [3,2,3,2]
    // Explanation: We have the following queries:
    // - Removing the subtree rooted at node with value 3. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 4).
    // - Removing the subtree rooted at node with value 2. The height of the tree becomes 2 (The path 5 -> 8 -> 1).
    // - Removing the subtree rooted at node with value 4. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 6).
    // - Removing the subtree rooted at node with value 8. The height of the tree becomes 2 (The path 5 -> 9 -> 3).
    tree2 := &TreeNode {
        5,
        &TreeNode { 8 , &TreeNode { 2 , &TreeNode { 4 ,nil, nil }, &TreeNode { 6 ,nil, nil }, },  &TreeNode { 1 ,nil, nil }, },
        &TreeNode { 9 , &TreeNode { 3 ,nil, nil },  &TreeNode { 7 ,nil, nil }, },
    }
    fmt.Println(treeQueries(tree2,[]int{3,2,4,8})) // [3,2,3,2]
}