package main

// 333. Largest BST Subtree
// Given the root of a binary tree, find the largest subtree , 
// which is also a Binary Search Tree (BST), where the largest means subtree has the largest number of nodes.

// A Binary Search Tree (BST) is a tree in which all the nodes follow the below-mentioned properties:
//     The left subtree values are less than the value of their parent (root) node's value.
//     The right subtree values are greater than the value of their parent (root) node's value.

// Note: A subtree must include all of its descendants.

// Example 1:
//         10
//        /  \
//      (5)   15
//     /  \    \
//   (1)  (8)   7
// <img src="https://assets.leetcode.com/uploads/2020/10/17/tmp.jpg" />
// Input: root = [10,5,15,1,8,null,7]
// Output: 3
// Explanation: The Largest BST Subtree in this case is the highlighted one. The return value is the subtree's size, which is 3.

// Example 2:
// Input: root = [4,2,7,2,3,5,null,2,null,null,null,null,null,1]
// Output: 2
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -10^4 <= Node.val <= 10^4

// Follow up: Can you figure out ways to solve it with O(n) time complexity?

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
func largestBSTSubtree(root *TreeNode) int {
    // 二叉搜索树的性质：中序遍历有序，所以可以中序遍历一遍树，选出连续有序的最大子树
    // 优化：若一个子树不是二叉搜索树，则其父节点为根的树也不是，所以很多节点就不必要遍历了。
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 返回是否为搜索二叉树，节点个数,及最大最小值用于上层确认是否满足
    var inorder func(node *TreeNode) (bool,int,int,int)
    inorder = func(node *TreeNode) (bool,int,int,int) {
        if node == nil {
            return true, 0,-1,-1
        }
        ok1, cnt1, min1, max1 := inorder(node.Left)
        ok2, cnt2, min2, max2 := inorder(node.Right)
        if (ok1 && (cnt1==0 || node.Val>max1)) && (ok2 && (cnt2==0 || node.Val<min2)) {
            if cnt1 == 0 { min1 = node.Val; }
            if cnt2 == 0 { max2 = node.Val; }
            res = max(cnt1 + cnt2 + 1, res)
            return true,cnt1 + cnt2 + 1, min1, max2
        } else {
            return false, -1,-1,-1 //false 的情况下其他值都无所谓了
        }
    }
    inorder(root)
    return res
}

func main() {
    // Example 1:
    //         10
    //        /  \
    //      (5)   15
    //     /  \    \
    //   (1)  (8)   7
    // <img src="https://assets.leetcode.com/uploads/2020/10/17/tmp.jpg" />
    // Input: root = [10,5,15,1,8,null,7]
    // Output: 3
    // Explanation: The Largest BST Subtree in this case is the highlighted one. The return value is the subtree's size, which is 3.
    tree1 := &TreeNode {
        10,
        &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{8, nil, nil}, },
        &TreeNode{15, nil, &TreeNode{7, nil, nil}, },
    }
    fmt.Println(largestBSTSubtree(tree1)) // 3
    // Example 2:
    // Input: root = [4,2,7,2,3,5,null,2,null,null,null,null,null,1]
    // Output: 2
    tree2 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{2, &TreeNode{ 2, &TreeNode{1, nil, nil}, nil}, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, &TreeNode{5, nil, nil}, nil, },
    }
    fmt.Println(largestBSTSubtree(tree2)) // 2
}