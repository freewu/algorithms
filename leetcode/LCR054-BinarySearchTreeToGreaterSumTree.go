package main

// LCR 054. 把二叉搜索树转换为累加树
// 给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。

// 提醒一下，二叉搜索树满足下列约束条件：
//     节点的左子树仅包含键 小于 节点键的节点。
//     节点的右子树仅包含键 大于 节点键的节点。
//     左右子树也必须是二叉搜索树。
    

// 示例 1：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/05/03/tree.png" />
// 输入：root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
// 输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

// 示例 2：
// 输入：root = [0,null,1]
// 输出：[1,null,1]

// 示例 3：
// 输入：root = [1,0,2]
// 输出：[3,3,2]

// 示例 4：
// 输入：root = [3,2,4,1]
// 输出：[7,9,4,10]

// 提示：
//     树中的节点数介于 0 和 10^4 之间。
//     每个节点的值介于 -10^4 和 10^4 之间。
//     树中的所有值 互不相同 。
//     给定的树为二叉搜索树。

import "fmt"

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
// func convertBST(root *TreeNode) *TreeNode {
//     sum := 0
//     var inOrder func(node *TreeNode,sum *int) 
//     inOrder = func (node *TreeNode,sum *int) {
//         if node == nil {
//             return
//         }
//         inOrder(node.Right,sum)
//         node.Val += *sum
//         *sum = node.Val
//         inOrder(node.Left,sum)
//     }
//     inOrder(root,&sum)
//     return root
// }

func convertBST(root *TreeNode) *TreeNode {
    var traverse func(root *TreeNode, sum int) int 
    traverse = func (root *TreeNode, sum int) int {
        if root == nil {
            return sum
        }
        // right
        sum = traverse(root.Right, sum)
        root.Val += sum
        // left
        return traverse(root.Left, root.Val)
    }
    traverse(root, 0)
    return root
}

func main() {
    // [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
    tree1 := &TreeNode {
        4,
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, &TreeNode{3, nil, nil}, }, },
        &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, &TreeNode{8, nil, nil}, }, },
    }
    fmt.Println(tree1.Val) // 4
    fmt.Println(convertBST(tree1)) // [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
    fmt.Println(tree1.Val) // 30

    // [0,null,1] 
    tree2 := &TreeNode {
        0,
        nil,
        &TreeNode{1, nil, nil},
    }
    fmt.Println(tree2.Val) // 0
    fmt.Println(convertBST(tree2)) // [1,null,1]
    fmt.Println(tree2.Val) // 1
}