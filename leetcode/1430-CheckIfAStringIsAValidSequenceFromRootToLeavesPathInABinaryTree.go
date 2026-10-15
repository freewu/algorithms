package main

// 1430. Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree
// Given a binary tree where each path going from the root to any leaf form a valid sequence, 
// check if a given string is a valid sequence in such binary tree. 

// We get the given string from the concatenation of an array of integers arr 
// and the concatenation of all values of the nodes along a path results in a sequence in the given binary tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_1.png" />
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,0,1]
// Output: true
// Explanation: 
// The path 0 -> 1 -> 0 -> 1 is a valid sequence (green color in the figure). 
// Other valid sequences are: 
// 0 -> 1 -> 1 -> 0 
// 0 -> 0 -> 0

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_2.png" />
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,0,1]
// Output: false 
// Explanation: The path 0 -> 0 -> 1 does not exist, therefore it is not even a sequence.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_3.png" />
// Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,1]
// Output: false
// Explanation: The path 0 -> 1 -> 1 is a sequence, but it is not a valid sequence.

// Constraints:
//     1 <= arr.length <= 5000
//     0 <= arr[i] <= 9
//     Each node's value is between [0 - 9].

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
// 递归
func isValidSequence(root *TreeNode, arr []int) bool {
    if root == nil || len(arr) == 0 || root.Val != arr[0] {
        return false
    }
    if root.Left == root.Right { // 叶子
        return len(arr) == 1
    }
    return isValidSequence(root.Left, arr[1:]) || isValidSequence(root.Right, arr[1:])
}

// dfs
func isValidSequence1(root *TreeNode, arr []int) bool {
    if root.Val != arr[0] {
        return false
    }
    res := false
    var dfs func(node *TreeNode, depth int, valid bool)
    dfs = func(node *TreeNode, depth int, valid bool) {
        if node == nil {
            return
        }
        cvalid := valid && depth < len(arr) && node.Val == arr[depth]
        if cvalid && node.Left == nil && node.Right == nil && depth == len(arr)-1 {
            res = true
        }
        // 如果当前已经无效, 则停止搜索其子节点
        if cvalid {
            depth++
            dfs(node.Left, depth,  cvalid)
            dfs(node.Right, depth, cvalid)
        }
    }
    dfs(root, 0, true)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_1.png" />
    // Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,0,1]
    // Output: true
    // Explanation: 
    // The path 0 -> 1 -> 0 -> 1 is a valid sequence (green color in the figure). 
    // Other valid sequences are: 
    // 0 -> 1 -> 1 -> 0 
    // 0 -> 0 -> 0
    tree1 := buildTree([]*int{ptr(0), ptr(1), ptr(0), ptr(0), ptr(1), ptr(0), nil, nil, ptr(1), ptr(0), ptr(0)})
    printTree(tree1)
    fmt.Println(isValidSequence(tree1, []int{0,1,0,1})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_2.png" />
    // Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,0,1]
    // Output: false 
    // Explanation: The path 0 -> 0 -> 1 does not exist, therefore it is not even a sequence.
    tree2 := buildTree([]*int{ptr(0), ptr(1), ptr(0), ptr(0), ptr(1), ptr(0), nil, nil, ptr(1), ptr(0), ptr(0)})
    printTree(tree2)
    fmt.Println(isValidSequence(tree2, []int{0,0,1})) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/12/18/leetcode_testcase_3.png" />
    // Input: root = [0,1,0,0,1,0,null,null,1,0,0], arr = [0,1,1]
    // Output: false
    // Explanation: The path 0 -> 1 -> 1 is a sequence, but it is not a valid sequence.
    tree3 := buildTree([]*int{ptr(0), ptr(1), ptr(0), ptr(0), ptr(1), ptr(0), nil, nil, ptr(1), ptr(0), ptr(0)})
    printTree(tree3)
    fmt.Println(isValidSequence(tree3, []int{0,1,1})) // false

    fmt.Println(isValidSequence1(tree1, []int{0,1,0,1})) // true
    fmt.Println(isValidSequence1(tree2, []int{0,0,1})) // false
    fmt.Println(isValidSequence1(tree3, []int{0,1,1})) // false
}