package main

// 2096. Step-By-Step Directions From a Binary Tree Node to Another
// You are given the root of a binary tree with n nodes. 
// Each node is uniquely assigned a value from 1 to n. 
// You are also given an integer startValue representing the value of the start node s, 
// and a different integer destValue representing the value of the destination node t.

// Find the shortest path starting from node s and ending at node t. 
// Generate step-by-step directions of such path as a string consisting of only the uppercase letters 'L', 'R', and 'U'. 
// Each letter indicates a specific direction:
//     'L' means to go from a node to its left child node.
//     'R' means to go from a node to its right child node.
//     'U' means to go from a node to its parent node.

// Return the step-by-step directions of the shortest path from node s to node t.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/15/eg1.png" />
// Input: root = [5,1,2,3,null,6,4], startValue = 3, destValue = 6
// Output: "UURL"
// Explanation: The shortest path is: 3 → 1 → 5 → 2 → 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/15/eg2.png" />
// Input: root = [2,1], startValue = 2, destValue = 1
// Output: "L"
// Explanation: The shortest path is: 2 → 1.

// Constraints:
//     The number of nodes in the tree is n.
//     2 <= n <= 10^5
//     1 <= Node.val <= n
//     All the values in the tree are unique.
//     1 <= startValue, destValue <= n
//     startValue != destValue

import "fmt"
import "strings"

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
// dfs + bfs
func getDirections(root *TreeNode, startValue, destValue int) string {
    queue, parents := []*TreeNode{nil}, map[*TreeNode]*TreeNode{}
    var dfs func(node, pa *TreeNode)
    dfs = func(node, pa *TreeNode) {
        if node == nil { return }
        parents[node] = pa
        if node.Val == startValue {
            queue[0] = node // 只有一个起点
        }
        dfs(node.Left, node)
        dfs(node.Right, node)
    }
    dfs(root, nil)
    res, vis := []byte{}, map[*TreeNode]bool{nil: true, queue[0]: true}
    type pair struct {
        from *TreeNode
        dir  byte
    }
    from := map[*TreeNode]pair{}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        if node.Val == destValue {
            for ; from[node].from != nil; node = from[node].from {
                res = append(res, from[node].dir)
            }
            break
        }
        if !vis[node.Left] {
            vis[node.Left] = true
            from[node.Left] = pair{node, 'L'}
            queue = append(queue, node.Left)
        }
        if !vis[node.Right] {
            vis[node.Right] = true
            from[node.Right] = pair{node, 'R'}
            queue = append(queue, node.Right)
        }
        if !vis[parents[node]] {
            vis[parents[node]] = true
            from[parents[node]] = pair{node, 'U'}
            queue = append(queue, parents[node])
        }
    }
    for i, n := 0, len(res); i < n/2; i++ {
        res[i], res[n-1-i] = res[n-1-i], res[i]
    }
    return string(res)
}

// 最近公共祖先
func getDirections1(root *TreeNode, startValue, destValue int) string {
    path := []byte{}
    var dfs func(*TreeNode, int) bool
    dfs = func(node *TreeNode, target int) bool {
        if node == nil { return false }
        if node.Val == target { return true }
        path = append(path, 'L')
        if dfs(node.Left, target) { return true }
        path[len(path)-1] = 'R'
        if dfs(node.Right, target) { return true }
        path = path[:len(path)-1]
        return false
    }
    dfs(root, startValue)
    pathToStart := path

    path = nil
    dfs(root, destValue)
    pathToDest := path

    for len(pathToStart) > 0 && len(pathToDest) > 0 && pathToStart[0] == pathToDest[0] {
        pathToStart = pathToStart[1:] // 去掉前缀相同的部分
        pathToDest = pathToDest[1:]
    }
    return strings.Repeat("U", len(pathToStart)) + string(pathToDest)
}

func getDirections2(root *TreeNode, startValue int, destValue int) string {
    var findTarget func(root *TreeNode, targetValue int, path *[]byte) bool
    findTarget = func(root *TreeNode, targetValue int, path *[]byte) bool {
        if root.Val == targetValue {
            return true
        }
        if root.Left != nil && findTarget(root.Left, targetValue, path) {
            *path = append(*path, 'L')
            return true
        }
        if root.Right != nil && findTarget(root.Right, targetValue, path) {
            *path = append(*path, 'R')
            return true
        }
        return false
    }
    reverse := func (path []byte) {
        left, right := 0, len(path)-1
        for left < right {
            path[left], path[right] = path[right], path[left]
            left++
            right--
        }
    }
    replace := func(path []byte) {
        for i := 0; i < len(path); i++ {
            path[i] = 'U'
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    sPath, dPath := make([]byte, 0), make([]byte, 0)
    findTarget(root, startValue, &sPath)
    findTarget(root, destValue, &dPath)
    size, i := min(len(sPath), len(dPath)), 0
    for i < size {
        if sPath[len(sPath)-1-i] == dPath[len(dPath)-1-i] {
            i++
        } else {
            break
        }
    }
    sPath = sPath[:len(sPath)-i]
    replace(sPath)
    dPath = dPath[:len(dPath)-i]
    reverse(dPath)
    sPath = append(sPath, dPath...)
    return string(sPath)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/15/eg1.png" />
    // Input: root = [5,1,2,3,null,6,4], startValue = 3, destValue = 6
    // Output: "UURL"
    // Explanation: The shortest path is: 3 → 1 → 5 → 2 → 6.
    tree1 := buildTree([]*int{ptr(5), ptr(1), ptr(2), ptr(3), nil, ptr(6), ptr(4)})
    printTree(tree1)
    fmt.Println(getDirections(tree1,3, 6)) // "UURL"
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/11/15/eg2.png" />
    // Input: root = [2,1], startValue = 2, destValue = 1
    // Output: "L"
    // Explanation: The shortest path is: 2 → 1.
    tree2 := buildTree([]*int{ptr(1), ptr(2)})
    printTree(tree2)
    fmt.Println(getDirections(tree2,2, 1)) // "L"

    fmt.Println(getDirections1(tree1,3, 6)) // "UURL"
    fmt.Println(getDirections1(tree2,2, 1)) // "L"

    fmt.Println(getDirections2(tree1,3, 6)) // "UURL"
    fmt.Println(getDirections2(tree2,2, 1)) // "L"
}