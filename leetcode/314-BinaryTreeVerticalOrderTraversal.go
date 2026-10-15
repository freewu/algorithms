package main

// 314. Binary Tree Vertical Order Traversal
// Given the root of a binary tree, return the vertical order traversal of its nodes' values. (i.e., from top to bottom, column by column).
// If two nodes are in the same row and column, the order should be from left to right.

// Example 1:
//         3
//       /   \
//      9    20
//          /  \
//         15   7
// <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree1.jpg" />
// Input: root = [3,9,20,null,null,15,7]
// Output: [[9],[3,15],[20],[7]]

// Example 2:
//         3
//       /   \
//      9     8
//     /  \  /  \
//    4    0 1   7
// <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree2-1.jpg" />
// Input: root = [3,9,8,4,0,1,7]
// Output: [[4],[9],[3,0,1],[8],[7]]

// Example 3:
//         3
//       /   \
//      9     8
//    /  \  /  \
//   4    0 1   7
//       /   \
//      5     2
// <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree2.jpg" />
// Input: root = [3,9,8,4,0,1,7,null,null,null,2,5]
// Output: [[4],[9,5],[3,0,1],[8,2],[7]]
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 100].
//     -100 <= Node.val <= 100

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalOrder(root *TreeNode) [][]int {
    // 想象在遍历过程中，路径会左右拐，从题意来看，左拐和右拐次数相抵消后剩余次数相同的节点为同一组；基于此我们可以在遍历过程中记录每个节点走过的路径和作为列号标号column，往左-1，往右+1
    // 另同一列的输出顺序为从上到下，那么先序/中序遍历不行，因为有可能出现左子树的某个后辈节点与右节点在同一列的情况，那样的话上下遍历顺序会不对；所以需要用层序遍历来保证上下顺序；
    // 不同列的顺序从左到右，正好可以利用column来排序;
    // 再观察一下可以发现一棵树每一列和列之间是连续的，不会出现空列，所以遍历过程中只需要记录最大最小column即可   
    res := [][]int{}
    if root == nil { return res } // 节点树为空，直接返回
    minCol, maxCol := 0,0 // 以根节点所在列为0列，则最大列不小于0，最小列不大于0,都给0即可
    groups := map[int][]int{} // 遍历的过程中可以根据列号分组

    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }

    // 为便于记录当前节点处于第几列，这里定义一个结构体用来存储层节点
    type Node struct {
        node *TreeNode
        col int
    }
    // 1. 层序遍历
    layer := []Node{ Node{node: root, col: 0} } // 第一层节点只有一个root，处在0列
    for len(layer) > 0 {
        nextLayer := []Node{}
        for _,n := range layer {
            maxCol, minCol = max(maxCol, n.col), min(minCol, n.col)
            groups[n.col] = append(groups[n.col], n.node.Val)
            if n.node.Left != nil {
                nextLayer = append(nextLayer, Node{ node: n.node.Left, col: n.col - 1} ) // 左 -1
            }
            if n.node.Right != nil {
                nextLayer = append(nextLayer, Node{ node: n.node.Right, col: n.col + 1}) // 右 + 1
            }
        }
        layer = nextLayer
    }
    for col := minCol; col <= maxCol; col++ { /// 输出
        res = append(res, groups[col])
    }
    return res
}

func main() {
    // Example 1:
    //         3
    //       /   \
    //      9    20
    //          /  \
    //         15   7
    // <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree1.jpg" />
    // Input: root = [3,9,20,null,null,15,7]
    // Output: [[9],[3,15],[20],[7]]
    tree1 := buildTree([]*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)})
    printTree(tree1)
    fmt.Println(verticalOrder(tree1)) // [[9],[3,15],[20],[7]]
    // Example 2:
    //         3
    //       /   \
    //      9     8
    //     /  \  /  \
    //    4    0 1   7
    // <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree2-1.jpg" />
    // Input: root = [3,9,8,4,0,1,7]
    // Output: [[4],[9],[3,0,1],[8],[7]]
    tree2 := buildTree([]*int{ptr(3), ptr(9), ptr(8), ptr(4), ptr(0), ptr(1), ptr(7)})
    printTree(tree2)
    fmt.Println(verticalOrder(tree2)) // [[4],[9],[3,0,1],[8],[7]]
    // Example 3:
    //         3
    //       /   \
    //      9     8
    //    /  \  /  \
    //   4    0 1   7
    //       /   \
    //      5     2
    // <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree2.jpg" />
    // Input: root = [3,9,8,4,0,1,7,null,null,null,2,5]
    // Output: [[4],[9,5],[3,0,1],[8,2],[7]]
    tree3 := buildTree([]*int{ptr(3), ptr(9), ptr(8), ptr(4), ptr(0), ptr(1), ptr(7), nil, nil, ptr(5), nil, nil, ptr(2)})
    printTree(tree3)
    fmt.Println(verticalOrder(tree3)) // [[4],[9,5],[3,0,1],[8,2],[7]]
}