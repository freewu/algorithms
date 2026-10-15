package main

// 1008. Construct Binary Search Tree from Preorder Traversal
// Given an array of integers preorder, which represents the preorder traversal of a BST (i.e., binary search tree), construct the tree and return its root.
// It is guaranteed that there is always possible to find a binary search tree with the given requirements for the given test cases.
// A binary search tree is a binary tree where for every node, any descendant of Node.left has a value strictly less than Node.val, and any descendant of Node.right has a value strictly greater than Node.val.
// A preorder traversal of a binary tree displays the value of the node first, then traverses Node.left, then traverses Node.right.

// Example 1:
//             8
//           /   \
//          5     10
//        /   \     \
//       1     7     12
// <img src="https://assets.leetcode.com/uploads/2019/03/06/1266.png" />
// Input: preorder = [8,5,1,7,10,12]
// Output: [8,5,10,1,7,null,12]

// Example 2:
//         1
//           \
//             3
// Input: preorder = [1,3]
// Output: [1,null,3]

// Constraints:
//     1 <= preorder.length <= 100
//     1 <= preorder[i] <= 1000
//     All the values of preorder are unique.

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
func bstFromPreorder(preorder []int) *TreeNode {
    less := func(n int, arr []int) []int {
        res := []int{}
        for _, v := range arr { if v < n { res = append(res, v) } }
        return res
    }
    greater := func(n int, arr []int) []int {
        res := []int{}
        for _, v := range arr {  if v > n { res = append(res, v) } }
        return res
    }
    var helper func(preorder []int) *TreeNode
    helper = func(preorder []int) *TreeNode {
        if len(preorder) == 0 { 
            return nil 
        }
        node := &TreeNode{
            Val: preorder[0],
            Left: helper(less(preorder[0], preorder)),
            Right: helper(greater(preorder[0], preorder)),
        }
        return node
    }
    return helper(preorder)
}

func bstFromPreorder1(preorder []int) *TreeNode {
    var build func(preorder []int, start, end int) *TreeNode
    build = func(preorder []int, start, end int) *TreeNode {
        if start > end {
            return nil
        }
        rootVal := preorder[start]
        root := &TreeNode{Val: rootVal}
        p := start + 1
        for p <= end && preorder[p] < rootVal {
            p++
        }
        root.Left, root.Right = build(preorder, start+1, p-1), build(preorder, p, end) // 左子树区间: [start+1, p-1], 右子树区间: [p, end]
        return root
    }
    return build(preorder, 0, len(preorder)-1)
}

func main() {
    // Example 1:
    //             8
    //           /   \
    //          5     10
    //        /   \     \
    //       1     7     12
    // <img src="https://assets.leetcode.com/uploads/2019/03/06/1266.png" />
    // Input: preorder = [8,5,1,7,10,12]
    // Output: [8,5,10,1,7,null,12]
    fmt.Println(bstFromPreorder([]int{8,5,1,7,10,12})) // &{8 0xc000008078 0xc0000080a8}
    // Example 2:
    //         1
    //           \
    //             3
    // Input: preorder = [1,3]
    // Output: [1,null,3]
    fmt.Println(bstFromPreorder([]int{1,3})) // &{1 <nil> 0xc0000080f0}

    fmt.Println(bstFromPreorder1([]int{8,5,1,7,10,12})) // &{8 0xc000008078 0xc0000080a8}
    fmt.Println(bstFromPreorder1([]int{1,3})) // &{1 <nil> 0xc0000080f0}
}
