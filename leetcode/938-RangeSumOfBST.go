package main

// 938. Range Sum of BST
// Given the root node of a binary search tree and two integers low and high, 
// return the sum of values of all nodes with a value in the inclusive range [low, high].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/bst1.jpg" />
// Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
// Output: 32
// Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/bst2.jpg" />
// Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
// Output: 23
// Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.
 
// Constraints:
//         The number of nodes in the tree is in the range [1, 2 * 10^4].
//         1 <= Node.val <= 10^5
//         1 <= low <= high <= 10^5
//         All Node.val are unique.

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
// 递归前序遍历
func rangeSumBST(root *TreeNode, low int, high int) int {
    res := 0
    var preOrder func(root *TreeNode, low, high int, res *int)
    preOrder = func(root *TreeNode, low, high int, res *int) {
        if root == nil {
            return
        }
        if low <= root.Val && root.Val <= high { // Val 为 [low,high] 的值累加到 res
            *res += root.Val
        }
        preOrder(root.Left, low, high, res)
        preOrder(root.Right, low, high, res)
    }
    preOrder(root, low, high, &res)
    return res
}

// best solution
func rangeSumBST1(root *TreeNode, low int, high int) int {
    if root==nil {
        return 0
    }
    if root.Val < low {
        return rangeSumBST(root.Right,low,high)
    } else if root.Val == low {
        return rangeSumBST(root.Right,low,high) + root.Val
    } else if root.Val < high {
        return rangeSumBST(root.Right,low,high) + root.Val + rangeSumBST(root.Left,low,high)
    } else if root.Val == high {
        return rangeSumBST(root.Left,low,high) + root.Val
    } else {
        return rangeSumBST(root.Left,low,high)
    }
}

func main() {
    tree1 := buildTree([]*int{ptr(10), ptr(5), ptr(15), ptr(3), ptr(7), nil, ptr(18)})
    printTree(tree1)
    tree2 := buildTree([]*int{ptr(10), ptr(5), ptr(15), ptr(3), ptr(7), ptr(13), ptr(18), ptr(1), nil, ptr(6)})
    printTree(tree2)

    fmt.Println(rangeSumBST(tree1,7,15)) // 32
    fmt.Println(rangeSumBST(tree2,6,10)) // 23

    fmt.Println(rangeSumBST1(tree1,7,15)) // 32
    fmt.Println(rangeSumBST1(tree2,6,10)) // 23
 }