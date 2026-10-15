package main

// 2415. Reverse Odd Levels of Binary Tree
// Given the root of a perfect binary tree, reverse the node values at each odd level of the tree.
//     For example, suppose the node values at level 3 are [2,1,3,4,7,11,29,18], then it should become [18,29,11,7,4,3,1,2].

// Return the root of the reversed tree.

// A binary tree is perfect if all parent nodes have two children and all leaves are on the same level.

// The level of a node is the number of edges along the path between it and the root node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/07/28/first_case1.png" />
// Input: root = [2,3,5,8,13,21,34]
// Output: [2,5,3,8,13,21,34]
// Explanation: 
// The tree has only one odd level.
// The nodes at level 1 are 3, 5 respectively, which are reversed and become 5, 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/07/28/second_case3.png" />
// Input: root = [7,13,11]
// Output: [7,11,13]
// Explanation: 
// The nodes at level 1 are 13, 11, which are reversed and become 11, 13.

// Example 3:
// Input: root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
// Output: [0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
// Explanation: 
// The odd levels have non-zero values.
// The nodes at level 1 were 1, 2, and are 2, 1 after the reversal.
// The nodes at level 3 were 1, 1, 1, 1, 2, 2, 2, 2, and are 2, 2, 2, 2, 1, 1, 1, 1 after the reversal.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 2^14].
//     0 <= Node.val <= 10^5
//     root is a perfect binary tree.

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
func reverseOddLevels(root *TreeNode) *TreeNode {
    queue := []*TreeNode{ root }
    for i := 0; len(queue) > 0; i++ {
        t := []*TreeNode{}
        for k := len(queue); k > 0; k-- {
            node := queue[0] // pop
            queue = queue[1:]
            if i % 2 == 1 { // 找出奇数层的节点
                t = append(t, node)
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
                queue = append(queue, node.Right)
            }
        }
        for l, r := 0, len(t) - 1; l < r; l, r = l + 1, r - 1 { // 交换 奇数层数值
            t[l].Val, t[r].Val = t[r].Val, t[l].Val 
        }
    }
    return root
}

func reverseOddLevels1(root *TreeNode) *TreeNode {
    if root == nil { return root }
    queue, flag := []*TreeNode{ root }, false
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            node := queue[i]
            if node.Left != nil  { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        if flag { // 奇数层才需要交换
            for i := 0; i < n / 2; i++ {
                a, b := queue[i], queue[n-i-1]
                a.Val, b.Val = b.Val, a.Val
            }
        }
        flag = !flag
        queue = queue[n:]
    }
    return root
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/07/28/first_case1.png" />
    // Input: root = [2,3,5,8,13,21,34]
    // Output: [2,5,3,8,13,21,34]
    // Explanation: 
    // The tree has only one odd level.
    // The nodes at level 1 are 3, 5 respectively, which are reversed and become 5, 3.
    tree1 := buildTree([]*int{ptr(2), ptr(3), ptr(5), ptr(8), ptr(13), ptr(21), ptr(34)})
    printTree(tree1)
    fmt.Println(reverseOddLevels(tree1))
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/07/28/second_case3.png" />
    // Input: root = [7,13,11]
    // Output: [7,11,13]
    // Explanation: 
    // The nodes at level 1 are 13, 11, which are reversed and become 11, 13.
    tree2 := buildTree([]*int{ptr(7), ptr(13), ptr(11)})
    printTree(tree2)
    fmt.Println(reverseOddLevels(tree2))
    // Example 3:
    // Input: root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
    // Output: [0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
    // Explanation: 
    // The odd levels have non-zero values.
    // The nodes at level 1 were 1, 2, and are 2, 1 after the reversal.
    // The nodes at level 3 were 1, 1, 1, 1, 2, 2, 2, 2, and are 2, 2, 2, 2, 1, 1, 1, 1 after the reversal.
    tree3 := buildTree([]*int{ptr(0), ptr(1), ptr(2), ptr(0), ptr(0), ptr(0), ptr(0), ptr(1), ptr(1), ptr(1), ptr(1), ptr(2), ptr(2), ptr(2), ptr(2)})
    printTree(tree3)
    fmt.Println(reverseOddLevels(tree3))

    tree11 := buildTree([]*int{ptr(2), ptr(3), ptr(5), ptr(8), ptr(13), ptr(21), ptr(34)})
    printTree(tree11)
    fmt.Println(reverseOddLevels1(tree11))
    tree12 := buildTree([]*int{ptr(7), ptr(13), ptr(11)})
    printTree(tree12)
    fmt.Println(reverseOddLevels1(tree12))
    tree13 := buildTree([]*int{ptr(0), ptr(1), ptr(2), ptr(0), ptr(0), ptr(0), ptr(0), ptr(1), ptr(1), ptr(1), ptr(1), ptr(2), ptr(2), ptr(2), ptr(2)})
    printTree(tree13)
    fmt.Println(reverseOddLevels1(tree13))
}