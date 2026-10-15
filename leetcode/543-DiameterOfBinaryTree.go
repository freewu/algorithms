package main

// 543. Diameter of Binary Tree
// Given the root of a binary tree, return the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. 
// This path may or may not pass through the root.
// The length of a path between two nodes is represented by the number of edges between them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/06/diamtree.jpg" />
// Input: root = [1,2,3,4,5]
// Output: 3
// Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

// Example 2:
// Input: root = [1,2]
// Output: 1
 
// Constraints:
//         The number of nodes in the tree is in the range [1, 10^4].
//         -100 <= Node.val <= 100

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
func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    var checkDiameter func(root *TreeNode, res *int) int 
    checkDiameter = func(root *TreeNode, res *int) int {
        if root == nil {
            return 0
        }
        max := func (x, y int) int { if x > y { return x; }; return y; }
        left := checkDiameter(root.Left, res)
        right := checkDiameter(root.Right, res)
        *res = max(*res, left + right)
        return max(left, right) + 1
    }
    checkDiameter(root, &res)
    return res
}

func diameterOfBinaryTree1(root *TreeNode) int {
    res := 0
    var dfs func(*TreeNode) int
    dfs = func (root *TreeNode) int {
        if root == nil {
            return 0
        }
        left := dfs(root.Left)
        right := dfs(root.Right)
        max := func (x, y int) int { if x > y { return x; }; return y; }
        res = max(res, left + right)
        return max(left, right) + 1
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/06/diamtree.jpg" />
    // Input: root = [1,2,3,4,5]
    // Output: 3
    // Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
    fmt.Println(diameterOfBinaryTree(
        &TreeNode {
            1,
            &TreeNode{
                2, 
                &TreeNode{4, nil, nil},
                &TreeNode{5, nil, nil},
            },
            &TreeNode{3, nil, nil},
        },
    )) // 3
    // Example 2:
    // Input: root = [1,2]
    // Output: 1
    fmt.Println(diameterOfBinaryTree(
        &TreeNode {
            1,
            &TreeNode{2, nil, nil},
            nil,
        },
    )) // 1

    fmt.Println(diameterOfBinaryTree1(
        &TreeNode {
            1,
            &TreeNode{
                2, 
                &TreeNode{4, nil, nil},
                &TreeNode{5, nil, nil},
            },
            &TreeNode{3, nil, nil},
        },
    )) // 3
    fmt.Println(diameterOfBinaryTree1(
        &TreeNode {
            1,
            &TreeNode{2, nil, nil},
            nil,
        },
    )) // 1
 }