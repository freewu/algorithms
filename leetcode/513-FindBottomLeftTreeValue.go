package main

// 513. Find Bottom Left Tree Value
// Given the root of a binary tree, return the leftmost value in the last row of the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/14/tree1.jpg" />
// Input: root = [2,1,3]
// Output: 1

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/14/tree2.jpg" />
// Input: root = [1,2,3,4,null,5,6,null,null,7]
// Output: 7

// Constraints:
//         The number of nodes in the tree is in the range [1, 10^4].
//         -2^31 <= Node.val <= 2^31 - 1

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
// DFS
func findBottomLeftValue(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res, maxHeight := 0, -1
    var dfs func(root *TreeNode, curHeight int) 
    dfs = func (root *TreeNode, curHeight int) {
        if curHeight > maxHeight && root.Left == nil && root.Right == nil { // 到达最深叶子
            maxHeight = curHeight
            res = root.Val
        }
        if root.Left != nil {
            dfs(root.Left, curHeight + 1)
        }
        if root.Right != nil {
            dfs(root.Right, curHeight + 1)
        }
    }
    dfs(root, 0)
    return res
}

// BFS
func findBottomLeftValue1(root *TreeNode) int {
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        next := []*TreeNode{}
        for _, node := range queue {
            if node.Left != nil {
                next = append(next, node.Left)
            }
            if node.Right != nil {
                next = append(next, node.Right)
            }
        }
        if len(next) == 0 {
            return queue[0].Val
        }
        queue = next
    }
    return 0
}

func main() {
    tree1 := buildTree([]*int{ptr(1), ptr(1), ptr(3)})
    printTree(tree1)
    fmt.Println(findBottomLeftValue(tree1)) // 1

    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), nil, ptr(5), ptr(6), nil, nil, ptr(7)})
    printTree(tree2)
    fmt.Println(findBottomLeftValue(tree2)) // 7

    tree3 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), nil, ptr(5), ptr(6), nil, nil, nil, ptr(7)})
    printTree(tree3)
    fmt.Println(findBottomLeftValue(tree3)) // 7

    fmt.Println(findBottomLeftValue1(tree1)) // 1

    fmt.Println(findBottomLeftValue1(tree2)) // 7

    fmt.Println(findBottomLeftValue1(tree3)) // 7
}