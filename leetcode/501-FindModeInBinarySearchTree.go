package main

// 501. Find Mode in Binary Search Tree
// Given the root of a binary search tree (BST) with duplicates, 
// return all the mode(s) (i.e., the most frequently occurred element) in it.

// If the tree has more than one mode, return them in any order.

// Assume a BST is defined as follows:
//     The left subtree of a node contains only nodes with keys less than or equal to the node's key.
//     The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
//     Both the left and right subtrees must also be binary search trees.
 
// Example 1:
//         1
//           \
//            2
//           /
//          2
// <img src="https://assets.leetcode.com/uploads/2021/03/11/mode-tree.jpg" />
// Input: root = [1,null,2,2]
// Output: [2]

// Example 2:
// Input: root = [0]
// Output: [0]

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5

// Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).

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
func findMode(root *TreeNode) []int {
    if root == nil { 
        return []int{} 
    }
    var dfs func(root *TreeNode, tmp map[int]int) map[int]int
    dfs = func(root *TreeNode, tmp map[int]int) map[int]int {
        if root == nil { 
            return tmp
        }
        tmp[root.Val]++
        dfs(root.Left, tmp)
        dfs(root.Right, tmp)
        return tmp
    }
    mx, res, mp := 0, []int{}, make(map[int]int)
    temp := dfs(root, mp)
    for _, v := range temp { // 找到最大的值
        if v > mx { 
            mx = v 
        }
    }
    for i, v := range temp {
        if v == mx { // 出现数量最大就加入到列表中
            res = append(res, i)
        }
    }
    return res
}

func findMode1(root *TreeNode) []int {
    var prev *TreeNode
    maxCount, currCount, res := 0, 0, []int{}
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        if prev != nil && prev.Val == node.Val {
            currCount++
        } else {
            currCount = 1
        }
        if currCount > maxCount { // 发现有更大的直接重置 res
            maxCount = currCount
            res = []int{ node.Val }
        } else if currCount == maxCount { // 发现一样的
            res = append(res, node.Val)
        }
        prev = node
        inorder(node.Right)
    }
    inorder(root)
    return res
}

func main() {
    // Example 1:
    //         1
    //           \
    //            2
    //           /
    //          2
    // <img src="https://assets.leetcode.com/uploads/2021/03/11/mode-tree.jpg" />
    // Input: root = [1,null,2,2]
    // Output: [2]
    tree1 := &TreeNode {
        1,
        nil,
        &TreeNode{2, &TreeNode{2, nil, nil, }, nil, },
    }
    fmt.Println(findMode(tree1)) // [2]
    // Example 2:
    // Input: root = [0]
    // Output: [0]
    tree2 := &TreeNode{0, nil, nil, }
    fmt.Println(findMode(tree2)) // [0]

    fmt.Println(findMode1(tree1)) // [2]
    fmt.Println(findMode1(tree2)) // [0]
}