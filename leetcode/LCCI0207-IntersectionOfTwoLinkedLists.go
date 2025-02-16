package main

// 面试题 02.07. Intersection of Two Linked Lists LCCI
// Given two (singly) linked lists, determine if the two lists intersect. 
// Return the inter­ secting node. Note that the intersection is defined based on reference, not value. 
// That is, if the kth node of the first linked list is the exact same node (by reference) as the jth node of the second linked list, then they are intersecting.

// Example 1:
// Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
// Output: Reference of the node with value = 8
// Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). 
// From the head of A, it reads as [4,1,8,4,5].
// From the head of B, it reads as [5,0,1,8,4,5]. 
// There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.

// Example 2:
// Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// Output: Reference of the node with value = 2
// Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). 
// From the head of A, it reads as [0,9,1,2,4]. 
// From the head of B, it reads as [3,2,4]. 
// There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.

// Example 3:
// Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// Output: null
// Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
// Explanation: The two lists do not intersect, so return null.

// Notes:
//     If the two linked lists have no intersection at all, return null.
//     The linked lists must retain their original structure after the function returns.
//     You may assume there are no cycles anywhere in the entire linked structure.
//     Your code should preferably run in O(n) time and use only O(1) memory.

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
    if headA == nil || headB == nil { return nil }
    pa, pb := headA, headB
    for pa != pb {
        if pa == nil {
            pa = headB
        } else {
            pa = pa.Next
        }
        if pb == nil {
            pb = headA
        } else {
            pb = pb.Next
        }
    }
    return pa
}

func main() {
    // Example 1:
    // Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
    // Output: Reference of the node with value = 8
    // Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). 
    // From the head of A, it reads as [4,1,8,4,5].
    // From the head of B, it reads as [5,0,1,8,4,5]. 
    // There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
    list11 := makeListNode([]int{4,1,8,4,5})
    list12 := makeListNode([]int{5,0,1,8,4,5})
    printListNode(list11) // 4 -> 1 -> 8 -> 4 -> 5
    printListNode(list12) // 5 -> 0 -> 1 -> 8 -> 4 -> 5
    printListNode(getIntersectionNode(list11, list12))
    // Example 2:
    // Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
    // Output: Reference of the node with value = 2
    // Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). 
    // From the head of A, it reads as [0,9,1,2,4]. 
    // From the head of B, it reads as [3,2,4]. 
    // There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
    list21 := makeListNode([]int{0,9,1,2,4})
    list22 := makeListNode([]int{3,2,4})
    printListNode(list21) // 0 -> 9 -> 1 -> 2 -> 4
    printListNode(list22) // 3 -> 2 -> 4
    printListNode(getIntersectionNode(list21, list22))
    // Example 3:
    // Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
    // Output: null
    // Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
    // Explanation: The two lists do not intersect, so return null.
    list31 := makeListNode([]int{2,6,4})
    list32 := makeListNode([]int{1,5})
    printListNode(list31) // 2 -> 6 -> 4
    printListNode(list32) // 1 -> 5
    printListNode(getIntersectionNode(list31, list32))
}