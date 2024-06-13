package main

// 142. Linked List Cycle II
// Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. 
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. 
// Note that pos is not passed as a parameter.
// Do not modify the linked list.

// Example 1:
// <img src="" />
// Input: head = [3,2,0,-4], pos = 1
// Output: tail connects to node index 1
// Explanation: There is a cycle in the linked list, where tail connects to the second node.

// Example 2:
// <img src="" />
// Input: head = [1,2], pos = 0
// Output: tail connects to node index 0
// Explanation: There is a cycle in the linked list, where tail connects to the first node.

// Example 3:
// Input: head = [1], pos = -1
// Output: no cycle
// Explanation: There is no cycle in the linked list.

// Constraints:
//         The number of the nodes in the list is in the range [0, 10^4].
//         -10^5 <= Node.val <= 10^5
//         pos is -1 or a valid index in the linked-list.
 
// Follow up: Can you solve it using O(1) (i.e. constant) memory?

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    hasCycle := func (head *ListNode) (bool, *ListNode) {
        fast := head
        slow := head
        for slow != nil && fast != nil && fast.Next != nil {
            fast = fast.Next.Next
            slow = slow.Next
            if fast == slow {
                return true, slow
            }
        }
        return false, nil
    }
    // 先判断是否存在环形
    isCycle, slow := hasCycle(head)
    if !isCycle {
        return nil
    }
    // 2 个指针相遇以后，如果 slow 继续往前走，fast 指针回到起点 head，
    // 两者都每次走一步，那么必定会在环的起点相遇，相遇以后输出这个点即是结果
    fast := head
    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }
    return fast
}

// best solution
// double pointer 
func detectCycle1(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    low, fast := head, head
    for fast != nil {
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        low = low.Next
        // 相遇则从答案头开始和 slow 都走一步 至到相同返回
        if low == fast {
            ans := head
            for ans != low {
                ans = ans.Next
                low = low.Next
            }
            return ans 
        }
    }
    return nil
}