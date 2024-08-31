package main

// 1430. Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree
// Given a binary tree where each path going from the root to any leaf form a valid sequence, 
// check if a given string is a valid sequence in such binary tree. 

// We get the given string from the concatenation of an array of integers arr 
// and the concatenation of all values of the nodes along a path results in a sequence in the given binary tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_1.png" />
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,0,1]
// Output: true
// Explanation: 
// The path 0 -> 1 -> 0 -> 1 is a valid sequence (green color in the figure). 
// Other valid sequences are: 
// 0 -> 1 -> 1 -> 0 
// 0 -> 0 -> 0

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_2.png" />
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,0,1]
// Output: false 
// Explanation: The path 0 -> 0 -> 1 does not exist, therefore it is not even a sequence.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_3.png" />
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,1]
// Output: false
// Explanation: The path 0 -> 1 -> 1 is a sequence, but it is not a valid sequence.

// Constraints:
//     1 <= arr.length <= 5000
//     0 <= arr[i] <= 9
//     Each node's value is between [0 - 9].

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
func isValidSequence(root *TreeNode, arr []int) bool {
    if root == nil || len(arr) == 0 || root.Val != arr[0] {
        return false
    }
    if root.Left == root.Right { // 叶子
        return len(arr) == 1
    }
    return isValidSequence(root.Left, arr[1:]) || isValidSequence(root.Right, arr[1:])
}

// dfs
func isValidSequence1(root *TreeNode, arr []int) bool {
    if root.Val != arr[0] {
        return false
    }
    res := false
    var dfs func(node *TreeNode, depth int, valid bool)
    dfs = func(node *TreeNode, depth int, valid bool) {
        if node == nil {
            return
        }
        cvalid := valid && depth < len(arr) && node.Val == arr[depth]
        if cvalid && node.Left == nil && node.Right == nil && depth == len(arr)-1 {
            res = true
        }
        // 如果当前已经无效, 则停止搜索其子节点
        if cvalid {
            depth++
            dfs(node.Left, depth,  cvalid)
            dfs(node.Right, depth, cvalid)
        }
    }
    dfs(root, 0, true)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_1.png" />
    // Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,0,1]
    // Output: true
    // Explanation: 
    // The path 0 -> 1 -> 0 -> 1 is a valid sequence (green color in the figure). 
    // Other valid sequences are: 
    // 0 -> 1 -> 1 -> 0 
    // 0 -> 0 -> 0
    tree1 := &TreeNode {
        0,
        &TreeNode{1, &TreeNode{0, nil, &TreeNode{1, nil, nil, }, }, &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{0, nil, nil, }, }, },
        &TreeNode{0, &TreeNode{0, nil, nil, }, nil, },
    }
    fmt.Println(isValidSequence(tree1, []int{0,1,0,1})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_2.png" />
    // Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,0,1]
    // Output: false 
    // Explanation: The path 0 -> 0 -> 1 does not exist, therefore it is not even a sequence.
    tree2 := &TreeNode {
        0,
        &TreeNode{1, &TreeNode{0, nil, &TreeNode{1, nil, nil, }, }, &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{0, nil, nil, }, }, },
        &TreeNode{0, &TreeNode{0, nil, nil, }, nil, },
    }
    fmt.Println(isValidSequence(tree2, []int{0,0,1})) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_3.png" />
    // Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,1]
    // Output: false
    // Explanation: The path 0 -> 1 -> 1 is a sequence, but it is not a valid sequence.
    tree3 := &TreeNode {
        0,
        &TreeNode{1, &TreeNode{0, nil, &TreeNode{1, nil, nil, }, }, &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{0, nil, nil, }, }, },
        &TreeNode{0, &TreeNode{0, nil, nil, }, nil, },
    }
    fmt.Println(isValidSequence(tree3, []int{0,1,1})) // false

    fmt.Println(isValidSequence1(tree1, []int{0,1,0,1})) // true
    fmt.Println(isValidSequence1(tree2, []int{0,0,1})) // false
    fmt.Println(isValidSequence1(tree3, []int{0,1,1})) // false
}