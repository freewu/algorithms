package main

// 面试题 04.06. Successor LCCI
// Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a binary search tree.

// Return null if there's no "next" node for the given node.

// Example 1:
// Input: root = [2,1,3], p = 1
//   2
//  / \
// 1   3
// Output: 2

// Example 2:
// Input: root = [5,3,6,2,4,null,null,1], p = 6
//       5
//      / \
//     3   6
//    / \
//   2   4
//  /   
// 1
// Output: null

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
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    found := false
    var res *TreeNode;
    var dfs func(node *TreeNode, p *TreeNode)
    dfs = func(node *TreeNode, p *TreeNode) {
        if node == nil { return }
        dfs(node.Left, p)
        if found && res == nil {
            res = node
            return // 找到下一个节点, 直接返回
        } else if node.Val == p.Val { // 找到了
            found = true
        }
        dfs(node.Right, p)
    }
    dfs(root, p)
    return res
}

func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
    var last, res *TreeNode
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil { return }
        dfs(node.Left)
        if last == p && res == nil {
            res = node
            return
        }
        last = node
        dfs(node.Right)
        return
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    // Input: root = [2,1,3], p = 1
    //   2
    //  / \
    // 1   3
    // Output: 2
    tree1 := &TreeNode {
        2,
        &TreeNode{1, nil, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(inorderSuccessor(tree1, &TreeNode{1, nil, nil, })) // &{2 0xc000110048 0xc000110060}
    // Example 2:
    // Input: root = [5,3,6,2,4,null,null,1], p = 6
    //       5
    //      / \
    //     3   6
    //    / \
    //   2   4
    //  /   
    // 1
    // Output: null
    tree2 := &TreeNode {
        5,
        &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, &TreeNode{4, nil, nil, }, },
        &TreeNode{6, nil, nil, },
    }
    fmt.Println(inorderSuccessor(tree2, &TreeNode{6, nil, nil, })) // <nil>

    fmt.Println(inorderSuccessor1(tree1, &TreeNode{1, nil, nil, })) // &{2 0xc000110048 0xc000110060}
    fmt.Println(inorderSuccessor1(tree2, &TreeNode{6, nil, nil, })) // <nil>
}