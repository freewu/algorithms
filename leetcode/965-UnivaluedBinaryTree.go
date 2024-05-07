package main

// 965. Univalued Binary Tree
// A binary tree is uni-valued if every node in the tree has the same value.
// Given the root of a binary tree, return true if the given tree is uni-valued, or false otherwise.

// Example 1:
//        1
//       /  \
//      1    1
//     /  \   \ 
//    1    1   1
// <img src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_1.png" />
// Input: root = [1,1,1,1,1,null,1]
// Output: true

// Example 2:
//        2
//       /  \
//      2    2
//     /  \ 
//    5    2 
// <img src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_2.png" />
// Input: root = [2,2,2,5,2]
// Output: false

// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
//     0 <= Node.val < 100

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
// bfs
func isUnivalTree(root *TreeNode) bool {
    path := []*TreeNode{root}
    s := root.Val
    for i,j := 0,0 ; i<len(path) ; i = j {
        h := len(path)
        for ;j < h; j++ {
            if path[j].Val != s { return false }
            if path[j].Left != nil {
                if path[j].Left.Val != s { return false }
                path = append(path, path[j].Left)
            }
            if path[j].Right != nil {
                if path[j].Right.Val != s { return false }
                path = append(path, path[j].Right)
            }
        }
    }
    return true
}

// dfs
func isUnivalTree1(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var dfs func(node *TreeNode, val int) bool
    dfs = func(node *TreeNode, val int) bool {
        if node == nil {
            return true
        } else if node.Val != val {
            return false
        }
        if node.Left != nil && !dfs(node.Left, val) {
            return false
        }
        if node.Right != nil && !dfs(node.Right, val) {
            return false
        }
        return true
    }
    if root.Left != nil && !dfs(root.Left, root.Val) {
        return false
    }
    if root.Right != nil && !dfs(root.Right, root.Val) {
        return false
    }
    return true
}

func main() {
    // Example 1:
    //        1
    //       /  \
    //      1    1
    //     /  \   \ 
    //    1    1   1
    // <img src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_1.png" />
    // Input: root = [1,1,1,1,1,null,1]
    // Output: true
    tree1 := &TreeNode {
        1,
        &TreeNode{
            1, 
            &TreeNode{1, nil, nil},
            &TreeNode{1, nil, nil},
        },
        &TreeNode{1, nil, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(isUnivalTree(tree1)) // true
    // Example 2:
    //        2
    //       /  \
    //      2    2
    //     /  \ 
    //    5    2 
    // <img src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_2.png" />
    // Input: root = [2,2,2,5,2]
    // Output: false
    tree2 := &TreeNode {
        2,
        &TreeNode{
            2, 
            &TreeNode{5, nil, nil },
            &TreeNode{2, nil, nil },
        },
        &TreeNode{2, nil, nil },
    }
    fmt.Println(isUnivalTree(tree2)) // true

    fmt.Println(isUnivalTree1(tree1)) // true
    fmt.Println(isUnivalTree1(tree2)) // true
}