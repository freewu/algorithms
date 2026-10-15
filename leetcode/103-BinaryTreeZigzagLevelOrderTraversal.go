package main

// 103. Binary Tree Zigzag Level Order Traversal
// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. 
// (i.e., from left to right, then right to left for the next level and alternate between).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" / >
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[20,9],[15,7]]

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 2000].
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
// 递归 dfs
func zigzagLevelOrder(root *TreeNode) [][]int {
    var res [][]int
    var search func(root *TreeNode, depth int, res *[][]int) 
    search = func (root *TreeNode, depth int, res *[][]int) {
        if root == nil {
            return
        }
        for len(*res) < depth+1 {
            *res = append(*res, []int{})
        }
        // 每一层的顺序是相互反转的，即上一层是从左往右，下一层就是从右往左
        if depth%2 == 0 {
            (*res)[depth] = append((*res)[depth], root.Val)
        } else {
            (*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
        }
        search(root.Left, depth+1, res)
        search(root.Right, depth+1, res)
    }
    // 按层序从上到下遍历一颗树
    search(root, 0, &res)
    return res
}

// BFS
func zigzagLevelOrder1(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    q := []*TreeNode{root}
    size, i, j, lay, tmp, flag := 0, 0, 0, []int{}, []*TreeNode{}, false
    for len(q) > 0 {
        size = len(q)
        tmp = []*TreeNode{}
        lay = make([]int, size)
        j = size - 1
        for i = 0; i < size; i++ {
            root = q[0]
            q = q[1:]
            if !flag {
                lay[i] = root.Val
            } else {
                lay[j] = root.Val
                j--
            }
            if root.Left != nil {
                tmp = append(tmp, root.Left)
            }
            if root.Right != nil {
                tmp = append(tmp, root.Right)
            }

        }
        res = append(res, lay)
        flag = !flag
        q = tmp
    }
    return res
}


func main() {
    tree1 := buildTree([]*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)})
    printTree(tree1)
    tree3 := buildTree([]*int{ptr(1)})
    printTree(tree3)
    fmt.Println(zigzagLevelOrder(tree1)) // [[3],[20,9],[15,7]]
    fmt.Println(zigzagLevelOrder(nil)) // []
    fmt.Println(zigzagLevelOrder(tree3)) // [1]

    fmt.Println(zigzagLevelOrder1(tree1)) // [[3],[20,9],[15,7]]
    fmt.Println(zigzagLevelOrder1(nil)) // []
    fmt.Println(zigzagLevelOrder1(tree3)) // [1]
}
