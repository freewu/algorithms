package main

// 988. Smallest String Starting From Leaf
// You are given the root of a binary tree where each node has a value in the range [0, 25] representing the letters 'a' to 'z'.

// Return the lexicographically smallest string that starts at a leaf of this tree and ends at the root.
// As a reminder, any shorter prefix of a string is lexicographically smaller.
//     For example, "ab" is lexicographically smaller than "aba".

// A leaf of a node is a node that has no children.

// Example 1:
//         0[a]
//        /     \
//      1[b]    2 c 
//     /  \     /  \
//    3[d] 4 e 3 d  4 e
// <img src="https://assets.leetcode.com/uploads/2019/01/30/tree1.png" />
// Input: root = [0,1,2,3,4,3,4]
// Output: "dba"

// Example 2:
//         25[z]
//        /     \
//      1 b     3 [d]
//     /  \     /   \
//    2 c 3 d  0[a]  2 c
// <img src="https://assets.leetcode.com/uploads/2019/01/30/tree2.png" />
// Input: root = [25,1,3,1,3,0,2]
// Output: "adz"

// Example 3:
//         2[c]
//        /     \
//       2 c     1 [b]
//        \     /   
//        1 b  0[a] 
//        /
//       0 a
// <img src="https://assets.leetcode.com/uploads/2019/02/01/tree3.png" />
// Input: root = [2,2,1,null,1,0,null,0]
// Output: "abc"
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 8500].
//     0 <= Node.val <= 25

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
func smallestFromLeaf(root *TreeNode) string {
    res := ""
    var dfs func(root *TreeNode, str string)
    dfs = func(root *TreeNode, str string) {
        if root == nil {
            return
        }
        str = string('a' + root.Val) + str // 把当前节点表示的字符添加到str头部
        if root.Left == nil && root.Right == nil { // 如果当前节点是叶子节点，把str和s进行比较
            if res == "" || str < res {
                res = str
            }
            return
        }
        dfs(root.Left, str)
        dfs(root.Right, str)
    }
    dfs(root, "")
    return res
}

// dfs
func smallestFromLeaf1(root *TreeNode) string {
    res := ""
    min := func (a, b string) string {
        if a == "" {
            return b
        }
        if b == "" {
            return a
        }
        if a > b {
            return b
        }
        return a
    }
    var dfs func(root *TreeNode, prev string, res *string)
    dfs = func(root *TreeNode, prev string, res *string) {
        prev = string('a' + root.Val) + prev
        if root.Left == nil && root.Right == nil {
            *res = min(*res, prev)
            return
        }
        if root.Left != nil {
            dfs(root.Left, prev, res)
        }
        if root.Right != nil {
            dfs(root.Right, prev, res)
        }
    }
    dfs(root, "", &res)
    return res
}

func main() {
    tree1 := &TreeNode {
        0,
        &TreeNode { 1, &TreeNode{3, nil, nil},  &TreeNode{4, nil, nil}, },
        &TreeNode { 2, &TreeNode{3, nil, nil},  &TreeNode{4, nil, nil}, },
    }
    // Example 1:
    //         0[a]
    //        /     \
    //      1[b]    2 c 
    //     /  \     /  \
    //    3[d] 4 e 3 d  4 e
    // <img src="https://assets.leetcode.com/uploads/2019/01/30/tree1.png" />
    // Input: root = [0,1,2,3,4,3,4]
    // Output: "dba"
    fmt.Println(smallestFromLeaf(tree1)) // dba

    tree2 := &TreeNode {
        25,
        &TreeNode { 1, &TreeNode{2, nil, nil},  &TreeNode{3, nil, nil}, },
        &TreeNode { 3, &TreeNode{0, nil, nil},  &TreeNode{2, nil, nil}, },
    }
    // Example 2:
    //         25[z]
    //        /     \
    //      1 b     3 [d]
    //     /  \     /   \
    //    2 c 3 d  0[a]  2 c
    // <img src="https://assets.leetcode.com/uploads/2019/01/30/tree2.png" />
    // Input: root = [25,1,3,1,3,0,2]
    // Output: "adz"
    fmt.Println(smallestFromLeaf(tree2)) // adz

    tree3 := &TreeNode {
        2,
        &TreeNode { 2, &TreeNode{1, &TreeNode{0, nil, nil}, nil},  nil, },
        &TreeNode { 1, &TreeNode{0, nil, nil},  nil, },
    }
    // Example 3:
    //         2[c]
    //        /     \
    //       2 c     1 [b]
    //        \     /   
    //        1 b  0[a] 
    //        /
    //       0 a
    // <img src="https://assets.leetcode.com/uploads/2019/02/01/tree3.png" />
    // Input: root = [2,2,1,null,1,0,null,0]
    // Output: "abc"
    fmt.Println(smallestFromLeaf(tree3)) // abc

    fmt.Println(smallestFromLeaf1(tree1)) // dba
    fmt.Println(smallestFromLeaf1(tree2)) // adz
    fmt.Println(smallestFromLeaf1(tree3)) // abc
}