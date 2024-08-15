package main

// LCR 155. 将二叉搜索树转化为排序的双向链表
// 将一个 二叉搜索树 就地转化为一个 已排序的双向循环链表 。

// 对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

// 特别地，我们希望可以 就地 完成转换操作。
// 当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。
// 还需要返回链表中最小元素的指针。

// 示例 1：
// 输入：root = [4,2,5,1,3] 
// <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdllreturndll.png" />
// 输出：[1,2,3,4,5]
// 解释：下图显示了转化后的二叉搜索树，实线表示后继关系，虚线表示前驱关系。
// <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdllreturnbst.png" />

// 示例 2：
// 输入：root = [2,1,3]
// 输出：[1,2,3]

// 示例 3：
// 输入：root = []
// 输出：[]
// 解释：输入是空树，所以输出也是空链表。

// 示例 4：
// 输入：root = [1]
// 输出：[1]

// 提示：
//     -1000 <= Node.val <= 1000
//     Node.left.val < Node.val < Node.right.val
//     Node.val 的所有值都是独一无二的
//     0 <= Number of Nodes <= 2000

import "fmt"

// Definition for a Node.
type Node struct {
    Val int
    Left *Node
    Right *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */
func treeToDoublyList(root *Node) *Node {
    //中序遍历
    // 把前后连接，每次返回前一个节点，并连接
    if root == nil {
        return nil
    }
    var inorder  func(root *Node) (left,right *Node)
    inorder = func(root *Node) (left,right *Node) {
        if root == nil {
            return nil, nil
        }
        l1,r1 := inorder(root.Left)
        l2,r2 := inorder(root.Right)
        if root.Left == nil && root.Right == nil {
            return root, root
        } else if root.Left == nil && root.Right != nil {
            root.Right = l2
            l2.Left = root
            return root, r2
        } else if root.Left != nil && root.Right == nil {
            root.Left = r1
            r1.Right = root
            return l1, root
        } else {
            root.Left = r1
            r1.Right = root
            root.Right = l2
            l2.Left = root
            return l1,  r2
        }
    }
    left, right := inorder(root)
    left.Left = right
    right.Right = left
    return left
}

// /*
// // Definition for a Node.
// class Node {
// public:
//     int val;
//     Node* left;
//     Node* right;

//     Node() {}

//     Node(int _val) {
//         val = _val;
//         left = NULL;
//         right = NULL;
//     }

//     Node(int _val, Node* _left, Node* _right) {
//         val = _val;
//         left = _left;
//         right = _right;
//     }
// };
// */

// class Solution {
// public:
//     Node* treeToDoublyList(Node* root) {
//         if(root == nullptr) {
//             return (Node*)nullptr;
//         }

//         // 中序遍历
//         stack<Node*> st;
//         Node* first = nullptr;
//         Node* pre = nullptr;

//         while(root != nullptr || !st.empty()) {
//             // 左子树全部入栈
//             while(root != nullptr) {
//                 st.push(root);
//                 root = root->left;
//             }

//             root = st.top(), st.pop();
//             if(first == nullptr) {
//                 first = root;
//             }

//             if(pre != nullptr) {
//                 pre->right = root;
//                 root->left = pre;
//             }

//             pre = root;
//             root = root->right;
//         }

//         first->left = pre;
//         pre->right = first;
//         return first;
//     }
// };

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdlloriginalbst.png" />
    // Input: root = [4,2,5,1,3]
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdllreturndll.png" />
    // Output: [1,2,3,4,5]
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdllreturnbst.png" />
    // Explanation: The figure below shows the transformed BST. The solid line indicates the successor relationship, while the dashed line means the predecessor relationship.
    tree1 := &Node {
        4,
        &Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}, },
        &Node{5, nil, nil},
    }
    t1 := treeToDoublyList(tree1)
    fmt.Println("t1 ", t1)
    fmt.Println("t1.Left ", t1.Left)
    fmt.Println("t1.Right ", t1.Right)
    // Example 2:
    // Input: root = [2,1,3]
    // Output: [1,2,3]
    tree2 := &Node {
        2,
        &Node{1, nil, nil},
        &Node{3, nil, nil},
    }
    t2 := treeToDoublyList(tree2)
    fmt.Println("t2 ", t2)
    fmt.Println("t2.Left ", t2.Left)
    fmt.Println("t2.Right ", t2.Right)
}