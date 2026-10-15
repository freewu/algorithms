package main

// 333. Largest BST Subtree
// Given the root of a binary tree, find the largest subtree , 
// which is also a Binary Search Tree (BST), where the largest means subtree has the largest number of nodes.

// A Binary Search Tree (BST) is a tree in which all the nodes follow the below-mentioned properties:
//     The left subtree values are less than the value of their parent (root) node's value.
//     The right subtree values are greater than the value of their parent (root) node's value.

// Note: A subtree must include all of its descendants.

// Example 1:
//         10
//        /  \
//      (5)   15
//     /  \    \
//   (1)  (8)   7
// <img src="https://assets.leetcode.com/uploads/2020/10/17/tmp.jpg" />
// Input: root = [10,5,15,1,8,null,7]
// Output: 3
// Explanation: The Largest BST Subtree in this case is the highlighted one. The return value is the subtree's size, which is 3.

// Example 2:
// Input: root = [4,2,7,2,3,5,null,2,null,null,null,null,null,1]
// Output: 2
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -10^4 <= Node.val <= 10^4

// Follow up: Can you figure out ways to solve it with O(n) time complexity?

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
func largestBSTSubtree(root *TreeNode) int {
    // 二叉搜索树的性质：中序遍历有序，所以可以中序遍历一遍树，选出连续有序的最大子树
    // 优化：若一个子树不是二叉搜索树，则其父节点为根的树也不是，所以很多节点就不必要遍历了。
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 返回是否为搜索二叉树，节点个数,及最大最小值用于上层确认是否满足
    var inorder func(node *TreeNode) (bool,int,int,int)
    inorder = func(node *TreeNode) (bool,int,int,int) {
        if node == nil {
            return true, 0,-1,-1
        }
        ok1, cnt1, min1, max1 := inorder(node.Left)
        ok2, cnt2, min2, max2 := inorder(node.Right)
        if (ok1 && (cnt1==0 || node.Val>max1)) && (ok2 && (cnt2==0 || node.Val<min2)) {
            if cnt1 == 0 { min1 = node.Val; }
            if cnt2 == 0 { max2 = node.Val; }
            res = max(cnt1 + cnt2 + 1, res)
            return true,cnt1 + cnt2 + 1, min1, max2
        } else {
            return false, -1,-1,-1 //false 的情况下其他值都无所谓了
        }
    }
    inorder(root)
    return res
}

func main() {
    // Example 1:
    //         10
    //        /  \
    //      (5)   15
    //     /  \    \
    //   (1)  (8)   7
    // <img src="https://assets.leetcode.com/uploads/2020/10/17/tmp.jpg" />
    // Input: root = [10,5,15,1,8,null,7]
    // Output: 3
    // Explanation: The Largest BST Subtree in this case is the highlighted one. The return value is the subtree's size, which is 3.
    tree1 := buildTree([]*int{ptr(10), ptr(5), ptr(15), ptr(1), ptr(8), nil, ptr(7)})
    printTree(tree1)
    fmt.Println(largestBSTSubtree(tree1)) // 3
    // Example 2:
    // Input: root = [4,2,7,2,3,5,null,2,null,null,null,null,null,1]
    // Output: 2
    tree2 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(2), ptr(3), ptr(5), nil, ptr(2), nil, nil, nil, nil, nil, ptr(1)})
    printTree(tree2)
    fmt.Println(largestBSTSubtree(tree2)) // 2
}