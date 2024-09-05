package main

// LCR 053. 二叉搜索树中的中序后继
// 给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
// 节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。

// 示例 1：
//         2
//       /   \
//      1     3
// <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_1.PNG" />
// 输入：root = [2,1,3], p = 1
// 输出：2
// 解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。

// 示例 2：
//         5
//       /   \
//      3     6
//    /   \
//   2    4
//  /
// 1
// <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_2.PNG" />
// 输入：root = [5,3,6,2,4,null,null,1], p = 6
// 输出：null
// 解释：因为给出的节点没有中序后继，所以答案就返回 null 了。

// 提示：
//     树中节点的数目在范围 [1, 10^4] 内。
//     -10^5 <= Node.val <= 10^5
//     树中各节点的值均保证唯一。
import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历
// 如果节点 p 是最后被访问的节点，则不存在节点 p 的中序后继，返回 nil
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    stack := []*TreeNode{}
    pre, cur := &TreeNode{}, root
    for len(stack) > 0 || cur != nil {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        cur = stack[len(stack)-1] // pop
        stack = stack[:len(stack)-1]
        if pre == p {
            return cur
        }
        pre = cur
        cur = cur.Right
    }
    return nil
}

// 二叉搜索树的一个性质是中序遍历序列单调递增，因此二叉搜索树中的节点 ppp 的中序后继满足以下条件：
//     中序后继的节点值大于 p 的节点值；
//     中序后继是节点值大于 p 的节点值的所有节点中节点值最小的一个节点。
// 利用二叉搜索树的性质，可以在不做中序遍历的情况下找到节点 ppp 的中序后继。
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
    var successor *TreeNode
    if p.Right != nil {
        successor = p.Right
        for successor.Left != nil {
            successor = successor.Left
        }
        return successor
    }
    node := root
    for node != nil {
        if node.Val > p.Val {
            successor = node
            node = node.Left
        } else {
            node = node.Right
        }
    }
    return successor
}

func main() {
    // Example 1:
    //         2
    //       /   \
    //      1     3
    // <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_1.PNG" />
    // Input: root = [2,1,3], p = 1
    // Output: 2
    // Explanation: 1's in-order successor node is 2. Note that both p and the return value is of TreeNode type.
    p1 := &TreeNode{1, nil, nil}
    tree1 := &TreeNode{
        2, 
        p1,
        &TreeNode{3, nil, nil},
    }
    t1 := inorderSuccessor(tree1, p1)
    fmt.Println("t1: ", t1) 
    fmt.Println("t1.Val: ", t1.Val) // 2 
    // Example 2:
    //           5
    //         /   \
    //        3     6
    //      /   \
    //     2     4
    //    /
    //   1
    // <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_2.PNG" />
    // Input: root = [5,3,6,2,4,null,null,1], p = 6
    // Output: null
    // Explanation: There is no in-order successor of the current node, so the answer is null.
    p2 := &TreeNode{6, nil, nil}
    tree2 := &TreeNode{
        5, 
        &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}, },
        p2,
    }
    t2 := inorderSuccessor(tree2, p2)
    fmt.Println("t2: ",t2)

    p11 := &TreeNode{1, nil, nil}
    tree11 := &TreeNode{
        2, 
        p11,
        &TreeNode{3, nil, nil},
    }
    t11 := inorderSuccessor1(tree11, p11)
    fmt.Println("t11: ", t11) 
    fmt.Println("t11.Val: ", t11.Val) // 2 

    p12 := &TreeNode{6, nil, nil}
    tree12 := &TreeNode{
        5, 
        &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}, },
        p12,
    }
    t12 := inorderSuccessor1(tree12, p12)
    fmt.Println("t12: ",t12) 
}