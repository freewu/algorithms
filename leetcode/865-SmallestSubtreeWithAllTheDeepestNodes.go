package main

// 865. Smallest Subtree with all the Deepest Nodes
// Given the root of a binary tree, the depth of each node is the shortest distance to the root.
// Return the smallest subtree such that it contains all the deepest nodes in the original tree.
// A node is called the deepest if it has the largest depth possible among any node in the entire tree.
// The subtree of a node is a tree consisting of that node, plus the set of all descendants of that node.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/01/sketch1.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4]
// Output: [2,7,4]
// Explanation: We return the node with value 2, colored in yellow in the diagram.
// The nodes coloured in blue are the deepest nodes of the tree.
// Notice that nodes 5, 3 and 2 contain the deepest nodes in the tree but node 2 is the smallest subtree among them, so we return it.

// Example 2:
// Input: root = [1]
// Output: [1]
// Explanation: The root is the deepest node in the tree.

// Example 3:
// Input: root = [0,1,3,null,2]
// Output: [2]
// Explanation: The deepest node in the tree is 2, the valid subtrees are the subtrees of nodes 2, 1 and 0 but the subtree of node 2 is the smallest.

// Constraints:
//     The number of nodes in the tree will be in the range [1, 500].
//     0 <= Node.val <= 500
//     The values of the nodes in the tree are unique.

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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    var traverse func(root *TreeNode) (*TreeNode, int) 
    traverse = func (root *TreeNode) (*TreeNode, int) {
        if root == nil { return nil, 0 }
        p1, n1 := traverse(root.Left)
        p2, n2 := traverse(root.Right)
        if p1 == nil && p2 == nil { return root, 1 }
        if p1 == nil { return p2, n2 + 1 } // 右边最深
        if p2 == nil { return p1, n1 + 1 } // 左边最深

        if n1 > n2 {
            return p1, n1 + 1
        } else if n1 < n2 {
            return p2, n2 + 1
        } else {
            return root, n1 + 1
        }
    }
    res, _ := traverse(root)
    return res
}

func subtreeWithAllDeepest1(root *TreeNode) *TreeNode {
    type Result struct {
        node  *TreeNode // record the recent common ancestor node
        depth int       // record the maximum depth of the binary tree with node as the root
    }
    var maxDepth func(root *TreeNode) Result 
    maxDepth = func(root *TreeNode) Result {
        if root == nil { return Result{nil, 0} }
        left := maxDepth(root.Left)
        right := maxDepth(root.Right)
        if left.depth == right.depth { // 当左右子树的最大深度相同时，这个根节点是新的最近公共祖先
            return Result{root, left.depth + 1}
        }
        res := left
        if right.depth > left.depth { // 左右子树的深度不同，则最近公共祖先在 depth 较大的一边
            res = right
        }
        res.depth++
        return res
    }
    res := maxDepth(root)
    return res.node
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/01/sketch1.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4]
    // Output: [2,7,4]
    // Explanation: We return the node with value 2, colored in yellow in the diagram.
    // The nodes coloured in blue are the deepest nodes of the tree.
    // Notice that nodes 5, 3 and 2 contain the deepest nodes in the tree but node 2 is the smallest subtree among them, so we return it.
    tree1 := buildTree([]*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(2), ptr(0), ptr(8), nil, nil, ptr(7), ptr(4)})
    printTree(tree1)
    fmt.Println(subtreeWithAllDeepest(tree1)) // &{2 0xc0000080a8 0xc0000080c0}
    // Example 2:
    // Input: root = [1]
    // Output: [1]
    // Explanation: The root is the deepest node in the tree.
    tree2 := buildTree([]*int{ptr(1)})
    printTree(tree2)
    fmt.Println(subtreeWithAllDeepest(tree2)) // &{1 <nil> <nil>}
    // Example 3:
    // Input: root = [0,1,3,null,2]
    // Output: [2]
    // Explanation: The deepest node in the tree is 2, the valid subtrees are the subtrees of nodes 2, 1 and 0 but the subtree of node 2 is the smallest.
    tree3 := buildTree([]*int{ptr(0), ptr(1), ptr(3), nil, ptr(2)})
    printTree(tree3)
    fmt.Println(subtreeWithAllDeepest(tree3)) // &{2 <nil> <nil>}

    tree11 := buildTree([]*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(2), ptr(0), ptr(8), nil, nil, ptr(7), ptr(4)})
    printTree(tree11)
    fmt.Println(subtreeWithAllDeepest1(tree11)) // &{2 0xc0000080a8 0xc0000080c0}
    tree12 := buildTree([]*int{ptr(1)})
    printTree(tree12)
    fmt.Println(subtreeWithAllDeepest1(tree12)) // &{1 <nil> <nil>}
    tree13 := buildTree([]*int{ptr(0), ptr(1), ptr(3), nil, ptr(2)})
    printTree(tree13)
    fmt.Println(subtreeWithAllDeepest1(tree13)) // &{2 <nil> <nil>}
}