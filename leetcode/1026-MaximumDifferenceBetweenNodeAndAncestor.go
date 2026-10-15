package main

// 1026. Maximum Difference Between Node and Ancestor
// Given the root of a binary tree, 
// find the maximum value v for which there exist different nodes a and b where v = |a.val - b.val| and a is an ancestor of b.
// A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/09/tmp-tree.jpg" />
// Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
// Output: 7
// Explanation: We have various ancestor-node differences, some of which are given below :
// |8 - 3| = 5
// |3 - 7| = 4
// |8 - 1| = 7
// |10 - 13| = 3
// Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/09/tmp-tree-1.jpg" />
// Input: root = [1,null,2,null,0,3]
// Output: 3
 
// Constraints:
//     The number of nodes in the tree is in the range [2, 5000].
//     0 <= Node.val <= 10^5

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
func maxAncestorDiff(root *TreeNode) int {
    var dfs func(*TreeNode, int, int) int
    dfs = func(node *TreeNode, pmax, pmin int) int {
        if node == nil {
            return pmax - pmin
        }
        max := func(x, y int) int { if x > y { return x; }; return y; }
        min := func(x, y int) int { if x < y { return x; }; return y; }
        mx, mi := max(node.Val, pmax), min(node.Val, pmin)
        return max(dfs(node.Left, mx, mi), dfs(node.Right, mx, mi))
    }
    return dfs(root, root.Val, root.Val)
}

func maxAncestorDiff1(root *TreeNode) int {
    var dfs func(root *TreeNode, mi, ma int) int
    dfs = func(root *TreeNode, mi, ma int) int {
        if root == nil {
            return 0
        }
        max := func(x, y int) int { if x > y { return x; }; return y; }
        min := func(x, y int) int { if x < y { return x; }; return y; }
        abs := func(x int) int { if x < 0 { return -x; }; return x; }

        diff := max(abs(root.Val - mi), abs(root.Val - ma))
        mi, ma = min(mi, root.Val), max(ma, root.Val)
        diff = max(diff, dfs(root.Left, mi, ma))
        diff = max(diff, dfs(root.Right, mi, ma))
        return diff
    }
    return dfs(root, root.Val, root.Val)
}

func main() {
    tree1 := buildTree([]*int{ptr(8), ptr(3), ptr(10), ptr(1), ptr(6), nil, ptr(14), nil, nil, ptr(4), ptr(7), ptr(13)})
    printTree(tree1)
    fmt.Println(maxAncestorDiff(tree1)) // 7

    tree2 := buildTree([]*int{ptr(1), nil, ptr(2), nil, ptr(0), ptr(3)})
    printTree(tree2)
    fmt.Println(maxAncestorDiff(tree2)) // 3

    fmt.Println(maxAncestorDiff1(tree1)) // 7
    fmt.Println(maxAncestorDiff1(tree2)) // 3
}
