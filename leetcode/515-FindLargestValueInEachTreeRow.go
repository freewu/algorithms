package main

// 515. Find Largest Value in Each Tree Row
// Given the root of a binary tree, return an array of the largest value in each row of the tree (0-indexed).

// Example 1:
//         (1)
//       /     \
//     (3)      2
//    /   \      \
//   5     3     (9)
// <img src="https://assets.leetcode.com/uploads/2020/08/21/largest_e1.jpg" />
// Input: root = [1,3,2,5,3,null,9]
// Output: [1,3,9]

// Example 2:
//      (1)
//     /   \
//    2    (3)
// Input: root = [1,2,3]
// Output: [1,3]
 
// Constraints:
//     The number of nodes in the tree will be in the range [0, 10^4].
//     -2^31 <= Node.val <= 2^31 - 1

import "fmt"
import "container/list"

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
// bfs
func largestValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    res, inf := []int{}, -1 << 32 - 1
    queue := []*TreeNode{ root }
    for len(queue) > 0 {
        size, mx := len(queue), inf
        for i := 0; i < size; i++ {
            node := queue[i]
            if node.Val > mx { // 找到更大的值了
                mx = node.Val
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[size:]
        res = append(res, mx)
    }
    return res
}

func largestValues1(root *TreeNode) []int {
    res, inf := []int{}, -1 << 32 - 1
    if root == nil {
        return res
    }
    queue := list.New()
    queue.PushFront(root) // 加入队列
    for queue.Len() > 0 {
        size, mx := queue.Len(), inf
        for i := 0; i < size; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
            if node.Val > mx {
                mx = node.Val
            }
        }
        res = append(res,mx)
    }
    return res
}

func main() {
    // Example 1:
    //         (1)
    //       /     \
    //     (3)      2
    //    /   \      \
    //   5     3     (9)
    // <img src="https://assets.leetcode.com/uploads/2020/08/21/largest_e1.jpg" />
    // Input: root = [1,3,2,5,3,null,9]
    // Output: [1,3,9]
    tree1 := &TreeNode {
        1,
        &TreeNode{3, &TreeNode{5, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{2, nil,                    &TreeNode{9, nil, nil}, },
    }
    fmt.Println(largestValues(tree1)) //  [1,3,9]
    // Example 2:
    //      (1)
    //     /   \
    //    2    (3)
    // Input: root = [1,2,3]
    // Output: [1,3]
    tree2 := &TreeNode {
        1,
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(largestValues(tree2)) // [1,3]

    fmt.Println(largestValues1(tree1)) //  [1,3,9]
    fmt.Println(largestValues1(tree2)) // [1,3]
}