package main

// 1644. Lowest Common Ancestor of a Binary Tree II
// Given the root of a binary tree, return the lowest common ancestor (LCA) of two given nodes, p and q. 
// If either node p or q does not exist in the tree, return null. All values of the nodes in the tree are unique.

// According to the definition of LCA on Wikipedia: 
// "The lowest common ancestor of two nodes p and q in a binary tree T is the lowest node that has both p and q as descendants (where we allow a node to be a descendant of itself)". 
// A descendant of a node x is a node y that is on the path from node x to some leaf node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// Output: 3
// Explanation: The LCA of nodes 5 and 1 is 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The LCA of nodes 5 and 4 is 5. A node can be a descendant of itself according to the definition of LCA.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 10
// Output: null
// Explanation: Node 10 does not exist in the tree, so return null.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^9 <= Node.val <= 10^9
//     All Node.val are unique.
//     p != q

// Follow up: Can you find the LCA traversing the tree, without checking nodes existence?

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    hasq, hasp := false, false
    var lca func(node *TreeNode) *TreeNode
    lca = func(node *TreeNode) *TreeNode {
        if node == nil { return node }
        // if node == p { hasp = true }
        if node.Val == p.Val { hasp = true }
        // if node == q { hasq = true }
        if node.Val == q.Val { hasq = true }
        res1, res2 := lca(node.Left), lca(node.Right)
        //if res1 != nil && res2 != nil || node == p || node == q {
        if res1 != nil && res2 != nil || node.Val == p.Val || node.Val == q.Val {
            return node
        }
        if res1 != nil { return res1 }
        return res2
    }
    res := lca(root)
    if hasq && hasp { return res }
    return nil
}

// /**
//  * Definition for a binary tree node.
//  * struct TreeNode {
//  *     int val;
//  *     TreeNode *left;
//  *     TreeNode *right;
//  *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
//  * };
//  */
// class Solution {
// public:
//     TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
//         bool hasq = false, hasp = false;
//         auto lca = [&](auto&& lca, TreeNode* node)->TreeNode* {
//             if(node == nullptr)
//                 return node;
//             if(node == p) 
//                 hasp = true;
//             if(node == q) 
//                 hasq = true;
//             auto res1 = lca(lca, node->left);
//             auto res2 = lca(lca, node->right);
//             if(res1 && res2 || node == p || node == q)
//                 return node;
//             return res1 ? res1 : res2;
//         };
//         auto res = lca(lca, root);
//         return (hasq && hasp) ? res : nullptr;
//     }
// };

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
    // Output: 3
    // Explanation: The LCA of nodes 5 and 1 is 3.
    tree1 := &TreeNode{
        3, 
        &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}, }, },
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}, },
    }
    p1 := &TreeNode{5, nil, nil, }
    q1 := &TreeNode{1, nil, nil, }
    fmt.Println(lowestCommonAncestor(tree1, p1, q1)) // &{3 0xc000114048 0xc0001140c0}
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
    // Output: 5
    // Explanation: The LCA of nodes 5 and 4 is 5. A node can be a descendant of itself according to the definition of LCA.
    p2 := &TreeNode{5, nil, nil, }
    q2 := &TreeNode{4, nil, nil, }
    fmt.Println(lowestCommonAncestor(tree1, p2, q2)) // &{5 0xc000114060 0xc000114078}
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 10
    // Output: null
    // Explanation: Node 10 does not exist in the tree, so return null.
    p3 := &TreeNode{5, nil, nil, }
    q3 := &TreeNode{10, nil, nil, }
    fmt.Println(lowestCommonAncestor(tree1, p3, q3)) // <nil>
}