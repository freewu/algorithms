package main

// LCR 021. 删除链表的倒数第 N 个结点
// 给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg" />
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]

// 示例 2：
// 输入：head = [1], n = 1
// 输出：[]

// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]
 
// 提示：
//     链表中结点的数目为 sz
//     1 <= sz <= 30
//     0 <= Node.val <= 100
//     1 <= n <= sz
 
// 进阶：能尝试使用一趟扫描实现吗？

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
    var l = (len(arr) - 1)
    var head = &ListNode{arr[l], nil}
    for i := l - 1; i >= 0; i--  {
        var n = &ListNode{arr[i], head}
        head = n
    }
    return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyHead := &ListNode{Next: head}
    // 设置 2 个指针，一个指针距离前一个指针 n 个距离
    preSlow, slow, fast := dummyHead, head, head
    for fast != nil {
        // 同时移动 2 个指针，2 个指针都移动相同的距离。当一个指针移动到了终点，那么前一个指针就是倒数第 n 个节点了
        if n <= 0 {
            preSlow = slow
            slow = slow.Next
        }
        // 重组链表
        n--
        fast = fast.Next
    }
    preSlow.Next = slow.Next
    return dummyHead.Next
}

// best solution 
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    first, second := dummy, head
    for i := 0; i < n; i++ {
        second = second.Next
    }
    for second != nil {
        second = second.Next
        first = first.Next
    }
    first.Next = first.Next.Next
    return dummy.Next
}

// 哑节点 + 链表长度
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
    getLength := func(head *ListNode) (length int) {
        for ; head != nil; head = head.Next {
            length++
        }
        return
    }
    // 先计算链表长度
    length := getLength(head)
    dummy := &ListNode{0, head}
    cur := dummy
    // 从哑节点开始遍历 L−n+1 个节点。
    for i := 0; i < length - n; i++ {
        cur = cur.Next
    }
    // 当遍历到第 L−n+1 个节点时，它的下一个节点就是我们需要删除的节点，这样我们只需要修改一次指针，就能完成删除操作
    cur.Next = cur.Next.Next
    return dummy.Next
}

// 栈
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
    nodes := []*ListNode{}
    dummy := &ListNode{0, head}
    // 在遍历链表的同时将所有节点依次入栈
    for node := dummy; node != nil; node = node.Next {
        nodes = append(nodes, node)
    }
    // 根据栈「先进后出」的原则，我们弹出栈的第 n 个节点就是需要删除的节点，并且目前栈顶的节点就是待删除节点的前驱节点
    prev := nodes[len(nodes)-1 - n]
    prev.Next = prev.Next.Next
    return dummy.Next
}

func main() {
    printListNode(makeListNode([]int{1,2,3,4,5}))
    // head = [1,2,3,4,5], n = 2
    printListNode(removeNthFromEnd(makeListNode([]int{1,2,3,4,5}), 2)) // 1 -> 2 -> 3 -> 5

    // head = [1], n = 1
    printListNode(makeListNode([]int{1}))
    printListNode(removeNthFromEnd(makeListNode([]int{1}), 1)) // nil

    // head = [1,2], n = 1
    printListNode(makeListNode([]int{1,2}))
    printListNode(removeNthFromEnd(makeListNode([]int{1,2}), 1)) // 1

    printListNode(makeListNode([]int{1,2,3,4,5}))
    // head = [1,2,3,4,5], n = 2
    printListNode(removeNthFromEnd1(makeListNode([]int{1,2,3,4,5}), 2)) // 1 -> 2 -> 3 -> 5

    // head = [1], n = 1
    printListNode(makeListNode([]int{1}))
    printListNode(removeNthFromEnd1(makeListNode([]int{1}), 1)) // nil

    // head = [1,2], n = 1
    printListNode(makeListNode([]int{1,2}))
    printListNode(removeNthFromEnd1(makeListNode([]int{1,2}), 1)) // 1

    fmt.Println("removeNthFromEnd2(makeListNode([]int{1,2,3,4,5}), 2)")
    printListNode(makeListNode([]int{1,2,3,4,5}))
    // head = [1,2,3,4,5], n = 2
    printListNode(removeNthFromEnd2(makeListNode([]int{1,2,3,4,5}), 2)) // 1 -> 2 -> 3 -> 5

    fmt.Println("removeNthFromEnd2(makeListNode([]int{1}), 1)")
    // head = [1], n = 1
    printListNode(makeListNode([]int{1}))
    printListNode(removeNthFromEnd2(makeListNode([]int{1}), 1)) // nil

    fmt.Println("removeNthFromEnd2(makeListNode([]int{1,2}), 1)")
    // head = [1,2], n = 1
    printListNode(makeListNode([]int{1,2}))
    printListNode(removeNthFromEnd2(makeListNode([]int{1,2}), 1)) // 1


    fmt.Println("removeNthFromEnd3(makeListNode([]int{1,2,3,4,5}), 2)")
    printListNode(makeListNode([]int{1,2,3,4,5}))
    // head = [1,2,3,4,5], n = 2
    printListNode(removeNthFromEnd3(makeListNode([]int{1,2,3,4,5}), 2)) // 1 -> 2 -> 3 -> 5

    fmt.Println("removeNthFromEnd3(makeListNode([]int{1}), 1)")
    // head = [1], n = 1
    printListNode(makeListNode([]int{1}))
    printListNode(removeNthFromEnd3(makeListNode([]int{1}), 1)) // nil

    fmt.Println("removeNthFromEnd3(makeListNode([]int{1,2}), 1)")
    // head = [1,2], n = 1
    printListNode(makeListNode([]int{1,2}))
    printListNode(removeNthFromEnd3(makeListNode([]int{1,2}), 1)) // 1
}