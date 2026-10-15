package main

import "fmt"

// 993. Cousins in Binary Tree
// Given the root of a binary tree with unique values and the values of two different nodes of the tree x and y, 
// return true if the nodes corresponding to the values x and y in the tree are cousins, or false otherwise.

// Two nodes of a binary tree are cousins if they have the same depth with different parents.
// Note that in a binary tree, the root node is at the depth 0, 
// and children of each depth k node are at the depth k + 1.

// Example 1:
// <img src = "https://assets.leetcode.com/uploads/2019/02/12/q1248-01.png" />
// Input: root = [1,2,3,4], x = 4, y = 3
// Output: false

// Example 2:
// <img src = "https://assets.leetcode.com/uploads/2019/02/12/q1248-02.png" />
// Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
// Output: true

// Example 3:
// <img src = "https://assets.leetcode.com/uploads/2019/02/12/q1248-03.png" />
// Input: root = [1,2,3,null,4], x = 2, y = 3
// Output: false

// Constraints:
//     The number of nodes in the tree is in the range [2, 100].
//     1 <= Node.val <= 100
//     Each node has a unique value.
//     x != y
//     x and y are exist in the tree.


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
// 解法一 递归
func isCousins(root *TreeNode, x int, y int) bool {
    if root == nil {
        return false
    }
    var findLevel func(root *TreeNode, x, level int) int
    findLevel = func(root *TreeNode, x, level int) int {
        if root == nil {
            return 0
        }
        if root.Val != x {
            leftLevel, rightLevel := findLevel(root.Left, x, level+1), findLevel(root.Right, x, level+1)
            if leftLevel == 0 {
                return rightLevel
            }
            return leftLevel
        }
        return level
    }
    var haveSameParents func(root *TreeNode, x, y int) bool
    haveSameParents = func (root *TreeNode, x, y int) bool {
        if root == nil {
            return false
        }
        if (root.Left != nil && root.Right != nil && root.Left.Val == x && root.Right.Val == y) ||
            (root.Left != nil && root.Right != nil && root.Left.Val == y && root.Right.Val == x) {
            return true
        }
        return haveSameParents(root.Left, x, y) || haveSameParents(root.Right, x, y)
    }
    levelX, levelY := findLevel(root, x, 1), findLevel(root, y, 1)
    if levelX != levelY {
        return false
    }
    return !haveSameParents(root, x, y)
}

// 解法二 BFS
func isCousinsBFS(root *TreeNode, x int, y int) bool {
    if root == nil {
        return false
    }
    type mark struct {
        prev  int
        depth int
    }
    queue := []*TreeNode{root}
    visited := [101]*mark{}
    visited[root.Val] = &mark{prev: -1, depth: 1}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        depth := visited[node.Val].depth
        if node.Left != nil {
            visited[node.Left.Val] = &mark{prev: node.Val, depth: depth + 1}
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            visited[node.Right.Val] = &mark{prev: node.Val, depth: depth + 1}
            queue = append(queue, node.Right)
        }
    }
    if visited[x] == nil || visited[y] == nil {
        return false
    }
    if visited[x].depth == visited[y].depth && visited[x].prev != visited[y].prev {
        return true
    }
    return false
}

// 解法三 DFS
func isCousinsDFS(root *TreeNode, x int, y int) bool {
	depth1, depth2, parent1, parent2 := 0, 0, 0, 0
    var dfs func(root *TreeNode, val, depth, last int, parent, res *int) 
    dfs = func(root *TreeNode, val, depth, last int, parent, res *int) {
        if root == nil {
            return
        }
        if root.Val == val {
            *res = depth
            *parent = last
            return
        }
        depth++
        dfs(root.Left, val, depth, root.Val, parent, res)
        dfs(root.Right, val, depth, root.Val, parent, res)
    }
	dfs(root, x, 0, -1, &parent1, &depth1)
	dfs(root, y, 0, -1, &parent2, &depth2)
	return depth1 > 1 && depth1 == depth2 && parent1 != parent2
}

func main() {
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4)})
    printTree(tree1)
    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, ptr(4), nil, ptr(5)})
    printTree(tree2)
    tree3 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, ptr(4)})
    printTree(tree3)
    fmt.Println(isCousins(tree1,4,3)) // false
    fmt.Println(isCousins(tree2,5,4)) // true
    fmt.Println(isCousins(tree3,2,3)) // false

    fmt.Println(isCousinsBFS(tree1,4,3)) // false
    fmt.Println(isCousinsBFS(tree2,5,4)) // true
    fmt.Println(isCousinsBFS(tree3,2,3)) // false

    fmt.Println(isCousinsDFS(tree1,4,3)) // false
    fmt.Println(isCousinsDFS(tree2,5,4)) // true
    fmt.Println(isCousinsDFS(tree3,2,3)) // false
}