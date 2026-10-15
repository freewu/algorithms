package main

// LCR 143. 子结构判断
// 给定两棵二叉树 tree1 和 tree2，判断 tree2 是否以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
// 注意，空树 不会是以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。

// 示例 1：
// <img src="https://pic.leetcode.cn/1694684670-vwyIgY-two_tree.png" />
// 输入：tree1 = [1,7,5], tree2 = [6,1]
// 输出：false
// 解释：tree2 与 tree1 的一个子树没有相同的结构和节点值。

// 示例 2：
// <img src="https://pic.leetcode.cn/1694685602-myWXCv-two_tree_2.png" />
// 输入：tree1 = [3,6,7,1,8], tree2 = [6,1]
// 输出：true
// 解释：tree2 与 tree1 的一个子树拥有相同的结构和节点值。即 6 - > 1。

// 提示：
//     0 <= 节点个数 <= 10000

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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if B == nil || A == nil {
        return false
    }
    return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(a, b *TreeNode) bool {
    if b == nil { return true } // b 可以有
    if a == nil { return false } // a 必须有
    if a.Val == b.Val { return recur(a.Left, b.Left) && recur(a.Right, b.Right) }
    return false
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode.cn/1694684670-vwyIgY-two_tree.png" />
    // 输入：tree1 = [1,7,5], tree2 = [6,1]
    // 输出：false
    // 解释：tree2 与 tree1 的一个子树没有相同的结构和节点值。
    tree11 := &TreeNode{
        1, 
        &TreeNode{7, nil, nil},
        &TreeNode{5, nil, nil},
    }
    tree12 := &TreeNode{
        6, 
        &TreeNode{1, nil, nil},
        nil,
    }
    fmt.Println(isSubStructure(tree11,tree12)) // false
    // 示例 2：
    // <img src="https://pic.leetcode.cn/1694685602-myWXCv-two_tree_2.png" />
    // 输入：tree1 = [3,6,7,1,8], tree2 = [6,1]
    // 输出：true
    // 解释：tree2 与 tree1 的一个子树拥有相同的结构和节点值。即 6 - > 1。
    tree21 := &TreeNode{
        3, 
        &TreeNode{6, &TreeNode{1, nil, nil}, &TreeNode{8, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    tree22 := &TreeNode{
        6, 
        &TreeNode{1, nil, nil},
        nil,
    }
    fmt.Println(isSubStructure(tree21,tree22)) // true

    tree31 := &TreeNode{
        10, 
        &TreeNode{12, &TreeNode{8, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{6,  &TreeNode{11, nil, nil}, nil},
    }
    tree32 := &TreeNode{
        10, 
        &TreeNode{12, &TreeNode{8, nil, nil}, nil, },
        &TreeNode{6,  nil, nil, },
    }
    fmt.Println(isSubStructure(tree31,tree32)) // true

    fmt.Println(isSubStructure(nil,nil)) // false
}