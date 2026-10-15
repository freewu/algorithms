package main

// 1367. Linked List in Binary Tree
// Given a binary tree root and a linked list with head as the first node. 
// Return True if all the elements in the linked list starting from the head correspond to some downward path connected in the binary tree otherwise return False.
// In this context downward path means a path that starts at some node and goes downwards.

// Example 1:
//             1
//          /     \
//         4      (4)
//          \     /
//           2   (2)
//          /    /  \
//         1    6   (8)
//                 /   \
//                1     3
// <img src="https://assets.leetcode.com/uploads/2020/02/12/sample_1_1720.png" />
// Input: head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
// Output: true
// Explanation: Nodes in blue form a subpath in the binary Tree.  

// Example 2:
//        (1)
//      /     \
//     4      (4)
//      \     /
//      2    (2)
//     /    /  \
//    1   (6)   8
//            /   \
//           1     3
// <img src="https://assets.leetcode.com/uploads/2020/02/12/sample_2_1720.png" />
// Input: head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
// Output: true

// Example 3:
// Input: head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
// Output: false
// Explanation: There is no path in the binary tree that contains all the elements of the linked list from head.

// Constraints:
//     The number of nodes in the tree will be in the range [1, 2500].
//     The number of nodes in the list will be in the range [1, 100].
//     1 <= Node.val <= 100 for each node in the linked list and binary tree.

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

type ListNode struct {
    Val int
    Next *ListNode
}

// 打印链表
func printListNode(l *ListNode) {
    if nil == l {
        return
    }
    for {
        if nil == l.Next {
            fmt.Print(l.Val)
            break
        } else {
            fmt.Print(l.Val, " -> ")
        }
        l = l.Next
    }
    fmt.Println()
}

// 数组创建链表
func makeListNode(arr []int) *ListNode {
    if (len(arr) == 0) {
        return nil
    }
    var l = (len(arr) - 1)
    var head = &ListNode{arr[l], nil}
    for i := l - 1; i >= 0; i--  {
        var n = &ListNode{arr[i], head}
        head = n
    }
    return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
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
func isSubPath1(head *ListNode, root *TreeNode) bool {
    var dfs func(head *ListNode, cur *ListNode, root *TreeNode) bool
    dfs = func(head *ListNode, cur *ListNode, root *TreeNode) bool {
       if cur == nil { return true } // reach the end of the linked list, return true (successful match)
       if root == nil { return false } // reach the end of a path in the tree without fully matching the list, return false
       // Match the current tree node with the current linked list node
       // If no match, but the tree node matches the head of the linked list, start a new match
       // Otherwise, reset `cur` to `head` to attempt matching the linked list from scratch
       if cur.Val == root.Val {
           cur = cur.Next
       } else if head.Val == root.Val {
           head = head.Next
       } else {
           cur = head
       }
       // Continue dfs down both left and right children
       return dfs(head, cur, root.Left) || dfs(head, cur, root.Right) 
    }
    return dfs(head, head, root) 
}

func isSubPath(head *ListNode, root *TreeNode) bool {
    // 以当前根节点开始是否有指定链表的路径
    var isPath func(head *ListNode, root *TreeNode) bool
    isPath = func(head *ListNode, root *TreeNode) bool {
        if head == nil { return true }
        if root == nil { return false }
        if root.Val != head.Val { return false }
        if isPath(head.Next, root.Left) { return true } // left
        return isPath(head.Next, root.Right) // right
    }
    if head == nil { return true  }
    if root == nil { return false }
    if isPath(head, root) { return true } // 深度优先遍历判断 以 每个节点为根是否符合
    if isSubPath(head, root.Left) { return true } // 左边能走就不用走右边了
    return isSubPath(head, root.Right)
}

func main() {
    // Example 1:
    //            1
    //          /     \
    //         4      (4)
    //          \     /
    //           2   (2)
    //          /    /  \
    //         1    6   (8)
    //                 /   \
    //                1     3
    // <img src="https://assets.leetcode.com/uploads/2020/02/12/sample_1_1720.png" />
    // Input: head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
    // Output: true
    // Explanation: Nodes in blue form a subpath in the binary Tree.  
    tree1 := buildTree([]*int{ptr(1), ptr(4), ptr(4), nil, ptr(2), ptr(2), nil, ptr(1), nil, ptr(6), ptr(8), nil, nil, nil, nil, ptr(1), ptr(3)})
    printTree(tree1)
    list1 := makeListNode([]int{4,2,8})
    printListNode(list1) // 4 -> 2 -> 8
    fmt.Println(isSubPath(list1,tree1)) // true
    // Example 2:
    //        (1)
    //      /     \
    //     4      (4)
    //      \     /
    //      2    (2)
    //     /    /  \
    //    1   (6)   8
    //            /   \
    //           1     3
    // <img src="https://assets.leetcode.com/uploads/2020/02/12/sample_2_1720.png" />
    // Input: head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
    // Output: true
    tree2 := buildTree([]*int{ptr(1), ptr(4), ptr(4), nil, ptr(2), ptr(2), nil, ptr(1), nil, ptr(6), ptr(8), nil, nil, nil, nil, ptr(1), ptr(3)})
    printTree(tree2)
    list2 := makeListNode([]int{1,4,2,6})
    printListNode(list2) // 1 -> 4 -> 2 -> 6
    fmt.Println(isSubPath(list2,tree2)) // true
    // Example 3:
    // Input: head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
    // Output: false
    // Explanation: There is no path in the binary tree that contains all the elements of the linked list from head.
    tree3 := buildTree([]*int{ptr(1), ptr(4), ptr(4), nil, ptr(2), ptr(2), nil, ptr(1), nil, ptr(6), ptr(8), nil, nil, nil, nil, ptr(1), ptr(3)})
    printTree(tree3)
    list3 := makeListNode([]int{1,4,2,6,8})
    printListNode(list3) // 1 -> 4 -> 2 -> 6 -> 8
    fmt.Println(isSubPath(list3,tree3)) // false

    fmt.Println(isSubPath1(list1,tree1)) // true
    fmt.Println(isSubPath1(list2,tree2)) // true
    fmt.Println(isSubPath1(list3,tree3)) // false
}