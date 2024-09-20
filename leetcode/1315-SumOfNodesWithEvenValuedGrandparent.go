package main

// 1315. Sum of Nodes with Even-Valued Grandparent
// Given the root of a binary tree, return the sum of values of nodes with an even-valued grandparent. 
// If there are no nodes with an even-valued grandparent, return 0.

// A grandparent of a node is the parent of its parent if it exists.

// Example 1:
//              (6)
//            /     \
//           7      (8)
//         /   \   /   \
//       [2]   [7][1]  [3]
//       /    /  \       \ 
//      9    1    4      [5]
// <img src="https://assets.leetcode.com/uploads/2021/08/10/even1-tree.jpg" />
// Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// Output: 18
// Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.

// Example 2:
// Input: root = [1]
// Output: 0
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     1 <= Node.val <= 100

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
func sumEvenGrandparent(root *TreeNode) int {
    res := 0
    if root == nil {
        return res
    }
    var dfs func(root, parent, grand *TreeNode)
    dfs = func(root, parent, grand *TreeNode)  {
        if root == nil { return }
        if grand != nil && grand.Val % 2 == 0 { // 是 祖父节点 且 值为 偶数
            res += root.Val // 累加孙子节点的值
        }
        dfs(root.Left, root, parent)
        dfs(root.Right, root, parent)
    }
    dfs(root, nil, nil)
    return res
}

func sumEvenGrandparent1(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil { return }
        if root.Val % 2 == 0 {
            if root.Left != nil{
                if root.Left.Left != nil {
                    res += root.Left.Left.Val
                }
                if root.Left.Right != nil {
                    res += root.Left.Right.Val
                }
            }
            if root.Right != nil {
                if root.Right.Left != nil {
                    res += root.Right.Left.Val
                }
                if root.Right.Right != nil {
                    res += root.Right.Right.Val
                }
            }
        }
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    return res
}


func main() {
    // Example 1:
    //              (6)
    //            /     \
    //           7      (8)
    //         /   \   /   \
    //       [2]   [7][1]  [3]
    //       /    /  \       \ 
    //      9    1    4      [5]
    // <img src="https://assets.leetcode.com/uploads/2021/08/10/even1-tree.jpg" />
    // Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
    // Output: 18
    // Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.
    tree1 := &TreeNode{
        6, 
        &TreeNode{7, &TreeNode{2, &TreeNode{9, nil, nil}, nil}, &TreeNode{7, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}, }, },
        &TreeNode{8, &TreeNode{1, nil, nil}, &TreeNode{ 3, nil, &TreeNode{5, nil, nil},}, },
    }
    fmt.Println(sumEvenGrandparent(tree1)) // 18
    // Example 2:
    // Input: root = [1]
    // Output: 0
    tree2 := &TreeNode{1, nil, nil}
    fmt.Println(sumEvenGrandparent(tree2)) // 0

    fmt.Println(sumEvenGrandparent1(tree1)) // 18
    fmt.Println(sumEvenGrandparent1(tree2)) // 0
}