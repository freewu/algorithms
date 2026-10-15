package main

// 450. Delete Node in a BST
// Given a root node reference of a BST and a key, delete the node with the given key in the BST. 
// Return the root node reference (possibly updated) of the BST.

// Basically, the deletion can be divided into two stages:
//     Search for a node to remove.
//     If the node is found, delete the node.
 
// Example 1:
//         5                 5
//       /   \              /   \
//      3     6     =>     4     6 
//     /  \    \          /       \
//    2    4    8        2         8
// <img src="https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg" />
// Input: root = [5,3,6,2,4,null,7], key = 3
// Output: [5,4,6,2,null,null,7]
// Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
// One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
// Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

// Example 2:
// Input: root = [5,3,6,2,4,null,7], key = 0
// Output: [5,3,6,2,4,null,7]
// Explanation: The tree does not contain a node with value = 0.

// Example 3:
// Input: root = [], key = 0
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -10^5 <= Node.val <= 10^5
//     Each node has a unique value.
//     root is a valid binary search tree.
//     -10^5 <= key <= 10^5

// Follow up: Could you solve it with time complexity O(height of tree)?

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
func deleteNode(root *TreeNode, key int) *TreeNode {
    var dfs func(*TreeNode) *TreeNode
    dfs = func(node *TreeNode) *TreeNode {
        if node == nil { return nil }
        if node.Val == key {
            if node.Left == nil { return node.Right }
            if node.Right == nil { return node.Left }
            curr := node.Right
            for curr.Left != nil { 
                curr = curr.Left
            }
            curr.Left = node.Left // 移除节点
            return node.Right
        }
        if key > node.Val { // 二分找到要删除的节点
            node.Right = dfs(node.Right)
        } else {
            node.Left = dfs(node.Left)
        }
        return node
    }
    return dfs(root)
}

func deleteNode1(root *TreeNode, key int) *TreeNode {
    var dfs func (cur **TreeNode,key int)
    dfs = func (cur **TreeNode,key int) {
        if (*cur) == nil { return }
        if key < (*cur).Val {
            dfs(&((*cur).Left),key)
        } else if key == (*cur).Val {
            if (*cur).Left == nil && (*cur).Right == nil {
                (*cur) = nil
            } else if (*cur).Left == nil && (*cur).Right != nil {
                (*cur) = (*cur).Right
            } else if (*cur).Left != nil && (*cur).Right == nil {
                (*cur) = (*cur).Left
            } else {
                tmp := (*cur).Right
                for tmp.Left != nil{
                    tmp = tmp.Left
                }
                (*cur).Val = tmp.Val
                dfs(&((*cur).Right),tmp.Val)
            }
        } else {
            dfs(&((*cur).Right),key)
        }
    }
    dfs(&root,key)
    return root
}

func main() {
    // Example 1:
    //         5                 5
    //       /   \              /   \
    //      3     6     =>     4     6 
    //     /  \    \          /       \
    //    2    4    7        2         7
    // <img src="https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg" />
    // Input: root = [5,3,6,2,4,null,7], key = 3
    // Output: [5,4,6,2,null,null,7]
    // Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
    // One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
    // Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.
    tree1 := buildTree([]*int{ptr(5), ptr(3), ptr(6), ptr(2), ptr(4), nil, ptr(7)})
    printTree(tree1)
    fmt.Println("tree1.Left: ", tree1.Left) 
    t1 := deleteNode(tree1,3)
    printTree(t1)
    fmt.Println("t1: ", t1) // t1:  &{5 0xc000008078 0xc000008090}
    fmt.Println("t1.Left: ", t1.Left) 
    // Example 2:
    // Input: root = [5,3,6,2,4,null,7], key = 0
    // Output: [5,3,6,2,4,null,7]
    // Explanation: The tree does not contain a node with value = 0.
    tree2 := buildTree([]*int{ptr(5), ptr(3), ptr(6), ptr(2), ptr(4), nil, ptr(7)})
    printTree(tree2)
    fmt.Println("tree2: ", tree2) 
    t2 := deleteNode(tree2, 0)
    printTree(t2)
    fmt.Println("t2: ", t2) 
    // Example 3:
    // Input: root = [], key = 0
    // Output: []
    t3 := deleteNode(nil, 0)
    printTree(t3)
    fmt.Println("t3: ", t3) 


    tree11 := buildTree([]*int{ptr(5), ptr(3), ptr(6), ptr(2), ptr(4), nil, ptr(7)})
    printTree(tree11)
    fmt.Println("tree11.Left: ", tree11.Left) 
    t11 := deleteNode1(tree11,3)
    printTree(t11)
    fmt.Println("t11: ", t11) // t11:  &{5 0xc000008078 0xc000008090}
    fmt.Println("t11.Left: ", t11.Left) 
    tree12 := buildTree([]*int{ptr(5), ptr(3), ptr(6), ptr(2), ptr(4), nil, ptr(7)})
    printTree(tree12)
    fmt.Println("tree12: ", tree12) 
    t12 := deleteNode1(tree12, 0)
    printTree(t12)
    fmt.Println("t12: ", t12) 
    t13 := deleteNode1(nil, 0)
    printTree(t13)
    fmt.Println("t13: ", t13) 
}