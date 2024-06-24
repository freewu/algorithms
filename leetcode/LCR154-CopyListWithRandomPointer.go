package main

// LCR 154. 复杂链表的复制
// 请实现 copyRandomList 函数，复制一个复杂链表。
// 在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

// 示例 1：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e1.png" />
// 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

// 示例 2：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e2.png" />
// 输入：head = [[1,1],[2,1]]
// 输出：[[1,1],[2,1]]

// 示例 3：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/09/e3.png" />
// 输入：head = [[3,null],[3,0],[3,null]]
// 输出：[[3,null],[3,0],[3,null]]

// 示例 4：
// 输入：head = []
// 输出：[]
// 解释：给定的链表为空（空指针），因此返回 null。

// 提示：
//     -10000 <= Node.val <= 10000
//     Node.random 为空（null）或指向链表中的节点。
//     节点数目不超过 1000 。

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
    Random *Node
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
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    oldToNew := make(map[*Node]*Node)
    curr := head
    for curr != nil {
        oldToNew[curr] = &Node{Val: curr.Val}
        curr = curr.Next
    }
    curr = head
    for curr != nil {
        oldToNew[curr].Next = oldToNew[curr.Next]
        oldToNew[curr].Random = oldToNew[curr.Random]
        curr = curr.Next
    }
    return oldToNew[head]
}

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    // 第一轮遍历
    // 注意：这里一次跳过两个节点，因为咱们在这里总是会插入一个新的节点
    for node := head; node != nil; node = node.Next.Next {
        // 创建一个新节点，
        // 1. 该节点的 Val 拷贝于当前值，Random 不初始化
        // 2. 该节点插入到当前节点之后
        node.Next = &Node{Val: node.Val, Next: node.Next}
    }
    // 第二轮遍历
    for node := head; node != nil; node = node.Next.Next {
        // 注意：有些节点的 Random 指针是没有赋值的，这里跳过处理这些节点
        if node.Random != nil {
            // 当前节点 Random 不为 nil，那么将其下一个节点的 Random 赋值为 node.Random.Next（后者是新复制出来的新节点）
            node.Next.Random = node.Random.Next
        }
    }
    // 第三轮遍历
    // 记录好要返回的新链表 head
    headNew := head.Next
    for node := head; node != nil; node = node.Next {
        // 得到新复制的节点
        nodeNew := node.Next
        // 老节点跨过新节点
        node.Next = node.Next.Next
        // 新节点跨过老节点
        if nodeNew.Next != nil {
            nodeNew.Next = nodeNew.Next.Next
        }
    }
    return headNew
}

func main() {
// Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]


// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/12/18/e2.png" />
// Input: head = [[1,1],[2,1]]
// Output: [[1,1],[2,1]]

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/12/18/e3.png" />
// Input: head = [[3,null],[3,0],[3,null]]
// Output: [[3,null],[3,0],[3,null]]
}