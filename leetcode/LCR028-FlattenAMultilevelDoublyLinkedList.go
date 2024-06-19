package main

// LCR 028. 扁平化多级双向链表
// 多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
// 这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

// 给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。

// 示例 1：
// 输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
// 输出：[1,2,3,7,8,11,12,9,10,4,5,6]
// 解释：
// 输入的多级列表如下图所示：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/multilevellinkedlist.png" />
// 扁平化后的链表如下图：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/multilevellinkedlistflattened.png" />

// 示例 2：
// 输入：head = [1,2,null,3]
// 输出：[1,3,2]
// 解释：
// 输入的多级列表如下图所示：
//   1---2---NULL
//   |
//   3---NULL

// 示例 3：
// 输入：head = []
// 输出：[]
 
// 如何表示测试用例中的多级链表？
// 以 示例 1 为例：

//  1---2---3---4---5---6--NULL
//          |
//          7---8---9---10--NULL
//              |
//              11--12--NULL

// 序列化其中的每一级之后：

// [1,2,3,4,5,6,null]
// [7,8,9,10,null]
// [11,12,null]
// 为了将每一级都序列化到一起，我们需要每一级中添加值为 null 的元素，以表示没有节点连接到上一级的上级节点。

// [1,2,3,4,5,6,null]
// [null,null,7,8,9,10,null]
// [null,11,12,null]
// 合并所有序列化结果，并去除末尾的 null 。

// [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
 
// 提示：
//     节点数目不超过 1000
//     1 <= Node.val <= 10^5

import "fmt"

type Node struct {
    Val int
    Prev *Node
    Next *Node
    Child *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */
// stack
func flatten(root *Node) *Node {
    stack, head := []*Node{}, root
    for root != nil || len(stack) != 0 {
        if root != nil && root.Child != nil {
            if root.Next != nil {
                stack = append(stack, root.Next)
            }
            child := root.Child
            root.Child = nil
            root.Next = child
            child.Prev = root
            root = child
            continue
        }
        if root != nil && root.Next != nil {
            root = root.Next
            continue
        }
        if len(stack) != 0 {
            node := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            root.Next = node
            node.Prev = root
            root = node
            continue
        }
        break
    }
    return head
}

// dfs
func flatten1(root *Node) *Node {
    var dfs func(*Node) *Node
    dfs = func(node *Node) (last *Node) {
        cur := node
        for cur != nil {
            next := cur.Next
            if cur.Child != nil {
                childLast := dfs(cur.Child)
                next = cur.Next
                cur.Next = cur.Child
                cur.Child.Prev = cur
                if next != nil {
                    childLast.Next = next
                    next.Prev = childLast
                }
                cur.Child = nil
                last = childLast
            } else {
                last = cur
            }
            cur = next
        }
        return
    }
    dfs(root)
    return root
}

func main() {
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten11.jpg" />
// Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
// Output: [1,2,3,7,8,11,12,9,10,4,5,6]
// Explanation: The multilevel linked list in the input is shown.
// After flattening the multilevel linked list it becomes:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten12.jpg" />

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten2.1jpg" />
// Input: head = [1,2,null,3]
// Output: [1,3,2]
// Explanation: The multilevel linked list in the input is shown.
// After flattening the multilevel linked list it becomes:
// <img src="https://assets.leetcode.com/uploads/2021/11/24/list.jpg" />

// Example 3:
// Input: head = []
// Output: []
// Explanation: There could be empty list in the input.
}