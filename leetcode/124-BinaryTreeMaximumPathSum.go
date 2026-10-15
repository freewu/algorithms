package main

// 124. Binary Tree Maximum Path Sum
// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. 
// Note that the path does not need to pass through the root.
// The path sum of a path is the sum of the node's values in the path.
// Given the root of a binary tree, return the maximum path sum of any non-empty path.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg" / >
//     1
//    /  \ 
//   2    3
// Input: root = [1,2,3]
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg" / >
//         -10
//         /  \
//        9    20
//            /  \
//           15   7
// Input: root = [-10,9,20,null,null,15,7]
// Output: 42
// Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

// Constraints:
//     The number of nodes in the tree is in the range [1, 3 * 10^4].
//     -1000 <= Node.val <= 1000

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
func maxPathSum(root *TreeNode) int {
    res := -1 << 63
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(root *TreeNode, res *int) int
    dfs = func(root *TreeNode, res *int) int {
        if root == nil {
            return 0
        }
        leftSum := dfs(root.Left, res)
        rightSum := dfs(root.Right, res)
        *res = max(*res, leftSum + rightSum + root.Val)
        return max(max(leftSum, rightSum) + root.Val, 0)
    }
    dfs(root, &res)
    return res
}

func maxPathSum1(root *TreeNode) int {
    res := -1 << 63
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        if r == nil {
            return 0
        }
        left := dfs(r.Left)
        right := dfs(r.Right)
        sum := left + right + r.Val
        res = max(res, sum)
        return max(max(left, right) + r.Val, 0)
    }
    dfs(root)
    return res
}

func maxPathSum2(root *TreeNode) int {
    res := -1 << 31
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0 }
        l, r := max(dfs(node.Left), 0), max(dfs(node.Right), 0)
        sum := node.Val + l + r
        res = max(res, sum)
        return node.Val + max(l, r)
    }
    dfs(root)
    return res
}

func main() {
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree1)
    tree2 := buildTree([]*int{ptr(-10), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)})
    printTree(tree2)
    // Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
    fmt.Println(maxPathSum(tree1)) // 6
    // Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
    fmt.Println(maxPathSum(tree2)) // 42

    fmt.Println(maxPathSum1(tree1)) // 6
    fmt.Println(maxPathSum1(tree2)) // 42
    
    fmt.Println(maxPathSum2(tree1)) // 6
    fmt.Println(maxPathSum2(tree2)) // 42
}
