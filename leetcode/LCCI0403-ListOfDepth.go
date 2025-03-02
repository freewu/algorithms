package main

// 面试题 04.03. List of Depth LCCI
// Given a binary tree, design an algorithm which creates a linked list of all the nodes at each depth 
// (e.g., if you have a tree with depth D, you'll have D linked lists). 
// Return a array containing all the linked lists.

// Example:
// Input: [1,2,3,4,5,null,7,8]
//         1
//        /  \ 
//       2    3
//      / \    \ 
//     4   5    7
//    /
//   8
// Output: [[1],[2,3],[4,5,7],[8]]

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type ListNode struct {
    Val int
    Next *ListNode
}

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

func makeNodeList(nums []int) *ListNode {
    var n = &ListNode{-1, nil}
    var b = &ListNode{-1, n}
    for i := 0; i < len(nums); i++ {
        n.Next = &ListNode{nums[i], nil}
        n = n.Next
    }
    return b.Next.Next
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// bfs
func listOfDepth(tree *TreeNode) []*ListNode {
    queue := make([]*TreeNode,1)
    queue[0] = tree
    res := []*ListNode{}
    for len(queue) > 0 {
        n := len(queue)
        dummyHead  := &ListNode{} 
        tail := dummyHead 
        for ;n != 0; n-- {
            temp := queue[0]
            queue = queue[1:]
            if temp.Left != nil {
                queue = append(queue,temp.Left)
            }
            if temp.Right != nil {
                queue = append(queue,temp.Right)
            }
            tail.Val = temp.Val
            if n > 1 {
                tail.Next = &ListNode{}
                tail = tail.Next
            }
        }
        res = append(res, dummyHead)
    }
    return res
}

func main() {
    // Example:
    // Input: [1,2,3,4,5,null,7,8]
    //         1
    //        /  \ 
    //       2    3
    //      / \    \ 
    //     4   5    7
    //    /
    //   8
    // Output: [[1],[2,3],[4,5,7],[8]]
    tree1 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil}, nil}, &TreeNode{5, nil, nil}, },
        &TreeNode{3, nil, &TreeNode{7, nil, nil}, },
    }
    for _, v :=  range listOfDepth(tree1) {
        printListNode(v) // [[1],[2,3],[4,5,7],[8]]
    }
}