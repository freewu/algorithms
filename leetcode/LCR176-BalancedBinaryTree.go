package main

// LCR 176. 判断是否为平衡二叉树
// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
// 如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

// 示例 1:
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true 
// 解释：如下图
// <img src="https://pic.leetcode.cn/1695102431-vbmWJn-image.png" />

// 示例 2:
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
// 解释：如下图
// <img src="https://pic.leetcode.cn/1695102434-WlaxCo-image.png" />

// 提示：
//     0 <= 树的结点个数 <= 10000

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
// 递归
func isBalanced1(root *TreeNode) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // iterative function to measure the maximum depth of a tree
    var maxDepth func(root *TreeNode) int 
    maxDepth = func(root *TreeNode) int {
        if root == nil { return 0 }
        return max(maxDepth(root.Left), maxDepth(root.Right))+1
    }
    var helper func(root *TreeNode) bool
    helper = func(root *TreeNode) bool {
        if root == nil {
            return true
        }
        // getting maximum depth possible from both left and right children
        // check if the difference is not more than 1, if it is then the tree is unbalanced and we can stop our check
        if abs(maxDepth(root.Left) -  maxDepth(root.Right)) > 1 {
            return false
        }
        // if both children are balanced then this tree is also balanced
        return helper(root.Left) && helper(root.Right)
    }
    return helper(root)
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    //iterative function to measure the maximum depth of a tree
    var maxDepth func(root *TreeNode) int
    maxDepth = func(root *TreeNode) int {
        if root == nil { return 0 }
        return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
    }
    // getting maximum depth possible from both left and right children
    // check if the difference is not more than 1, if it is then the tree is unbalanced and we can stop our check
    if abs(maxDepth(root.Left) - maxDepth(root.Right)) > 1 {
        return false
    }
    // if both children are balanced then this tree is also balanced
    return isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {
    // Example 1:
    //         3
    //        /  \
    //       9   20
    //          /  \
    //         15   7
    // <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg" />
    // Input: root = [3,9,20,null,null,15,7]
    // Output: true
    tree1 := &TreeNode {
        3,
        &TreeNode{ 9,  nil, nil},
        &TreeNode{ 20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}, },
    }
    fmt.Println(isBalanced(tree1)) // true
    // Example 2:
    //             1 
    //           /   \
    //          2     2
    //        /   \
    //       3     3
    //     /   \
    //    4     4
    // <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg" />
    // Input: root = [1,2,2,3,3,null,null,4,4]
    // Output: false
    tree2 := &TreeNode {
        1,
        &TreeNode {
            2,
            &TreeNode{ 3,  &TreeNode{4, nil, nil },  &TreeNode{4, nil, nil },  },
            &TreeNode{3, nil, nil },
        },
        &TreeNode{2, nil, nil},
    }
    fmt.Println(isBalanced(tree2)) // false
    // Example 3:
    // <img src="" />
    // Input: root = []
    // Output: true
    fmt.Println(isBalanced(nil)) // false

    fmt.Println(isBalanced1(tree1)) // true
    fmt.Println(isBalanced1(tree2)) // false
    fmt.Println(isBalanced1(nil)) // false
}