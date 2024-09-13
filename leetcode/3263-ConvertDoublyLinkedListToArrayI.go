package main

// 3263. Convert Doubly Linked List to Array I
// You are given the head of a doubly linked list, 
// which contains nodes that have a next pointer and a previous pointer.

// Return an integer array which contains the elements of the linked list in order.

// Example 1:
// Input: head = [1,2,3,4,3,2,1]
// Output: [1,2,3,4,3,2,1]

// Example 2:
// Input: head = [2,2,2,2,2]
// Output: [2,2,2,2,2]

// Example 3:
// Input: head = [3,2,3,2,3,2]
// Output: [3,2,3,2,3,2]

// Constraints:
//     The number of nodes in the given list is in the range [1, 50].
//     1 <= Node.val <= 50

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
    res := []int{}
    for head != nil {
        res = append(res, head.Val)
        head = head.Next
    }
    return res
}

func main() {
    // Example 1:
    // Input: head = [1,2,3,4,3,2,1]
    // Output: [1,2,3,4,3,2,1]
    list1 := makeListNode([]int{1,2,3,4,3,2,1})
    printListNode(list1) // 1 -> 2 -> 3 -> 4 -> 3 -> 2 -> 1
    fmt.Println(toArray(list1)) // [1,2,3,4,3,2,1]
    // Example 2:
    // Input: head = [2,2,2,2,2]
    // Output: [2,2,2,2,2]
    list2 := makeListNode([]int{2,2,2,2,2})
    printListNode(list2) // 2 -> 2 -> 2 -> 2 -> 2
    fmt.Println(toArray(list2)) // [2,2,2,2,2]
    // Example 3:
    // Input: head = [3,2,3,2,3,2]
    // Output: [3,2,3,2,3,2]
    list3 := makeListNode([]int{3,2,3,2,3,2})
    printListNode(list3) // 3 -> 2 -> 3 -> 2 -> 3 -> 2
    fmt.Println(toArray(list3)) // [3,2,3,2,3,2]
}