package main

// 662. Maximum Width of Binary Tree
// Given the root of a binary tree, return the maximum width of the given tree.
// The maximum width of a tree is the maximum width among all levels.
// The width of one level is defined as the length between the end-nodes (the leftmost and rightmost non-null nodes), 
// where the null nodes between the end-nodes that would be present in a complete binary tree extending down to that level are also counted into the length calculation.
// It is guaranteed that the answer will in the range of a 32-bit signed integer.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/width1-tree.jpg" />
// Input: root = [1,3,2,5,3,null,9]
// Output: 4
// Explanation: The maximum width exists in the third level with length 4 (5,3,null,9).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/14/maximum-width-of-binary-tree-v3.jpg" />
// Input: root = [1,3,2,5,null,null,9,6,null,7]
// Output: 7
// Explanation: The maximum width exists in the fourth level with length 7 (6,null,null,null,null,null,7).

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/width3-tree.jpg" />
// Input: root = [1,3,2,5]
// Output: 2
// Explanation: The maximum width exists in the second level with length 2 (3,2).

// Constraints:
//     The number of nodes in the tree is in the range [1, 3000].
//     -100 <= Node.val <= 100

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
func widthOfBinaryTree(root *TreeNode) int {
    res := 0
    lst := []int{}
    var dfs func(root *TreeNode, id, depth int, lst *[]int)
    dfs = func (root *TreeNode, id, depth int, lst *[]int) {
        if root == nil {
            return
        }
        if depth >= len(*lst) {
            *lst = append(*lst, id)
        }
        if (id + 1 - (*lst)[depth]) > res {
            res = id + 1 - (*lst)[depth]
        }
        if root.Left != nil  { dfs(root.Left, id*2, depth+1, lst)    }
        if root.Right != nil { dfs(root.Right, id*2+1, depth+1, lst) }
    }
    dfs(root, 1, 0, &lst)
    return res
}

// bfs
func widthOfBinaryTree1(root *TreeNode) int {
    type Node struct {
        id int
        tree *TreeNode
    }
    s, h, t, res := make([]*Node, 1), 0, 1, 1
    s[0] = &Node{
        id: 1,
        tree: root,
    }
    n := s[0]
    for h < t {
        x := s[h]
        if x.tree.Left != nil {
            s = append(s, &Node{
                id: x.id*2,
                tree: x.tree.Left,
            })
            t++
        }
        if x.tree.Right != nil {
            s = append(s, &Node{
                id: x.id*2+1,
                tree: x.tree.Right,
            })
            t++
        }
        if x == n && h < t-1 {
            w := s[t-1].id - s[h+1].id + 1
            if w > res {
                res = w
            }
            n = s[t-1]
        }
        h++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/width1-tree.jpg" />
    // Input: root = [1,3,2,5,3,null,9]
    // Output: 4
    // Explanation: The maximum width exists in the third level with length 4 (5,3,null,9).
    tree1 := &TreeNode {
        1,
        &TreeNode{3, &TreeNode{5, nil, nil, }, &TreeNode{3, nil, nil, }, },
        &TreeNode{2, nil,                      &TreeNode{9, nil, nil, }, },
    }
    fmt.Println(widthOfBinaryTree(tree1)) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/14/maximum-width-of-binary-tree-v3.jpg" />
    // Input: root = [1,3,2,5,null,null,9,6,null,7]
    // Output: 7
    // Explanation: The maximum width exists in the fourth level with length 7 (6,null,null,null,null,null,7).
    tree2 := &TreeNode {
        1,
        &TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil, }, nil, }, nil,                      },
        &TreeNode{2, nil,                                           &TreeNode{9, &TreeNode{7, nil, nil, }, nil, }, },
    }
    fmt.Println(widthOfBinaryTree(tree2)) // 7
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/width3-tree.jpg" />
    // Input: root = [1,3,2,5]
    // Output: 2
    // Explanation: The maximum width exists in the second level with length 2 (3,2).
    tree3 := &TreeNode {
        1,
        &TreeNode{3, &TreeNode{5, nil, nil, }, nil, },
        &TreeNode{2, nil,                      nil, },
    }
    fmt.Println(widthOfBinaryTree(tree3)) // 2

    fmt.Println(widthOfBinaryTree1(tree1)) // 4
    fmt.Println(widthOfBinaryTree1(tree2)) // 7
    fmt.Println(widthOfBinaryTree1(tree3)) // 2
}