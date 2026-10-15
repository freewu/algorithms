package main

// 1028. Recover a Tree From Preorder Traversal
// We run a preorder depth-first search (DFS) on the root of a binary tree.

// At each node in this traversal, we output D dashes (where D is the depth of this node), 
// then we output the value of this node.  
// If the depth of a node is D, the depth of its immediate child is D + 1.  
// The depth of the root node is 0.

// If a node has only one child, that child is guaranteed to be the left child.

// Given the output traversal of this traversal, recover the tree and return its root.

// Example 1:
//             1
//           /   \
//          2     5
//        /   \  /  \
//       3    4 6    7
// <img src="https://assets.leetcode.com/uploads/2019/04/08/recover-a-tree-from-preorder-traversal.png" />
// Input: traversal = "1-2--3--4-5--6--7"
// Output: [1,2,5,3,4,6,7]

// Example 2:
//             1
//          /     \
//         2       5
//        /       /
//       3       6
//      /       /
//     4       7
// <img src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114101-pm.png" />
// Input: traversal = "1-2--3---4-5--6---7"
// Output: [1,2,5,3,null,6,null,4,null,7]

// Example 3:
//             1
//           /
//         401
//       /    \
//     349    88
//     /
//    90
// <img src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114955-pm.png" />
// Input: traversal = "1-401--349---90--88"
// Output: [1,401,null,349,88,90]
 
// Constraints:
//     The number of nodes in the original tree is in the range [1, 1000].
//     1 <= Node.val <= 10^9

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
func recoverFromPreorder(traversal string) *TreeNode {
    parent := make(map[*TreeNode]*TreeNode) // 用 map 存当前节点的父节点
    res := &TreeNode{ Val: -1, Left: nil, Right: nil, } // 定义一个虚拟节点
    parent[res] = nil
    cur, parentLevel, curLevel, p, n := res,  -1, 0, 0, len(traversal)
    for p < n {
        for ; traversal[p] == '-'; p++ {
            curLevel++
        }
        start := p
        for ; p < n && traversal[p] != '-'; p++ { }
        curVal, _ := strconv.Atoi(traversal[start:p])
        if curLevel > parentLevel {
            cur.Left = &TreeNode{ Val: curVal, Left: nil, Right: nil, }
            parentLevel = curLevel
            curLevel = 0
            parent[cur.Left] = cur
            cur = cur.Left
        } else {
            pa := cur
            for curLevel <= parentLevel {
                pa = parent[pa]
                parentLevel--
            }
            pa.Right = &TreeNode{ Val: curVal, Left: nil, Right: nil, }
            parent[pa.Right] = pa
            cur = pa.Right
            parentLevel = curLevel
        }
        curLevel = 0
    }
    return res.Left
}

// stack
func recoverFromPreorder1(traversal string) *TreeNode {
    type Item struct {
        node *TreeNode
        level int
    }
    n, d, index := len(traversal), 0, 0
    for ; index < n && traversal[index] >= '0' && traversal[index] <= '9'; index++ { // 先取第1个值
        ch := traversal[index]
        d = d * 10 + int(ch - '0')
    }
    root := &TreeNode{}
    root.Val = d
    stack := []Item{ Item{node: root, level: 0}}
    for index < n {
        d = 0
        level := 0
        for ; index < n && traversal[index] == '-'; index++ {
            level++
        }
        for ; index < n && traversal[index] != '-'; index++ {
            d = d * 10+ int(traversal[index] - '0')
        }
        tmpNode := &TreeNode{ Val : d }
        for len(stack) > 0 && stack[len(stack)-1].level != level-1 {
            stack = stack[:len(stack)-1]
        }
        parent := stack[len(stack)-1].node
        if parent.Left == nil {
            parent.Left = tmpNode
        } else {
            parent.Right = tmpNode
        }
        stack = append(stack, Item{node: tmpNode, level: level})
    }
    return root
}

func main() {
    // Example 1:
    //             1
    //           /   \
    //          2     5
    //        /   \  /  \
    //       3    4 6    7
    // <img src="https://assets.leetcode.com/uploads/2019/04/08/recover-a-tree-from-preorder-traversal.png" />
    // Input: traversal = "1-2--3--4-5--6--7"
    // Output: [1,2,5,3,4,6,7]
    fmt.Println(recoverFromPreorder("1-2--3--4-5--6--7")) // &{1 0xc000110060 0xc0001100a8}
    // Example 2:
    //             1
    //          /     \
    //         2       5
    //        /       /
    //       3       6
    //      /       /
    //     4       7
    // <img src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114101-pm.png" />
    // Input: traversal = "1-2--3---4-5--6---7"
    // Output: [1,2,5,3,null,6,null,4,null,7]
    fmt.Println(recoverFromPreorder("1-2--3---4-5--6---7")) // &{1 0xc000110138 0xc000110180}
    // Example 3:
    //             1
    //           /
    //         401
    //       /    \
    //     349    88
    //     /
    //    90
    // <img src="https://assets.leetcode.com/uploads/2019/04/11/screen-shot-2019-04-10-at-114955-pm.png" />
    // Input: traversal = "1-401--349---90--88"
    // Output: [1,401,null,349,88,90]
    fmt.Println(recoverFromPreorder("1-401--349---90--88")) // &{1 0xc000110210 <nil>}

    fmt.Println(recoverFromPreorder1("1-2--3--4-5--6--7")) // &{1 0xc000110060 0xc0001100a8}
    fmt.Println(recoverFromPreorder1("1-2--3---4-5--6---7")) // &{1 0xc000110138 0xc000110180}
    fmt.Println(recoverFromPreorder1("1-401--349---90--88")) // &{1 0xc000110210 <nil>}
}
