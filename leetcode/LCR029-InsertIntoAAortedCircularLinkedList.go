package main

// LCR 029. 循环有序列表的插入
// 给定循环单调非递减列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。
// 给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
// 如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
// 如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2019/01/19/example_1_before_65p.jpg" />
// 输入：head = [3,4,1], insertVal = 2
// 输出：[3,4,1,2]
// 解释：在上图中，有一个包含三个元素的循环有序列表，你获得值为 3 的节点的指针，我们需要向表中插入元素 2 。新插入的节点应该在 1 和 3 之间，插入之后，整个列表如上图所示，最后返回节点 3 。
// <img src="https://assets.leetcode.com/uploads/2019/01/19/example_1_after_65p.jpg" />

// 示例 2：
// 输入：head = [], insertVal = 1
// 输出：[1]
// 解释：列表为空（给定的节点是 null），创建一个循环有序列表并返回这个节点。

// 示例 3：
// 输入：head = [1], insertVal = 0
// 输出：[1,0]

// 提示：
//     0 <= Number of Nodes <= 5 * 10^4
//     -10^6 <= Node.val <= 10^6
//     -10^6 <= insertVal <= 10^6

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