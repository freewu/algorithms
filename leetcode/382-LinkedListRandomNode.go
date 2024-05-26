package main

// 382. Linked List Random Node
// Given a singly linked list, return a random node's value from the linked list. 
// Each node must have the same probability of being chosen.

// Implement the Solution class:
//     Solution(ListNode head) 
//         Initializes the object with the head of the singly-linked list head.
//     int getRandom() 
//         Chooses a node randomly from the list and returns its value. 
//         All the nodes of the list should be equally likely to be chosen.

// Example 1:
// <img src="" />
// Input
// ["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
// [[[1, 2, 3]], [], [], [], [], []]
// Output
// [null, 1, 3, 2, 2, 3]
// Explanation
// Solution solution = new Solution([1, 2, 3]);
// solution.getRandom(); // return 1
// solution.getRandom(); // return 3
// solution.getRandom(); // return 2
// solution.getRandom(); // return 2
// solution.getRandom(); // return 3
// // getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
 
// Constraints:
//     The number of nodes in the linked list will be in the range [1, 10^4].
//     -10^4 <= Node.val <= 10^4
//     At most 104 calls will be made to getRandom.
 
// Follow up:
//     What if the linked list is extremely large and its length is unknown to you?
//     Could you solve this efficiently without using extra space?

import "fmt"
import "math/rand"

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
type Solution struct {
    data []int
}

func Constructor(head *ListNode) Solution {
    arr := []int{}
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    return Solution{ data: arr }
}

func (this *Solution) GetRandom() int {
    return this.data[rand.Intn(len(this.data))]
}
 
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
    // Solution solution = new Solution([1, 2, 3]);]
    obj := Constructor(makeListNode([]int{1, 2, 3}))
    // solution.getRandom(); // return 1
    fmt.Println(obj.GetRandom())
    // solution.getRandom(); // return 3
    fmt.Println(obj.GetRandom())
    // solution.getRandom(); // return 2
    fmt.Println(obj.GetRandom())
    // solution.getRandom(); // return 2
    fmt.Println(obj.GetRandom())
    // solution.getRandom(); // return 3
    fmt.Println(obj.GetRandom())
    // // getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
}
