package main

// 160. Intersection of Two Linked Lists
// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. 
// If the two linked lists have no intersection at all, return null.

// For example, the following two linked lists begin to intersect at node c1:
// <img src="https://assets.leetcode.com/uploads/2021/03/05/160_statement.png" />
//         (a1) -> (a2) ->
//                         (c1) -> (c2) -> (c3)
// (b1) -> (b2) -> (b3) -> 

// The test cases are generated such that there are no cycles anywhere in the entire linked structure.
// Note that the linked lists must retain their original structure after the function returns.

// Custom Judge:
// The inputs to the judge are given as follows (your program is not given these inputs):
//     intersectVal - The value of the node where the intersection occurs. This is 0 if there is no intersected node.
//     listA - The first linked list.
//     listB - The second linked list.
//     skipA - The number of nodes to skip ahead in listA (starting from the head) to get to the intersected node.
//     skipB - The number of nodes to skip ahead in listB (starting from the head) to get to the intersected node.

// The judge will then create the linked structure based on these inputs and pass the two heads, headA and headB to your program. 
// If you correctly return the intersected node, then your solution will be accepted.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/05/160_example_1_1.png" />
//        (4) -> (1) ->
//                       (8) -> (4) -> (5)
// (5) -> (6) -> (1) -> 
// Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
// Output: Intersected at '8'
// Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
// From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
// - Note that the intersected node's value is not 1 because the nodes with value 1 in A and B (2nd node in A and 3rd node in B) are different node references. In other words, they point to two different locations in memory, while the nodes with value 8 in A and B (3rd node in A and 4th node in B) point to the same location in memory.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/05/160_example_2.png" />
// (1) -> (9) -> (1) ->
//                     (2) -> (4)
//               (3) -> 
// Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// Output: Intersected at '2'
// Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect).
// From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/03/05/160_example_3.png" />
// (2) -> (6) -> (4)

// (1) -> (5)
// Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// Output: No intersection
// Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
// Explanation: The two lists do not intersect, so return null.
 
// Constraints:
//     The number of nodes of listA is in the m.
//     The number of nodes of listB is in the n.
//     1 <= m, n <= 3 * 10^4
//     1 <= Node.val <= 10^5
//     0 <= skipA < m
//     0 <= skipB < n
//     intersectVal is 0 if listA and listB do not intersect.
//     intersectVal == listA[skipA] == listB[skipB] if listA and listB intersect.

import "fmt"

type ListNode struct {
    Val  int
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
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    a, b := headA, headB
    // AB不相交时，AB走的路径都是两条链表。
    // AB相交时，转换为环形链表的入环点。
    for a != b {
        if a != nil { a = a.Next } else { a = headB }
        if b != nil { b = b.Next } else { b = headA }
    }
    return a
 }

func main() {
    fmt.Println("Example 1:")
    l11 := makeListNode([]int{4,1,8,4,5})
    l12 := makeListNode([]int{5,6,1,8,4,5})
    printListNode(l11)
    printListNode(l12)
    printListNode(getIntersectionNode(l11,l12))

    fmt.Println("Example 2:")
    l21 := makeListNode([]int{1,9,1,2,4})
    l22 := makeListNode([]int{3,2,4})
    printListNode(l21)
    printListNode(l22)
    printListNode(getIntersectionNode(l21,l22))

    fmt.Println("Example 3:")
    l31 := makeListNode([]int{2,6,4})
    l32 := makeListNode([]int{1,5})
    printListNode(l31)
    printListNode(l32)
    printListNode(getIntersectionNode(l31, l32))
}