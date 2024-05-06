package main

// 2487. Remove Nodes From Linked List
// You are given the head of a linked list.
// Remove every node which has a node with a greater value anywhere to the right side of it.
// Return the head of the modified linked list.

// Example 1:
// (5) -> (2) -> 13 -> (3) -> 8 => 13 -> 8
// <img src="https://assets.leetcode.com/uploads/2022/10/02/drawio.png" />
// Input: head = [5,2,13,3,8]
// Output: [13,8]
// Explanation: The nodes that should be removed are 5, 2 and 3.
// - Node 13 is to the right of node 5.
// - Node 13 is to the right of node 2.
// - Node 8 is to the right of node 3.

// Example 2:
// Input: head = [1,1,1,1]
// Output: [1,1,1,1]
// Explanation: Every node has value 1, so no nodes are removed.
 
// Constraints:
//     The number of the nodes in the given list is in the range [1, 10^5].
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
// 维护一个单调栈，栈底到栈顶的值为递减
func removeNodes(head *ListNode) *ListNode {
    dummy := &ListNode{ Val:  1e6, Next: head } // 哑巴节点
    stack := []*ListNode{ dummy } // 栈
    for node := head; node != nil; node = node.Next {
        for node.Val > stack[len(stack) - 1].Val { // 更新栈 移除每个右侧有一个更大数值的节点
            stack = stack[:len(stack)-1]
        }
        stack[len(stack)-1].Next = node
        stack = append(stack, node) // 入栈
    }
    return dummy.Next
}

// 倒序遍历维护一个递增序列，求这个的过程和【递归】过程很相似，浑然天成，递归递归，先递后归
// 分解子问题，比较当前节点的值和以 next 节点为 head 递归后的最大值，从而进行递归的转移
func removeNodes1(head *ListNode) *ListNode {
    if head.Next == nil { // 边界
        return head
    }
    node := removeNodes1(head.Next) // 状态转移 & 答案
    if head.Val < node.Val { // 删除head
        return node
    }
    head.Next = node 
    return head
}

// 倒序遍历维护一个递增序列
// 那么我们可先【反转链表】，遍历筛选仅保留递增序列，而后再一次【反转链表】回来即可
func removeNodes2(head *ListNode) *ListNode {
    reverseListNode := func (head *ListNode) *ListNode { // 反转链表
        var pre *ListNode
        cur := head
        for cur != nil {  // 双指针 遍历
            nxt := cur.Next
            cur.Next = pre
            pre, cur = cur, nxt
        }
        return pre
    }
    head = reverseListNode(head) // 反转
    node := head // 维护筛选保留递增节点
    for node.Next != nil {
        if node.Next.Val < node.Val {
            node.Next = node.Next.Next
        } else {
            node = node.Next
        }
    }
    // 反转
    return reverseListNode(head)
}

func main() {
    // Example 1:
    // (5) -> (2) -> 13 -> (3) -> 8 => 13 -> 8
    // <img src="https://assets.leetcode.com/uploads/2022/10/02/drawio.png" />
    // Input: head = [5,2,13,3,8]
    // Output: [13,8]
    // Explanation: The nodes that should be removed are 5, 2 and 3.
    // - Node 13 is to the right of node 5.
    // - Node 13 is to the right of node 2.
    // - Node 8 is to the right of node 3.
    l1 := makeListNode([]int{5,2,13,3,8})
    printListNode(l1)
    printListNode(removeNodes(l1))
    // Example 2:
    // Input: head = [1,1,1,1]
    // Output: [1,1,1,1]
    // Explanation: Every node has value 1, so no nodes are removed.
    l2 := makeListNode([]int{1,1,1,1})
    printListNode(l2)
    printListNode(removeNodes(l2))

    l11 := makeListNode([]int{5,2,13,3,8})
    printListNode(l11)
    printListNode(removeNodes1(l11))
    l12 := makeListNode([]int{1,1,1,1})
    printListNode(l12)
    printListNode(removeNodes(l12))

    l21 := makeListNode([]int{5,2,13,3,8})
    printListNode(l21)
    printListNode(removeNodes1(l21))
    l22 := makeListNode([]int{1,1,1,1})
    printListNode(l22)
    printListNode(removeNodes(l22))
}