package main

// 1676. Lowest Common Ancestor of a Binary Tree IV
// Given the root of a binary tree and an array of TreeNode objects nodes, return the lowest common ancestor (LCA) of all the nodes in nodes. 
// All the nodes will exist in the tree, and all values of the tree's nodes are unique.

// Extending the definition of LCA on Wikipedia: 
// "The lowest common ancestor of n nodes p1, p2, ..., pn in a binary tree T is the lowest node that has every pi as a descendant (where we allow a node to be a descendant of itself) for every valid i". 
// A descendant of a node x is a node y that is on the path from node x to some leaf node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], nodes = [4,7]
// Output: 2
// Explanation: The lowest common ancestor of nodes 4 and 7 is node 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], nodes = [1]
// Output: 1
// Explanation: The lowest common ancestor of a single node is the node itself.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], nodes = [7,6,2,4]
// Output: 5
// Explanation: The lowest common ancestor of the nodes 7, 6, 2, and 4 is node 5.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^9 <= Node.val <= 10^9
//     All Node.val are unique.
//     All nodes[i] will exist in the tree.
//     All nodes[i] are distinct.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, nodes []*TreeNode) *TreeNode {
    exists := func(node *TreeNode) bool {
        for _, v := range nodes {
            if v.Val == node.Val {
                return true
            }
        }
        return false
    }
    var dfs func(node *TreeNode) *TreeNode 
    dfs = func(node *TreeNode) *TreeNode {
        // if (root == 0) return 0;
        if node == nil { return nil }
        if exists(node) {
            return node
        }
        left, right := dfs(node.Left), dfs(node.Right)
        if left != nil && right != nil {
            return node
        }
        if left != nil { return left }
        return right
    }
    return dfs(root)
}

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
// class Solution {
// unordered_set<TreeNode*> st;
// public:
//     TreeNode* lowestCommonAncestor(TreeNode* root, vector<TreeNode*> &nodes) {
//         for (auto node: nodes) {
//             st.insert(node);
//         }
//         return traverse(root);
//     }

//     TreeNode* traverse(TreeNode* root) {
//         if (root == 0) return 0;
//         if (st.count(root)) return root;
//         TreeNode* left = traverse(root->left);
//         TreeNode* right = traverse(root->right);
//         if (left && right) return root;
//         return left ? left : right;
//     }
// };

// class Solution {
// public:
//     TreeNode *res = nullptr;
//     int traverse(TreeNode* r, unordered_set<TreeNode*> &ns) {
//         int match = r == nullptr ? 0 : ns.count(r) + traverse(r->left, ns) + traverse(r->right, ns);
//         if (match == ns.size() && res == nullptr) // res == nullptr make sure inner process update res by LCA
//             res = r;
//         return match;
//     }
//     TreeNode* lowestCommonAncestor(TreeNode* root, vector<TreeNode*> &nodes) {
//         unordered_set<TreeNode*> ns(begin(nodes), end(nodes));
//         traverse(root, ns);
//         return res;
//     }
// };

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], nodes = [4,7]
    // Output: 2
    // Explanation: The lowest common ancestor of nodes 4 and 7 is node 2.
    tree1 := &TreeNode{
        3, 
        &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}, }, },
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}, },
    }
    nodes1 := []*TreeNode{ 
        &TreeNode{4, nil, nil}, 
        &TreeNode{7, nil, nil},
    }
    fmt.Println(lowestCommonAncestor(tree1, nodes1)) // &{2 0xc0000080a8 0xc0000080c0}
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], nodes = [1]
    // Output: 1
    // Explanation: The lowest common ancestor of a single node is the node itself.
    nodes2 := []*TreeNode{ 
        &TreeNode{1, nil, nil}, 
    }
    fmt.Println(lowestCommonAncestor(tree1, nodes2)) // &{1 0xc0000080f0 0xc000008108}
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], nodes = [7,6,2,4]
    // Output: 5
    // Explanation: The lowest common ancestor of the nodes 7, 6, 2, and 4 is node 5.
    nodes3 := []*TreeNode{ 
        &TreeNode{7, nil, nil}, 
        &TreeNode{6, nil, nil}, 
        &TreeNode{2, nil, nil}, 
        &TreeNode{4, nil, nil}, 
    }
    fmt.Println(lowestCommonAncestor(tree1, nodes3)) // &{5 0xc000008078 0xc000008090}
}