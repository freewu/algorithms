package main

// 1457. Pseudo-Palindromic Paths in a Binary Tree
// Given a binary tree where node values are digits from 1 to 9. 
// A path in the binary tree is said to be pseudo-palindromic if at least one permutation of the node values in the path is a palindrome.

// Return the number of pseudo-palindromic paths going from the root node to leaf nodes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/05/06/palindromic_paths_1.png" />
// Input: root = [2,3,1,3,1,null,1]
// Output: 2 
// Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the red path [2,3,3], the green path [2,1,1], and the path [2,3,1]. Among these paths only red path and green path are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in [3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/05/07/palindromic_paths_2.png" />
// Input: root = [2,1,1,1,3,null,null,null,null,null,1]
// Output: 1 
// Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the green path [2,1,1], the path [2,1,3,1], and the path [2,1]. Among these paths only the green path is pseudo-palindromic since [2,1,1] can be rearranged in [1,2,1] (palindrome).

// Example 3:
// Input: root = [9]
// Output: 1

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 9

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths(root *TreeNode) int {
    var dfs func(*TreeNode, int) int
    dfs = func(node *TreeNode, mask int) int {
        if node == nil { return 0 }
        mask ^= 1 << node.Val
        if node.Left == node.Right {
            if mask & (mask - 1) == 0 { return 1 }
            return 0
        }
        left, right := dfs(node.Left, mask), dfs(node.Right, mask)
        return left + right
    }
    return dfs(root, 0)
}

func pseudoPalindromicPaths1(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode, count int)
    dfs = func(node *TreeNode, count int) {
        if node == nil { return }
        if node.Left == nil && node.Right == nil {
            count ^= 1 << node.Val
            if count&(count-1) == 0 { res++ }
            count ^= 1 << node.Val
        }
        count ^= 1 << node.Val
        dfs(node.Left, count) // travel left
        dfs(node.Right, count) // travel right
    }
    dfs(root, 0)
    return res
}
 
func pseudoPalindromicPaths2(root *TreeNode) int {
    numberOfOnes := func(n int) int{
        res := 0
        for n > 0 {
           res += n & 1
           n >>= 1
        }
        return res
    }
    var preorder func(root *TreeNode, count int) int
    preorder = func(root *TreeNode, count int) int {
        if root == nil { return 0 }
        count ^= (1 << root.Val)
        if root.Left == nil && root.Right == nil && numberOfOnes(count) <= 1  { return 1 }
        return preorder(root.Left, count) + preorder(root.Right, count)
    }
    return preorder(root, 0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/05/06/palindromic_paths_1.png" />
    // Input: root = [2,3,1,3,1,null,1]
    // Output: 2 
    // Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the red path [2,3,3], the green path [2,1,1], and the path [2,3,1]. Among these paths only red path and green path are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in [3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).
    tree1 := buildTree([]*int{ptr(2), ptr(3), ptr(1), ptr(3), ptr(1), nil, ptr(1)})
    printTree(tree1)
    fmt.Println(pseudoPalindromicPaths(tree1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/05/07/palindromic_paths_2.png" />
    // Input: root = [2,1,1,1,3,null,null,null,null,null,1]
    // Output: 1 
    // Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the green path [2,1,1], the path [2,1,3,1], and the path [2,1]. Among these paths only the green path is pseudo-palindromic since [2,1,1] can be rearranged in [1,2,1] (palindrome).
    tree2 := buildTree([]*int{ptr(2), ptr(1), ptr(1), ptr(1), ptr(3), nil, nil, nil, nil, nil, ptr(1)})
    printTree(tree2)
    fmt.Println(pseudoPalindromicPaths(tree2)) // 1
    // Example 3:
    // Input: root = [9]
    // Output: 1
    tree3 := buildTree([]*int{ptr(9)})
    printTree(tree3)
    fmt.Println(pseudoPalindromicPaths(tree3)) // 1

    fmt.Println(pseudoPalindromicPaths1(tree1)) // 2
    fmt.Println(pseudoPalindromicPaths1(tree2)) // 1
    fmt.Println(pseudoPalindromicPaths1(tree3)) // 1
    
    fmt.Println(pseudoPalindromicPaths2(tree1)) // 2
    fmt.Println(pseudoPalindromicPaths2(tree2)) // 1
    fmt.Println(pseudoPalindromicPaths2(tree3)) // 1
}