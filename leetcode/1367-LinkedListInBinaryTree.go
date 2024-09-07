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
    tree1 := &TreeNode{
        1, 
        &TreeNode{4, nil, &TreeNode{2, &TreeNode{1, nil, nil,}, nil,}, },
        &TreeNode{4, &TreeNode{2, &TreeNode{6, nil, nil,}, &TreeNode{8, &TreeNode{1, nil, nil,}, &TreeNode{3, nil, nil,},},}, nil, },
        
    }
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
    tree2 := &TreeNode{
        1, 
        &TreeNode{4, nil, &TreeNode{2, &TreeNode{1, nil, nil,}, nil,}, },
        &TreeNode{4, &TreeNode{2, &TreeNode{6, nil, nil,}, &TreeNode{8, &TreeNode{1, nil, nil,}, &TreeNode{3, nil, nil,},},}, nil, },
        
    }
    list2 := makeListNode([]int{1,4,2,6})
    printListNode(list2) // 1 -> 4 -> 2 -> 6
    fmt.Println(isSubPath(list2,tree2)) // true
    // Example 3:
    // Input: head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
    // Output: false
    // Explanation: There is no path in the binary tree that contains all the elements of the linked list from head.
    tree3 := &TreeNode{
        1, 
        &TreeNode{4, nil, &TreeNode{2, &TreeNode{1, nil, nil,}, nil,}, },
        &TreeNode{4, &TreeNode{2, &TreeNode{6, nil, nil,}, &TreeNode{8, &TreeNode{1, nil, nil,}, &TreeNode{3, nil, nil,},},}, nil, },
        
    }
    list3 := makeListNode([]int{1,4,2,6,8})
    printListNode(list3) // 1 -> 4 -> 2 -> 6 -> 8
    fmt.Println(isSubPath(list3,tree3)) // false

    fmt.Println(isSubPath1(list1,tree1)) // true
    fmt.Println(isSubPath1(list2,tree2)) // true
    fmt.Println(isSubPath1(list3,tree3)) // false
}