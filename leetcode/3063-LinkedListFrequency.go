package main

// 3063. Linked List Frequency
// Given the head of a linked list containing k distinct elements, 
// return the head to a linked list of length k containing the frequency of each distinct element in the given linked list in any order.

// Example 1:
// Input: head = [1,1,2,1,2,3]
// Output: [3,2,1]
// Explanation: There are 3 distinct elements in the list. The frequency of 1 is 3, the frequency of 2 is 2 and the frequency of 3 is 1. Hence, we return 3 -> 2 -> 1.
// Note that 1 -> 2 -> 3, 1 -> 3 -> 2, 2 -> 1 -> 3, 2 -> 3 -> 1, and 3 -> 1 -> 2 are also valid answers.

// Example 2:
// Input: head = [1,1,2,2,2]
// Output: [2,3]
// Explanation: There are 2 distinct elements in the list. The frequency of 1 is 2 and the frequency of 2 is 3. Hence, we return 2 -> 3.

// Example 3:
// Input: head = [6,5,4,3,2,1]
// Output: [1,1,1,1,1,1]
// Explanation: There are 6 distinct elements in the list. The frequency of each of them is 1. Hence, we return 1 -> 1 -> 1 -> 1 -> 1 -> 1.

// Constraints:
//     The number of nodes in the list is in the range [1, 10^5].
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
func frequenciesOfElements(head *ListNode) *ListNode {
    mp := make(map[int]int)
    for ; head != nil; head = head.Next { // 遍历统计每个值出现的频率
        if _,ok := mp[head.Val]; !ok{
            mp[head.Val] =1
        } else {
            mp[head.Val]++
        }
    }
    var res, dummy *ListNode
    dummy = &ListNode{ Val: 0 }
    for _, v := range mp {
        node := &ListNode{ Val: v }
        if res == nil { // 第一次进入循环
            res = node
            dummy.Next = res
        } else {
            res.Next = node
            res = res.Next
        }
    }
    return dummy.Next
}

func frequenciesOfElements1(head *ListNode) *ListNode {
    mp := make(map[int]int)
    for cur := head; cur != nil; cur = cur.Next {
        mp[cur.Val]++
    }

    dummy := &ListNode{}
    cur := dummy
    for _, v := range mp {
        cur.Next = &ListNode{ Val: v }
        cur = cur.Next
    }
    return dummy.Next
}

func main() {
    // Example 1:
    // Input: head = [1,1,2,1,2,3]
    // Output: [3,2,1]  is 2 and the frequency of 3 is 1. Hence, we return 3 -> 2 -> 1.
    // Note that 1 -> 2 -> 3, 1 -> 3 -> 2, 2 -> 1 -> 3, 2 -> 3 -> 1, and 3 -> 1 -> 2 are also valid answers.
    list1 := makeListNode([]int{1,1,2,1,2,3}) 
    printListNode(list1) // 1 -> 1 -> 2 -> 1 -> 2 -> 3
    printListNode(frequenciesOfElements(list1)) // 3 -> 2 -> 1
    // Example 2:
    // Input: head = [1,1,2,2,2]
    // Output: [2,3]
    // Explanation: There are 2 distinct elements in the list. The frequency of 1 is 2 and the frequency of 2 is 3. Hence, we return 2 -> 3.
    list2 := makeListNode([]int{1,1,2,2,2}) 
    printListNode(list2) // 1 -> 1 -> 2 -> 2 -> 2
    printListNode(frequenciesOfElements(list2)) // 2 -> 3
    // Example 3:
    // Input: head = [6,5,4,3,2,1]
    // Output: [1,1,1,1,1,1]
    // Explanation: There are 6 distinct elements in the list. The frequency of each of them is 1. Hence, we return 1 -> 1 -> 1 -> 1 -> 1 -> 1.
    list3 := makeListNode([]int{6,5,4,3,2,1}) 
    printListNode(list3) // 6 -> 5 -> 4 -> 3 -> 2 -> 1
    printListNode(frequenciesOfElements(list3)) // 1 -> 1 -> 1 -> 1 -> 1 -> 1

    list11 := makeListNode([]int{1,1,2,1,2,3}) 
    printListNode(list11) // 1 -> 1 -> 2 -> 1 -> 2 -> 3
    printListNode(frequenciesOfElements1(list11)) // 3 -> 2 -> 1
    list12 := makeListNode([]int{1,1,2,2,2}) 
    printListNode(list12) // 1 -> 1 -> 2 -> 2 -> 2
    printListNode(frequenciesOfElements1(list12)) // 2 -> 3
    list13 := makeListNode([]int{6,5,4,3,2,1}) 
    printListNode(list13) // 6 -> 5 -> 4 -> 3 -> 2 -> 1
    printListNode(frequenciesOfElements1(list13)) // 1 -> 1 -> 1 -> 1 -> 1 -> 1
}