package main

// 987. Vertical Order Traversal of a Binary Tree
// Given the root of a binary tree, calculate the vertical order traversal of the binary tree.
// For each node at position (row, col), its left and right children will be at positions (row + 1, col - 1) and (row + 1, col + 1) respectively. 
// The root of the tree is at (0, 0).

// The vertical order traversal of a binary tree is a list of top-to-bottom orderings for each column index starting from the leftmost column and ending on the rightmost column. 
// There may be multiple nodes in the same row and same column. In such a case, sort these nodes by their values.
// Return the vertical order traversal of the binary tree.
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/29/vtree1.jpg" />
// Input: root = [3,9,20,null,null,15,7]
// Output: [[9],[3,15],[20],[7]]
// Explanation:
// Column -1: Only node 9 is in this column.
// Column 0: Nodes 3 and 15 are in this column in that order from top to bottom.
// Column 1: Only node 20 is in this column.
// Column 2: Only node 7 is in this column.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/29/vtree2.jpg" />
// Input: root = [1,2,3,4,5,6,7]
// Output: [[4],[2],[1,5,6],[3],[7]]
// Explanation:
// Column -2: Only node 4 is in this column.
// Column -1: Only node 2 is in this column.
// Column 0: Nodes 1, 5, and 6 are in this column.
//           1 is at the top, so it comes first.
//           5 and 6 are at the same position (2, 0), so we order them by their value, 5 before 6.
// Column 1: Only node 3 is in this column.
// Column 2: Only node 7 is in this column.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/01/29/vtree3.jpg" />
// Input: root = [1,2,3,4,6,5,7]
// Output: [[4],[2],[1,5,6],[3],[7]]
// Explanation:
// This case is the exact same as example 2, but with nodes 5 and 6 swapped.
// Note that the solution remains the same since 5 and 6 are in the same location and should be ordered by their values.
 
// Constraints:
// 		The number of nodes in the tree is in the range [1, 1000].
// 		0 <= Node.val <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "fmt"
import "math"
import "sort"


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 用于存储元素的坐标
type node struct {
    x, y, val int
}

func verticalTraversal(root *TreeNode) [][]int {
    var dfs func(root *TreeNode, x, y int)
    var nodes []node
    dfs = func(root *TreeNode, x, y int) {
        if root == nil {
            return
        }
        // 按照先序遍历，就可以将这些结点的二维坐标计算出来
        nodes = append(nodes, node{x, y, root.Val})
        dfs(root.Left, x+1, y-1)
        dfs(root.Right, x+1, y+1)
    }
    // 根结点是 (0，0) ，即根结点是坐标原点
    // 它的左子树的 x 坐标都是负数，它的右子树的 x 坐标都是正数
    dfs(root, 0, 0)
    // 进行一次排序，按照 x 坐标从小到大排序，坐标相同的情况对应着结点摞起来的情况
    sort.Slice(nodes, func(i, j int) bool {
        a, b := nodes[i], nodes[j]
        return a.y < b.y || a.y == b.y &&
            (a.x < b.x || a.x == b.x && a.val < b.val)
    })
    // 扫描一遍排好序的数组，按照列的顺序，依次将同一列的结点打包至一个一维数组
    var res [][]int
    lastY := math.MinInt32
    for _, node := range nodes {
        if lastY != node.y {
            res = append(res, []int{node.val})
            lastY = node.y
        } else {
            res[len(res)-1] = append(res[len(res)-1], node.val)
        }
    }
    return res
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode {
            9,
            nil,
            nil,
        },
        &TreeNode {
            30,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(verticalTraversal(tree1)) // [[9],[3,15],[20],[7]]

    tree2 := &TreeNode {
        1,
        &TreeNode {
            2,
            &TreeNode{4, nil, nil},
            &TreeNode{5, nil, nil},
        },
        &TreeNode {
            3,
            &TreeNode{6, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(verticalTraversal(tree2)) // [[4],[2],[1,5,6],[3],[7]]

    tree3 := &TreeNode {
        1,
        &TreeNode {
            2,
            &TreeNode{4, nil, nil},
            &TreeNode{6, nil, nil},
        },
        &TreeNode {
            3,
            &TreeNode{5, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(verticalTraversal(tree3)) // [[4],[2],[1,5,6],[3],[7]]
}