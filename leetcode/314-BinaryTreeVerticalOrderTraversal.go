package main

// 314. Binary Tree Vertical Order Traversal
// Given the root of a binary tree, return the vertical order traversal of its nodes' values. (i.e., from top to bottom, column by column).
// If two nodes are in the same row and column, the order should be from left to right.

// Example 1:
//         3
//       /   \
//      9    20
//          /  \
//         15   7
// <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree1.jpg" />
// Input: root = [3,9,20,null,null,15,7]
// Output: [[9],[3,15],[20],[7]]

// Example 2:
//         3
//       /   \
//      9     8
//     /  \  /  \
//    4    0 1   7
// <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree2-1.jpg" />
// Input: root = [3,9,8,4,0,1,7]
// Output: [[4],[9],[3,0,1],[8],[7]]

// Example 3:
//         3
//       /   \
//      9     8
//    /  \  /  \
//   4    0 1   7
//       /   \
//      5     2
// <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree2.jpg" />
// Input: root = [3,9,8,4,0,1,7,null,null,null,2,5]
// Output: [[4],[9,5],[3,0,1],[8,2],[7]]
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 100].
//     -100 <= Node.val <= 100

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalOrder(root *TreeNode) [][]int {
    // 想象在遍历过程中，路径会左右拐，从题意来看，左拐和右拐次数相抵消后剩余次数相同的节点为同一组；基于此我们可以在遍历过程中记录每个节点走过的路径和作为列号标号column，往左-1，往右+1
    // 另同一列的输出顺序为从上到下，那么先序/中序遍历不行，因为有可能出现左子树的某个后辈节点与右节点在同一列的情况，那样的话上下遍历顺序会不对；所以需要用层序遍历来保证上下顺序；
    // 不同列的顺序从左到右，正好可以利用column来排序;
    // 再观察一下可以发现一棵树每一列和列之间是连续的，不会出现空列，所以遍历过程中只需要记录最大最小column即可   
    res := [][]int{}
    if root == nil { return res } // 节点树为空，直接返回
    minCol, maxCol := 0,0 // 以根节点所在列为0列，则最大列不小于0，最小列不大于0,都给0即可
    groups := map[int][]int{} // 遍历的过程中可以根据列号分组

    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }

    // 为便于记录当前节点处于第几列，这里定义一个结构体用来存储层节点
    type Node struct {
        node *TreeNode
        col int
    }
    // 1. 层序遍历
    layer := []Node{ Node{node: root, col: 0} } // 第一层节点只有一个root，处在0列
    for len(layer) > 0 {
        nextLayer := []Node{}
        for _,n := range layer {
            maxCol, minCol = max(maxCol, n.col), min(minCol, n.col)
            groups[n.col] = append(groups[n.col], n.node.Val)
            if n.node.Left != nil {
                nextLayer = append(nextLayer, Node{ node: n.node.Left, col: n.col - 1} ) // 左 -1
            }
            if n.node.Right != nil {
                nextLayer = append(nextLayer, Node{ node: n.node.Right, col: n.col + 1}) // 右 + 1
            }
        }
        layer = nextLayer
    }
    for col := minCol; col <= maxCol; col++ { /// 输出
        res = append(res, groups[col])
    }
    return res
}

func main() {
    // Example 1:
    //         3
    //       /   \
    //      9    20
    //          /  \
    //         15   7
    // <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree1.jpg" />
    // Input: root = [3,9,20,null,null,15,7]
    // Output: [[9],[3,15],[20],[7]]
    tree1 := &TreeNode{
        3, 
        &TreeNode{9, nil, nil, },
        &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}, },
    }
    fmt.Println(verticalOrder(tree1)) // [[9],[3,15],[20],[7]]
    // Example 2:
    //         3
    //       /   \
    //      9     8
    //     /  \  /  \
    //    4    0 1   7
    // <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree2-1.jpg" />
    // Input: root = [3,9,8,4,0,1,7]
    // Output: [[4],[9],[3,0,1],[8],[7]]
    tree2 := &TreeNode{
        3, 
        &TreeNode{9, &TreeNode{4, nil, nil}, &TreeNode{0, nil, nil}, },
        &TreeNode{8, &TreeNode{1, nil, nil}, &TreeNode{7, nil, nil}, },
    }
    fmt.Println(verticalOrder(tree2)) // [[4],[9],[3,0,1],[8],[7]]
    // Example 3:
    //         3
    //       /   \
    //      9     8
    //    /  \  /  \
    //   4    0 1   7
    //       /   \
    //      5     2
    // <img src="https://assets.leetcode.com/uploads/2021/01/28/vtree2.jpg" />
    // Input: root = [3,9,8,4,0,1,7,null,null,null,2,5]
    // Output: [[4],[9,5],[3,0,1],[8,2],[7]]
    tree3 := &TreeNode{
        3, 
        &TreeNode{9, &TreeNode{4, nil, nil}, &TreeNode{0, &TreeNode{5, nil, nil}, nil}, },
        &TreeNode{8, &TreeNode{1, nil, &TreeNode{2, nil, nil}, }, &TreeNode{7, nil, nil}, },
    }
    fmt.Println(verticalOrder(tree3)) // [[4],[9,5],[3,0,1],[8,2],[7]]
}