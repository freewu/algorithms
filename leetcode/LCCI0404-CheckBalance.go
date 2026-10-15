package main

// 面试题 04.04. Check Balance LCCI
// Implement a function to check if a binary tree is balanced. 
// For the purposes of this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any node never differ by more than one.

// Example 1:
// Given tree [3,9,20,null,null,15,7]
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return true.

// Example 2:
// Given [1,2,2,3,3,null,null,4,4]
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
// return false.

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
func isBalanced(root *TreeNode) bool { // 自顶向下的递归
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var height func(node *TreeNode) int // 获取树的高度
    height = func(node *TreeNode) int {
        if node == nil { return 0 }
        return max(height(node.Left), height(node.Right)) + 1
    }
    var dfs func(node *TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil { return true }
        return abs(height(node.Left) - height(node.Right)) <= 1 && dfs(node.Left) && dfs(node.Right)
    }
    return dfs(root)
}

// 自底向上的递归
func isBalanced1(root *TreeNode) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var height func(node *TreeNode) int
    height = func(node *TreeNode) int {
        if node == nil { return 0 }
        leftHeight, rightHeight := height(node.Left), height(node.Right)
        if leftHeight == -1 || rightHeight == -1 || abs(leftHeight - rightHeight) > 1 {
            return -1
        }
        return max(leftHeight, rightHeight) + 1
    }
    return height(root) >= 0
}

func main() {
    // Example 1:
    // Given tree [3,9,20,null,null,15,7]
    //     3
    //    / \
    //   9  20
    //     /  \
    //    15   7
    // return true.
    tree1 := buildTree([]*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)})
    printTree(tree1)
    fmt.Println(isBalanced(tree1)) // true
    // Example 2:
    // Given [1,2,2,3,3,null,null,4,4]
    //       1
    //      / \
    //     2   2
    //    / \
    //   3   3
    //  / \
    // 4   4
    // return false.
    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(2), ptr(3), ptr(3), nil, nil, ptr(4), ptr(4)})
    printTree(tree2)
    fmt.Println(isBalanced(tree2)) // false

    tree11 := buildTree([]*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)})
    printTree(tree11)
    fmt.Println(isBalanced1(tree11)) // true
    tree12 := buildTree([]*int{ptr(1), ptr(2), ptr(2), ptr(3), ptr(3), nil, nil, ptr(4), ptr(4)})
    printTree(tree12)
    fmt.Println(isBalanced1(tree12)) // false
}