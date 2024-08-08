package main

// LCR 124. 推理二叉树
// 某二叉树的先序遍历结果记录于整数数组 preorder，它的中序遍历结果记录于整数数组 inorder。
// 请根据 preorder 和 inorder 的提示构造出这棵二叉树并返回其根节点。

// 注意：preorder 和 inorder 中均不含重复数字。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2021/02/19/tree.jpg" />
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]

// 示例 2:
// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]

// 提示:
//     1 <= preorder.length <= 3000
//     inorder.length == preorder.length
//     -3000 <= preorder[i], inorder[i] <= 3000
//     inorder 均出现在 preorder
//     preorder 保证 为二叉树的前序遍历序列
//     inorder 保证 为二叉树的中序遍历序列

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
func deduceTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: preorder[0]}
    for pos, node := range inorder {
        if node == root.Val {
            root.Left = deduceTree(preorder[1:pos+1], inorder[:pos])
            root.Right = deduceTree(preorder[pos+1:], inorder[pos+1:])
        }
    }
    return root
}

// dfs
func deduceTree1(preorder []int, inorder []int) *TreeNode {
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
func deduceTree2(preorder []int, inorder []int) *TreeNode {
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
        node.Left = deduceTree2(preorder[1:index+1], inorder[0:index])
        node.Right = deduceTree2(preorder[index+1:], inorder[index+1:])
        return node
    }
}

func deduceTree3(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    index := 0
    for preorder[0] != inorder[index] {
        index++
    }
    return &TreeNode{
        Val:   preorder[0],
        Left:  deduceTree3(preorder[1:index+1], inorder[:index]),
        Right: deduceTree3(preorder[index+1:], inorder[index+1:]),
    }
}

func main() {
    fmt.Printf("%v\n",deduceTree([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
    fmt.Printf("%v\n",deduceTree1([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
    fmt.Printf("%v\n",deduceTree2([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
    fmt.Printf("%v\n",deduceTree3([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
}