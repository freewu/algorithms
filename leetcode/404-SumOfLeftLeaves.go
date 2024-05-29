package main

// 404. Sum of Left Leaves
// Given the root of a binary tree, return the sum of all left leaves.
// A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.

// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 24
// Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.

// Example 2:
// Input: root = [1]
// Output: 0
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     -1000 <= Node.val <= 1000

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// dfs
func sumOfLeftLeaves(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        // root.Left.Left == nil &&root.Left.Right == nil 左边为叶的判断
        if root.Left != nil && root.Left.Left == nil &&root.Left.Right == nil { // root.Left 存在且为叶子节点 
            res += root.Left.Val
        }
        dfs(root.Left)
        dfs(root.Right)
        return
    }
    dfs(root)
    return res
}

// 采用后序遍历
// 确定递归的出口：在最后一个非叶子结点就应该结束递归，因为遍历到叶子结点不能判断该结点为左结点还是右结点
// 确定本次递归需要做的事情：遍历左、右子树
// 递归的返回值：sum
func sumOfLeftLeaves1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 0
    }
    lSum := sumOfLeftLeaves(root.Left)
    rSum := sumOfLeftLeaves(root.Right)
    // 左叶子结点
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        lSum = root.Left.Val
    }
    return lSum + rSum
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode{9, nil, nil},
        &TreeNode {
            20,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    tree2 := &TreeNode{1, nil, nil}
    // Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
    fmt.Println(sumOfLeftLeaves(tree1)) // 24  9 + 15
    fmt.Println(sumOfLeftLeaves(tree2)) // 0

    fmt.Println(sumOfLeftLeaves1(tree1)) // 24  9 + 15
    fmt.Println(sumOfLeftLeaves1(tree2)) // 0
}