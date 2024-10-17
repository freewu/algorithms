package main

// 1457. Pseudo-Palindromic Paths in a Binary Tree
// Given a binary tree where node values are digits from 1 to 9. 
// A path in the binary tree is said to be pseudo-palindromic if at least one permutation of the node values in the path is a palindrome.

// Return the number of pseudo-palindromic paths going from the root node to leaf nodes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/05/06/palindromic_paths_1.png" />
// Input: root = [2,3,1,3,1,null,1]
// Output: 2 
// Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the red path [2,3,3], the green path [2,1,1], and the path [2,3,1]. Among these paths only red path and green path are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in [3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/05/07/palindromic_paths_2.png" />
// Input: root = [2,1,1,1,3,null,null,null,null,null,1]
// Output: 1 
// Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the green path [2,1,1], the path [2,1,3,1], and the path [2,1]. Among these paths only the green path is pseudo-palindromic since [2,1,1] can be rearranged in [1,2,1] (palindrome).

// Example 3:
// Input: root = [9]
// Output: 1

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 9

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
func pseudoPalindromicPaths(root *TreeNode) int {
    var dfs func(*TreeNode, int) int
    dfs = func(node *TreeNode, mask int) int {
        if node == nil { return 0 }
        mask ^= 1 << node.Val
        if node.Left == node.Right {
            if mask & (mask - 1) == 0 { return 1 }
            return 0
        }
        left, right := dfs(node.Left, mask), dfs(node.Right, mask)
        return left + right
    }
    return dfs(root, 0)
}

func pseudoPalindromicPaths1(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode, count int)
    dfs = func(node *TreeNode, count int) {
        if node == nil { return }
        if node.Left == nil && node.Right == nil {
            count ^= 1 << node.Val
            if count&(count-1) == 0 { res++ }
            count ^= 1 << node.Val
        }
        count ^= 1 << node.Val
        dfs(node.Left, count) // travel left
        dfs(node.Right, count) // travel right
    }
    dfs(root, 0)
    return res
}
 
func pseudoPalindromicPaths2(root *TreeNode) int {
    numberOfOnes := func(n int) int{
        res := 0
        for n > 0 {
           res += n & 1
           n >>= 1
        }
        return res
    }
    var preorder func(root *TreeNode, count int) int
    preorder = func(root *TreeNode, count int) int {
        if root == nil { return 0 }
        count ^= (1 << root.Val)
        if root.Left == nil && root.Right == nil && numberOfOnes(count) <= 1  { return 1 }
        return preorder(root.Left, count) + preorder(root.Right, count)
    }
    return preorder(root, 0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/05/06/palindromic_paths_1.png" />
    // Input: root = [2,3,1,3,1,null,1]
    // Output: 2 
    // Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the red path [2,3,3], the green path [2,1,1], and the path [2,3,1]. Among these paths only red path and green path are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in [3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).
    tree1 := &TreeNode{
        2, 
        &TreeNode{3, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil, }, },
        &TreeNode{1, nil,                    &TreeNode{1, nil, nil, }, },
    }
    fmt.Println(pseudoPalindromicPaths(tree1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/05/07/palindromic_paths_2.png" />
    // Input: root = [2,1,1,1,3,null,null,null,null,null,1]
    // Output: 1 
    // Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the green path [2,1,1], the path [2,1,3,1], and the path [2,1]. Among these paths only the green path is pseudo-palindromic since [2,1,1] can be rearranged in [1,2,1] (palindrome).
    tree2 := &TreeNode{
        2, 
        &TreeNode{1, &TreeNode{1, nil, nil, }, &TreeNode{3, nil, &TreeNode{1, nil, nil, }, }, },
        &TreeNode{1, nil, nil, },
    }
    fmt.Println(pseudoPalindromicPaths(tree2)) // 1
    // Example 3:
    // Input: root = [9]
    // Output: 1
    tree3 := &TreeNode{9, nil, nil, }
    fmt.Println(pseudoPalindromicPaths(tree3)) // 1

    fmt.Println(pseudoPalindromicPaths1(tree1)) // 2
    fmt.Println(pseudoPalindromicPaths1(tree2)) // 1
    fmt.Println(pseudoPalindromicPaths1(tree3)) // 1
    
    fmt.Println(pseudoPalindromicPaths2(tree1)) // 2
    fmt.Println(pseudoPalindromicPaths2(tree2)) // 1
    fmt.Println(pseudoPalindromicPaths2(tree3)) // 1
}