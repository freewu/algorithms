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