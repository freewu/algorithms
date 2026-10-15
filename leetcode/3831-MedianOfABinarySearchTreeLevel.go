package main

// 3831. Median of a Binary Search Tree Level
// You are given the root of a Binary Search Tree (BST) and an integer level.

// The root node is at level 0. 
// Each level represents the distance from the root.

// Return the median value of all node values present at the given level. 
// If the level does not exist or contains no nodes, return -1.

// The median is defined as the middle element after sorting the values at that level in non-decreasing order. 
// If the number of values at that level is even, return the upper median (the larger of the two middle elements after sorting).

// Example 1:
// ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-20801pm.png" />
// Input: root = [4,null,5,null,7], level = 2
// Output: 7
// Explanation:
// The nodes at level = 2 are [7]. The median value is 7.

// Example 2:
// ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-20926pm.png" />
// Input: root = [6,3,8], level = 1
// Output: 8
// Explanation:
// The nodes at level = 1 are [3, 8]. There are two possible median values, so the larger one 8 is the answer.

// Example 3:
// ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-21001pm.png" />
// Input: root = [2,1], level = 2
// Output: -1
// Explanation:
// There is no node present at level = 2​​​​​​​, so the answer is -1.

// Constraints:
//     The number of nodes in the tree is in the range [1, 2 * 10^5].
//     1 <= Node.val <= 10^6
//     0 <= level <= 2 * 10^​​​​​​​5

import "fmt"
import "sort"

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
func levelMedian(root *TreeNode, level int) int {
    if root == nil { return -1 } // 处理根节点为空的边界情况
    queue := []*TreeNode{root} // 使用队列进行层序遍历，初始时加入根节点
    curr := 0
    // 遍历每一层
    for len(queue) > 0 {
        n := len(queue) // 当前层的节点数量
        arr := make([]int, 0, n) // 存储当前层的所有节点值
        for i := 0; i < n; i++ { // 处理当前层的所有节点 
            // 取出队首节点
            node := queue[0]
            queue = queue[1:]
            arr = append(arr, node.Val) // 记录当前节点值
            if node.Left != nil { // 将子节点加入队列（下一层）
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        if curr == level { // 如果当前层级是目标层级，计算中位数
            if len(arr) == 0 {  return -1 }
            sort.Ints(arr) // 替换为标准库排序
            return arr[len(arr)/2]
        }
        curr++ // 进入下一层
    }
    return -1 // 目标层级不存在，返回-1
}

func levelMedian1(r *TreeNode, l int) int {
    if l == 0 { 
        return r.Val 
    }
    queue := []*TreeNode{ r }
    for len(queue) > 0 && l > 0 {
        nq := []*TreeNode{}
        for k := range queue {
            if queue[k].Left != nil {
                nq = append(nq, queue[k].Left)
            }
            if queue[k].Right != nil {
                nq = append(nq, queue[k].Right)
            }
        }
        if len(nq) > 0 {
            queue = nq
            l --
        } else { 
            return -1 
        }
    }
    v := queue[len(queue) / 2]
    return v.Val
}

func main() {
    // Example 1:
    // ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-20801pm.png" />
    // Input: root = [4,null,5,null,7], level = 2
    // Output: 7
    // Explanation:
    // The nodes at level = 2 are [7]. The median value is 7.
    tree1 := &TreeNode { 
        4, 
        nil, 
        &TreeNode { 5, nil, &TreeNode { 7, nil, nil, }, },
    }
    fmt.Println(levelMedian(tree1, 2)) // 7
    // Example 2:
    // ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-20926pm.png" />
    // Input: root = [6,3,8], level = 1
    // Output: 8
    // Explanation:
    // The nodes at level = 1 are [3, 8]. There are two possible median values, so the larger one 8 is the answer.
    tree2 := &TreeNode { 
        6, 
        &TreeNode { 3, nil, nil, }, 
        &TreeNode { 8, nil, nil, },
    }
    fmt.Println(levelMedian(tree2, 1)) // 8
    // Example 3:
    // ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-21001pm.png" />
    // Input: root = [2,1], level = 2
    // Output: -1
    // Explanation:
    // There is no node present at level = 2​​​​​​​, so the answer is -1.
    tree3 := &TreeNode { 
        2, 
        &TreeNode { 1, nil, nil, }, 
        nil,
    }
    fmt.Println(levelMedian(tree3, 2)) // -1

    fmt.Println(levelMedian1(tree1, 2)) // 7
    fmt.Println(levelMedian1(tree2, 1)) // 8
    fmt.Println(levelMedian1(tree3, 2)) // -1
}