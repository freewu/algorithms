package main

// 285. Inorder Successor in BST
// Given the root of a binary search tree and a node p in it, return the in-order successor of that node in the BST. 
// If the given node has no in-order successor in the tree, return null.
// The successor of a node p is the node with the smallest key greater than p.val.

// Example 1:
//         2
//       /   \
//      1     3
// <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_1.PNG" />
// Input: root = [2,1,3], p = 1
// Output: 2
// Explanation: 1's in-order successor node is 2. Note that both p and the return value is of TreeNode type.

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
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5
//     All Nodes will have unique values.

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
    st := []*TreeNode{}
    var pre, cur *TreeNode = nil, root
    for len(st) > 0 || cur != nil {
        for cur != nil {
            st = append(st, cur)
            cur = cur.Left
        }
        cur = st[len(st)-1]
        st = st[:len(st)-1]
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