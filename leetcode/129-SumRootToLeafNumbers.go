package main

// 129. Sum Root to Leaf Numbers
// You are given the root of a binary tree containing digits from 0 to 9 only.
// Each root-to-leaf path in the tree represents a number.
//     For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.

// Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.
// A leaf node is a node with no children.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/num1tree.jpg" />
// Input: root = [1,2,3]
// Output: 25
// Explanation:
// The root-to-leaf path 1->2 represents the number 12.
// The root-to-leaf path 1->3 represents the number 13.
// Therefore, sum = 12 + 13 = 25.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/num2tree.jpg" />
// Input: root = [4,9,0,5,1]
// Output: 1026
// Explanation:
// The root-to-leaf path 4->9->5 represents the number 495.
// The root-to-leaf path 4->9->1 represents the number 491.
// The root-to-leaf path 4->0 represents the number 40.
// Therefore, sum = 495 + 491 + 40 = 1026.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     0 <= Node.val <= 9
//     The depth of the tree will not exceed 10.

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
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

func printTreeNode(t *TreeNode) {
    if nil == t {
        return
    }
    fmt.Println()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var dfs func(root *TreeNode, sum int) int 
    dfs = func(root *TreeNode, sum int) int {
        if root == nil {
            return 0
        }
        sum = sum * 10 + root.Val // 每递归一层 * 10
        if root.Left == nil && root.Right == nil {
            return sum
        } else {
            return dfs(root.Left, sum) + dfs(root.Right, sum)
        }
    }
    return dfs(root, 0)
}

func sumNumbers1(root *TreeNode) int {
    var dfs func(root *TreeNode, num int) int
    dfs = func(root *TreeNode, num int) int {
        if root.Left == nil && root.Right == nil {
            return num * 10 + root.Val
        }
        sum := 0
        if root.Left != nil  {
            sum += dfs(root.Left, num * 10 + root.Val)
        }
        if root.Right != nil {
            sum += dfs(root.Right, num * 10 + root.Val)
        }
        return sum
    }
    return dfs(root, 0)
}

func main() {
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree1)
    tree2 := buildTree([]*int{ptr(4), ptr(9), ptr(0), ptr(5), ptr(1)})
    printTree(tree2)
    // The root-to-leaf path 1->2 represents the number 12.
    // The root-to-leaf path 1->3 represents the number 13.
    // Therefore, sum = 12 + 13 = 25.
    fmt.Println(sumNumbers(tree1)) // 25
    // The root-to-leaf path 4->9->5 represents the number 495.
    // The root-to-leaf path 4->9->1 represents the number 491.
    // The root-to-leaf path 4->0 represents the number 40.
    // Therefore, sum = 495 + 491 + 40 = 1026.
    fmt.Println(sumNumbers(tree2)) // 1026

    fmt.Println(sumNumbers1(tree1)) // 25
    fmt.Println(sumNumbers1(tree2)) // 1026
}
