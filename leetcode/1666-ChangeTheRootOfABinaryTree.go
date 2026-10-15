package main

// 1666. Change the Root of a Binary Tree
// Given the root of a binary tree and a leaf node, reroot the tree so that the leaf is the new root.

// You can reroot the tree with the following steps for each node cur on the path starting from the leaf up to the root​​​ excluding the root:
//     1. If cur has a left child, then that child becomes cur's right child.
//     2. cur's original parent becomes cur's left child. Note that in this process the original parent's pointer to cur becomes null, making it have at most one child.

// Return the new root of the rerooted tree.

// Note: Ensure that your solution sets the Node.parent pointers correctly after rerooting or you will receive "Wrong Answer".

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/09/21/bt_image_1.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], leaf = 7
// Output: [7,2,null,5,4,3,6,null,null,null,1,null,null,0,8]

// Example 2:
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], leaf = 0
// Output: [0,1,null,3,8,5,null,null,null,6,2,null,null,7,4]
 
// Constraints:
//     The number of nodes in the tree is in the range [2, 100].
//     -10^9 <= Node.val <= 10^9
//     All Node.val are unique.
//     leaf exist in the tree.

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

/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* parent;
};
*/
// class Solution {
// public:
//     Node* flipBinaryTree(Node* root, Node * leaf) {
//         flip(leaf);
//         return leaf;
//     }
//     void flip(Node* cur)
//     {
//         if(!cur) return;
//         Node* parent = cur->parent;
//         if(!parent) return;
//         if(parent->left == cur)
//             parent->left = nullptr;
//         else if(parent->right == cur)
//             parent->right = nullptr;
//         cur->parent = nullptr;
//         flip(parent);
//         if(cur->left!=nullptr)
//             cur->right = cur->left;
//         cur->left = parent;
//         parent->parent = cur;
//     }
// };

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/09/21/bt_image_1.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], leaf = 7
    // Output: [7,2,null,5,4,3,6,null,null,null,1,null,null,0,8]

    // Example 2:
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], leaf = 0
    // Output: [0,1,null,3,8,5,null,null,null,6,2,null,null,7,4]
    fmt.Println()
}