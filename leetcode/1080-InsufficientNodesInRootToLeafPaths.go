package main

// 1080. Insufficient Nodes in Root to Leaf Paths
// Given the root of a binary tree and an integer limit, 
// delete all insufficient nodes in the tree simultaneously, 
// and return the root of the resulting binary tree.

// A node is insufficient if every root to leaf path intersecting this node has a sum strictly less than limit.

// A leaf is a node with no children.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/06/05/insufficient-11.png" />
// Input: root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
// Output: [1,2,3,4,null,null,7,8,9,null,14]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/06/05/insufficient-3.png" />
// Input: root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
// Output: [5,4,8,11,null,17,4,7,null,null,null,5]

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/06/11/screen-shot-2019-06-11-at-83301-pm.png" />
// Input: root = [1,2,-3,-5,null,4,null], limit = -1
// Output: [1,null,-3,4]

// Constraints:
//     The number of nodes in the tree is in the range [1, 5000].
//     -10^5 <= Node.val <= 10^5
//     -10^9 <= limit <= 10^9

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
// // 递归
// func sufficientSubset(root *TreeNode, limit int) *TreeNode {
//     if root.Left == nil && root.Right == nil {
//         if limit - root.Val > 0 {
//             return nil
//         }
//         return root
//     }
//     if root.Left != nil {
//         root.Left = sufficientSubset(root.Left, limit - root.Val)
//     }
//     if root.Right != nil {
//         root.Right = sufficientSubset(root.Right, limit - root.Val)
//     }
//     if root.Left == nil && root.Right == nil {
//         return nil
//     }
//     return root
// }

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    if root == nil {
        return nil
    }
    limit -= root.Val
    if root.Left == root.Right { // root 是叶子
        if limit > 0 { // 从根到叶子的路径和小于 limit，删除叶子
            return nil
        }
        return root // 否则不删除
    }
    root.Left, root.Right = sufficientSubset(root.Left, limit), sufficientSubset(root.Right, limit)
    if root.Left ==  root.Right { // 如果儿子都被删除，就删 root
        return nil
    }
    return root // 否则不删 root
}

// dfs
func sufficientSubset1(root *TreeNode, limit int) *TreeNode {
    var dfs func(root *TreeNode, limit int) *TreeNode
    dfs = func(root *TreeNode, limit int) *TreeNode {
        if root.Left == nil && root.Right == nil {
            if limit - root.Val > 0 {
                return nil
            }
            return root
        }
        if root.Left != nil {
            root.Left = dfs(root.Left, limit - root.Val)
        }
        if root.Right != nil {
            root.Right = dfs(root.Right, limit - root.Val)
        }
        if root.Left == nil && root.Right == nil {
            return nil
        }
        return root
    }
    return dfs(root, limit)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/06/05/insufficient-11.png" />
    // Input: root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
    // Output: [1,2,3,4,null,null,7,8,9,null,14]
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(-99), ptr(-99), ptr(7), ptr(9), ptr(9), ptr(-99), ptr(-99), ptr(12), ptr(13), ptr(-99), ptr(14)})
    printTree(tree1)
    fmt.Println(sufficientSubset(tree1, 1)) // &{1 0xc000094048 0xc0000940f0} [1,2,3,4,null,null,7,8,9,null,14]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/06/05/insufficient-3.png" />
    // Input: root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
    // Output: [5,4,8,11,null,17,4,7,null,null,null,5]
    tree2 := buildTree([]*int{ptr(5), ptr(4), ptr(8), ptr(11), nil, ptr(17), ptr(4), ptr(7), ptr(1), nil, nil, ptr(5), ptr(3)})
    printTree(tree2)
    fmt.Println(sufficientSubset(tree2, 22)) // &{5 0xc0000941c8 0xc000094228} [5,4,8,11,null,17,4,7,null,null,null,5]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/06/11/screen-shot-2019-06-11-at-83301-pm.png" />
    // Input: root = [1,2,-3,-5,null,4,null], limit = -1
    // Output: [1,null,-3,4]
    tree3 := buildTree([]*int{ptr(1), ptr(2), ptr(-3), ptr(-5), nil, ptr(4)})
    printTree(tree3)
    fmt.Println(sufficientSubset(tree3, -1)) // &{1 <nil> 0xc000094300} [1,null,-3,4]

    fmt.Println(sufficientSubset1(tree1, 1)) // &{1 0xc000094048 0xc0000940f0} [1,2,3,4,null,null,7,8,9,null,14]
    fmt.Println(sufficientSubset1(tree2, 22)) // &{5 0xc0000941c8 0xc000094228} [5,4,8,11,null,17,4,7,null,null,null,5]
    fmt.Println(sufficientSubset1(tree3, -1)) // &{1 <nil> 0xc000094300} [1,null,-3,4]
}
