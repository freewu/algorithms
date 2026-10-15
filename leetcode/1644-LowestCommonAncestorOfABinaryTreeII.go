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
    tree1 := buildTree([]*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(2), ptr(0), ptr(8), nil, nil, ptr(7), ptr(4)})
    printTree(tree1)
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