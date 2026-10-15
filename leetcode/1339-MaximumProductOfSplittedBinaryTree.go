package main

// 1339. Maximum Product of Splitted Binary Tree
// Given the root of a binary tree, split the binary tree into two subtrees by removing one edge 
// such that the product of the sums of the subtrees is maximized.

// Return the maximum product of the sums of the two subtrees. Since the answer may be too large, return it modulo 10^9 + 7.
// Note that you need to maximize the answer before taking the mod and not after taking it.

// Example 1:
//         1                               1
//       /   \                               \
//      2      3  =>          2               3 
//    /   \   /             /   \            /
//   4    5   6            4     5          6
//                           11             10       =  110
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/02/sample_1_1699.png" />
// Input: root = [1,2,3,4,5,6]
// Output: 110
// Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11*10)

// Example 2:
//       1 
//         \ 
//          2                 1
//         /  \       =>       \               4
//        3    4                2            /   \
//           /   \             /            5     6
//          5     6           3  
//                            6      *         15       =  90
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/02/sample_2_1699.png" />
// Input: root = [1,null,2,3,4,null,null,5,6]
// Output: 90
// Explanation: Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15*6)

// Constraints:
//     The number of nodes in the tree is in the range [2, 5 * 10^4].
//     1 <= Node.val <= 10^4

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
func maxProduct(root *TreeNode) int {
    dp := []int{}
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0 }
        res := node.Val + dfs(node.Left) + dfs(node.Right)
        dp = append(dp, res)
        return res
    }
    res, sum := 0, dfs(root)
    for _, v := range dp {
        cand := v * (sum - v)
        if cand > res { 
            res = cand 
        }
    }
    return res % 1_000_000_007
}

func maxProduct1(root *TreeNode) int {
    var sumTree func(root *TreeNode) int
    sumTree = func(root *TreeNode) int {
        if root == nil { return 0 }
        return root.Val + sumTree(root.Left) + sumTree(root.Right)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, total := 0, sumTree(root)
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0}
        left, right := dfs(node.Left), dfs(node.Right)
        res = max(res, (total - left) * left)
        res = max(res, (total - right) * right)
        return node.Val + left + right
    }
    dfs(root)
    return res % 1000000007
}

func main() {
    // Example 1:
    //         1                               1
    //       /   \                               \
    //      2      3  =>          2               3 
    //    /   \   /             /   \            /
    //   4    5  6             4     5          6
    //                           11             10       =  110
    // <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/02/sample_1_1699.png" />
    // Input: root = [1,2,3,4,5,6]
    // Output: 110
    // Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11*10)
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6)})
    printTree(tree1)
    fmt.Println(maxProduct(tree1)) // 110
    // Example 2:
    //       1 
    //         \ 
    //          2                 1
    //         /  \       =>       \               4
    //        3    4                2            /   \
    //           /   \             /            5     6
    //          5     6           3  
    //                            6      *         15       =  90
    // <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/02/sample_2_1699.png" />
    // Input: root = [1,null,2,3,4,null,null,5,6]
    // Output: 90
    // Explanation: Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15*6)
    tree2 := buildTree([]*int{ptr(1), nil, ptr(2), ptr(3), ptr(4), nil, nil, ptr(5), ptr(6)})
    printTree(tree2)
    fmt.Println(maxProduct(tree2)) // 90

    tree11 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6)})
    printTree(tree11)
    fmt.Println(maxProduct1(tree11)) // 110
    tree12 := buildTree([]*int{ptr(1), nil, ptr(2), ptr(3), ptr(4), nil, nil, ptr(5), ptr(6)})
    printTree(tree12)
    fmt.Println(maxProduct1(tree12)) // 90
}