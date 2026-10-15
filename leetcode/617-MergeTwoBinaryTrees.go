package main

// 617. Merge Two Binary Trees
// You are given two binary trees root1 and root2.
// Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. 
// You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

// Return the merged tree.
// Note: The merging process must start from the root nodes of both trees.

// Example 1: 
//     1             2                    3
//    / \           / \                  / \
//   3   2         1   3         =>     4   5
//  /               \   \              / \   \
// 5                 4   7            5   4   7
// <img src="https://assets.leetcode.com/uploads/2021/02/05/merge.jpg" />
// Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// Output: [3,4,5,5,4,null,7]

// Example 2:
//      1      1                2
//            /      =>        /
//           2                2
// Input: root1 = [1], root2 = [1,2]
// Output: [2,2]
 
// Constraints:
//     The number of nodes in both trees is in the range [0, 2000].
//     -10^4 <= Node.val <= 10^4

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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil { return root2; }
    if root2 == nil { return root1; }
    var merge func(root1 *TreeNode, root2 *TreeNode) 
    merge = func (root1 *TreeNode, root2 *TreeNode) {
        root1.Val += root2.Val
        if root1.Left != nil && root2.Left != nil {
            merge(root1.Left,root2.Left)
        } else if root1.Left == nil {
            root1.Left = root2.Left
        }
        if root1.Right != nil && root2.Right != nil {
            merge(root1.Right,root2.Right)
        } else if root1.Right == nil{
            root1.Right = root2.Right
        } 
    }
    merge(root1,root2)
    return root1   
}

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 != nil {
        return root2
    } else if root1 != nil && root2 == nil {
        return root1
    } else if root1 != nil && root2 != nil {
        root1.Val += root2.Val
        root1.Left = mergeTrees(root1.Left, root2.Left)
        root1.Right = mergeTrees(root1.Right, root2.Right)
    }
    return root1
}

func main() {
    tree11 := buildTree([]*int{ptr(1), ptr(3), ptr(2), ptr(5)})
    printTree(tree11)
    tree12 := buildTree([]*int{ptr(2), ptr(1), ptr(3), nil, ptr(4), nil, ptr(7)})
    printTree(tree12)
    fmt.Println(mergeTrees(tree11,tree12).Val) // 3

    tree21 := buildTree([]*int{ptr(1)})
    printTree(tree21)
    tree22 := buildTree([]*int{ptr(1), ptr(2)})
    printTree(tree22)
    fmt.Println(mergeTrees(tree21,tree22).Val) // 2

    tree111 := buildTree([]*int{ptr(1), ptr(3), ptr(2), ptr(5)})
    printTree(tree111)
    tree112 := buildTree([]*int{ptr(2), ptr(1), ptr(3), nil, ptr(4), nil, ptr(7)})
    printTree(tree112)
    fmt.Println(mergeTrees1(tree111,tree112).Val) // 3
    tree121 := buildTree([]*int{ptr(1)})
    printTree(tree121)
    tree122 := buildTree([]*int{ptr(1), ptr(2)})
    printTree(tree122)
    fmt.Println(mergeTrees1(tree121,tree122).Val) // 2
}