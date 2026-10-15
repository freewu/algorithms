package main

// 965. Univalued Binary Tree
// A binary tree is uni-valued if every node in the tree has the same value.
// Given the root of a binary tree, return true if the given tree is uni-valued, or false otherwise.

// Example 1:
//        1
//       /  \
//      1    1
//     /  \   \ 
//    1    1   1
// <img src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_1.png" />
// Input: root = [1,1,1,1,1,null,1]
// Output: true

// Example 2:
//        2
//       /  \
//      2    2
//     /  \ 
//    5    2 
// <img src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_2.png" />
// Input: root = [2,2,2,5,2]
// Output: false

// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
//     0 <= Node.val < 100

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
// bfs
func isUnivalTree(root *TreeNode) bool {
    path := []*TreeNode{root}
    s := root.Val
    for i,j := 0,0 ; i<len(path) ; i = j {
        h := len(path)
        for ;j < h; j++ {
            if path[j].Val != s { return false }
            if path[j].Left != nil {
                if path[j].Left.Val != s { return false }
                path = append(path, path[j].Left)
            }
            if path[j].Right != nil {
                if path[j].Right.Val != s { return false }
                path = append(path, path[j].Right)
            }
        }
    }
    return true
}

// dfs
func isUnivalTree1(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var dfs func(node *TreeNode, val int) bool
    dfs = func(node *TreeNode, val int) bool {
        if node == nil {
            return true
        } else if node.Val != val {
            return false
        }
        if node.Left != nil && !dfs(node.Left, val) {
            return false
        }
        if node.Right != nil && !dfs(node.Right, val) {
            return false
        }
        return true
    }
    if root.Left != nil && !dfs(root.Left, root.Val) {
        return false
    }
    if root.Right != nil && !dfs(root.Right, root.Val) {
        return false
    }
    return true
}

func main() {
    // Example 1:
    //        1
    //       /  \
    //      1    1
    //     /  \   \ 
    //    1    1   1
    // <img src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_1.png" />
    // Input: root = [1,1,1,1,1,null,1]
    // Output: true
    tree1 := buildTree([]*int{ptr(1), ptr(1), ptr(1), ptr(1), ptr(1), nil, ptr(1)})
    printTree(tree1)
    fmt.Println(isUnivalTree(tree1)) // true
    // Example 2:
    //        2
    //       /  \
    //      2    2
    //     /  \ 
    //    5    2 
    // <img src="https://assets.leetcode.com/uploads/2018/12/28/unival_bst_2.png" />
    // Input: root = [2,2,2,5,2]
    // Output: false
    tree2 := buildTree([]*int{ptr(2), ptr(2), ptr(2), ptr(5), ptr(2)})
    printTree(tree2)
    fmt.Println(isUnivalTree(tree2)) // true

    fmt.Println(isUnivalTree1(tree1)) // true
    fmt.Println(isUnivalTree1(tree2)) // true
}