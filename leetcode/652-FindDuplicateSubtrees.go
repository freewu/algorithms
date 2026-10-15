package main

// 652. Find Duplicate Subtrees
// Given the root of a binary tree, return all duplicate subtrees.
// For each kind of duplicate subtrees, you only need to return the root node of any one of them.
// Two trees are duplicate if they have the same structure with the same node values.

// Example 1:
//             1
//           /   \
//         (2)     3
//         /     /   \
//       (4)   (2)     4
//             /
//           (4)
// <img src="https://assets.leetcode.com/uploads/2020/08/16/e1.jpg" />
// Input: root = [1,2,3,4,null,2,4,null,null,4]
// Output: [[2,4],[4]]

// Example 2:
//         2
//       /   \
//     (1)    (1)
// <img src="https://assets.leetcode.com/uploads/2020/08/16/e2.jpg" />
// Input: root = [2,1,1]
// Output: [[1]]

// Example 3:
//         2
//       /   \
//     (2)   (2)
//     /     /
//   (4)   (4)
// <img src="https://assets.leetcode.com/uploads/2020/08/16/e33.jpg" />
// Input: root = [2,2,2,3,null,3,null]
// Output: [[2,3],[3]]
 
// Constraints:
//     The number of the nodes in the tree will be in the range [1, 5000]
//     -200 <= Node.val <= 200

import "fmt"
import "strconv"

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    exist := map[string]int{}
    res := []*TreeNode{}
    var dfs func(root *TreeNode) string
    dfs = func(root *TreeNode) string {
        if root == nil {
            return "nil"
        }
        tree := strconv.Itoa(root.Val)
        tree += "," + dfs(root.Left)
        tree += "," + dfs(root.Right)
        if v, _ := exist[tree]; v == 1 { // 存在一样的key 说有重复的子树
            res = append(res, root)
        }
        exist[tree]++
        return tree
    }
    dfs(root)
    return res
}

func findDuplicateSubtrees1(root *TreeNode) []*TreeNode {
    type pair struct {
        node   *TreeNode
        index  int
    }
    repeat, seen, index := map[*TreeNode]struct{}{}, map[[3]int]pair{}, 0
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        tri := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
        if p, ok := seen[tri]; ok {
            repeat[p.node] = struct{}{}
            return p.index
        }
        index++
        seen[tri] = pair{node, index}
        return index
    }
    dfs(root)
    res := make([]*TreeNode, 0, len(repeat))
    for node := range repeat {
        res = append(res, node)
    }
    return res
}

func main() {
    // Example 1:
    //             1
    //           /   \
    //         (2)     3
    //         /     /   \
    //       (4)   (2)     4
    //             /
    //           (4)
    // <img src="https://assets.leetcode.com/uploads/2020/08/16/e1.jpg" />
    // Input: root = [1,2,3,4,null,2,4,null,null,4]
    // Output: [[2,4],[4]]
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), nil, ptr(2), ptr(4), nil, nil, ptr(4)})
    printTree(tree1)
    fmt.Println(findDuplicateSubtrees(tree1)) //[[2,4],[4]]
    // Example 2:
    //         2
    //       /   \
    //     (1)    (1)
    // <img src="https://assets.leetcode.com/uploads/2020/08/16/e2.jpg" />
    // Input: root = [2,1,1]
    // Output: [[1]]
    tree2 := buildTree([]*int{ptr(2), ptr(1), ptr(1)})
    printTree(tree2)
    fmt.Println(findDuplicateSubtrees(tree2)) //[[1]]
    // Example 3:
    //         2
    //       /   \
    //     (2)   (2)
    //     /     /
    //   (3)   (3)
    // <img src="https://assets.leetcode.com/uploads/2020/08/16/e33.jpg" />
    // Input: root = [2,2,2,3,null,3,null]
    // Output: [[2,3],[3]]
    tree3 := buildTree([]*int{ptr(2), ptr(2), ptr(2), ptr(3), nil, ptr(3)})
    printTree(tree3)
    fmt.Println(findDuplicateSubtrees(tree3)) // [[2,3],[3]]

    fmt.Println(findDuplicateSubtrees1(tree1)) //[[2,4],[4]]
    fmt.Println(findDuplicateSubtrees1(tree2)) //[[1]]
    fmt.Println(findDuplicateSubtrees1(tree3)) // [[2,3],[3]]
}