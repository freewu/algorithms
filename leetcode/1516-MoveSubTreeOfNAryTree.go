package main

// 1516. Move Sub-Tree of N-Ary Tree
// Given the root of an N-ary tree of unique values, and two nodes of the tree p and q.

// You should move the subtree of the node p to become a direct child of node q. 
// If p is already a direct child of q, do not change anything. 
// Node p must be the last child in the children list of node q.

// Return the root of the tree after adjusting it.

// There are 3 cases for nodes p and q:
//     1. Node q is in the sub-tree of node p.
//     2. Node p is in the sub-tree of node q.
//     3. Neither node p is in the sub-tree of node q nor node q is in the sub-tree of node p.

// In cases 2 and 3, you just need to move p (with its sub-tree) to be a child of q, but in case 1 the tree may be disconnected, thus you need to reconnect the tree again. Please read the examples carefully before solving this problem.

// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// For example, the above tree is serialized as [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/13/move_e1.jpg" />
// Input: root = [1,null,2,3,null,4,5,null,6,null,7,8], p = 4, q = 1
// Output: [1,null,2,3,4,null,5,null,6,null,7,8]
// Explanation: This example follows the second case as node p is in the sub-tree of node q. We move node p with its sub-tree to be a direct child of node q.
// Notice that node 4 is the last child of node 1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/07/13/move_e2.jpg" />
// Input: root = [1,null,2,3,null,4,5,null,6,null,7,8], p = 7, q = 4
// Output: [1,null,2,3,null,4,5,null,6,null,7,8]
// Explanation: Node 7 is already a direct child of node 4. We don't change anything.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/07/13/move_e3.jpg" />
// Input: root = [1,null,2,3,null,4,5,null,6,null,7,8], p = 3, q = 8
// Output: [1,null,2,null,4,5,null,7,8,null,null,null,3,null,6]
// Explanation: This example follows case 3 because node p is not in the sub-tree of node q and vice-versa. We can move node 3 with its sub-tree and make it as node 8's child.
 
// Constraints:
//     The total number of nodes is between [2, 1000].
//     Each node has a unique value.
//     p != null
//     q != null
//     p and q are two different nodes (i.e. p != q).

import "fmt"

type Node struct {
    Val int
    Children []*Node
}

// == N 叉树构建 & 打印输出 start ==================================================================
func ptr(v int) *int {
    return &v
}

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
func moveSubTree(root *Node, p *Node, q *Node) *Node {
    var find func(root, p *Node) *Node
    find = func(root, p *Node) *Node {
        if root == p {
            return nil
        }
        for _, child := range root.Children {
            if child == p  {
                return root
            }
            if temp := find(child, p); temp != nil {
                return temp
            }
        }
        return nil
    }
    remove := func(parent, target *Node) {
        if parent == nil { return }
        for i, child := range parent.Children {
            if child == target {
                copy(parent.Children[i:], parent.Children[i+1:])
                parent.Children = parent.Children[:len(parent.Children)-1]
                return
            }
        }
    }
    if parent := find(p, q); parent != nil {
        remove(parent, q)
        pp := find(root, p)
        if pp != nil {
            for i, child := range pp.Children {
                if child == p{
                    pp.Children[i] = q
                }
            }
        } else {
            root = q
        }
        q.Children = append(q.Children, p)
        return root
    }
    parent := find(root, p)
    if parent != q {
        remove(parent, p)
        q.Children = append(q.Children, p)
    }
    return root
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/13/move_e1.jpg" />
    // Input: root = [1,null,2,3,null,4,5,null,6,null,7,8], p = 4, q = 1
    // Output: [1,null,2,3,4,null,5,null,6,null,7,8]
    // Explanation: This example follows the second case as node p is in the sub-tree of node q. We move node p with its sub-tree to be a direct child of node q.
    // Notice that node 4 is the last child of node 1.
    tree1 := buildNTree([]*int{ptr(1), nil, ptr(2), ptr(3), nil, ptr(4), ptr(5), nil, ptr(6), nil, ptr(7), ptr(8)})
    printNTree(tree1)
    p1, q1 := &Node{5, nil},  &Node{1, nil}
    r1 := moveSubTree(tree1, p1, q1)
    fmt.Println(r1)
    printNTree(r1)
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/07/13/move_e2.jpg" />
    // Input: root = [1,null,2,3,null,4,5,null,6,null,7,8], p = 7, q = 4
    // Output: [1,null,2,3,null,4,5,null,6,null,7,8]
    // Explanation: Node 7 is already a direct child of node 4. We don't change anything.
    tree2 := buildNTree([]*int{ptr(1), nil, ptr(2), ptr(3), nil, ptr(4), ptr(5), nil, ptr(6), nil, ptr(7), ptr(8)})
    printNTree(tree2)
    p2, q2 := &Node{7, nil},  &Node{7, nil}
    r2 := moveSubTree(tree2, p2, q2)
    fmt.Println(r2)
    printNTree(r2)
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/07/13/move_e3.jpg" />
    // Input: root = [1,null,2,3,null,4,5,null,6,null,7,8], p = 3, q = 8
    // Output: [1,null,2,null,4,5,null,7,8,null,null,null,3,null,6]
    // Explanation: This example follows case 3 because node p is not in the sub-tree of node q and vice-versa. We can move node 3 with its sub-tree and make it as node 8's child.
    tree3 := buildNTree([]*int{ptr(1), nil, ptr(2), ptr(3), nil, ptr(4), ptr(5), nil, ptr(6), nil, ptr(7), ptr(8)})
    printNTree(tree3)
    p3, q3 := &Node{3, nil},  &Node{8, nil}
    r3 := moveSubTree(tree3, p3, q3)
    fmt.Println(r3)
    printNTree(r3)
}