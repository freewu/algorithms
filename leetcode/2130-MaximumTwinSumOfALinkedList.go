package main

// 2130. Maximum Twin Sum of a Linked List
// In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.
//     For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. 
//     These are the only nodes with twins for n = 4.

// The twin sum is defined as the sum of a node and its twin.
// Given the head of a linked list with even length, return the maximum twin sum of the linked list.

// Example 1:
// 5 -> (4) -> (2) -> 1
// 0     1      2     3
// <img src="https://assets.leetcode.com/uploads/2021/12/03/eg1drawio.png" />
// Input: head = [5,4,2,1]
// Output: 6
// Explanation:
// Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
// There are no other nodes with twins in the linked list.
// Thus, the maximum twin sum of the linked list is 6. 

// Example 2:
// 4 -> (2) -> (2) -> 3
// 0     1      2     3
// <img src="https://assets.leetcode.com/uploads/2021/12/03/eg2drawio.png" />
// Input: head = [4,2,2,3]
// Output: 7
// Explanation:
// The nodes with twins present in this linked list are:
// - Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
// - Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
// Thus, the maximum twin sum of the linked list is max(7, 4) = 7. 

// Example 3:
// 1 -> 100000
// 0     1    
// <img src="https://assets.leetcode.com/uploads/2021/12/03/eg3drawio.png" />
// Input: head = [1,100000]
// Output: 100001
// Explanation:
// There is only one node with a twin in the linked list having twin sum of 1 + 100000 = 100001.

// Constraints:
//     The number of nodes in the list is an even integer in the range [2, 10^5].
//     1 <= Node.val <= 10^5

import "fmt"

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
    l := len(arr) - 1
    head := &ListNode{arr[l], nil}
    for i := l - 1; i >= 0; i-- {
        n := &ListNode{arr[i], head}
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
func pairSum(head *ListNode) int {
    getLinkListLenght := func (head *ListNode) int { // 获取链表长度
        sum := 0
        for curr := head; curr != nil; {
            curr, sum = curr.Next, sum + 1
        }
        return sum
    }
    reverse := func (head *ListNode) *ListNode { // 反转链表
        var prev, curr *ListNode = nil, head
        for curr != nil {
            prev, curr, curr.Next = curr, curr.Next, prev
        }
        return prev
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    lenght := getLinkListLenght(head)
    res, middle, curr := 0, lenght / 2, head
    for i := 0; i < middle; i++ {
        curr = curr.Next
    }
    first, second := head, reverse(curr)
    for i := 0; i < middle; i++ {
        res = max(res, first.Val + second.Val)
        first, second = first.Next, second.Next
    }
    return res
}

func pairSum1(head *ListNode) int {
    res, f, s := 0, head, head
    var rev * ListNode = nil
    for f != nil && f.Next != nil {
        f = f.Next.Next
        nxt := s.Next
        s.Next = rev
        rev = s
        s = nxt
    }
    if f != nil {
        s = s.Next
    }
    for rev != nil && s != nil {
        temp := rev.Val + s.Val
        if res < temp {
            res = temp
        }
        rev = rev.Next
        s = s.Next
    }
    return res
}

func main() {
    // Example 1:
    // 5 -> (4) -> (2) -> 1
    // 0     1      2     3
    // Input: head = [5,4,2,1]
    // Output: 6
    // Explanation:
    // Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
    // There are no other nodes with twins in the linked list.
    // Thus, the maximum twin sum of the linked list is 6. 
    l1 := makeListNode([]int{5,4,2,1})
    printListNode(l1)
    fmt.Println(pairSum(l1)) // 6
    // Example 2:
    // 4 -> (2) -> (2) -> 3
    // 0     1      2     3
    // Input: head = [4,2,2,3]
    // Output: 7
    // Explanation:
    // The nodes with twins present in this linked list are:
    // - Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
    // - Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
    // Thus, the maximum twin sum of the linked list is max(7, 4) = 7. 
    l2 := makeListNode([]int{5,4,2,1})
    printListNode(l2)
    fmt.Println(pairSum(l2)) // 7
    // Example 3:
    // 1 -> 100000
    // 0     1    
    // Input: head = [1,100000]
    // Output: 100001
    // Explanation:
    // There is only one node with a twin in the linked list having twin sum of 1 + 100000 = 100001.
    l3 := makeListNode([]int{1,100000})
    printListNode(l3)
    fmt.Println(pairSum(l3)) // 100001

    l11 := makeListNode([]int{5,4,2,1})
    printListNode(l11)
    fmt.Println(pairSum(l11)) // 6
    l12 := makeListNode([]int{5,4,2,1})
    printListNode(l12)
    fmt.Println(pairSum(l12)) // 7
    l13 := makeListNode([]int{1,100000})
    printListNode(l13)
    fmt.Println(pairSum(l13)) // 100001
}