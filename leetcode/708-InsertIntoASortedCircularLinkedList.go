package main

// 708. Insert into a Sorted Circular Linked List
// Given a Circular Linked List node, which is sorted in non-descending order,
// write a function to insert a value insertVal into the list such that it remains a sorted circular list. 
// The given node can be a reference to any single node in the list and may not necessarily be the smallest value in the circular list.

// If there are multiple suitable places for insertion, you may choose any place to insert the new value. 
// After the insertion, the circular list should remain sorted.

// If the list is empty (i.e., the given node is null), you should create a new single circular list 
// and return the reference to that single node. Otherwise, you should return the originally given node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/01/19/example_1_before_65p.jpg" />
// Input: head = [3,4,1], insertVal = 2
// Output: [3,4,1,2]
// Explanation: 
// In the figure above, there is a sorted circular list of three elements. 
// You are given a reference to the node with value 3, and we need to insert 2 into the list. 
// The new node should be inserted between node 1 and node 3. 
// After the insertion, the list should look like this, and we should still return node 3.
// <img src="https://assets.leetcode.com/uploads/2019/01/19/example_1_after_65p.jpg" />

// Example 2:
// Input: head = [], insertVal = 1
// Output: [1]
// Explanation: 
// The list is empty (given head is null). 
// We create a new single circular list and return the reference to that single node.

// Example 3:
// Input: head = [1], insertVal = 0
// Output: [1,0]

// Constraints:
//     The number of nodes in the list is in the range [0, 5 * 10^4].
//     -10^6 <= Node.val, insertVal <= 10^6

import "fmt"

type Node struct {
    Val  int
    Next *Node
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
    var l = (len(arr) - 1)
    var head = &Node{arr[l], nil}
    for i := l - 1; i >= 0; i--  {
        var n = &Node{arr[i], head}
        head = n
    }
    return head
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */
// func insert(aNode *Node, x int) *Node {
//     newNode := &Node{ Val: x }
//     if aNode == nil {
//         newNode.Next = newNode
//         return newNode
//     }
//     cur, mx := aNode, aNode
//     for {
//         if cur.Val >= mx.Val { // 找到最大节点位置
//             mx = cur
//         }
//         if (cur.Val <= x && cur.Next.Val > x) || cur.Next == aNode {
//             break
//         }
//         cur = cur.Next
//     }
//     if cur.Val <= x && cur.Next.Val > x {
//         // 找到一个节点 大于当前节点，小于下个节点，并插入新节点
//         cur.Next, newNode.Next = newNode, cur.Next
//         return aNode
//     }
//     mx.Next, newNode.Next = newNode, mx.Next // 插入到最大节点之后
//     return aNode
// }

func insert(head *Node, insertVal int) *Node {
    node := &Node{Val: insertVal}
    if head == nil {
        node.Next = node
        return node
    }
    if head.Next == head {
        head.Next = node
        node.Next = head
        return head
    }
    curr, next := head, head.Next
    for next != head {
        if insertVal >= curr.Val && insertVal <= next.Val {
            break
        }
        if curr.Val > next.Val {
            if insertVal > curr.Val || insertVal < next.Val {
                break
            }
        }
        curr = curr.Next
        next = next.Next
    }
    curr.Next = node
    node.Next = next
    return head
}

func insert1(aNode *Node, x int) *Node {
    n := &Node { Val: x, }
    if aNode == nil {
        n.Next = n
        return n
    }
    self, p:= aNode, aNode.Next
    for  {
        if p.Val <= x && p.Next.Val >= x || (p.Val <= x || x <= p.Next.Val) && p.Next.Val < p.Val || p == self  {
            n.Next = p.Next
            p.Next = n
            break
        }
        p = p.Next
    }
    return aNode
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/01/19/example_1_before_65p.jpg" />
    // Input: head = [3,4,1], insertVal = 2
    // Output: [3,4,1,2]
    // Explanation: 
    // In the figure above, there is a sorted circular list of three elements. 
    // You are given a reference to the node with value 3, and we need to insert 2 into the list. 
    // The new node should be inserted between node 1 and node 3. 
    // After the insertion, the list should look like this, and we should still return node 3.
    // <img src="https://assets.leetcode.com/uploads/2019/01/19/example_1_after_65p.jpg" />
    l1 := makeListNode([]int{3,4,1})
    fmt.Println("before: ")
    printListNode(l1)
    fmt.Println("after: ")
    r1 := insert(l1, 2)
    printListNode(r1)
    // Example 2:
    // Input: head = [], insertVal = 1
    // Output: [1]
    // Explanation: 
    // The list is empty (given head is null). 
    // We create a new single circular list and return the reference to that single node.
    l2 := makeListNode([]int{})
    fmt.Println("before: ")
    printListNode(l2)
    fmt.Println("after: ")
    r2 := insert(l2, 1)
    printListNode(r2)
    // Example 3:
    // Input: head = [1], insertVal = 0
    // Output: [1,0]
    l3 := makeListNode([]int{1})
    fmt.Println("before: ")
    printListNode(l3)
    fmt.Println("after: ")
    r3 := insert(l3, 0)
    printListNode(r3)
}