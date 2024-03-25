package main

// 82. Remove Duplicates from Sorted List II
// Given the head of a sorted linked list, delete all nodes that have duplicate numbers, 
// leaving only distinct numbers from the original list. Return the linked list sorted as well.

// Example 1:
// 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
//             |
// 1 -> 2 -> 5
// Input: head = [1,2,3,3,4,4,5]
// Output: [1,2,5]

// Example 2:
// 1 -> 1 -> 1 -> 2 -> 3
//             |
// 2 -> 3
// Input: head = [1,1,1,2,3]
// Output: [2,3]
	
// Constraints:
//     The number of nodes in the list is in the range [0, 300].
//     -100 <= Node.val <= 100
//     The list is guaranteed to be sorted in ascending order.

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

// 双循环简单解法 O(n*m)
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    nilNode := &ListNode{Val: 0, Next: head}
    head = nilNode
    lastVal := 0
    for head.Next != nil && head.Next.Next != nil {
        if head.Next.Val == head.Next.Next.Val {
            lastVal = head.Next.Val
            for head.Next != nil && lastVal == head.Next.Val {
                head.Next = head.Next.Next
            }
        } else {
            head = head.Next
        }
    }
    return nilNode.Next
}

// 双指针+删除标志位，单循环解法 O(n)
func deleteDuplicates1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    nilNode := &ListNode{Val: 0, Next: head}
    // 上次遍历有删除操作的标志位
    lastIsDel := false
    // 虚拟空结点
    head = nilNode
    // 前后指针用于判断
    pre, back := head.Next, head.Next.Next
    // 每次只删除前面的一个重复的元素，留一个用于下次遍历判重
    // pre, back 指针的更新位置和值比较重要和巧妙
    for head.Next != nil && head.Next.Next != nil {
        // 如果
        if pre.Val != back.Val && lastIsDel {
            head.Next = head.Next.Next
            pre, back = head.Next, head.Next.Next
            lastIsDel = false
            continue
        }

        // 如果和前值一样
        if pre.Val == back.Val {
            head.Next = head.Next.Next
            pre, back = head.Next, head.Next.Next
            lastIsDel = true
        } else {
            head = head.Next
            pre, back = head.Next, head.Next.Next
            lastIsDel = false
        }
    }
    // 处理 [1,1] 这种删除还剩一个的情况
    if lastIsDel && head.Next != nil {
        head.Next = nil
    }
    return nilNode.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{
        Next: head,
    }
    curr,prev := head,dummy
    for curr != nil && curr.Next != nil {
        if curr.Val == curr.Next.Val {
            for curr.Next != nil && curr.Val == curr.Next.Val {
                curr = curr.Next
            }
            prev.Next = curr.Next
        }else{
            prev = prev.Next
        }
        curr = curr.Next
    }
    return dummy.Next
}

func main() {
    l11 := makeListNode([]int{1,2,3,3,4,4,5})
    fmt.Println("before: ")
    printListNode(l11) // [1,2,3,3,4,4,5] 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
    printListNode(deleteDuplicates(l11)) // [1,2,5] 1 -> 2 -> 5

    l12 := makeListNode([]int{1,1,1,2,3})
    fmt.Println("before: ")
    printListNode(l12) // [1,1,1,2,3] 1 -> 1 -> 1 -> 2 -> 3
    printListNode(deleteDuplicates(l12)) // [2,3] 2 -> 3

    l21 := makeListNode([]int{1,2,3,3,4,4,5})
    fmt.Println("before: ")
    printListNode(l21) // [1,2,3,3,4,4,5] 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
    printListNode(deleteDuplicates1(l21)) // [1,2,5] 1 -> 2 -> 5

    l22 := makeListNode([]int{1,1,1,2,3})
    fmt.Println("before: ")
    printListNode(l22) // [1,1,1,2,3] 1 -> 1 -> 1 -> 2 -> 3
    printListNode(deleteDuplicates1(l22)) // [2,3] 2 -> 3

    l31 := makeListNode([]int{1,2,3,3,4,4,5})
    fmt.Println("before: ")
    printListNode(l31) // [1,2,3,3,4,4,5] 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
    printListNode(deleteDuplicates2(l31)) // [1,2,5] 1 -> 2 -> 5

    l32 := makeListNode([]int{1,1,1,2,3})
    fmt.Println("before: ")
    printListNode(l32) // [1,1,1,2,3] 1 -> 1 -> 1 -> 2 -> 3
    printListNode(deleteDuplicates2(l32)) // [2,3] 2 -> 3
}

