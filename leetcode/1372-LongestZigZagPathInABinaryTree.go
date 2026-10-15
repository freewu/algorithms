package main

// 1372. Longest ZigZag Path in a Binary Tree
// You are given the root of a binary tree.
// A ZigZag path for a binary tree is defined as follow:
//     Choose any node in the binary tree and a direction (right or left).
//     If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
//     Change the direction from right to left or from left to right.
//     Repeat the second and third steps until you can't move in the tree.

// Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).
// Return the longest ZigZag path contained in that tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/01/22/sample_1_1702.png" />
// Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
// Output: 3
// Explanation: Longest ZigZag path in blue nodes (right -> left -> right).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/01/22/sample_2_1702.png" />
// Input: root = [1,1,1,null,1,null,null,1,1,null,1]
// Output: 4
// Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).

// Example 3:
// Input: root = [1]
// Output: 0
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 5 * 10^4].
//     1 <= Node.val <= 100


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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    res := 0
    var dfs func(n *TreeNode, max *int, count int, parentRight bool)
    dfs = func(n *TreeNode, max *int, count int, parentRight bool) {
        if n == nil { return; }
        if count > *max {
            *max = count
        }
        if parentRight {
            dfs(n.Left, max, count+1, false)
            dfs(n.Right, max, 1, true)
        } else {
            dfs(n.Left, max, 1, false)
            dfs(n.Right, max, count+1, true)
        }
    }
    dfs(root, &res, 0 , false)
    return res
}

func longestZigZag1(root *TreeNode) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node *TreeNode, lastIsLeft bool, len int)
    dfs = func(node *TreeNode, lastIsLeft bool, len int) {
        if node == nil { return; }
        res = max(res, len)
        if lastIsLeft {
            dfs(node.Left, true, 1) // 方向不变，从头开始计算
            dfs(node.Right, false, len+1) // 方向改变，长度加1
            return
        }
        dfs(node.Left, true, len+1) // 方向改变，长度加1
        dfs(node.Right, false, 1) // 方向不变，从头开始计算
    }
    dfs(root, false, 0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/01/22/sample_1_1702.png" />
    // Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
    // Output: 3
    // Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
    tree1 := buildTree([]*int{ptr(1), nil, ptr(1), ptr(1), ptr(1), nil, nil, ptr(1), ptr(1), nil, ptr(1), nil, nil, nil, ptr(1)})
    printTree(tree1)
    fmt.Println(longestZigZag(tree1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/01/22/sample_2_1702.png" />
    // Input: root = [1,1,1,null,1,null,null,1,1,null,1]
    // Output: 4
    // Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
    tree2 := buildTree([]*int{ptr(1), ptr(1), ptr(1), nil, ptr(1), nil, nil, ptr(1), ptr(1), nil, ptr(1)})
    printTree(tree2)
    fmt.Println(longestZigZag(tree2)) // 4
    // Example 3:
    // Input: root = [1]
    // Output: 0
    tree3 := buildTree([]*int{ptr(1)})
    printTree(tree3)
    fmt.Println(longestZigZag(tree3)) // 0

    fmt.Println(longestZigZag1(tree1)) // 3
    fmt.Println(longestZigZag1(tree2)) // 4
    fmt.Println(longestZigZag1(tree3)) // 0
}