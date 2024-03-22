package main

// 234. Palindrome Linked List
// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
// Example 1:
// (1) -> (2) -> (2) -> (1)
// Input: head = [1,2,2,1]
// Output: true

// Example 2:
// (1) -> (2) 
// Input: head = [1,2]
// Output: false
 
// Constraints:
//     The number of nodes in the list is in the range [1, 10^5].
//     0 <= Node.val <= 9

// Follow up: Could you do it in O(n) time and O(1) space?

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    arr := make([]int,0) // 创建一个数组保存链表的值
    // 遍历链表取值 
    for {
        arr = append(arr, head.Val)
        if nil == head.Next {
            break
        }
        head = head.Next
    }
    // 判断是否是回文
    i,j := 0,len(arr) - 1 
    for {
        if i > j {
            break
        }
        if arr[i] != arr[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func isPalindrome1(head *ListNode) bool {
    getListTail := func (node *ListNode) *ListNode {
        slow, fast := node, node
        for nil != fast.Next && nil != fast.Next.Next {
            fast = fast.Next.Next
            slow = slow.Next
        }
        return slow
    }
    reverseList := func (node *ListNode) *ListNode {
        var pre *ListNode
        for nil != node {
            temp := node.Next
            node.Next = pre
            pre = node
            node = temp
        }
        return pre
    }
    // 找到前半部分链表的尾节点
    tail := getListTail(head)
    // 反转后半部分链表
    list := reverseList(tail.Next)
    // 判断是否回文
    p1, p2 := head, list
    for nil != p2 {
        if p1.Val != p2.Val {
            return false
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    // 恢复链表
    tail.Next = reverseList(list)
    return true
}

// 快慢指针
func isPalindrome2(head *ListNode) bool {
    reverse := func(head *ListNode) *ListNode{
        var prev *ListNode
        var futr *ListNode
        for head != nil{
            futr = head.Next
            head.Next = prev
            prev = head
            head = futr
        }
        return prev
    } 
    
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    second := slow.Next
    slow.Next = nil
    second = reverse(second)
    
    for second != nil && head != nil {
        if second.Val != head.Val {
            return false
        }
        second = second.Next
        head = head.Next
    }
    return true
}


func main() {
    fmt.Println("isPalindrome: ")
    l1 := makeListNode([]int{1,2,2,1}) 
    printListNode(l1)
    fmt.Println(isPalindrome(l1)) // true
    l2 := makeListNode([]int{1,2}) 
    printListNode(l2)
    fmt.Println(isPalindrome(l2)) // false
    l3 := makeListNode([]int{1,2,3,2,1}) 
    printListNode(l3)
    fmt.Println(isPalindrome(l3)) // true

    fmt.Println("isPalindrome1: ")
    l1 = makeListNode([]int{1,2,2,1}) 
    printListNode(l1)
    fmt.Println(isPalindrome1(l1)) // true
    l2 = makeListNode([]int{1,2}) 
    printListNode(l2)
    fmt.Println(isPalindrome1(l2)) // false
    l3 = makeListNode([]int{1,2,3,2,1}) 
    printListNode(l3)
    fmt.Println(isPalindrome1(l3)) // true

    fmt.Println("isPalindrome2: ")
    l1 = makeListNode([]int{1,2,2,1}) 
    printListNode(l1)
    fmt.Println(isPalindrome2(l1)) // true
    l2 = makeListNode([]int{1,2}) 
    printListNode(l2)
    fmt.Println(isPalindrome2(l2)) // false
    l3 = makeListNode([]int{1,2,3,2,1}) 
    printListNode(l3)
    fmt.Println(isPalindrome2(l3)) // true
}