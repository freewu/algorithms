package main

// 1171. Remove Zero Sum Consecutive Nodes from Linked List
// Given the head of a linked list, we repeatedly delete consecutive sequences of nodes that sum to 0 until there are no such sequences.
// After doing so, return the head of the final linked list.  
// You may return any such answer.
// (Note that in the examples below, all sequences are serializations of ListNode objects.)

// Example 1:
// Input: head = [1,2,-3,3,1]
// Output: [3,1]
// Note: The answer [1,2,1] would also be accepted.

// Example 2:
// Input: head = [1,2,3,-3,4]
// Output: [1,2,4]

// Example 3:
// Input: head = [1,2,3,-3,-2]
// Output: [1]
 
// Constraints:
//         The given linked list will contain between 1 and 1000 nodes.
//         Each node in the linked list has -1000 <= node.val <= 1000.

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
 func removeZeroSumSublists(head *ListNode) *ListNode {
    // 计算累加和，和作为 key 存在 map 中，value 存那个节点的指针。如果字典中出现了重复的和，代表出现了和为 0 的段。
	sum, sumMap, cur := 0, make(map[int]*ListNode), head
	// 字典中增加 0 这个特殊值，是为了防止最终链表全部消除完
	sumMap[0] = nil
	for cur != nil {
		sum = sum + cur.Val
		if ptr, ok := sumMap[sum]; ok {
			// 在字典中找到了重复的和，代表 [ptr, tmp] 中间的是和为 0 的段，要删除的就是这一段。
			// 同时删除 map 中中间这一段的和
			if ptr != nil {
				iter := ptr.Next
				tmpSum := sum + iter.Val
				for tmpSum != sum {
					// 删除中间为 0 的那一段，tmpSum 不断的累加删除 map 中的和
					delete(sumMap, tmpSum)
					iter = iter.Next
					tmpSum = tmpSum + iter.Val
				}
				ptr.Next = cur.Next
			} else {
				head = cur.Next
				sumMap = make(map[int]*ListNode)
				sumMap[0] = nil
			}
		} else {
			sumMap[sum] = cur
		}
		cur = cur.Next
	}
	return head
 }

// best solution
func removeZeroSumSublists1(head *ListNode) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	curr := newHead
	sum := 0
    // 把 链表的值 累加 并到写到 Map  中  
	m := make(map[int]*ListNode)
	for curr != nil {
		sum += curr.Val
		m[sum] = curr
		curr = curr.Next
	}
	sum = 0
	curr = newHead
	for curr != nil {
		sum += curr.Val
		curr.Next = m[sum].Next
		curr = curr.Next
	}
	return newHead.Next
}

 func main() {
    printListNode(makeListNode([]int{1,2,-3,3,1})) // [1] -> [2] -> [-3] -> [3] -> [1]
	printListNode(removeZeroSumSublists(makeListNode([]int{1,2,-3,3,1}))) // [3] -> [1]

    printListNode(makeListNode([]int{1,2,3,-3,4})) // [1] -> [2] -> [3] -> [-3] -> [4]
	printListNode(removeZeroSumSublists(makeListNode([]int{1,2,3,-3,4}))) // [1] -> [2] -> [4]

    printListNode(makeListNode([]int{1,2,3,-3,-2})) // [1] -> [2] -> [3] -> [-3] -> [-2]
	printListNode(removeZeroSumSublists(makeListNode([]int{1,2,3,-3,-2}))) // [1]

    printListNode(makeListNode([]int{1,2,-3,3,1})) // [1] -> [2] -> [-3] -> [3] -> [1]
	printListNode(removeZeroSumSublists1(makeListNode([]int{1,2,-3,3,1}))) // [3] -> [1]

    printListNode(makeListNode([]int{1,2,3,-3,4})) // [1] -> [2] -> [3] -> [-3] -> [4]
	printListNode(removeZeroSumSublists1(makeListNode([]int{1,2,3,-3,4}))) // [1] -> [2] -> [4]

    printListNode(makeListNode([]int{1,2,3,-3,-2})) // [1] -> [2] -> [3] -> [-3] -> [-2]
	printListNode(removeZeroSumSublists1(makeListNode([]int{1,2,3,-3,-2}))) // [1]
}
