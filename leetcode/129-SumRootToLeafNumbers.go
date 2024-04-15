package main

// 129. Sum Root to Leaf Numbers
// You are given the root of a binary tree containing digits from 0 to 9 only.
// Each root-to-leaf path in the tree represents a number.
//     For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.

// Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.
// A leaf node is a node with no children.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/num1tree.jpg" />
// Input: root = [1,2,3]
// Output: 25
// Explanation:
// The root-to-leaf path 1->2 represents the number 12.
// The root-to-leaf path 1->3 represents the number 13.
// Therefore, sum = 12 + 13 = 25.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/num2tree.jpg" />
// Input: root = [4,9,0,5,1]
// Output: 1026
// Explanation:
// The root-to-leaf path 4->9->5 represents the number 495.
// The root-to-leaf path 4->9->1 represents the number 491.
// The root-to-leaf path 4->0 represents the number 40.
// Therefore, sum = 495 + 491 + 40 = 1026.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     0 <= Node.val <= 9
//     The depth of the tree will not exceed 10.

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func printTreeNode(t *TreeNode) {
    if nil == t {
        return
    }
    fmt.Println()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var dfs func(root *TreeNode, sum int) int 
    dfs = func(root *TreeNode, sum int) int {
        if root == nil {
            return 0
        }
        sum = sum * 10 + root.Val // 每递归一层 * 10
        if root.Left == nil && root.Right == nil {
            return sum
        } else {
            return dfs(root.Left, sum) + dfs(root.Right, sum)
        }
    }
    return dfs(root, 0)
}

func sumNumbers1(root *TreeNode) int {
    var dfs func(root *TreeNode, num int) int
    dfs = func(root *TreeNode, num int) int {
        if root.Left == nil && root.Right == nil {
            return num * 10 + root.Val
        }
        sum := 0
        if root.Left != nil  {
            sum += dfs(root.Left, num * 10 + root.Val)
        }
        if root.Right != nil {
            sum += dfs(root.Right, num * 10 + root.Val)
        }
        return sum
    }
    return dfs(root, 0)
}

func main() {
    tree1 := &TreeNode {
        1,
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    tree2 := &TreeNode {
        4,
        &TreeNode{
            9, 
            &TreeNode{5, nil, nil}, 
            &TreeNode{1, nil, nil},
        },
        &TreeNode{0, nil, nil},
    }
    // The root-to-leaf path 1->2 represents the number 12.
    // The root-to-leaf path 1->3 represents the number 13.
    // Therefore, sum = 12 + 13 = 25.
    fmt.Println(sumNumbers(tree1)) // 25
    // The root-to-leaf path 4->9->5 represents the number 495.
    // The root-to-leaf path 4->9->1 represents the number 491.
    // The root-to-leaf path 4->0 represents the number 40.
    // Therefore, sum = 495 + 491 + 40 = 1026.
    fmt.Println(sumNumbers(tree2)) // 1026

    fmt.Println(sumNumbers1(tree1)) // 25
    fmt.Println(sumNumbers1(tree2)) // 1026
}