package main

// 988. Smallest String Starting From Leaf
// You are given the root of a binary tree where each node has a value in the range [0, 25] representing the letters 'a' to 'z'.

// Return the lexicographically smallest string that starts at a leaf of this tree and ends at the root.
// As a reminder, any shorter prefix of a string is lexicographically smaller.
//     For example, "ab" is lexicographically smaller than "aba".

// A leaf of a node is a node that has no children.

// Example 1:
//         0[a]
//        /     \
//      1[b]    2 c 
//     /  \     /  \
//    3[d] 4 e 3 d  4 e
// <img src="https://assets.leetcode.com/uploads/2019/01/30/tree1.png" />
// Input: root = [0,1,2,3,4,3,4]
// Output: "dba"

// Example 2:
//         25[z]
//        /     \
//      1 b     3 [d]
//     /  \     /   \
//    2 c 3 d  0[a]  2 c
// <img src="https://assets.leetcode.com/uploads/2019/01/30/tree2.png" />
// Input: root = [25,1,3,1,3,0,2]
// Output: "adz"

// Example 3:
//         2[c]
//        /     \
//       2 c     1 [b]
//        \     /   
//        1 b  0[a] 
//        /
//       0 a
// <img src="https://assets.leetcode.com/uploads/2019/02/01/tree3.png" />
// Input: root = [2,2,1,null,1,0,null,0]
// Output: "abc"
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 8500].
//     0 <= Node.val <= 25

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
func smallestFromLeaf(root *TreeNode) string {
    res := ""
    var dfs func(root *TreeNode, str string)
    dfs = func(root *TreeNode, str string) {
        if root == nil {
            return
        }
        str = string('a' + root.Val) + str // 把当前节点表示的字符添加到str头部
        if root.Left == nil && root.Right == nil { // 如果当前节点是叶子节点，把str和s进行比较
            if res == "" || str < res {
                res = str
            }
            return
        }
        dfs(root.Left, str)
        dfs(root.Right, str)
    }
    dfs(root, "")
    return res
}

// dfs
func smallestFromLeaf1(root *TreeNode) string {
    res := ""
    min := func (a, b string) string {
        if a == "" {
            return b
        }
        if b == "" {
            return a
        }
        if a > b {
            return b
        }
        return a
    }
    var dfs func(root *TreeNode, prev string, res *string)
    dfs = func(root *TreeNode, prev string, res *string) {
        prev = string('a' + root.Val) + prev
        if root.Left == nil && root.Right == nil {
            *res = min(*res, prev)
            return
        }
        if root.Left != nil {
            dfs(root.Left, prev, res)
        }
        if root.Right != nil {
            dfs(root.Right, prev, res)
        }
    }
    dfs(root, "", &res)
    return res
}

func main() {
    tree1 := buildTree([]*int{ptr(0), ptr(1), ptr(2), ptr(3), ptr(4), ptr(3), ptr(4)})
    printTree(tree1)
    // Example 1:
    //         0[a]
    //        /     \
    //      1[b]    2 c 
    //     /  \     /  \
    //    3[d] 4 e 3 d  4 e
    // <img src="https://assets.leetcode.com/uploads/2019/01/30/tree1.png" />
    // Input: root = [0,1,2,3,4,3,4]
    // Output: "dba"
    fmt.Println(smallestFromLeaf(tree1)) // dba

    tree2 := buildTree([]*int{ptr(25), ptr(1), ptr(3), ptr(2), ptr(3), ptr(0), ptr(2)})
    printTree(tree2)
    // Example 2:
    //         25[z]
    //        /     \
    //      1 b     3 [d]
    //     /  \     /   \
    //    2 c 3 d  0[a]  2 c
    // <img src="https://assets.leetcode.com/uploads/2019/01/30/tree2.png" />
    // Input: root = [25,1,3,1,3,0,2]
    // Output: "adz"
    fmt.Println(smallestFromLeaf(tree2)) // adz

    tree3 := buildTree([]*int{ptr(2), ptr(2), ptr(1), ptr(1), nil, ptr(0), nil, ptr(0)})
    printTree(tree3)
    // Example 3:
    //         2[c]
    //        /     \
    //       2 c     1 [b]
    //        \     /   
    //        1 b  0[a] 
    //        /
    //       0 a
    // <img src="https://assets.leetcode.com/uploads/2019/02/01/tree3.png" />
    // Input: root = [2,2,1,null,1,0,null,0]
    // Output: "abc"
    fmt.Println(smallestFromLeaf(tree3)) // abc

    fmt.Println(smallestFromLeaf1(tree1)) // dba
    fmt.Println(smallestFromLeaf1(tree2)) // adz
    fmt.Println(smallestFromLeaf1(tree3)) // abc
}