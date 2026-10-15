package main

// 222. Count Complete Tree Nodes
// Given the root of a complete binary tree, return the number of the nodes in the tree.

// According to Wikipedia, every level, except possibly the last, 
// is completely filled in a complete binary tree, 
// and all nodes in the last level are as far left as possible. 
// It can have between 1 and 2h nodes inclusive at the last level h.

// Design an algorithm that runs in less than O(n) time complexity.

// Example 1:
//         1
//       /   \
//      2     3
//     /  \  /  
//    4    5 6   
// <img src="https://assets.leetcode.com/uploads/2021/01/14/complete.jpg" />
// Input: root = [1,2,3,4,5,6]
// Output: 6

// Example 2:
// Input: root = []
// Output: 0

// Example 3:
// Input: root = [1]
// Output: 1
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 5 * 10^4].
//     0 <= Node.val <= 5 * 10^4
//     The tree is guaranteed to be complete.

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
func countNodes(root *TreeNode) int {
    res := 1 
    if root == nil {
        return 0
    }
    var dfs func(root *TreeNode) 
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }
        if root.Left != nil {
            res++
            dfs(root.Left)
        }
        if root.Right != nil {
            res++
            dfs(root.Right)
        }
    }
    dfs(root)
    return res 
}

func countNodes1(root *TreeNode) int {
    // Define var to track count
    res := 0
    var dfs func (root *TreeNode, currSum *int)
    dfs = func (root *TreeNode, currSum *int) {
        // Cut condition. If node is nil, return
        if root == nil {
            return
        }
        
        *currSum++ // As node is not nil, count it
        if root.Left != nil { // If there's left subtree to check, go and count nodes there
            dfs(root.Left, currSum) 
        }
        if root.Right != nil { // If there's right subtree to check, go and count nodes there
            dfs(root.Right, currSum) 
        }
    }
    // Pass it as reference
    dfs(root, &res)
    return res
}

func main() {
    // Example 1:
    //         1
    //       /   \
    //      2     3
    //     /  \  /  
    //    4    5 6   
    // <img src="https://assets.leetcode.com/uploads/2021/01/14/complete.jpg" />
    // Input: root = [1,2,3,4,5,6]
    // Output: 6
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6)})
    printTree(tree1)
    fmt.Println(countNodes(tree1)) // 6
    // Example 2:
    // Input: root = []
    // Output: 0
    fmt.Println(countNodes(nil)) // 0
    // Example 3:
    // Input: root = [1]
    // Output: 1
    fmt.Println(countNodes(&TreeNode { 1, nil, nil })) // 1

    fmt.Println(countNodes1(tree1)) // 6
    fmt.Println(countNodes1(nil)) // 0
    fmt.Println(countNodes1(&TreeNode { 1, nil, nil })) // 1
}
