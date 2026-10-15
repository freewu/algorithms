package main

// 971. Flip Binary Tree To Match Preorder Traversal
// You are given the root of a binary tree with n nodes, where each node is uniquely assigned a value from 1 to n. 
// You are also given a sequence of n values voyage, which is the desired pre-order traversal of the binary tree.

// Any node in the binary tree can be flipped by swapping its left and right subtrees. 
// For example, flipping node 1 will have the following effect:
// <img src="https://assets.leetcode.com/uploads/2021/02/15/fliptree.jpg" />
//         1                   1
//       /   \                /  \
//      2     3       =>     3    2
//           /  \          /   \
//          4    5        4     5

// Flip the smallest number of nodes so that the pre-order traversal of the tree matches voyage.

// Return a list of the values of all flipped nodes. You may return the answer in any order. 
// If it is impossible to flip the nodes in the tree to make the pre-order traversal match voyage, return the list [-1].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-01.png" />
//      1
//     /
//    2
// Input: root = [1,2], voyage = [2,1]
// Output: [-1]
// Explanation: It is impossible to flip the nodes such that the pre-order traversal matches voyage.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-02.png" />
//      1
//     /  \
//    2    3
// Input: root = [1,2,3], voyage = [1,3,2]
// Output: [1]
// Explanation: Flipping node 1 swaps nodes 2 and 3, so the pre-order traversal matches voyage.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-02.png" />
//      1
//     /  \
//    2    3
// Input: root = [1,2,3], voyage = [1,2,3]
// Output: []
// Explanation: The tree's pre-order traversal already matches voyage, so no nodes need to be flipped.

// Constraints:
//     The number of nodes in the tree is n.
//     n == voyage.length
//     1 <= n <= 100
//     1 <= Node.val, voyage[i] <= n
//     All the values in the tree are unique.
//     All the values in voyage are unique.

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
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    stack := []*TreeNode{ root }
    res, i := []int{}, 0
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if voyage[i] != node.Val { return []int{-1} }
        i++
        if node.Right != nil && node.Left != nil && node.Left.Val != voyage[i] {
            node.Left, node.Right = node.Right, node.Left
            res = append(res, node.Val)
        }
        if node.Right != nil { stack = append(stack, node.Right) }
        if node.Left != nil { stack = append(stack, node.Left) }
    }
    return res
}

// dfs
func flipMatchVoyage1(root *TreeNode, voyage []int) []int {
    res, i := []int{}, 0
    var dfs func(*TreeNode) bool
    dfs = func(root *TreeNode) bool {
        if root == nil {
            return true
        } else if root.Val != voyage[i] {
            return false
        }
        i++
        if root.Left != nil && root.Left.Val != voyage[i] {
            res = append(res, root.Val)
            return dfs(root.Right) && dfs(root.Left)
        }
        return dfs(root.Left) && dfs(root.Right)
    }
    if dfs(root) {
        return res
    }
    return []int{-1}
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-01.png" />
    //      1
    //     /
    //    2
    // Input: root = [1,2], voyage = [2,1]
    // Output: [-1]
    // Explanation: It is impossible to flip the nodes such that the pre-order traversal matches voyage.
    tree1 := buildTree([]*int{ptr(1), ptr(2)})
    printTree(tree1)
    fmt.Println(flipMatchVoyage(tree1,[]int{2,1})) // [-1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-02.png" />
    //      1
    //     /  \
    //    2    3
    // Input: root = [1,2,3], voyage = [1,3,2]
    // Output: [1]
    // Explanation: Flipping node 1 swaps nodes 2 and 3, so the pre-order traversal matches voyage.
    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree2)
    fmt.Println(flipMatchVoyage(tree2,[]int{1,3,2})) // [1]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-02.png" />
    //      1
    //     /  \
    //    2    3
    // Input: root = [1,2,3], voyage = [1,2,3]
    // Output: []
    // Explanation: The tree's pre-order traversal already matches voyage, so no nodes need to be flipped.
    tree3 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree3)
    fmt.Println(flipMatchVoyage(tree3,[]int{1,2,3})) // []

    tree11 := buildTree([]*int{ptr(1), ptr(2)})
    printTree(tree11)
    fmt.Println(flipMatchVoyage1(tree11,[]int{2,1})) // [-1]
    tree12 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree12)
    fmt.Println(flipMatchVoyage1(tree12,[]int{1,3,2})) // [1]
    tree13 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree13)
    fmt.Println(flipMatchVoyage1(tree13,[]int{1,2,3})) // []
}