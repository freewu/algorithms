package main

// 1902. Depth of BST Given Insertion Order
// You are given a 0-indexed integer array order of length n, a permutation of integers from 1 to n representing the order of insertion into a binary search tree.

// A binary search tree is defined as follows:
//     The left subtree of a node contains only nodes with keys less than the node's key.
//     The right subtree of a node contains only nodes with keys greater than the node's key.
//     Both the left and right subtrees must also be binary search trees.

// The binary search tree is constructed as follows:
//     order[0] will be the root of the binary search tree.
//     All subsequent elements are inserted as the child of any existing node such that the binary search tree properties hold.

// Return the depth of the binary search tree.
// A binary tree's depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/15/1.png" />
// Input: order = [2,1,4,3]
// Output: 3
// Explanation: The binary search tree has a depth of 3 with path 2->3->4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/15/2.png" />
// Input: order = [2,1,3,4]
// Output: 3
// Explanation: The binary search tree has a depth of 3 with path 2->3->4.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/15/3.png" />
// Input: order = [1,2,3,4]
// Output: 4
// Explanation: The binary search tree has a depth of 4 with path 1->2->3->4.

// Constraints:
//     n == order.length
//     1 <= n <= 10^5
//     order is a permutation of integers between 1 and n.

import "fmt"
import "github.com/emirpasic/gods/treemap"

// 超出时间限制 56 / 70 
func maxDepthBST(order []int) int {
    res, n := 0, len(order)
    var dfs func(i, low, high, depth int)
    dfs = func(i, low, high, depth int) {
        for j := i + 1; j < n; j++ {
            if order[j] > low && order[j] < high {
                dfs(j, low, order[j], depth + 1)
                dfs(j, order[j], high, depth + 1)
                return
            }
        }
        if depth > res {
            res = depth
        }
    }
    dfs(0, order[0], 100001, 1) // 1 <= n <= 10^5
    dfs(0, 0, order[0], 1)
    return res
}

func maxDepthBST1(order []int) int {
    res := 0
    m := treemap.NewWithIntComparator()
    m.Put(0, 0)
    for i := 0; i < len(order); i++ {
        k,v := m.Floor(order[i])
        key,value:= k.(int),v.(int)+1
        m.Put(order[i]+1, value)
        m.Put(key,value)
        if res < value {
            res = value
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/15/1.png" />
    // Input: order = [2,1,4,3]
    // Output: 3
    // Explanation: The binary search tree has a depth of 3 with path 2->3->4.
    fmt.Println(maxDepthBST([]int{2,1,4,3})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/15/2.png" />
    // Input: order = [2,1,3,4]
    // Output: 3
    // Explanation: The binary search tree has a depth of 3 with path 2->3->4.
    fmt.Println(maxDepthBST([]int{2,1,3,4})) // 3
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/15/3.png" />
    // Input: order = [1,2,3,4]
    // Output: 4
    // Explanation: The binary search tree has a depth of 4 with path 1->2->3->4.
    fmt.Println(maxDepthBST([]int{1,2,3,4})) // 4

    fmt.Println(maxDepthBST([]int{2,1,4,3})) // 3
    fmt.Println(maxDepthBST([]int{2,1,3,4})) // 3
    fmt.Println(maxDepthBST([]int{1,2,3,4})) // 4
}