package main

import "fmt"

// 2641. Cousins in Binary Tree II
// Given the root of a binary tree, replace the value of each node in the tree with the sum of all its cousins' values.
// Two nodes of a binary tree are cousins if they have the same depth with different parents.
// Return the root of the modified tree.
// Note that the depth of a node is the number of edges in the path from the root node to it.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/01/11/example11.png" />
// Input: root = [5,4,9,1,10,null,7]
// Output: [0,0,0,7,7,null,11]
// Explanation: 
// 		The diagram above shows the initial binary tree and the binary tree after changing the value of each node.
// 		- Node with value 5 does not have any cousins so its sum is 0.
// 		- Node with value 4 does not have any cousins so its sum is 0.
// 		- Node with value 9 does not have any cousins so its sum is 0.
// 		- Node with value 1 has a cousin with value 7 so its sum is 7.
// 		- Node with value 10 has a cousin with value 7 so its sum is 7.
// 		- Node with value 7 has cousins with values 1 and 10 so its sum is 11.
	
// Example 2:
// Input: root = [3,1,2]
// Output: [0,0,0]
// Explanation: 
// 		The diagram above shows the initial binary tree and the binary tree after changing the value of each node.
// 		- Node with value 3 does not have any cousins so its sum is 0.
// 		- Node with value 1 does not have any cousins so its sum is 0.
// 		- Node with value 2 does not have any cousins so its sum is 0.
 
// Constraints:
// 		The number of nodes in the tree is in the range [1, 10^5].
// 		1 <= Node.val <= 10^4

// 如果两个节点在树中有相同的深度且它们的父节点不同，那么它们互为 堂兄弟 

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    children := []*TreeNode{root}
    root.Val = 0
    // Each iteration children list contains all elements with the same depth
    for len(children) > 0 {
        parents := children
        children = []*TreeNode{}
        // Find sum of all children in this depth
        sum := 0
        for _, parent := range parents {
            if parent.Left != nil { 
                sum += parent.Left.Val 
            }
            if parent.Right != nil { 
                sum += parent.Right.Val 
            } 
        }
        // Update children values:
        // Calculated sum minus the current childer values (if they exist)
        // Also add all children to the list for the next iteration (depth+1)
        for _, parent := range parents {
            reduce := 0
            if parent.Left != nil { 
                reduce += parent.Left.Val 
            }
            if parent.Right != nil { 
                reduce += parent.Right.Val 
            }
            if parent.Left != nil {
                parent.Left.Val = sum - reduce
                children = append(children, parent.Left)
            }
            if parent.Right != nil {
                parent.Right.Val = sum - reduce
                children = append(children, parent.Right)
            }
        }
    }
    return root
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
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
 
func printBinaryTree(root *TreeNode) {
    if root == nil {
        return
    }
    
    fmt.Printf("%d ", root.Val) // 先输出当前节点的值
    
    printBinaryTree(root.Left)  // 再遍历左子树
    printBinaryTree(root.Right) // 最后遍历右子树
}
 
func main() {
    // 创建一个二叉树
    tree := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6), ptr(7)})
    printTree(tree)
    printBinaryTree(tree)
}