package main

// 437. Path Sum III
// Given the root of a binary tree and an integer targetSum, 
// return the number of paths where the sum of the values along the path equals targetSum.

// The path does not need to start or end at the root or a leaf,
// but it must go downwards (i.e., traveling only from parent nodes to child nodes).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/pathsum3-1-tree.jpg" / >
// Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// Output: 3
// Explanation: The paths that sum to 8 are shown.

// Example 2:
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: 3
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 1000].
//     -10^9 <= Node.val <= 10^9
//     -1000 <= targetSum <= 1000

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
// dfs 
func pathSum(root *TreeNode, targetSum int) int {
    res := 0
    var dfs func(tree *TreeNode, targetSum int, pathStack []*TreeNode)
    dfs = func (tree *TreeNode, targetSum int, pathStack []*TreeNode) {
        if tree == nil {
            return
        }
        sum := 0
        pathStack = append(pathStack, tree)
        for i := len(pathStack) - 1; i >= 0; i-- {
            sum += pathStack[i].Val
            if sum == targetSum {
                res++
            }
        }
        dfs(tree.Left, targetSum, pathStack)
        dfs(tree.Right, targetSum, pathStack)
    }
    dfs(root, targetSum, []*TreeNode{})
    return res
}

func pathSum1(root *TreeNode, targetSum int) int {
    res, m := 0, map[int]int{0:1}
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, cur int) {
        if root == nil {
            return
        }
        cur += root.Val
        res += m[cur - targetSum]
        m[cur]++
        defer func() {m[cur]--}()
        dfs(root.Left, cur)
        dfs(root.Right, cur)
    }
    dfs(root, 0)
    return res
}

func main() {
    tree1 := buildTree([]*int{ptr(10), ptr(5), ptr(-3), ptr(3), ptr(2), nil, ptr(11), ptr(3), ptr(-2), nil, ptr(1)})
    printTree(tree1)
    tree2 := buildTree([]*int{ptr(5), ptr(4), ptr(8), ptr(11), nil, ptr(13), ptr(4), ptr(7), ptr(2), nil, nil, ptr(5), ptr(1)})
    printTree(tree2)
    fmt.Println(pathSum(tree1,8)) // 3
    fmt.Println(pathSum(tree2,22)) // 3

    fmt.Println(pathSum1(tree1,8)) // 3
    fmt.Println(pathSum1(tree2,22)) // 3
}