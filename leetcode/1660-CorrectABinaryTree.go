package main

// 1660. Correct a Binary Tree
// You have a binary tree with a small defect. 
// There is exactly one invalid node where its right child incorrectly points to another node at the same depth but to the invalid node's right.

// Given the root of the binary tree with this defect, root, 
// return the root of the binary tree after removing this invalid node and every node underneath it (minus the node it incorrectly points to).

// Custom testing:

// The test input is read as 3 lines:
//     TreeNode root
//     int fromNode (not available to correctBinaryTree)
//     int toNode (not available to correctBinaryTree)

// After the binary tree rooted at root is parsed, 
// the TreeNode with value of fromNode will have its right child pointer pointing to the TreeNode with a value of toNode. 
// Then, root is passed to correctBinaryTree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/22/ex1v2.png" />
// Input: root = [1,2,3], fromNode = 2, toNode = 3
// Output: [1,null,3]
// Explanation: The node with value 2 is invalid, so remove it.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/22/ex2v3.png" />
// Input: root = [8,3,1,7,null,9,4,2,null,null,null,5,6], fromNode = 7, toNode = 4
// Output: [8,3,1,null,null,9,4,null,null,5,6]
// Explanation: The node with value 7 is invalid, so remove it and the node underneath it, node 2.

// Constraints:
//     The number of nodes in the tree is in the range [3, 10^4].
//     -10^9 <= Node.val <= 10^9
//     All Node.val are unique.
//     fromNode != toNode
//     fromNode and toNode will exist in the tree and will be on the same depth.
//     toNode is to the right of fromNode.
//     fromNode.right is null in the initial tree from the test data.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
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
// public:
//     TreeNode* correctBinaryTree(TreeNode* root) {
//         using ptt = pair<TreeNode*, TreeNode*>;
//         vector<ptt> q;
//         q.emplace_back(root, nullptr);
//         unordered_set<TreeNode*> s;
//         while(q.size()) {
//             vector<ptt> tmp;
//             s.clear();
//             for(auto [node, _] : q) 
//                 s.insert(node);
//             for(auto [node, pa] : q) {
//                 if(s.contains(node->right)) {
//                     if(node == pa->left)
//                         pa->left = nullptr;
//                     else
//                         pa->right = nullptr;
//                     return root;
//                 }
//                 if(node->left)
//                     tmp.emplace_back(node->left, node);
//                 if(node->right)
//                     tmp.emplace_back(node->right, node);
//             }
//             q = std::move(tmp);
//         }
//         return nullptr;
//     }
// };

// class Solution {
// public:
//     TreeNode* correctBinaryTree(TreeNode* root) {
//         // Queue for BFS. Every element stores [node, parent]
//         queue<pair<TreeNode*, TreeNode*>> q;
//         // node, parent
//         q.push({root, nullptr});
//         // Traverse Level by Level
//         while (!q.empty()) {
//             // Nodes in the current level
//             int n = q.size();
//             // Hash Set to store nodes of the current level
//             unordered_set<TreeNode*> visited;
//             // Traverse all nodes in the current level
//             for (int i = 0; i < n; i++) {
//                 // Pop the node and its parent from the queue
//                 auto [node, parent] = q.front();
//                 q.pop();   
//                 // If node.right is already visited, then the node is defective
//                 if (visited.count(node->right)) {
//                     // Replace the child of the node's parent with null and return the root
//                     if (parent->left == node) {
//                         parent->left = nullptr;
//                     } else {
//                         parent->right = nullptr;
//                     }
//                     return root;
//                 }
//                 // Add node to visited
//                 visited.insert(node);
//                 // Add child in queue for traversal in next level
//                 // They won't get popped in this level because of "n"
//                 // Add the right child first, so that we can explore right to left
//                 if (node->right) {
//                     q.push({node->right, node});
//                 }
//                 if (node->left) {
//                     q.push({node->left, node});
//                 }
//             }
//         }
//         // For the sake of compilation
//         return root;  
//     }
// };
    

func main() {
    // Example 1:
    // Input: root = [1,2,3], fromNode = 2, toNode = 3
    // Output: [1,null,3]
    // Explanation: The node with value 2 is invalid, so remove it.

    // Example 2:
    // Input: root = [8,3,1,7,null,9,4,2,null,null,null,5,6], fromNode = 7, toNode = 4
    // Output: [8,3,1,null,null,9,4,null,null,5,6]
    // Explanation: The node with value 7 is invalid, so remove it and the node underneath it, node 2.
    fmt.Println()
}