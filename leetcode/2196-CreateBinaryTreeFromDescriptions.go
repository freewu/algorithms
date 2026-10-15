package main

// 2196. Create Binary Tree From Descriptions
// You are given a 2D integer array descriptions where descriptions[i] = [parenti, childi, isLefti] indicates that parenti is the parent of childi in a binary tree of unique values. Furthermore,
//     If isLefti == 1, then childi is the left child of parenti.
//     If isLefti == 0, then childi is the right child of parenti.

// Construct the binary tree described by descriptions and return its root.
// The test cases will be generated such that the binary tree is valid.

// Example 1:
//          50
//        /    \
//      20      80
//     /  \     /
//   15    17  19
// <img src="https://assets.leetcode.com/uploads/2022/02/09/example1drawio.png" />
// Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
// Output: [50,20,80,15,17,19]
// Explanation: The root node is the node with value 50 since it has no parent.
// The resulting binary tree is shown in the diagram.

// Example 2:
//      1
//     /
//    2
//     \ 
//      3
//     /
//    4
// <img src="https://assets.leetcode.com/uploads/2022/02/09/example2drawio.png" />
// Input: descriptions = [[1,2,1],[2,3,0],[3,4,1]]
// Output: [1,2,null,null,3,4]
// Explanation: The root node is the node with value 1 since it has no parent.
// The resulting binary tree is shown in the diagram.

// Constraints:
//     1 <= descriptions.length <= 10^4
//     descriptions[i].length == 3
//     1 <= parenti, childi <= 10^5
//     0 <= isLefti <= 1
//     The binary tree described by descriptions is valid.

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
func createBinaryTree(descriptions [][]int) *TreeNode {
    n, res := len(descriptions), &TreeNode{}
    if len(descriptions) == 0 { return res }
    cache := map[int]*TreeNode{}
    for i := 0; i < n; i++ {
        cache[descriptions[i][1]] = &TreeNode{ Val: descriptions[i][1], }
    }
    for i := 0;i < n; i++ {
        p, c, child := descriptions[i][0], descriptions[i][1], descriptions[i][2]
        pNode, isExist := cache[p]
        if !isExist {
            pNode = &TreeNode{  Val: p, }
            res = pNode
            cache[p] = res
        }
        cNode := cache[c]
        if child == 1 {
            pNode.Left = cNode
        } else {
            pNode.Right = cNode
        }
    }
    return res
}

func createBinaryTree1(descriptions [][]int) *TreeNode {
    nodes := make(map[int]*TreeNode, len(descriptions) + 1) // 预分配空间
    root := 0
    for _, d := range descriptions {
        x, y := d[0], d[1]
        if nodes[x] == nil {
            nodes[x] = &TreeNode{ Val: x }
            root ^= x
        }
        if nodes[y] == nil {
            nodes[y] = &TreeNode{Val: y }
            root ^= y
        }
        if d[2] == 1 {
            nodes[x].Left = nodes[y]
        } else {
            nodes[x].Right = nodes[y]
        }
        root ^= y
    }
    return nodes[root]
}

func main() {
    // Example 1:
    //          50
    //        /    \
    //      20      80
    //     /  \     /
    //   15    17  19
    // <img src="https://assets.leetcode.com/uploads/2022/02/09/example1drawio.png" />
    // Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
    // Output: [50,20,80,15,17,19]
    // Explanation: The root node is the node with value 50 since it has no parent.
    // The resulting binary tree is shown in the diagram.
    fmt.Println(createBinaryTree([][]int{{20,15,1},{20,17,0},{50,20,1},{50,80,0},{80,19,1}})) // &{50 0xc000008090 0xc0000080a8}
    // Example 2:
    //      1
    //     /
    //    2
    //     \ 
    //      3
    //     /
    //    4
    // <img src="https://assets.leetcode.com/uploads/2022/02/09/example2drawio.png" />
    // Input: descriptions = [[1,2,1],[2,3,0],[3,4,1]]
    // Output: [1,2,null,null,3,4]
    // Explanation: The root node is the node with value 1 since it has no parent.
    // The resulting binary tree is shown in the diagram.
    fmt.Println(createBinaryTree([][]int{{1,2,1},{2,3,0},{3,4,1}})) // &{1 0xc000008120 <nil>}

    fmt.Println(createBinaryTree1([][]int{{20,15,1},{20,17,0},{50,20,1},{50,80,0},{80,19,1}})) // &{50 0xc000008090 0xc0000080a8}
    fmt.Println(createBinaryTree1([][]int{{1,2,1},{2,3,0},{3,4,1}})) // &{1 0xc000008120 <nil>}
}