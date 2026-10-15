package main

// 987. Vertical Order Traversal of a Binary Tree
// Given the root of a binary tree, calculate the vertical order traversal of the binary tree.
// For each node at position (row, col), its left and right children will be at positions (row + 1, col - 1) and (row + 1, col + 1) respectively. 
// The root of the tree is at (0, 0).

// The vertical order traversal of a binary tree is a list of top-to-bottom orderings for each column index starting from the leftmost column and ending on the rightmost column. 
// There may be multiple nodes in the same row and same column. In such a case, sort these nodes by their values.
// Return the vertical order traversal of the binary tree.
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/29/vtree1.jpg" />
// Input: root = [3,9,20,null,null,15,7]
// Output: [[9],[3,15],[20],[7]]
// Explanation:
// Column -1: Only node 9 is in this column.
// Column 0: Nodes 3 and 15 are in this column in that order from top to bottom.
// Column 1: Only node 20 is in this column.
// Column 2: Only node 7 is in this column.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/29/vtree2.jpg" />
// Input: root = [1,2,3,4,5,6,7]
// Output: [[4],[2],[1,5,6],[3],[7]]
// Explanation:
// Column -2: Only node 4 is in this column.
// Column -1: Only node 2 is in this column.
// Column 0: Nodes 1, 5, and 6 are in this column.
//           1 is at the top, so it comes first.
//           5 and 6 are at the same position (2, 0), so we order them by their value, 5 before 6.
// Column 1: Only node 3 is in this column.
// Column 2: Only node 7 is in this column.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/01/29/vtree3.jpg" />
// Input: root = [1,2,3,4,6,5,7]
// Output: [[4],[2],[1,5,6],[3],[7]]
// Explanation:
// This case is the exact same as example 2, but with nodes 5 and 6 swapped.
// Note that the solution remains the same since 5 and 6 are in the same location and should be ordered by their values.
 
// Constraints:
// 		The number of nodes in the tree is in the range [1, 1000].
// 		0 <= Node.val <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "fmt"
import "math"
import "sort"


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

// 用于存储元素的坐标
type node struct {
    x, y, val int
}

func verticalTraversal(root *TreeNode) [][]int {
    var dfs func(root *TreeNode, x, y int)
    var nodes []node
    dfs = func(root *TreeNode, x, y int) {
        if root == nil {
            return
        }
        // 按照先序遍历，就可以将这些结点的二维坐标计算出来
        nodes = append(nodes, node{x, y, root.Val})
        dfs(root.Left, x+1, y-1)
        dfs(root.Right, x+1, y+1)
    }
    // 根结点是 (0，0) ，即根结点是坐标原点
    // 它的左子树的 x 坐标都是负数，它的右子树的 x 坐标都是正数
    dfs(root, 0, 0)
    // 进行一次排序，按照 x 坐标从小到大排序，坐标相同的情况对应着结点摞起来的情况
    sort.Slice(nodes, func(i, j int) bool {
        a, b := nodes[i], nodes[j]
        return a.y < b.y || a.y == b.y &&
            (a.x < b.x || a.x == b.x && a.val < b.val)
    })
    // 扫描一遍排好序的数组，按照列的顺序，依次将同一列的结点打包至一个一维数组
    var res [][]int
    lastY := math.MinInt32
    for _, node := range nodes {
        if lastY != node.y {
            res = append(res, []int{node.val})
            lastY = node.y
        } else {
            res[len(res)-1] = append(res[len(res)-1], node.val)
        }
    }
    return res
}

func main() {
    tree1 := buildTree([]*int{ptr(3), ptr(9), ptr(30), nil, nil, ptr(15), ptr(7)})
    printTree(tree1)
    fmt.Println(verticalTraversal(tree1)) // [[9],[3,15],[20],[7]]

    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6), ptr(7)})
    printTree(tree2)
    fmt.Println(verticalTraversal(tree2)) // [[4],[2],[1,5,6],[3],[7]]

    tree3 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(6), ptr(5), ptr(7)})
    printTree(tree3)
    fmt.Println(verticalTraversal(tree3)) // [[4],[2],[1,5,6],[3],[7]]
}