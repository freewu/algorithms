package main

// 897. Increasing Order Search Tree
// Given the root of a binary search tree, 
// rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, 
// and every node has no left child and only one right child.

// Example 1:
//             5                           1
//           /    \                          \
//          3      6                          2 
//        /   \      \             =>          \
//       2     4      8                          3
//     /             /  \                         \
//    1             7    9                         4
//                                                 [\]
//                                                   8
//                                                     \
//                                                      9
// <img src="https://assets.leetcode.com/uploads/2020/11/17/ex1.jpg" />
// Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/17/ex2.jpg" />
//         5                   1
//       /   \        =>         \
//      1     7                    5
//                                  \
//                                    7
// Input: root = [5,1,7] 
// Output: [1,null,5,null,7]

// Constraints:
//     The number of nodes in the given tree will be in the range [1, 100].
//     0 <= Node.val <= 1000

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
func increasingBST1(root *TreeNode) *TreeNode {
    nodes := make([]int, 0)
    var inorder func (root *TreeNode)
    inorder = func (root *TreeNode) { // 中序遍历
        if root != nil {
            inorder(root.Left)
            nodes = append(nodes, root.Val)
            inorder(root.Right)
        }
    }
    inorder(root)
    res := new(TreeNode)
    dummy := res
    for _, val := range nodes { // 重建树
        dummy.Right = &TreeNode{ Val: val, Left: nil, Right: nil }
        dummy = dummy.Right
    }
    return res.Right
}


// 递归
func increasingBST(root *TreeNode) *TreeNode {
    if root == nil { return nil }
    left := increasingBST(root.Left) // 先把左右子树拉平
    root.Left = nil // 然后把根节点左子树置空
    right := increasingBST(root.Right) // 再把右子树拉平
    root.Right = right // 连上根节点
    if left == nil { // 左子树为空的话，就不用处理了
        return root
    }
    p := left
    for p != nil && p.Right != nil { // 左子树非空，把根节点和右子树接到左子树末尾
        p = p.Right
    }
    p.Right = root
    return left
}

func main() {
    // Example 1:
    //             5                           1
    //           /    \                          \
    //          3      6                          2 
    //        /   \      \             =>          \
    //       2     4      8                          3
    //     /             /  \                         \
    //    1             7    9                         4
    //                                                 [\]
    //                                                   8
    //                                                     \
    //                                                      9
    // <img src="https://assets.leetcode.com/uploads/2020/11/17/ex1.jpg" />
    // Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
    // Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
    tree1 := &TreeNode{
        5, 
        &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, &TreeNode{4, nil, nil, }, },
        &TreeNode{6, nil, &TreeNode{8, &TreeNode{7, nil, nil, }, &TreeNode{9, nil, nil, }, }, },
    }
    fmt.Println(increasingBST(tree1)) // &{1 <nil> 0xc000094120}
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/11/17/ex2.jpg" />
    //         5                   1
    //       /   \        =>         \
    //      1     7                    5
    //                                  \
    //                                    7
    // Input: root = [5,1,7] 
    // Output: [1,null,5,null,7]
    tree2 := &TreeNode{
        5, 
        &TreeNode{1, nil, nil, },
        &TreeNode{7, nil, nil, },
    }
    fmt.Println(increasingBST(tree2)) // &{1 <nil> 0xc000094258}

    tree11 := &TreeNode{
        5, 
        &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, &TreeNode{4, nil, nil, }, },
        &TreeNode{6, nil, &TreeNode{8, &TreeNode{7, nil, nil, }, &TreeNode{9, nil, nil, }, }, },
    }
    fmt.Println(increasingBST1(tree11)) // &{1 <nil> 0xc000094120}
    tree12 := &TreeNode{
        5, 
        &TreeNode{1, nil, nil, },
        &TreeNode{7, nil, nil, },
    }
    fmt.Println(increasingBST1(tree12)) // &{1 <nil> 0xc000094258}
}