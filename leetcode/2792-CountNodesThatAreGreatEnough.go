package main

// 2792. Count Nodes That Are Great Enough
// You are given a root to a binary tree and an integer k. 
// A node of this tree is called great enough if the followings hold:
//     1. Its subtree has at least k nodes.
//     2. Its value is greater than the value of at least k nodes in its subtree.

// Return the number of nodes in this tree that are great enough.

// The node u is in the subtree of the node v, if u == v or v is an ancestor of u.

// Example 1:
// Input: root = [7,6,5,4,3,2,1], k = 2
// Output: 3
// Explanation: Number the nodes from 1 to 7.
// The values in the subtree of node 1: {1,2,3,4,5,6,7}. Since node.val == 7, there are 6 nodes having a smaller value than its value. So it's great enough.
// The values in the subtree of node 2: {3,4,6}. Since node.val == 6, there are 2 nodes having a smaller value than its value. So it's great enough.
// The values in the subtree of node 3: {1,2,5}. Since node.val == 5, there are 2 nodes having a smaller value than its value. So it's great enough.
// It can be shown that other nodes are not great enough.
// See the picture below for a better understanding.
// <img src="https://assets.leetcode.com/uploads/2023/07/25/1.png" />
//         7
//      /     \
//     6       5
//   /   \   /   \
//  4     3 2      1  

// Example 2:
// Input: root = [1,2,3], k = 1
// Output: 0
// Explanation: Number the nodes from 1 to 3.
// The values in the subtree of node 1: {1,2,3}. Since node.val == 1, there are no nodes having a smaller value than its value. So it's not great enough.
// The values in the subtree of node 2: {2}. Since node.val == 2, there are no nodes having a smaller value than its value. So it's not great enough.
// The values in the subtree of node 3: {3}. Since node.val == 3, there are no nodes having a smaller value than its value. So it's not great enough.
// See the picture below for a better understanding.
// <img src="https://assets.leetcode.com/uploads/2023/07/25/2.png" />
//         1
//      /     \
//     2       3

// Example 3:
// Input: root = [3,2,2], k = 2
// Output: 1
// Explanation: Number the nodes from 1 to 3.
// The values in the subtree of node 1: {2,2,3}. Since node.val == 3, there are 2 nodes having a smaller value than its value. So it's great enough.
// The values in the subtree of node 2: {2}. Since node.val == 2, there are no nodes having a smaller value than its value. So it's not great enough.
// The values in the subtree of node 3: {2}. Since node.val == 2, there are no nodes having a smaller value than its value. So it's not great enough.
// See the picture below for a better understanding.
// <img src="https://assets.leetcode.com/uploads/2023/07/25/3.png" />
//         3
//      /     \
//     2       2

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     1 <= Node.val <= 10^4
//     1 <= k <= 10

import "fmt"
import "sort"

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
func countGreatEnoughNodes(root *TreeNode, k int) int {
    res, memo := 0, make(map[*TreeNode][]int)
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil { return }
        if _, ok := memo[root]; ok { return }
        dfs(root.Left)
        dfs(root.Right)
        memo[root] = append(memo[root], memo[root.Left]...)
        memo[root] = append(memo[root], memo[root.Right]...)
        memo[root] = append(memo[root], root.Val)
        sort.Ints(memo[root])
        if len(memo[root]) >= k {
            memo[root] = memo[root][:k]
            id := sort.SearchInts(memo[root], root.Val)
            if id == k {
                res++
            }
        }
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    // Input: root = [7,6,5,4,3,2,1], k = 2
    // Output: 3
    // Explanation: Number the nodes from 1 to 7.
    // The values in the subtree of node 1: {1,2,3,4,5,6,7}. Since node.val == 7, there are 6 nodes having a smaller value than its value. So it's great enough.
    // The values in the subtree of node 2: {3,4,6}. Since node.val == 6, there are 2 nodes having a smaller value than its value. So it's great enough.
    // The values in the subtree of node 3: {1,2,5}. Since node.val == 5, there are 2 nodes having a smaller value than its value. So it's great enough.
    // It can be shown that other nodes are not great enough.
    // See the picture below for a better understanding.
    // <img src="https://assets.leetcode.com/uploads/2023/07/25/1.png" />
    //         7
    //      /     \
    //     6       5
    //   /   \   /   \
    //  4     3 2      1  
    tree1 := &TreeNode {
        7,
        &TreeNode { 6 , &TreeNode { 4 , nil , nil, }, &TreeNode { 3 , nil , nil, }, },
        &TreeNode { 5 , &TreeNode { 2 , nil , nil, }, &TreeNode { 1 , nil , nil, }, },
    }
    fmt.Println(countGreatEnoughNodes(tree1, 2)) // 3
    // Example 2:
    // Input: root = [1,2,3], k = 1
    // Output: 0
    // Explanation: Number the nodes from 1 to 3.
    // The values in the subtree of node 1: {1,2,3}. Since node.val == 1, there are no nodes having a smaller value than its value. So it's not great enough.
    // The values in the subtree of node 2: {2}. Since node.val == 2, there are no nodes having a smaller value than its value. So it's not great enough.
    // The values in the subtree of node 3: {3}. Since node.val == 3, there are no nodes having a smaller value than its value. So it's not great enough.
    // See the picture below for a better understanding.
    // <img src="https://assets.leetcode.com/uploads/2023/07/25/2.png" />
    //         1
    //      /     \
    //     2       3
    tree2 := &TreeNode {
        1,
        &TreeNode { 2 , nil , nil, },
        &TreeNode { 3 , nil , nil, },
    }
    fmt.Println(countGreatEnoughNodes(tree2, 1)) // 0
    // Example 3:
    // Input: root = [3,2,2], k = 2
    // Output: 1
    // Explanation: Number the nodes from 1 to 3.
    // The values in the subtree of node 1: {2,2,3}. Since node.val == 3, there are 2 nodes having a smaller value than its value. So it's great enough.
    // The values in the subtree of node 2: {2}. Since node.val == 2, there are no nodes having a smaller value than its value. So it's not great enough.
    // The values in the subtree of node 3: {2}. Since node.val == 2, there are no nodes having a smaller value than its value. So it's not great enough.
    // See the picture below for a better understanding.
    // <img src="https://assets.leetcode.com/uploads/2023/07/25/3.png" />
    //         3
    //      /     \
    //     2       2
    tree3 := &TreeNode {
        3,
        &TreeNode { 2 , nil , nil, },
        &TreeNode { 2 , nil , nil, },
    }
    fmt.Println(countGreatEnoughNodes(tree3, 2)) // 1
}