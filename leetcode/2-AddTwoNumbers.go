package main

// 2. Add Two Numbers
// You are given two non-empty linked lists representing two non-negative integers. 
// The digits are stored in reverse order, and each of their nodes contains a single digit. 
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg" />
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

// Example 2:
// Input: l1 = [0], l2 = [0]
// Output: [0]

// Example 3:
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
 
// Constraints:
//     The number of nodes in each linked list is in the range [1, 100].
//     0 <= Node.val <= 9
//     It is guaranteed that the list represents a number that does not have leading zeros.

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var t = &ListNode{-1, nil}
    var l3 = &ListNode{-1, t}

    var flag = 0 // 进位符
    for {
        // 如果循环到任意节点为空直接跳出
        if nil == l1 || nil == l2 {
            break
        }

        var s = l1.Val + l2.Val + flag
        if s >= 10 {
            // 如果和大于10 取模型进位
            t.Next = &ListNode{s % 10, nil}
            flag = 1
        } else {
            t.Next = &ListNode{s, nil}
            flag = 0
        }

        t = t.Next
        l1 = l1.Next
        l2 = l2.Next
    }
    // 循环 l1 的剩余节点
    for {
        if nil == l1 {
            break
        }
        if flag == 1 {
            if (l1.Val + 1) >= 10 {
                flag = 1
                l1.Val = (l1.Val + 1) % 10
            } else {
                flag = 0
                l1.Val = l1.Val + 1
            }
        }
        t.Next = l1
        l1 = l1.Next
        t = t.Next
    }
    // 循环 l2 的剩余节点
    for {
        if nil == l2 {
            break
        }
        if flag == 1 {
            if (l2.Val + 1) >= 10 {
                flag = 1
                l2.Val = (l2.Val + 1) % 10
            } else {
                flag = 0
                l2.Val = l2.Val + 1
            }
        }
        t.Next = l2
        l2 = l2.Next
        t = t.Next
    }
    // 如果还存在进位
    if 1 == flag {
        t.Next = &ListNode{1, nil}
    }
    return l3.Next.Next
}

// best speed solution 这个方案好像上有问题的
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    var head *ListNode
    var currentNode *ListNode
    // 都为空时跳出循环
    for (l1 != nil) || (l2 != nil) || (carry != 0) {
        l1_value := 0
        if l1 != nil {
            l1_value = l1.Val
            l1 = l1.Next
        }

        l2_value := 0
        if l2 != nil {
            l2_value = l2.Val
            l2 = l2.Next
        }

        // fmt.Println(l1_value, l2_value, carry)
        current_value := carry + l1_value + l2_value
        carry = current_value / 10         // int
        current_value = current_value % 10 // 取余

        if head == nil {
            head = &ListNode{Val: current_value, Next: nil}
            currentNode = head
        } else {
            next := ListNode{Val: current_value, Next: nil}
            currentNode.Next = &next
            currentNode = &next
        }
    }
    return head
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        return nil
    }
    mergeListHead := &ListNode{}
    cur := mergeListHead
    carray := 0
    for (l1 != nil || l2 != nil) {
        a, b := 0, 0
        if l1 != nil {
            a = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            b = l2.Val
            l2 = l2.Next
        }
        sum := a + b + carray
        carray = sum/10
        sum = sum%10

        newNode := &ListNode{Val:sum}
        cur.Next = newNode
        cur = cur.Next
    }
    if carray > 0 {
        newNode := &ListNode{Val:carray}
        cur.Next = newNode
    }
    return mergeListHead.Next
}

func main() {
    // Explanation: 342 + 465 = 807.
    l11 := makeListNode([]int{2,4,3})
    l12 := makeListNode([]int{5,6,4})
    printListNode(l11) // 2 -> 4 -> 3
    printListNode(l12) // 5 -> 6 -> 4
    fmt.Println("addTwoNumbers: ")
    printListNode(addTwoNumbers(l11, l12)) // 7 -> 0 -> 8
    fmt.Println()

    l21 := makeListNode([]int{0})
    l22 := makeListNode([]int{0})
    printListNode(l21) // 0
    printListNode(l22) // 0
    fmt.Println("addTwoNumbers: ")
    printListNode(addTwoNumbers(l21, l22)) // 0
    fmt.Println()

    l31 := makeListNode([]int{9,9,9,9,9,9,9})
    l32 := makeListNode([]int{9,9,9,9})
    printListNode(l31) // 9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9
    printListNode(l32) // 9 -> 9 -> 9 -> 9
    fmt.Println("addTwoNumbers: ")
    printListNode(addTwoNumbers(l31, l32)) // 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
    fmt.Println()


    // Explanation: 342 + 465 = 807.
    l111 := makeListNode([]int{2,4,3})
    l112 := makeListNode([]int{5,6,4})
    printListNode(l111) // 2 -> 4 -> 3
    printListNode(l112) // 5 -> 6 -> 4
    fmt.Println("addTwoNumbers1: ")
    printListNode(addTwoNumbers1(l111, l112)) // 7 -> 0 -> 8
    fmt.Println()

    l121 := makeListNode([]int{0})
    l122 := makeListNode([]int{0})
    printListNode(l121) // 0
    printListNode(l122) // 0
    fmt.Println("addTwoNumbers1: ")
    printListNode(addTwoNumbers1(l121, l122)) // 0
    fmt.Println()

    l131 := makeListNode([]int{9,9,9,9,9,9,9})
    l132 := makeListNode([]int{9,9,9,9})
    printListNode(l131) // 9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9
    printListNode(l132) // 9 -> 9 -> 9 -> 9
    fmt.Println("addTwoNumbers1: ")
    printListNode(addTwoNumbers1(l131, l132)) // 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
    fmt.Println()


    // Explanation: 342 + 465 = 807.
    l211 := makeListNode([]int{2,4,3})
    l212 := makeListNode([]int{5,6,4})
    printListNode(l211) // 2 -> 4 -> 3
    printListNode(l212) // 5 -> 6 -> 4
    fmt.Println("addTwoNumbers2: ")
    printListNode(addTwoNumbers2(l211, l212)) // 7 -> 0 -> 8
    fmt.Println()

    l221 := makeListNode([]int{0})
    l222 := makeListNode([]int{0})
    printListNode(l221) // 0
    printListNode(l222) // 0
    fmt.Println("addTwoNumbers2: ")
    printListNode(addTwoNumbers2(l221, l222)) // 0
    fmt.Println()

    l231 := makeListNode([]int{9,9,9,9,9,9,9})
    l232 := makeListNode([]int{9,9,9,9})
    printListNode(l231) // 9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9
    printListNode(l232) // 9 -> 9 -> 9 -> 9
    fmt.Println("addTwoNumbers2: ")
    printListNode(addTwoNumbers2(l231, l232)) // 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
    fmt.Println()


    // // Explanation: 342 + 465 = 807.
    // l311 := makeListNode([]int{2,4,3})
    // l312 := makeListNode([]int{5,6,4})
    // printListNode(l311) // 2 -> 4 -> 3
    // printListNode(l312) // 5 -> 6 -> 4
    // fmt.Println("addTwoNumbers3: ")
    // printListNode(addTwoNumbers3(l311, l312)) // 7 -> 0 -> 8
    // fmt.Println()

    // l321 := makeListNode([]int{0})
    // l322 := makeListNode([]int{0})
    // printListNode(l321) // 0
    // printListNode(l222) // 0
    // fmt.Println("addTwoNumbers3: ")
    // printListNode(addTwoNumbers3(l321, l322)) // 0
    // fmt.Println()

    // l331 := makeListNode([]int{9,9,9,9,9,9,9})
    // l332 := makeListNode([]int{9,9,9,9})
    // printListNode(l331) // 9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9
    // printListNode(l332) // 9 -> 9 -> 9 -> 9
    // fmt.Println("addTwoNumbers3: ")
    // printListNode(addTwoNumbers3(l331, l332)) // 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
    // fmt.Println()
    
}
