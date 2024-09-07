package main

// 3062. Winner of the Linked List Game
// You are given the head of a linked list of even length containing integers.

// Each odd-indexed node contains an odd integer and each even-indexed node contains an even integer.

// We call each even-indexed node and its next node a pair, 
// e.g., the nodes with indices 0 and 1 are a pair, the nodes with indices 2 and 3 are a pair, and so on.

// For every pair, we compare the values of the nodes in the pair:
//     If the odd-indexed node is higher, the "Odd" team gets a point.
//     If the even-indexed node is higher, the "Even" team gets a point.

// Return the name of the team with the higher points, if the points are equal, return "Tie".

// Example 1:
// Input: head = [2,1]
// Output: "Even"
// Explanation: There is only one pair in this linked list and that is (2,1). Since 2 > 1, the Even team gets the point.
// Hence, the answer would be "Even".

// Example 2:
// Input: head = [2,5,4,7,20,5]
// Output: "Odd"
// Explanation: There are 3 pairs in this linked list. Let's investigate each pair individually:
// (2,5) -> Since 2 < 5, The Odd team gets the point.
// (4,7) -> Since 4 < 7, The Odd team gets the point.
// (20,5) -> Since 20 > 5, The Even team gets the point.
// The Odd team earned 2 points while the Even team got 1 point and the Odd team has the higher points.
// Hence, the answer would be "Odd".

// Example 3:
// Input: head = [4,5,2,1]
// Output: "Tie"s
// Explanation: There are 2 pairs in this linked list. Let's investigate each pair individually:
// (4,5) -> Since 4 < 5, the Odd team gets the point.
// (2,1) -> Since 2 > 1, the Even team gets the point.
// Both teams earned 1 point.
// Hence, the answer would be "Tie".

// Constraints:
//     The number of nodes in the list is in the range [2, 100].
//     The number of nodes in the list is even.
//     1 <= Node.val <= 100
//     The value of each odd-indexed node is odd.
//     The value of each even-indexed node is even.

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
func gameResult(head *ListNode) string {
    index, even, odd := 0, 0, 0
    for nil != head {
        if index % 2 == 0 { // 两个比对一次
            if head.Next != nil {
                v1, v2 := head.Val, head.Next.Val
                if v1 > v2 {
                    odd++
                } else if v1 < v2 {
                    even++
                }
            }
        }
        index++
        head = head.Next
    }
    if even == odd { return "Tie" }
    if even > odd { return "Odd" }
    return "Even"
}

func gameResult1(head *ListNode) string {
    score := 0
    for head != nil {
        if head.Val > head.Next.Val {
            score += 1
        }else {
            score -= 1
        }
        head = head.Next.Next
    }
    if score > 0 {
        return "Even"
    } else if score < 0 {
        return "Odd"
    }
    return "Tie"
}

func main() {
    // Example 1:
    // Input: head = [2,1]
    // Output: "Even"
    // Explanation: There is only one pair in this linked list and that is (2,1). Since 2 > 1, the Even team gets the point.
    // Hence, the answer would be "Even".
    list1 := makeListNode([]int{2,1})
    printListNode(list1)
    fmt.Println(gameResult(list1)) // Even
    // Example 2:
    // Input: head = [2,5,4,7,20,5]
    // Output: "Odd"
    // Explanation: There are 3 pairs in this linked list. Let's investigate each pair individually:
    // (2,5) -> Since 2 < 5, The Odd team gets the point.
    // (4,7) -> Since 4 < 7, The Odd team gets the point.
    // (20,5) -> Since 20 > 5, The Even team gets the point.
    // The Odd team earned 2 points while the Even team got 1 point and the Odd team has the higher points.
    // Hence, the answer would be "Odd".
    list2 := makeListNode([]int{2,5,4,7,20,5})
    printListNode(list2)
    fmt.Println(gameResult(list2)) // Odd
    // Example 3:
    // Input: head = [4,5,2,1]
    // Output: "Tie"s
    // Explanation: There are 2 pairs in this linked list. Let's investigate each pair individually:
    // (4,5) -> Since 4 < 5, the Odd team gets the point.
    // (2,1) -> Since 2 > 1, the Even team gets the point.
    // Both teams earned 1 point.
    // Hence, the answer would be "Tie".
    list3 := makeListNode([]int{4,5,2,1})
    printListNode(list3)
    fmt.Println(gameResult(list3)) // Tie

    list11 := makeListNode([]int{2,1})
    printListNode(list11)
    fmt.Println(gameResult1(list11)) // Even
    list12 := makeListNode([]int{2,5,4,7,20,5})
    printListNode(list12)
    fmt.Println(gameResult1(list12)) // Odd
    list13 := makeListNode([]int{4,5,2,1})
    printListNode(list13)
    fmt.Println(gameResult1(list13)) // Tie
}