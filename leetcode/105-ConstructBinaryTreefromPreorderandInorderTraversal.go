package main

// 105. Construct Binary Tree from Preorder and Inorder Traversal
// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree 
// and inorder is the inorder traversal of the same tree, construct and return the binary tree.

// Example 1:
// Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// Output: [3,9,20,null,null,15,7]
//     3
//    / \
//   9  20
//     /  \
//    15   7

// Example 2:
// Input: preorder = [-1], inorder = [-1]
// Output: [-1]

// Constraints:
//     1 <= preorder.length <= 3000
//     inorder.length == preorder.length
//     -3000 <= preorder[i], inorder[i] <= 3000
//     preorder and inorder consist of unique values.
//     Each value of inorder also appears in preorder.
//     preorder is guaranteed to be the preorder traversal of the tree.
//     inorder is guaranteed to be the inorder traversal of the tree.

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

// 直接传入需要的 slice 范围作为输入, 可以避免申请对应 inorder 索引的内存
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: preorder[0]}
    for pos, node := range inorder {
        if node == root.Val {
            root.Left = buildTree(preorder[1:pos+1], inorder[:pos])
            root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])
        }
    }
    return root
}

// dfs
func buildTree1(preorder []int, inorder []int) *TreeNode {
    inPos := make(map[int]int)
    for i := 0; i < len(inorder); i++ {
        inPos[inorder[i]] = i
    }
    var dfs func(pre []int, preStart int, preEnd int, inStart int, inPos map[int]int) *TreeNode
    dfs = func(pre []int, preStart int, preEnd int, inStart int, inPos map[int]int) *TreeNode {
        if preStart > preEnd {
            return nil
        }
        root := &TreeNode{Val: pre[preStart]}
        rootIdx := inPos[pre[preStart]]
        leftLen := rootIdx - inStart
        root.Left = dfs(pre, preStart + 1, preStart+leftLen, inStart, inPos)
        root.Right = dfs(pre, preStart + leftLen + 1, preEnd, rootIdx + 1, inPos)
        return root
    }
    return dfs(preorder, 0, len(preorder)-1, 0, inPos)
}

// best solution
func buildTree2(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    if len(preorder) == 1 && len(inorder) == 1 && preorder[0] == inorder[0] {
        return &TreeNode{Val: preorder[0]}
    } else {
        node := &TreeNode{Val: preorder[0]}
        var index int
        for index = 0; index < len(inorder); index++ {
            if inorder[index] == node.Val {
                break
            }
        }
        node.Left = buildTree2(preorder[1:index+1], inorder[0:index])
        node.Right = buildTree2(preorder[index+1:], inorder[index+1:])
        return node
    }
}

func buildTree3(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    index := 0
    for preorder[0] != inorder[index] {
        index++
    }
    return &TreeNode{
        Val:   preorder[0],
        Left:  buildTree3(preorder[1:index+1], inorder[:index]),
        Right: buildTree3(preorder[index+1:], inorder[index+1:]),
    }
}

func main() {
    fmt.Printf("%v\n",buildTree([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
    fmt.Printf("%v\n",buildTree1([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
    fmt.Printf("%v\n",buildTree2([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
    fmt.Printf("%v\n",buildTree3([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
}
