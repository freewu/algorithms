package main

// 701. Insert into a Binary Search Tree
// You are given the root node of a binary search tree (BST) and a value to insert into the tree. 
// Return the root node of the BST after the insertion. 
// It is guaranteed that the new value does not exist in the original BST.

// Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. 
// You can return any of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/05/insertbst.jpg" />
// Input: root = [4,2,7,1,3], val = 5
// Output: [4,2,7,1,3,5]
// Explanation: Another accepted tree is:

// Example 2:
// Input: root = [40,20,60,10,30,50,70], val = 25
// Output: [40,20,60,10,30,50,70,null,null,25]

// Example 3:
// Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
// Output: [4,2,7,1,3,5]

// Constraints:
//     The number of nodes in the tree will be in the range [0, 10^4].
//     -10^8 <= Node.val <= 10^8
//     All the values Node.val are unique.
//     -10^8 <= val <= 10^8
//     It's guaranteed that val does not exist in the original BST.

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
// 递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}

// bfs
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    node := root
    for {
        if node.Val > val {
            if node.Left != nil {
                node = node.Left
            } else {
                node.Left = &TreeNode{Val: val}
                break
            }
        }
        if node.Val < val {
            if node.Right != nil {
                node = node.Right
            } else {
                node.Right = &TreeNode{Val: val}
                break
            }
        }
    }
    return root
}

// dfs
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
    var dfs func (root *TreeNode, val int) *TreeNode 
    dfs = func (root *TreeNode, val int) *TreeNode {
        if root == nil {
            return &TreeNode{Val: val}
        }
        if val < root.Val {
            root.Left = dfs(root.Left, val)
        } else {
            root.Right = dfs(root.Right, val)
        }
        return root
    }
    return dfs(root, val)
}

func insertIntoBST3(root *TreeNode, val int) *TreeNode {
    p := root
    var prev *TreeNode = nil
    for p != nil {
        prev = p
        if p.Val == val {
            return root
        } else if p.Val > val {
            p = p.Left
        } else {
            p = p.Right
        }
    }
    if prev == nil {
        prev = &TreeNode{Val: val}
        root = prev
    } else if prev.Val > val {
        prev.Left = &TreeNode{Val: val}
    } else {
        prev.Right = &TreeNode{Val: val}
    }
    return root
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/05/insertbst.jpg" />
    // Input: root = [4,2,7,1,3], val = 5
    // Output: [4,2,7,1,3,5]
    // Explanation: Another accepted tree is:
    tree1 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), ptr(3)})
    printTree(tree1)
    fmt.Println(insertIntoBST(tree1, 5))
    // Example 2:
    // Input: root = [40,20,60,10,30,50,70], val = 25
    // Output: [40,20,60,10,30,50,70,null,null,25]
    tree2 := buildTree([]*int{ptr(40), ptr(20), ptr(60), ptr(10), ptr(30), ptr(50), ptr(70)})
    printTree(tree2)
    fmt.Println(insertIntoBST(tree2, 25))
    // Example 3:
    // Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
    // Output: [4,2,7,1,3,5]
    tree3 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), ptr(3)})
    printTree(tree3)
    fmt.Println(insertIntoBST(tree3, 5))


    tree11 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), ptr(3)})
    printTree(tree11)
    fmt.Println(insertIntoBST1(tree11, 5))
    tree12 := buildTree([]*int{ptr(40), ptr(20), ptr(60), ptr(10), ptr(30), ptr(50), ptr(70)})
    printTree(tree12)
    fmt.Println(insertIntoBST1(tree12, 25))
    tree13 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), ptr(3)})
    printTree(tree13)
    fmt.Println(insertIntoBST1(tree13, 5))

    tree21 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), ptr(3)})
    printTree(tree21)
    fmt.Println(insertIntoBST2(tree21, 5))
    tree22 := buildTree([]*int{ptr(40), ptr(20), ptr(60), ptr(10), ptr(30), ptr(50), ptr(70)})
    printTree(tree22)
    fmt.Println(insertIntoBST2(tree22, 25))
    tree23 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), ptr(3)})
    printTree(tree23)
    fmt.Println(insertIntoBST2(tree23, 5))

    tree31 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), ptr(3)})
    printTree(tree31)
    fmt.Println(insertIntoBST3(tree31, 5))
    tree32 := buildTree([]*int{ptr(40), ptr(20), ptr(60), ptr(10), ptr(30), ptr(50), ptr(70)})
    printTree(tree32)
    fmt.Println(insertIntoBST3(tree32, 25))
    tree33 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), ptr(3)})
    printTree(tree33)
    fmt.Println(insertIntoBST3(tree33, 5))
}