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
    tree1 := buildTree([]*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(2), ptr(0), ptr(8), nil, nil, ptr(7), ptr(4)})
    printTree(tree1)
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