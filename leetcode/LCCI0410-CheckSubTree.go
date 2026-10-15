package main

// 面试题 04.10. Check SubTree LCCI
// T1 and T2 are two very large binary trees. Create an algorithm to determine if T2 is a subtree of T1.

// A tree T2 is a subtree of T1 if there exists a node n in T1 such that the subtree of n is identical to T2. That is, if you cut off the tree at node n, the two trees would be identical.

// Note: This problem is slightly different from the original problem.

// Example1:
// Input: t1 = [1, 2, 3], t2 = [2]
// Output: true

// Example2:
// Input: t1 = [1, null, 2, 4], t2 = [3, 2]
// Output: false

// Note:
//     The node numbers of both tree are in [0, 20000].

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
func checkSubTree(root, subRoot *TreeNode) bool {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var getHeight func(root *TreeNode) int // 二叉树的最大深度
    getHeight = func(root *TreeNode) int {
        if root == nil { return 0 }
        leftH, rightH := getHeight(root.Left), getHeight(root.Right)
        return max(leftH, rightH) + 1
    }
    var isSameTree func(p, q *TreeNode) bool
    isSameTree = func(p, q *TreeNode) bool { // 是否相同的树
        if p == nil || q == nil { return p == q } // 必须都是 nil
        return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    }
    high := getHeight(subRoot)
    var dfs func(*TreeNode) (int, bool) // 返回 node 的高度，以及是否找到了 subRoot
    dfs = func(node *TreeNode) (int, bool) {
        if node == nil { return 0, false }
        leftH, leftFound := dfs(node.Left)
        rightH, rightFound := dfs(node.Right)
        if leftFound || rightFound { return 0, true }
        nodeH := max(leftH, rightH) + 1
        return nodeH, (nodeH == high && isSameTree(node, subRoot))
    }
    _, found := dfs(root)
    return found
}

func main() {
    // Example1:
    // Input: t1 = [1, 2, 3], t2 = [2]
    // Output: true
    tree11 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree11)
    tree12 := buildTree([]*int{ptr(2)})
    printTree(tree12)
    fmt.Println(checkSubTree(tree11, tree12)) // true
    // Example2:
    // Input: t1 = [1, null, 2, 4], t2 = [3, 2]
    // Output: false
    tree21 := buildTree([]*int{ptr(1), nil, ptr(2), ptr(4)})
    printTree(tree21)
    tree22 := buildTree([]*int{ptr(3), ptr(2)})
    printTree(tree22)
    fmt.Println(checkSubTree(tree21, tree22)) // false
}