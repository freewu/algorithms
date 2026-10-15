package main

// 2773. Height of Special Binary Tree
// You are given a root, which is the root of a special binary tree with n nodes. T
// he nodes of the special binary tree are numbered from 1 to n. 
// Suppose the tree has k leaves in the following order: b1 < b2 < ... < bk.

// The leaves of this tree have a special property! That is, for every leaf bi, the following conditions hold:
//     The right child of bi is bi + 1 if i < k, and b1 otherwise.
//     The left child of bi is bi - 1 if i > 1, and bk otherwise.

// Return the height of the given tree.

// Note: The height of a binary tree is the length of the longest path from the root to any other node.

// Example 1:
// Input: root = [1,2,3,null,null,4,5]
// Output: 2
// Explanation: The given tree is shown in the following picture. Each leaf's left child is the leaf to its left (shown with the blue edges). Each leaf's right child is the leaf to its right (shown with the red edges). We can see that the graph has a height of 2.
// <img src="https://assets.leetcode.com/uploads/2023/07/12/1.png" />

// Example 2:
// Input: root = [1,2]
// Output: 1
// Explanation: The given tree is shown in the following picture. There is only one leaf, so it doesn't have any left or right child. We can see that the graph has a height of 1.
// <img src="https://assets.leetcode.com/uploads/2023/07/12/2.png" />

// Example 3:
// Input: root = [1,2,3,null,null,4,null,5,6]
// Output: 3
// Explanation: The given tree is shown in the following picture. Each leaf's left child is the leaf to its left (shown with the blue edges). Each leaf's right child is the leaf to its right (shown with the red edges). We can see that the graph has a height of 3.
// <img src="https://assets.leetcode.com/uploads/2023/07/12/3.png" />

// Constraints:
//     n == number of nodes in the tree
//     2 <= n <= 10^4
//     1 <= node.val <= n
//     The input is generated such that each node.val is unique.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// == 二叉树构建 & 打印输出 start ==================================================================
func ptr(v int) *int {
    return &v
}

// buildTree 按层序数组（LeetCode 格式，nil 表示 null）创建二叉树
func buildTree(nums []*int) *TreeNode {
    if len(nums) == 0 || nums[0] == nil {
        return nil
    }
    root := &TreeNode{Val: *nums[0]}
    queue := []*TreeNode{root}
    i := 1
    for len(queue) > 0 && i < len(nums) {
        node := queue[0]
        queue = queue[1:]
        if nums[i] != nil {
            node.Left = &TreeNode{Val: *nums[i]}
            queue = append(queue, node.Left)
        }
        i++
        if i < len(nums) && nums[i] != nil {
            node.Right = &TreeNode{Val: *nums[i]}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

// printTree 自顶向下打印二叉树，父子之间用 / \ 连线
// 第 r 层（根为第 1 层）有 2^(r-1) 个位置，每个节点占 block 列，
// 因此节点中心列 = (2j+1) * 2^(h-r) * block / 2（j 为该层从左往右的序号）
func printTree(root *TreeNode) {
    if root == nil {
        fmt.Println("(empty tree)")
        return
    }
    // itoa 手写整数转字符串，替代 strconv.Itoa
    itoa := func(v int) string {
        if v == 0 {
            return "0"
        }
        neg := v < 0
        u := uint(v)
        if neg {
            u = -u // 补码取绝对值，v 为最小整数时同样安全
        }
        buf := [24]byte{}
        i := len(buf)
        for u > 0 {
            i--
            buf[i] = byte('0' + u%10)
            u /= 10
        }
        if neg {
            i--
            buf[i] = '-'
        }
        return string(buf[i:])
    }
    // blankLine 生成 n 个空格的画布行（替代 strings.Repeat）
    blankLine := func(n int) []byte {
        line := make([]byte, n)
        for i := range line {
            line[i] = ' '
        }
        return line
    }
    // treeHeight 返回二叉树的层数（空树为 0，根节点为第 1 层）
    var treeHeight func(root *TreeNode) int
    treeHeight = func (root *TreeNode) int {
        if root == nil {
            return 0
        }
        l, r := treeHeight(root.Left), treeHeight(root.Right)
        if l < r {
            l = r
        }
        return l + 1
    }
    h := treeHeight(root)
    // block：每个节点占用的列宽，取最长值的宽度 + 1，并保证为偶数使中心落在整数列
    block := 2
    var scan func(node *TreeNode)
    scan = func(node *TreeNode) {
        if node == nil {
            return
        }
        if w := len(itoa(node.Val)) + 1; w > block {
            block = w
        }
        scan(node.Left)
        scan(node.Right)
    }
    scan(root)
    if block%2 != 0 {
        block++
    }

    // 画布：值占 1 行、连线占 1 行，共 2h-1 行
    width := (1 << h) * block
    rows := make([][]byte, 2*h-1)
    for i := range rows {
        rows[i] = blankLine(width)
    }
    put := func(row, col int, s string) {
        start := col - len(s)/2
        for i := 0; i < len(s); i++ {
            if p := start + i; p >= 0 && p < width {
                rows[row][p] = s[i]
            }
        }
    }
    var draw func(node *TreeNode, r, j int)
    draw = func(node *TreeNode, r, j int) {
        if node == nil {
            return
        }
        col := (2*j + 1) * (1 << (h - r)) * block / 2 // 当前节点中心列
        put(2*(r-1), col, itoa(node.Val))
        if r == h {
            return
        }
        d := (1 << (h - r - 1)) * block / 2 // 父子之间的列距
        off := (d + 1) / 2                  // 连线相对父节点中心的偏移（取行间中点）
        if node.Left != nil {
            put(2*r-1, col-off, "/")
            draw(node.Left, r+1, 2*j)
        }
        if node.Right != nil {
            put(2*r-1, col+off, "\\")
            draw(node.Right, r+1, 2*j+1)
        }
    }
    draw(root, 1, 0)

    for _, line := range rows {
        end := len(line)
        for end > 0 && line[end-1] == ' ' {
            end--
        }
        fmt.Println(string(line[:end]))
    }
}
// == 二叉树构建 & 打印输出 end ====================================================================

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func heightOfTree(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode, depht int)
    dfs = func(node *TreeNode, depht int) {
        if depht > res {
            res = depht
        }
        if node.Left != nil && node.Left.Right != node {
            dfs(node.Left, depht+1)
        }
        if node.Right != nil && node.Right.Left != node {
            dfs(node.Right, depht+1)
        }
    } 
    dfs(root, 0)
    return res
}

func main() {
    // Example 1:
    // Input: root = [1,2,3,null,null,4,5]
    // Output: 2
    // Explanation: The given tree is shown in the following picture. Each leaf's left child is the leaf to its left (shown with the blue edges). Each leaf's right child is the leaf to its right (shown with the red edges). We can see that the graph has a height of 2.
    // <img src="https://assets.leetcode.com/uploads/2023/07/12/1.png" />
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, nil, ptr(4), ptr(5)})
    printTree(tree1)
    fmt.Println(heightOfTree(tree1)) // 2
    // Example 2:
    // Input: root = [1,2]
    // Output: 1
    // Explanation: The given tree is shown in the following picture. There is only one leaf, so it doesn't have any left or right child. We can see that the graph has a height of 1.
    // <img src="https://assets.leetcode.com/uploads/2023/07/12/2.png" />
    tree2 := buildTree([]*int{ptr(1), ptr(2)})
    printTree(tree2)
    fmt.Println(heightOfTree(tree2)) // 1
    // Example 3:
    // Input: root = [1,2,3,null,null,4,null,5,6]
    // Output: 3
    // Explanation: The given tree is shown in the following picture. Each leaf's left child is the leaf to its left (shown with the blue edges). Each leaf's right child is the leaf to its right (shown with the red edges). We can see that the graph has a height of 3.
    // <img src="https://assets.leetcode.com/uploads/2023/07/12/3.png" />
    tree3 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, nil, ptr(4), nil, ptr(5), ptr(6)})
    printTree(tree3)
    fmt.Println(heightOfTree(tree3)) // 3
}