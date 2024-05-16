package main

// 2331. Evaluate Boolean Binary Tree
// You are given the root of a full binary tree with the following properties:
//     Leaf nodes have either the value 0 or 1, where 0 represents False and 1 represents True.
//     Non-leaf nodes have either the value 2 or 3, where 2 represents the boolean OR and 3 represents the boolean AND.

// The evaluation of a node is as follows:
//     If the node is a leaf node, the evaluation is the value of the node, i.e. True or False.
//     Otherwise, evaluate the node's two children and apply the boolean operation of its value with the children's evaluations.

// Return the boolean result of evaluating the root node.
// A full binary tree is a binary tree where each node has either 0 or 2 children.
// A leaf node is a node that has zero children.

// Example 1:
//         2(or)                          2(or)                 1(true)
//         /   \                          /   \
//     1(true) 3(and)         =>    1(true)   0(false)   => 
//             /   \
//         0(false) 1(true)
// <img src="https://assets.leetcode.com/uploads/2022/05/16/example1drawio1.png" />
// Input: root = [2,1,3,null,null,0,1]
// Output: true
// Explanation: The above diagram illustrates the evaluation process.
// The AND node evaluates to False AND True = False.
// The OR node evaluates to True OR False = True.
// The root node evaluates to True, so we return true.

// Example 2:
// Input: root = [0]
// Output: false
// Explanation: The root node is a leaf node and it evaluates to false, so we return false.

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     0 <= Node.val <= 3
//     Every node has either 0 or 2 children.
//     Leaf nodes have a value of 0 or 1.
//     Non-leaf nodes have a value of 2 or 3.

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
func evaluateTree1(root *TreeNode) bool {
    var dfs func(root *TreeNode) bool 
    dfs = func(node *TreeNode) bool {
        if node.Left == nil && node.Right == nil { // 叶子节点
            return node.Val == 1 // 1 true / 0 false
        }
        if node.Val == 2 { // or 的处理
            return dfs(node.Left) || dfs(node.Right)
        }
        return dfs(node.Left) && dfs(node.Right) // 为 3 and 的处理
    }
    return dfs(root)
}

func evaluateTree(root *TreeNode) bool {
    switch root.Val {
    case 0: // 叶子节点
        return false
    case 1: // 叶子节点
        return true
    case 2: // or
        return evaluateTree(root.Left) || evaluateTree(root.Right)
    case 3: // and
        return evaluateTree(root.Left) && evaluateTree(root.Right)
    }
    return false
}

func main() {
    // Example 1:
    //         2(or)                          2(or)                 1(true)
    //         /   \                          /   \
    //     1(true) 3(and)         =>    1(true)   0(false)   => 
    //             /   \
    //         0(false) 1(true)
    // <img src="https://assets.leetcode.com/uploads/2022/05/16/example1drawio1.png" />
    // Input: root = [2,1,3,null,null,0,1]
    // Output: true
    // Explanation: The above diagram illustrates the evaluation process.
    // The AND node evaluates to False AND True = False.
    // The OR node evaluates to True OR False = True.
    // The root node evaluates to True, so we return true.
    tree1 := &TreeNode {
        2,
        &TreeNode { 1, nil, nil },
        &TreeNode { 3, &TreeNode { 1, nil, nil }, &TreeNode { 0, nil, nil } },
    }
    fmt.Println(evaluateTree(tree1)) // true
    // Example 2:
    // Input: root = [0]
    // Output: false
    // Explanation: The root node is a leaf node and it evaluates to false, so we return false.
    tree2 := &TreeNode { 0, nil, nil }
    fmt.Println(evaluateTree(tree2)) // false

    fmt.Println(evaluateTree1(tree1)) // true
    fmt.Println(evaluateTree1(tree2)) // false
}