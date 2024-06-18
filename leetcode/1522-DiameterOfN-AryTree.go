package main

// 1522. Diameter of N-Ary Tree
// Given a root of an N-ary tree, you need to compute the length of the diameter of the tree.

// The diameter of an N-ary tree is the length of the longest path between any two nodes in the tree. 
// This path may or may not pass through the root.

// (Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value.)

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/19/sample_2_1897.png" />
// Input: root = [1,null,3,2,4,null,5,6]
// Output: 3
// Explanation: Diameter is shown in red color.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/07/19/sample_1_1897.png" />
// Input: root = [1,null,2,null,3,4,null,5,null,6]
// Output: 4

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/07/19/sample_3_1897.png" />
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: 7

// Constraints:
//     The depth of the n-ary tree is less than or equal to 1000.
//     The total number of nodes is between [1, 10^4].

import "fmt"

type Node struct {
    Val int
    Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func diameter(root *Node) int {
    //从经过一个节点的所有路径长度
    // 对每个节点取出发自该节点的最长路径和次长路径加和，并返回最长路径给上层
    res := 0
    var dfs func(node *Node) int
    dfs = func(node *Node) int {
        if node == nil {
            return -1
        }
        maxL, secondL := -1, -1
        for _, c := range node.Children {
            p := dfs(c)
            if p > maxL {
                secondL, maxL = maxL, p
            } else if p > secondL {
                secondL = p
            }
        }
        res = max(res, maxL + secondL + 2)
        return maxL + 1 
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/19/sample_2_1897.png" />
    // Input: root = [1,null,3,2,4,null,5,6]
    // Output: 3
    // Explanation: Diameter is shown in red color.

    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/07/19/sample_1_1897.png" />
    // Input: root = [1,null,2,null,3,4,null,5,null,6]
    // Output: 4

    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/07/19/sample_3_1897.png" />
    // Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    // Output: 7
    fmt.Println()
}