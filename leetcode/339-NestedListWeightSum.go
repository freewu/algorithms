package main

// 339. Nested List Weight Sum
// You are given a nested list of integers nestedList. 
// Each element is either an integer or a list whose elements may also be integers or other lists.

// The depth of an integer is the number of lists that it is inside of. 
// For example, the nested list [1,[2,2],[[3],2],1] has each integer's value set to its depth.

// Return the sum of each integer in nestedList multiplied by its depth.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/14/nestedlistweightsumex1.png" />
// Input: nestedList = [[1,1],2,[1,1]]
// Output: 10
// Explanation: Four 1's at depth 2, one 2 at depth 1. 1*2 + 1*2 + 2*1 + 1*2 + 1*2 = 10.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/14/nestedlistweightsumex2.png" />
// Input: nestedList = [1,[4,[6]]]
// Output: 27
// Explanation: One 1 at depth 1, one 4 at depth 2, and one 6 at depth 3. 1*1 + 4*2 + 6*3 = 27.

// Example 3:
// Input: nestedList = [0]
// Output: 0

// Constraints:
//     1 <= nestedList.length <= 50
//     The values of the integers in the nested list is in the range [-100, 100].
//     The maximum depth of any integer is less than or equal to 50.


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
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func depthSum(nestedList []*NestedInteger) int {
    var dfs func(nesteds []*NestedInteger,depth int) int
    dfs = func(nesteds []*NestedInteger,depth int) int {
        sum := 0
        for _, v := range nesteds {
            if v.IsInteger() { 
                sum += v.GetInteger() * depth
            } else { // 为数据组则递归
                sum += dfs(v.GetList(), depth + 1)
            }
        }
        return sum
    }
    return dfs(nestedList,1)
}

func main() {
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/14/nestedlistweightsumex1.png" />
// Input: nestedList = [[1,1],2,[1,1]]
// Output: 10
// Explanation: Four 1's at depth 2, one 2 at depth 1. 1*2 + 1*2 + 2*1 + 1*2 + 1*2 = 10.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/14/nestedlistweightsumex2.png" />
// Input: nestedList = [1,[4,[6]]]
// Output: 27
// Explanation: One 1 at depth 1, one 4 at depth 2, and one 6 at depth 3. 1*1 + 4*2 + 6*3 = 27.

// Example 3:
// Input: nestedList = [0]
// Output: 0
fmt.Println()
}