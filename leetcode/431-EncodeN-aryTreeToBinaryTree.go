package main

// 431. Encode N-ary Tree to Binary Tree
// Design an algorithm to encode an N-ary tree into a binary tree and decode the binary tree to get the original N-ary tree. 
// An N-ary tree is a rooted tree in which each node has no more than N children. 
// Similarly, a binary tree is a rooted tree in which each node has no more than 2 children. 
// There is no restriction on how your encode/decode algorithm should work. 
// You just need to ensure that an N-ary tree can be encoded to a binary tree and this binary tree can be decoded to the original N-nary tree structure.

// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See following example).
// For example, you may encode the following 3-ary tree to a binary tree in this way:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreebinarytreeexample.png" />

// Input: root = [1,null,3,2,4,null,5,6]
// Note that the above is just an example which might or might not work. 
// You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

// Example 1:
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,null,3,2,4,null,5,6]

// Example 2:
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     0 <= Node.val <= 10^4
//     The height of the n-ary tree is less than or equal to 1000
//     Do not use class member/global/static variables to store states. Your encode and decode algorithms should be stateless.

import "fmt"

type Node struct {
    Val int
    Children []*Node
}

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

// == N 叉树构建 & 打印输出 start ==================================================================

// buildNTree 按层序数组（LeetCode 格式）创建 N 叉树：
// nil 是分组符，它之后到下一个 nil 之间的值，都是队首节点的孩子
func buildNTree(nums []*int) *Node {
    if len(nums) == 0 || nums[0] == nil {
        return nil
    }
    root := &Node{Val: *nums[0]}
    queue := []*Node{root}
    i := 1
    for i < len(nums) && len(queue) > 0 {
        if nums[i] == nil { // 跳过分组符
            i++
        }
        if i >= len(nums) {
            break
        }
        parent := queue[0]
        queue = queue[1:]
        for i < len(nums) && nums[i] != nil {
            child := &Node{Val: *nums[i]}
            parent.Children = append(parent.Children, child)
            queue = append(queue, child)
            i++
        }
    }
    return root
}

// printNTree 自顶向下打印 N 叉树：根在上、孩子在下铺开，父子之间用 / | \ 连线
func printNTree(root *Node) {
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
    const gap = 4 // 兄弟子树之间的最小列间隔
    // 画布一格一个字符、用 rune 存，制表符才只占一列
    blank := func(n int) []rune {
        line := make([]rune, n)
        for i := range line {
            line[i] = ' '
        }
        return line
    }

    type canvas struct {
        lines   []string // 子树画布：值行 + 连线行 + 孩子画布
        width   int      // 画布宽度
        rootCol int      // 根节点的中心列
    }
    // render 自底向上组装：先量出每一棵子树的宽度，再横向铺开、父节点居中
    var render func(node *Node) canvas
    render = func(node *Node) canvas {
        val := []rune(itoa(node.Val))
        kids := make([]canvas, 0, len(node.Children))
        for _, child := range node.Children {
            kids = append(kids, render(child))
        }
        kidsWidth, kidsHeight := 0, 0
        for i, kid := range kids {
            if i > 0 {
                kidsWidth += gap
            }
            kidsWidth += kid.width
            if len(kid.lines) > kidsHeight {
                kidsHeight = len(kid.lines)
            }
        }
        width := len(val)
        if kidsWidth > width {
            width = kidsWidth
        }
        kidsStart := 0 // 第一个孩子画布的起始列
        if kidsWidth > 0 {
            kidsStart = (width - kidsWidth) / 2
        }
        // 每个孩子在画布里的中心列；父节点中心对齐到首尾孩子的中间
        centers := make([]int, 0, len(kids))
        col := kidsStart
        for _, kid := range kids {
            centers = append(centers, col+kid.rootCol)
            col += kid.width + gap
        }
        rootCol := width / 2 // 没有孩子时直接居中
        if len(centers) > 0 {
            rootCol = (centers[0] + centers[len(centers)-1]) / 2
        }
        row := blank(width)
        start := rootCol - len(val)/2 // 值按中心列居中
        if start < 0 {
            start = 0
        }
        if start > width-len(val) {
            start = width - len(val)
        }
        copy(row[start:], val)
        lines := []string{string(row)}
        if len(kids) == 0 {
            return canvas{lines: lines, width: width, rootCol: rootCol}
        }

        // 连线用制表符拼：父节点一根竖线 + 一条横向总线
        left, right := rootCol, rootCol // 总线的左右端
        for _, c := range centers {
            if c < left {
                left = c
            }
            if c > right {
                right = c
            }
        }
        stem := blank(width)
        stem[rootCol] = '│' // 父节点向下的竖线
        lines = append(lines, string(stem))
        if left != right { // 孩子正对父节点时总线只有一个点，一根竖线就够
            bus := blank(width)
            for c := left; c <= right; c++ {
                bus[c] = '─'
            }
            switch { // 父节点接口：向上接竖线，左右接总线
            case rootCol == left:
                bus[rootCol] = '└'
            case rootCol == right:
                bus[rootCol] = '┘'
            default:
                bus[rootCol] = '┴'
            }
            for _, c := range centers { // 孩子接口：向下接竖线
                switch {
                case c == left:
                    bus[c] = '┌'
                case c == right:
                    bus[c] = '┐'
                default:
                    bus[c] = '┬'
                }
                if c == rootCol { // 孩子正好在父节点正下方：上下左右都接上
                    switch {
                    case rootCol == left:
                        bus[c] = '├'
                    case rootCol == right:
                        bus[c] = '┤'
                    default:
                        bus[c] = '┼'
                    }
                }
            }
            lines = append(lines, string(bus))
        }

        for r := 0; r < kidsHeight; r++ { // 孩子画布：根节点行对齐，矮画布下方留空
            row := blank(width)
            col := kidsStart
            for _, kid := range kids {
                if r < len(kid.lines) {
                    copy(row[col:], []rune(kid.lines[r]))
                }
                col += kid.width + gap
            }
            lines = append(lines, string(row))
        }
        return canvas{lines: lines, width: width, rootCol: rootCol}
    }

    for _, s := range render(root).lines {
        line := []rune(s)
        end := len(line)
        for end > 0 && line[end-1] == ' ' {
            end--
        }
        fmt.Println(string(line[:end]))
    }
}
// == N 叉树构建 & 打印输出 end ====================================================================

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
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
type Codec struct {
}

func Constructor() *Codec {
    return &Codec{}
}

func (this *Codec) encode(root *Node) *TreeNode {
    if root == nil {
        return nil
    }
    tree := &TreeNode{ Val: root.Val }
    if 0 < len(root.Children) {
        tree.Left = this.encode(root.Children[0])
        cur := tree.Left
        for _, c := range root.Children[1:] {
            cur.Right = this.encode(c)
            cur = cur.Right
        }
    }
    return tree
}

func (this *Codec) decode(root *TreeNode) *Node {
    if root == nil {
        return nil
    }
    node := &Node{Val: root.Val}
    cur := root.Left
    for cur != nil {
        node.Children = append(node.Children, this.decode(cur))
        cur = cur.Right
    }
    return node
}
 
 /**
  * Your Codec object will be instantiated and called as such:
  * obj := Constructor();
  * bst := obj.encode(root);
  * ans := obj.decode(bst);
  */

func main() {
    // Example 1:
    // Input: root = [1,null,3,2,4,null,5,6]
    // Output: [1,null,3,2,4,null,5,6]
    tree1 := buildNTree([]*int{ptr(1), nil, ptr(3), ptr(2), ptr(4), nil, ptr(5), ptr(6)})
    printNTree(tree1)
    // Example 2:
    // Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    // Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    tree2 := buildNTree([]*int{ptr(1), nil, ptr(2), ptr(3), ptr(4), ptr(5), nil, nil, ptr(6), nil, ptr(7), nil, ptr(8), nil, ptr(9), ptr(10), nil, nil, ptr(11), nil, ptr(12), nil, ptr(13), nil, nil, ptr(14)})
    printNTree(tree2)
    // Example 3:
    // Input: root = []
    // Output: []
    tree3 := buildNTree([]*int{})
    printNTree(tree3)
}
