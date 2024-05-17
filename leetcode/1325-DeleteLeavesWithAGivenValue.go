package main

// 1325. Delete Leaves With a Given Value
// Given a binary tree root and an integer target, delete all the leaf nodes with value target.
// Note that once you delete a leaf node with value target, 
// if its parent node becomes a leaf node and has the value target, 
// it should also be deleted (you need to continue doing that until you cannot).

// Example 1:
//        1                   1           1
//      /    \               /  \          \
//     2      3      =>    (2)   3   =>     3
//    /     /   \                 \          \
//  (2)    (2)  (4)                4          4
// <img src="https://assets.leetcode.com/uploads/2020/01/09/sample_1_1684.png" />
// Input: root = [1,2,3,2,null,2,4], target = 2
// Output: [1,null,3,null,4]
// Explanation: Leaf nodes in green with value (target = 2) are removed (Picture in left). 
// After removing, new nodes become leaf nodes with value (target = 2) (Picture in center).

// Example 2:
//        1                 1
//      /   \              /
//     3    (3)     =>    3
//    /  \                 \ 
//  (3)   2                  2
// <img src="https://assets.leetcode.com/uploads/2020/01/09/sample_2_1684.png" />
// Input: root = [1,3,3,3,2], target = 3
// Output: [1,3,null,null,2]

// Example 3:
//        1         1         1       1
//       /         /         /
//      2   =>    2   =>   (2)  =>  
//     /         /
//    2        (2)
//   /
// (2)
// <img src="https://assets.leetcode.com/uploads/2020/01/15/sample_3_1684.png" />
// Input: root = [1,2,null,2,null,2], target = 2
// Output: [1]
// Explanation: Leaf nodes in green with value (target = 2) are removed at each step.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 3000].
//     1 <= Node.val, target <= 1000

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
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    var dfsDelete func(root, node *TreeNode, l, target int) bool
    dfsDelete = func(root, node *TreeNode, l, target int) bool{
        if node == nil { return false; }
        if node.Left == nil && node.Right == nil && node.Val == target { // 如果是叶子（node.Left == nil && node.Right == nil）且是删除目标 做删除操作
            if l == 0 { root.Left = nil; } else { root.Right = nil; }
            return true
        }
        return dfsDelete(node, node.Left, 0, target) || dfsDelete(node, node.Right, 1, target)
    }
    for {
        if dfsDelete(root, root.Left, 0, target) || dfsDelete(root, root.Right, 1, target) { 
            continue
        } else {
            break
        }
    }
    if root.Left == nil && root.Right == nil { // 处理整棵树都是 target 的情况
        if root.Val == target { return nil; }  // 如果根节点也是要删除的值
    }
    return root
}

func removeLeafNodes1(root *TreeNode, target int) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left = removeLeafNodes1(root.Left, target)
    root.Right = removeLeafNodes1(root.Right, target)
    if root.Val == target && root.Left == nil && root.Right == nil {
        return nil
    }
    return root
}

func main() {
    // Example 1:
    //        1                   1           1
    //      /    \               /  \          \
    //     2      3      =>    (2)   3   =>     3
    //    /     /   \                 \          \
    //  (2)    (2)  (4)                4          4
    // <img src="https://assets.leetcode.com/uploads/2020/01/09/sample_1_1684.png" />
    // Input: root = [1,2,3,2,null,2,4], target = 2
    // Output: [1,null,3,null,4]
    // Explanation: Leaf nodes in green with value (target = 2) are removed (Picture in left). 
    // After removing, new nodes become leaf nodes with value (target = 2) (Picture in center).
    tree1 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode { 2, nil, nil }, nil },
        &TreeNode { 3, &TreeNode { 2, nil, nil }, &TreeNode { 4, nil, nil } },
    }
    fmt.Println("tree1.Left: ",tree1.Left)
    t1 := removeLeafNodes(tree1, 2)
    fmt.Println("t1.Left: ",t1.Left)
    // Example 2:
    //        1                 1
    //      /   \              /
    //     3    (3)     =>    3
    //    /  \                 \ 
    //  (3)   2                  2
    // <img src="https://assets.leetcode.com/uploads/2020/01/09/sample_2_1684.png" />
    // Input: root = [1,3,3,3,2], target = 3
    // Output: [1,3,null,null,2]
    tree2 := &TreeNode {
        1,
        &TreeNode { 3, &TreeNode { 3, nil, nil }, &TreeNode { 3, nil, nil } },
        &TreeNode { 3, nil, nil },
        
    }
    fmt.Println("tree2.Left: ", tree2.Left)
    t2 := removeLeafNodes(tree2, 3)
    fmt.Println("t2.Left: ",t2.Left)
    // Example 3:
    //        1         1         1       1
    //       /         /         /
    //      2   =>    2   =>   (2)  =>  
    //     /         /
    //    2        (2)
    //   /
    // (2)
    // <img src="https://assets.leetcode.com/uploads/2020/01/15/sample_3_1684.png" />
    // Input: root = [1,2,null,2,null,2], target = 2
    // Output: [1]
    // Explanation: Leaf nodes in green with value (target = 2) are removed at each step.
    tree3 := &TreeNode {
        1,
        &TreeNode{ 2, &TreeNode{ 2, &TreeNode{ 2, nil, nil }, nil }, nil },
        nil,
    }
    fmt.Println("tree3.Left: ",tree3.Left)
    t3 := removeLeafNodes(tree3, 2)
    fmt.Println("t3.Left: ",t3.Left)


    tree11 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode { 2, nil, nil }, nil },
        &TreeNode { 3, &TreeNode { 2, nil, nil }, &TreeNode { 4, nil, nil } },
    }
    fmt.Println("tree11.Left: ",tree11.Left)
    t11 := removeLeafNodes(tree11, 2)
    fmt.Println("t11.Left: ",t11.Left)
    tree12 := &TreeNode {
        1,
        &TreeNode { 3, &TreeNode { 3, nil, nil }, &TreeNode { 3, nil, nil } },
        &TreeNode { 3, nil, nil },
        
    }
    fmt.Println("tree12.Left: ", tree12.Left)
    t12 := removeLeafNodes(tree12, 3)
    fmt.Println("t12.Left: ",t12.Left)
    tree13 := &TreeNode {
        1,
        &TreeNode{ 2, &TreeNode{ 2, &TreeNode{ 2, nil, nil }, nil }, nil },
        nil,
    }
    fmt.Println("tree13.Left: ",tree13.Left)
    t13 := removeLeafNodes(tree13, 2)
    fmt.Println("t13.Left: ",t13.Left)
}