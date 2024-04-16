package main

// 623. Add One Row to Tree
// Given the root of a binary tree and two integers val and depth, add a row of nodes with value val at the given depth depth.
// Note that the root node is at depth 1.

// The adding rule is:
//     Given the integer depth, for each not null tree node cur at the depth depth - 1, create two tree nodes with value val as cur's left subtree root and right subtree root.
//     cur's original left subtree should be the left subtree of the new left subtree root.
//     cur's original right subtree should be the right subtree of the new right subtree root.
//     If depth == 1 that means there is no depth depth - 1 at all, then create a tree node with value val as the new root of the whole original tree, and the original tree is the new root's left subtree.
    
// Example 1:
//         4                     4
//        /  \                  /  \
//       2    6    ==>        [1]  [1]
//     /  \  /                /      \
//    3    1 5               2        6
//                          / \      /
//                         3   1     5
// <img src="https://assets.leetcode.com/uploads/2021/03/15/addrow-tree.jpg" />
// Input: root = [4,2,6,3,1,5], val = 1, depth = 2
// Output: [4,1,1,2,null,null,6,3,1,5]

// Example 2:
//         4                     4
//        /                     /   
//       2         ==>         2     
//     /  \                   /  \     
//    3    1                [1]  [1]     
//                          /      \      
//                         3        1      
// <img src="https://assets.leetcode.com/uploads/2021/03/11/add2-tree.jpg" />
// Input: root = [4,2,null,3,1], val = 1, depth = 3
// Output: [4,2,null,1,1,3,null,null,1]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     The depth of the tree is in the range [1, 10^4].
//     -100 <= Node.val <= 100
//     -10^5 <= val <= 10^5
//     1 <= depth <= the depth of tree + 1

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
// dfs
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if root == nil {
        return nil
    }
    // 如果 depth == 1 意味着 depth - 1 根本没有深度，
    // 那么创建一个树节点，值 val 作为整个原始树的新根，而原始树就是新根的左子树。
    if depth == 1 { // rule 4
        newRoot := TreeNode{val, root, nil}
        return &newRoot
    }
    // 给定整数 depth，对于深度为 depth - 1 的每个非空树节点 cur ，
    // 创建两个值为 val 的树节点作为 cur 的左子树根和右子树根
    if depth == 2 { // rule 1 
        newLeft := TreeNode{val, root.Left, nil}
        newRight := TreeNode{val, nil, root.Right}
        root.Left = &newLeft
        root.Right = &newRight
        return root
    }
    root.Left = addOneRow(root.Left, val, depth - 1) // rule 2   cur 原来的左子树应该是新的左子树根的左子树
    root.Right = addOneRow(root.Right, val, depth - 1) // rule 3 cur 原来的右子树应该是新的右子树根的右子树。
    return root
}

// bfs
func addOneRow1(root *TreeNode, val int, depth int) *TreeNode {
    // 如果 depth == 1 意味着 depth - 1 根本没有深度，
    // 那么创建一个树节点，值 val 作为整个原始树的新根，而原始树就是新根的左子树。
    if depth == 1 {
        return &TreeNode{
            Val: val,
            Left: root,
        }
    }
    queue := []*TreeNode{root}
    currDepth := 1
    for len(queue) > 0 {
        size := len(queue)
        // 给定整数 depth，对于深度为 depth - 1 的每个非空树节点 cur ，
        // 创建两个值为 val 的树节点作为 cur 的左子树根和右子树根
        // 作为depth-1 的子节点
        if currDepth == depth - 1 {
            for i := 0; i < size; i++ {
                currNode := queue[i]
                newLeft := currNode.Left
                newRight := currNode.Right
                currNode.Left = &TreeNode{ Val: val, Left: newLeft,}
                currNode.Right = &TreeNode{ Val: val, Right: newRight,}
            }
            break
        }
        for i := 0; i < size; i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left) // rule 2   cur 原来的左子树应该是新的左子树根的左子树
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right) // rule 3 cur 原来的右子树应该是新的右子树根的右子树。
            }
        }
        // 移除整行
        queue = queue[size:]
        currDepth++
    }
    return root
}

func main() {
    tree1 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}, },
        &TreeNode {6, &TreeNode{5, nil, nil}, nil, },
    }
    fmt.Println(addOneRow(tree1,1,2)) // [4,1,1,2,null,null,6,3,1,5]
    tree2 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}, },
        nil,
    }
    fmt.Println(addOneRow(tree2,1,3)) // [4,2,null,1,1,3,null,null,1]

    tree11 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}, },
        &TreeNode {6, &TreeNode{5, nil, nil}, nil, },
    }
    fmt.Println(addOneRow1(tree11,1,2)) // [4,1,1,2,null,null,6,3,1,5]
    tree12 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}, },
        nil,
    }
    fmt.Println(addOneRow1(tree12,1,3)) // [4,2,null,1,1,3,null,null,1]
}