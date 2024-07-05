package main

// 2058. Find the Minimum and Maximum Number of Nodes Between Critical Points
// A critical point in a linked list is defined as either a local maxima or a local minima.
// A node is a local maxima if the current node has a value strictly greater than the previous node and the next node.
// A node is a local minima if the current node has a value strictly smaller than the previous node and the next node.
// Note that a node can only be a local maxima/minima if there exists both a previous node and a next node.
// Given a linked list head, return an array of length 2 containing [minDistance, maxDistance] where minDistance is the minimum distance between any two distinct critical points and maxDistance is the maximum distance between any two distinct critical points. If there are fewer than two critical points, return [-1, -1].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/13/a1.png" />
// Input: head = [3,1]
// Output: [-1,-1]
// Explanation: There are no critical points in [3,1].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/13/a2.png" />
// Input: head = [5,3,1,2,5,1,2]
// Output: [1,3]
// Explanation: There are three critical points:
// - [5,3,1,2,5,1,2]: The third node is a local minima because 1 is less than 3 and 2.
// - [5,3,1,2,5,1,2]: The fifth node is a local maxima because 5 is greater than 2 and 1.
// - [5,3,1,2,5,1,2]: The sixth node is a local minima because 1 is less than 5 and 2.
// The minimum distance is between the fifth and the sixth node. minDistance = 6 - 5 = 1.
// The maximum distance is between the third and the sixth node. maxDistance = 6 - 3 = 3.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/10/14/a5.png" />
// Input: head = [1,3,2,2,3,2,2,2,7]
// Output: [3,3]
// Explanation: There are two critical points:
// - [1,3,2,2,3,2,2,2,7]: The second node is a local maxima because 3 is greater than 1 and 2.
// - [1,3,2,2,3,2,2,2,7]: The fifth node is a local maxima because 3 is greater than 2 and 2.
// Both the minimum and maximum distances are between the second and the fifth node.
// Thus, minDistance and maxDistance is 5 - 2 = 3.
// Note that the last node is not considered a local maxima because it does not have a next node.
 
// Constraints:
//     The number of nodes in the list is in the range [2, 10^5].
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
func nodesBetweenCriticalPoints(head *ListNode) []int {
    var prev *ListNode
    next, curr := head.Next, head
    criticalPoints, nodeIndex := []int{}, 1
    for curr != nil && curr.Next != nil {
        if (prev != nil && prev.Val > curr.Val && curr.Val < next.Val) ||
           (prev != nil && prev.Val < curr.Val && curr.Val > next.Val ) {
            criticalPoints = append(criticalPoints, nodeIndex)
        }
        prev, curr = curr, curr.Next
        next = curr.Next
        nodeIndex++
    }
    if len(criticalPoints) < 2 {
        return []int{ -1, -1}
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    mx, mn := criticalPoints[len(criticalPoints) - 1] - criticalPoints[0], 1 << 32 - 1
    for i := 0; i < len(criticalPoints) - 1 ; i++ {
        dif := criticalPoints[i + 1] - criticalPoints[i]
        mn = min(mn, dif) 
    }
    return []int{ mn, mx }
}

func nodesBetweenCriticalPoints1(head *ListNode) []int {
    mn, first, last := 1 << 32 - 1, -1, -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for p, i := head, 0; p != nil && p.Next != nil && p.Next.Next != nil; p, i = p.Next, i+1 {
        if (p.Next.Val - p.Val) * (p.Next.Val - p.Next.Next.Val) > 0 {
            if last != -1 {
                mn = min(mn,  i - last)
            }
            if first == -1 {
                first = i
            }
            last = i
        }
    } 
    if first == last {
        return []int{-1, -1 }
    } 
    return []int{mn, last - first}
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/13/a1.png" />
    // Input: head = [3,1]
    // Output: [-1,-1]
    // Explanation: There are no critical points in [3,1].
    list1 := makeListNode([]int{3, 1})
    printListNode(list1) // 3 -> 1
    fmt.Println(nodesBetweenCriticalPoints(list1)) // [-1,-1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/13/a2.png" />
    // Input: head = [5,3,1,2,5,1,2]
    // Output: [1,3]
    // Explanation: There are three critical points:
    // - [5,3,1,2,5,1,2]: The third node is a local minima because 1 is less than 3 and 2.
    // - [5,3,1,2,5,1,2]: The fifth node is a local maxima because 5 is greater than 2 and 1.
    // - [5,3,1,2,5,1,2]: The sixth node is a local minima because 1 is less than 5 and 2.
    // The minimum distance is between the fifth and the sixth node. minDistance = 6 - 5 = 1.
    // The maximum distance is between the third and the sixth node. maxDistance = 6 - 3 = 3.
    list2 := makeListNode([]int{5,3,1,2,5,1,2})
    printListNode(list2) // 5 -> 3 -> 1 -> 2 -> 5 -> 1 -> 2
    fmt.Println(nodesBetweenCriticalPoints(list2)) // [1,3]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/10/14/a5.png" />
    // Input: head = [1,3,2,2,3,2,2,2,7]
    // Output: [3,3]
    // Explanation: There are two critical points:
    // - [1,3,2,2,3,2,2,2,7]: The second node is a local maxima because 3 is greater than 1 and 2.
    // - [1,3,2,2,3,2,2,2,7]: The fifth node is a local maxima because 3 is greater than 2 and 2.
    // Both the minimum and maximum distances are between the second and the fifth node.
    // Thus, minDistance and maxDistance is 5 - 2 = 3.
    // Note that the last node is not considered a local maxima because it does not have a next node.
    list3 := makeListNode([]int{1,3,2,2,3,2,2,2,7})
    printListNode(list3) // 1 -> 3 -> 2 -> 2 -> 3 -> 2 -> 2 -> 2 -> 7
    fmt.Println(nodesBetweenCriticalPoints(list3)) // [3,3]

    fmt.Println(nodesBetweenCriticalPoints1(list1)) // [-1,-1]
    fmt.Println(nodesBetweenCriticalPoints1(list2)) // [1,3]
    fmt.Println(nodesBetweenCriticalPoints1(list3)) // [3,3]
}