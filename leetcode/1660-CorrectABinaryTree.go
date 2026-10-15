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

// == 二叉树构建 & 打印输出 start ==================================================================
func ptr(v int) *int {
    return &v
}

// buildTree 按层序数组（LeetCode 格式，nil 表示 null）创建二叉树
func buildTree(nums []*int) *TreeNode {
    if len(nums) == 0 || nums[0] == nil {
        return nil
    }
    root := &TreeNode{Val: *nums[0]}
    queue := []*TreeNode{root}
    i := 1
    for len(queue) > 0 && i < len(nums) {
        node := queue[0]
        queue = queue[1:]
        if nums[i] != nil {
            node.Left = &TreeNode{Val: *nums[i]}
            queue = append(queue, node.Left)
        }
        i++
        if i < len(nums) && nums[i] != nil {
            node.Right = &TreeNode{Val: *nums[i]}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

// printTree 自顶向下打印二叉树，父子之间用 / \ 连线
// 第 r 层（根为第 1 层）有 2^(r-1) 个位置，每个节点占 block 列，
// 因此节点中心列 = (2j+1) * 2^(h-r) * block / 2（j 为该层从左往右的序号）
func printTree(root *TreeNode) {
    if root == nil {
        fmt.Println("(empty tree)")
        return
    }
    // itoa 手写整数转字符串，替代 strconv.Itoa
    itoa := func(v int) string {
        if v == 0 {
            return "0"
        }
        neg := v < 0
        u := uint(v)
        if neg {
            u = -u // 补码取绝对值，v 为最小整数时同样安全
        }
        buf := [24]byte{}
        i := len(buf)
        for u > 0 {
            i--
            buf[i] = byte('0' + u%10)
            u /= 10
        }
        if neg {
            i--
            buf[i] = '-'
        }
        return string(buf[i:])
    }
    // blankLine 生成 n 个空格的画布行（替代 strings.Repeat）
    blankLine := func(n int) []byte {
        line := make([]byte, n)
        for i := range line {
            line[i] = ' '
        }
        return line
    }
    // treeHeight 返回二叉树的层数（空树为 0，根节点为第 1 层）
    var treeHeight func(root *TreeNode) int
    treeHeight = func (root *TreeNode) int {
        if root == nil {
            return 0
        }
        l, r := treeHeight(root.Left), treeHeight(root.Right)
        if l < r {
            l = r
        }
        return l + 1
    }
    h := treeHeight(root)
    // block：每个节点占用的列宽，取最长值的宽度 + 1，并保证为偶数使中心落在整数列
    block := 2
    var scan func(node *TreeNode)
    scan = func(node *TreeNode) {
        if node == nil {
            return
        }
        if w := len(itoa(node.Val)) + 1; w > block {
            block = w
        }
        scan(node.Left)
        scan(node.Right)
    }
    scan(root)
    if block%2 != 0 {
        block++
    }

    // 画布：值占 1 行、连线占 1 行，共 2h-1 行
    width := (1 << h) * block
    rows := make([][]byte, 2*h-1)
    for i := range rows {
        rows[i] = blankLine(width)
    }
    put := func(row, col int, s string) {
        start := col - len(s)/2
        for i := 0; i < len(s); i++ {
            if p := start + i; p >= 0 && p < width {
                rows[row][p] = s[i]
            }
        }
    }
    var draw func(node *TreeNode, r, j int)
    draw = func(node *TreeNode, r, j int) {
        if node == nil {
            return
        }
        col := (2*j + 1) * (1 << (h - r)) * block / 2 // 当前节点中心列
        put(2*(r-1), col, itoa(node.Val))
        if r == h {
            return
        }
        d := (1 << (h - r - 1)) * block / 2 // 父子之间的列距
        off := (d + 1) / 2                  // 连线相对父节点中心的偏移（取行间中点）
        if node.Left != nil {
            put(2*r-1, col-off, "/")
            draw(node.Left, r+1, 2*j)
        }
        if node.Right != nil {
            put(2*r-1, col+off, "\\")
            draw(node.Right, r+1, 2*j+1)
        }
    }
    draw(root, 1, 0)

    for _, line := range rows {
        end := len(line)
        for end > 0 && line[end-1] == ' ' {
            end--
        }
        fmt.Println(string(line[:end]))
    }
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