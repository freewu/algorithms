package main

// 3294. Convert Doubly Linked List to Array II
// You are given an arbitrary node from a doubly linked list, 
// which contains nodes that have a next pointer and a previous pointer.

// Return an integer array which contains the elements of the linked list in order.

// Example 1:
// Input: head = [1,2,3,4,5], node = 5
// Output: [1,2,3,4,5]

// Example 2:
// Input: head = [4,5,6,7,8], node = 8
// Output: [4,5,6,7,8]

// Constraints:
//     The number of nodes in the given list is in the range [1, 500].
//     1 <= Node.val <= 1000
//     All nodes have unique Node.val.

import "fmt"

type Node struct {
    Val int
    Next *Node
    Prev *Node
}

// 打印链表
func printListNode(l *Node) {
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
func makeListNode(arr []int) *Node {
    if (len(arr) == 0) {
        return nil
    }
    l := len(arr) - 1
    head := &Node{arr[l], nil, nil }
    var pre *Node
    for i := l - 1; i >= 0; i-- {
        n := &Node{arr[i], head, pre}
        head = n
        pre = n
    }
    return head
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Prev *Node
 * }
 */
func toArray(head *Node) []int {
    for head != nil && head.Prev != nil {
        head = head.Prev
    }
    res := []int{}
    for ; head != nil; head = head.Next {
        res = append(res, head.Val)
    }
    return res
}

func main() {
    // Example 1:
    // Input: head = [1,2,3,4,5], node = 5
    // Output: [1,2,3,4,5]
    list1 := makeListNode([]int{1,2,3,4,5})
    printListNode(list1)
    fmt.Println(toArray(list1))
    // Example 2:
    // Input: head = [4,5,6,7,8], node = 8
    // Output: [4,5,6,7,8]
    list2 := makeListNode([]int{4,5,6,7,8})
    printListNode(list2)
    fmt.Println(toArray(list2))
}