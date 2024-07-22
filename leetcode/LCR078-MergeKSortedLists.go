package main

// LCR 078. 合并 K 个升序链表
// 给定一个链表数组，每个链表都已经按升序排列。
// 请将所有链表合并到一个升序链表中，返回合并后的链表。

// 示例 1：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6

// 示例 2：
// 输入：lists = []
// 输出：[]

// 示例 3：
// 输入：lists = [[]]
// 输出：[]

// 提示：
//     k == lists.length
//     0 <= k <= 10^4
//     0 <= lists[i].length <= 500
//     -10^4 <= lists[i][j] <= 10^4
//     lists[i] 按 升序 排列
//     lists[i].length 的总和不超过 10^4

import "fmt"
import "container/heap"

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
// 优先队列（小根堆）
func mergeKLists(lists []*ListNode) *ListNode {
    // 创建一个小根堆来 pq 维护所有链表的头节点
    pq := hp{}
    for _, head := range lists {
        if head != nil {
            pq = append(pq, head)
        }
    }
    heap.Init(&pq)
    dummy := &ListNode{}
    cur := dummy
    for len(pq) > 0 {
        // 每次从小根堆中取出值最小的节点，
        cur.Next = heap.Pop(&pq).(*ListNode)
        cur = cur.Next
        if cur.Next != nil {
            // 添加到结果链表的末尾，然后将该节点的下一个节点加入堆中
            heap.Push(&pq, cur.Next)
        }
        // 重复上述步骤直到堆为空
    }
    return dummy.Next
}

type hp []*ListNode
func (h hp)  Len() int           { return len(h) }
func (h hp)  Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 递归合并
func mergeKLists1(lists []*ListNode) *ListNode {
    length := len(lists)
    if length < 1 {
        return nil
    }
    if length == 1 {
        return lists[0]
    }
    num := length / 2
    left := mergeKLists1(lists[:num])
    right := mergeKLists1(lists[num:])
    return mergeTwoLists1(left, right)
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists1(l1.Next, l2)
        return l1
    }
    l2.Next = mergeTwoLists1(l1, l2.Next)
    return l2
}

func main() {
    printListNode(mergeKLists(
        []*ListNode{
            makeListNode([]int{1,4,5}),
            makeListNode([]int{1,3,4}),
            makeListNode([]int{2,6}),
        },
    )) // 1->1->2->3->4->4->5->6

    printListNode(mergeKLists(
        []*ListNode{},
    )) // 
    printListNode(mergeKLists(
        []*ListNode{
            makeListNode([]int{}),
        },
    )) // 

    printListNode(mergeKLists1(
        []*ListNode{
            makeListNode([]int{1,4,5}),
            makeListNode([]int{1,3,4}),
            makeListNode([]int{2,6}),
        },
    )) // 1->1->2->3->4->4->5->6

    printListNode(mergeKLists1(
        []*ListNode{},
    )) // 
    printListNode(mergeKLists1(
        []*ListNode{
            makeListNode([]int{}),
        },
    )) //
 }