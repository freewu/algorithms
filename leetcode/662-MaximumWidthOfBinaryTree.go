package main

// 662. Maximum Width of Binary Tree
// Given the root of a binary tree, return the maximum width of the given tree.
// The maximum width of a tree is the maximum width among all levels.
// The width of one level is defined as the length between the end-nodes (the leftmost and rightmost non-null nodes), 
// where the null nodes between the end-nodes that would be present in a complete binary tree extending down to that level are also counted into the length calculation.
// It is guaranteed that the answer will in the range of a 32-bit signed integer.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/width1-tree.jpg" />
// Input: root = [1,3,2,5,3,null,9]
// Output: 4
// Explanation: The maximum width exists in the third level with length 4 (5,3,null,9).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/14/maximum-width-of-binary-tree-v3.jpg" />
// Input: root = [1,3,2,5,null,null,9,6,null,7]
// Output: 7
// Explanation: The maximum width exists in the fourth level with length 7 (6,null,null,null,null,null,7).

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/width3-tree.jpg" />
// Input: root = [1,3,2,5]
// Output: 2
// Explanation: The maximum width exists in the second level with length 2 (3,2).

// Constraints:
//     The number of nodes in the tree is in the range [1, 3000].
//     -100 <= Node.val <= 100

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
func widthOfBinaryTree(root *TreeNode) int {
    res := 0
    lst := []int{}
    var dfs func(root *TreeNode, id, depth int, lst *[]int)
    dfs = func (root *TreeNode, id, depth int, lst *[]int) {
        if root == nil {
            return
        }
        if depth >= len(*lst) {
            *lst = append(*lst, id)
        }
        if (id + 1 - (*lst)[depth]) > res {
            res = id + 1 - (*lst)[depth]
        }
        if root.Left != nil  { dfs(root.Left, id*2, depth+1, lst)    }
        if root.Right != nil { dfs(root.Right, id*2+1, depth+1, lst) }
    }
    dfs(root, 1, 0, &lst)
    return res
}

// bfs
func widthOfBinaryTree1(root *TreeNode) int {
    type Node struct {
        id int
        tree *TreeNode
    }
    s, h, t, res := make([]*Node, 1), 0, 1, 1
    s[0] = &Node{
        id: 1,
        tree: root,
    }
    n := s[0]
    for h < t {
        x := s[h]
        if x.tree.Left != nil {
            s = append(s, &Node{
                id: x.id*2,
                tree: x.tree.Left,
            })
            t++
        }
        if x.tree.Right != nil {
            s = append(s, &Node{
                id: x.id*2+1,
                tree: x.tree.Right,
            })
            t++
        }
        if x == n && h < t-1 {
            w := s[t-1].id - s[h+1].id + 1
            if w > res {
                res = w
            }
            n = s[t-1]
        }
        h++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/width1-tree.jpg" />
    // Input: root = [1,3,2,5,3,null,9]
    // Output: 4
    // Explanation: The maximum width exists in the third level with length 4 (5,3,null,9).
    tree1 := buildTree([]*int{ptr(1), ptr(3), ptr(2), ptr(5), ptr(3), nil, ptr(9)})
    printTree(tree1)
    fmt.Println(widthOfBinaryTree(tree1)) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/14/maximum-width-of-binary-tree-v3.jpg" />
    // Input: root = [1,3,2,5,null,null,9,6,null,7]
    // Output: 7
    // Explanation: The maximum width exists in the fourth level with length 7 (6,null,null,null,null,null,7).
    tree2 := buildTree([]*int{ptr(1), ptr(3), ptr(2), ptr(5), nil, nil, ptr(9), ptr(6), nil, ptr(7)})
    printTree(tree2)
    fmt.Println(widthOfBinaryTree(tree2)) // 7
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/width3-tree.jpg" />
    // Input: root = [1,3,2,5]
    // Output: 2
    // Explanation: The maximum width exists in the second level with length 2 (3,2).
    tree3 := buildTree([]*int{ptr(1), ptr(3), ptr(2), ptr(5)})
    printTree(tree3)
    fmt.Println(widthOfBinaryTree(tree3)) // 2

    fmt.Println(widthOfBinaryTree1(tree1)) // 4
    fmt.Println(widthOfBinaryTree1(tree2)) // 7
    fmt.Println(widthOfBinaryTree1(tree3)) // 2
}