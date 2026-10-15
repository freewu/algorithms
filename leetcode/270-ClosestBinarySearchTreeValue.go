package main

// 270. Closest Binary Search Tree Value
// Given the root of a binary search tree and a target value, 
// return the value in the BST that is closest to the target.
// If there are multiple answers, print the smallest.

// Example 1:
//         4
//       /   \ 
//      2     5
//    /   \
//   1     3
// <img src="https://assets.leetcode.com/uploads/2021/03/12/closest1-1-tree.jpg" />
// Input: root = [4,2,5,1,3], target = 3.714286
// Output: 4

// Example 2:
// Input: root = [1], target = 4.428571
// Output: 1
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     0 <= Node.val <= 10^9
//     -10^9 <= target <= 10^9

import "fmt"
import "math"

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
func closestValue(root *TreeNode, target float64) int {
    closest, cur := root, root
    for cur != nil {
        if target == float64(cur.Val) {
            return cur.Val
        }
        if math.Abs(target - float64(cur.Val)) < math.Abs(target - float64(closest.Val)) {
            closest = cur
        } else if (math.Abs(target - float64(cur.Val)) == math.Abs(target - float64(closest.Val))) { // 处理值差值一样的情况
            if cur.Val < closest.Val { // 谁小选择谁
                closest = cur
            }
        }
        if target < float64(cur.Val) {
            cur = cur.Left
        } else {
            cur = cur.Right
        }
    }
    return closest.Val
}

// dfs 
func closestValue1(root *TreeNode, target float64) int {
    closest := root
    var dfs func(*TreeNode,float64)
    dfs = func(node *TreeNode,target float64) {
        if node == nil {
            return 
        }
        dfs(node.Left,target)
        if math.Abs(target - float64(node.Val)) < math.Abs(target - float64(closest.Val)) {
            closest = node
        } else if (math.Abs(target - float64(node.Val)) == math.Abs(target - float64(closest.Val))) { // 处理值差值一样的情况
            if node.Val < closest.Val { // 谁小选择谁
                closest = node
            }
        }
        dfs(node.Right,target)
    }
    dfs(root,target)
    return closest.Val
}

func closestValue2(root *TreeNode, target float64) int {
    diff := func ( a, b float64) float64 {
        if a > b {
            return a - b
        } else {
            return b - a
        }
    }
    closest, bl, a := root, diff(float64(root.Val), target), 0.0
    for root != nil {
        a = diff(float64(root.Val), target)
        if a < bl {
            closest = root
            bl = a
        }
        if target > float64(root.Val) {
            root = root.Right
        } else {
            root = root.Left
        }
    } 
    return closest.Val 
}


func main() {
    // Example 1:
    //         4
    //       /   \ 
    //      2     5
    //    /   \
    //   1     3
    // <img src="https://assets.leetcode.com/uploads/2021/03/12/closest1-1-tree.jpg" />
    // Input: root = [4,2,5,1,3], target = 3.714286
    // Output: 4
    tree1 := buildTree([]*int{ptr(4), ptr(2), ptr(5), ptr(1), ptr(3)})
    printTree(tree1)
    fmt.Println(closestValue(tree1, 3.714286)) // 4
    // Example 2:
    // Input: root = [1], target = 4.428571
    // Output: 1
    tree2 := buildTree([]*int{ptr(1)})
    printTree(tree2)
    fmt.Println(closestValue(tree2, 4.428571)) // 1

    // tree3 := &TreeNode {
    //     4,
    //     &TreeNode{2, &TreeNode{1, nil, nil }, &TreeNode{3, nil, nil }, },
    //     &TreeNode{5, nil, nil },
    // }
    fmt.Println(closestValue(tree1, 3.5)) // 3
    fmt.Println(closestValue(tree1, 4.5)) // 5

    fmt.Println(closestValue1(tree1, 3.714286)) // 4
    fmt.Println(closestValue1(tree2, 4.428571)) // 1
    fmt.Println(closestValue1(tree1, 3.5)) // 3
    fmt.Println(closestValue1(tree1, 4.5)) // 5

    fmt.Println(closestValue2(tree1, 3.714286)) // 4
    fmt.Println(closestValue2(tree2, 4.428571)) // 1
    fmt.Println(closestValue2(tree1, 3.5)) // 3
    fmt.Println(closestValue2(tree1, 4.5)) // 5
}