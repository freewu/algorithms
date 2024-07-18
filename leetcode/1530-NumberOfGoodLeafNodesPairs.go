package main

// 1530. Number of Good Leaf Nodes Pairs
// You are given the root of a binary tree and an integer distance. 
// A pair of two different leaf nodes of a binary tree is said to be good if the length of the shortest path between them is less than or equal to distance.

// Return the number of good leaf node pairs in the tree.

// Example 1:
//         1
//       /   \
//      2     3
//       \ 
//         4
// <img src="https://assets.leetcode.com/uploads/2020/07/09/e1.jpg" />
// Input: root = [1,2,3,null,4], distance = 3
// Output: 1
// Explanation: The leaf nodes of the tree are 3 and 4 and the length of the shortest path between them is 3. This is the only good pair.

// Example 2:
//         1
//       /   \
//      2     3
//    /  \   /  \
//   4    5 6    7
// <img src="https://assets.leetcode.com/uploads/2020/07/09/e2.jpg" />
// Input: root = [1,2,3,4,5,6,7], distance = 3
// Output: 2
// Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. The pair [4,6] is not good because the length of ther shortest path between them is 4.

// Example 3:
// Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
// Output: 1
// Explanation: The only good pair is [2,5].

// Constraints:
//     The number of nodes in the tree is in the range [1, 2^10].
//     1 <= Node.val <= 100
//     1 <= distance <= 10

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
func countPairs(root *TreeNode, distance int) int {
    res := 0
    var dfs func(*TreeNode) []int
    dfs = func(node *TreeNode) []int {
        if node == nil { return []int{} }
        if node.Left == nil && node.Right == nil { return []int{0} }
        arr := make([]int, 0)
        left := dfs(node.Left)
        for _, l := range left {
            if l++; l <= distance {
                arr = append(arr, l)
            }
        }
        right := dfs(node.Right)
        for _, r := range right {
            if r++; r <= distance {
                arr = append(arr, r)
            }
        }
        for _, l := range left {
            for _, r := range right {
                if l+r+2 <= distance { // a与left b与right+2
                    res++
                }
            }
        }
        return arr
    }
    dfs(root)
    return res
}

func countPairs1(root *TreeNode, distance int) int {
    var dfs func(*TreeNode) ([]int, int)
    dfs = func (node *TreeNode) ([]int, int) {
        if node == nil {
            return []int{}, 0
        }
        if node.Left == nil && node.Right == nil {
            return []int{1}, 0
        }
        lLeaves, lCount := dfs(node.Left)
        rLeaves, rCount := dfs(node.Right)
        count := lCount + rCount
        for _, left := range lLeaves {
            for _, right := range rLeaves {
                if left + right <= distance {
                    count++
                }
            }
        }
        leaves := make([]int, 0, len(lLeaves) + len(rLeaves))
        for _, leaf := range lLeaves {
            if leaf < 9 {
                leaves = append(leaves, leaf + 1)
            }
        }
        for _, leaf := range rLeaves {
            if leaf < 9 {
                leaves = append(leaves, leaf + 1)
            }
        }
        return leaves, count
    }
    _, res := dfs(root)
    return res
}
 
func main() {
    // Example 1:
    //         1
    //       /   \
    //      2     3
    //       \ 
    //         4
    // <img src="https://assets.leetcode.com/uploads/2020/07/09/e1.jpg" />
    // Input: root = [1,2,3,null,4], distance = 3
    // Output: 1
    // Explanation: The leaf nodes of the tree are 3 and 4 and the length of the shortest path between them is 3. This is the only good pair.
    tree1 := &TreeNode {
        1,
        &TreeNode { 2, nil, &TreeNode { 4, nil, nil }, },
        &TreeNode { 3, nil, nil },
    }
    fmt.Println(countPairs(tree1, 3)) // 1
    // Example 2:
    //         1
    //       /   \
    //      2     3
    //    /  \   /  \
    //   4    5 6    7
    // <img src="https://assets.leetcode.com/uploads/2020/07/09/e2.jpg" />
    // Input: root = [1,2,3,4,5,6,7], distance = 3
    // Output: 2
    // Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. The pair [4,6] is not good because the length of ther shortest path between them is 4.
    tree2 := &TreeNode{
        1,
        &TreeNode{ 2, &TreeNode{ 4, nil, nil }, &TreeNode{ 5, nil, nil }, },
        &TreeNode{ 3, &TreeNode{ 6, nil, nil }, &TreeNode{ 7, nil, nil }, },
    }
    fmt.Println(countPairs(tree2, 3)) // 2
    // Example 3:
    // Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
    // Output: 1
    // Explanation: The only good pair is [2,5].
    tree3 := &TreeNode{
        7,
        &TreeNode{ 1, &TreeNode{ 6, nil, nil }, nil, },
        &TreeNode{ 4, &TreeNode{ 5, nil, nil }, &TreeNode{ 3, nil, &TreeNode{ 2, nil, nil }, }, },
    }
    fmt.Println(countPairs(tree3, 3)) // 1


    tree11 := &TreeNode {
        1,
        &TreeNode { 2, nil, &TreeNode { 4, nil, nil }, },
        &TreeNode { 3, nil, nil },
    }
    fmt.Println(countPairs1(tree11, 3)) // 1
    tree12 := &TreeNode{
        1,
        &TreeNode{ 2, &TreeNode{ 4, nil, nil }, &TreeNode{ 5, nil, nil }, },
        &TreeNode{ 3, &TreeNode{ 6, nil, nil }, &TreeNode{ 7, nil, nil }, },
    }
    fmt.Println(countPairs1(tree12, 3)) // 2
    // Example 3:
    // Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
    // Output: 1
    // Explanation: The only good pair is [2,5].
    tree13 := &TreeNode{
        7,
        &TreeNode{ 1, &TreeNode{ 6, nil, nil }, nil, },
        &TreeNode{ 4, &TreeNode{ 5, nil, nil }, &TreeNode{ 3, nil, &TreeNode{ 2, nil, nil }, }, },
    }
    fmt.Println(countPairs1(tree13, 3)) // 1
}