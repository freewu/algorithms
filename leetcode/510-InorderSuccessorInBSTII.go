package main

// 510. Inorder Successor in BST II
// Given a node in a binary search tree, return the in-order successor of that node in the BST. 
// If that node has no in-order successor, return null.

// The successor of a node is the node with the smallest key greater than node.val.

// You will have direct access to the node but not to the root of the tree. 
// Each node will have a reference to its parent node. Below is the definition for Node:
//     class Node {
//         public int val;
//         public Node left;
//         public Node right;
//         public Node parent;
//     }

// Example 1:
// <img src="" />
// Input: tree = [2,1,3], node = 1
// Output: 2
// Explanation: 1's in-order successor node is 2. Note that both the node and the return value is of Node type.

// Example 2:
// <img src="" />
// Input: tree = [5,3,6,2,4,null,null,1], node = 6
// Output: null
// Explanation: There is no in-order successor of the current node, so the answer is null.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5
//     All Nodes will have unique values.

// Follow up: Could you solve it without looking up any of the node's values?

import "fmt"

type Node struct {
    Val int
    Left *Node
    Right *Node
    Parent *Node
}

/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */
func inorderSuccessor(node *Node) *Node {
    var findRoot func(cur *Node)*Node
    findRoot = func(cur *Node) *Node {
        if cur.Parent == nil {
            return cur
        }
        return findRoot(cur.Parent)
    }
    root := findRoot(node) // 递归找到根节点
    var findMostLeft func(cur *Node) *Node
    findMostLeft = func(cur *Node) *Node { // 寻找最左叶子
        if cur == nil || cur.Left == nil { 
            return cur 
        }
        return findMostLeft(cur.Left)
    }
    var dfs func(r *Node,p *Node) *Node
    dfs = func(r *Node,p *Node) *Node {
        if r == nil {
            return r
        }
        if r.Val == p.Val { // 如果 p 就是 r，那么p的后继节点就是p的右子树的最左叶子
            return findMostLeft(r.Right)
        }
        if r.Val < p.Val { // 如果 p 大于r，则调用自身，在 r 的右子树上寻找
            return dfs(r.Right,p)
        }
        t := dfs(r.Left, p) // 如果 p 小于r，则调用自身，在r的左子树上寻找。
        if t == nil {
            return r
        }
        return t
    }
    return dfs(root, node)
}

func inorderSuccessor1(node *Node) *Node {
    if node == nil {
        return nil
    }
    cur := node.Right
    if cur != nil {
        for cur.Left != nil {
            cur = cur.Left
        }
        return cur
    }
    cur = node
    for cur.Parent != nil {
        cur = cur.Parent
        if cur.Val > node.Val {
            return cur
        }
    }
    return nil
}

func main() {
    // Example 1:
    //      2
    //    /   \
    //   1     3
    // <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_1.PNG" />
    // Input: tree = [2,1,3], node = 1
    // Output: 2
    // Explanation: 1's in-order successor node is 2. Note that both the node and the return value is of Node type.

    // Example 2:
    //            5
    //          /   \
    //         3     6
    //       /   \
    //      2     4
    //     /
    //    1
    // <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_2.PNG" />
    // Input: tree = [5,3,6,2,4,null,null,1], node = 6
    // Output: null
    // Explanation: There is no in-order successor of the current node, so the answer is null.
    fmt.Println()
}

