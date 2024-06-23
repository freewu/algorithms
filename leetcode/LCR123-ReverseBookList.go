package main

// LCR 123. 图书整理 I
// 书店店员有一张链表形式的书单，每个节点代表一本书，节点中的值表示书的编号。
// 为更方便整理书架，店员需要将书单倒过来排列，就可以从最后一本书开始整理，逐一将书放回到书架上。请倒序返回这个书单链表。

// 示例 1：
// 输入：head = [3,6,4,1]
// 输出：[1,4,6,3]

// 提示：
//     0 <= 链表长度 <= 10000

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
func reverseBookList(head *ListNode) []int {
    res := []int{}
    if nil == head {
        return res
    }
    reverse := func(head *ListNode) *ListNode {
        var res *ListNode
        for head != nil {
            next := head.Next  // 备份head.Next
            head.Next = res // 更新  head.Next
            res = head      // 移动 new_head
            head = next
            //fmt.Println("new_head: ",new_head,"head: ",head)
        }
        return res
    }
    head = reverse(head)
    for {
        res = append(res, head.Val)
        if nil == head.Next {
            break
        }
        head = head.Next
    }
    return res
}

func main() {
    l1 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("l1: ")
    printListNode(l1) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println("reverseBookList(l1): ",reverseBookList(l1)) // [5 4 3 2 1]

    l2 := makeListNode([]int{1,2}) 
    fmt.Println("l2: ")
    printListNode(l2) // 1 -> 2
    fmt.Println("reverseBookList(l2): ", reverseBookList(l2)) // [2 1]

    l3 := makeListNode([]int{})
    fmt.Println("l3: ")
    printListNode(l3)
    fmt.Println("reverseBookList(l3): ", reverseBookList(l3)) // []
}