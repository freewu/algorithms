package main

// LCR 022. 环形链表 II
// 给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。如果链表无环，则返回 null。

// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

// 说明：不允许修改给定的链表。

// 示例 1：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png" />
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。

// 示例 2：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png" />
// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点
// 解释：链表中有一个环，其尾部连接到第一个节点。

// 示例 3：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png" />
// 输入：head = [1], pos = -1
// 输出：返回 null
// 解释：链表中没有环。
 
// 提示：
//     链表中节点的数目范围在范围 [0, 10^4] 内
//     -10^5 <= Node.val <= 10^5
//     pos 的值为 -1 或者链表中的一个有效索引

// 进阶：是否可以使用 O(1) 空间解决此题？

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