package main

// 270. Closest Binary Search Tree Value
// Given the root of a binary search tree and a target value, 
// return the value in the BST that is closest to the target.
// If there are multiple answers, print the smallest.

// Example 1:
//         4
//       /   \ 
//      2     5
//    /   \
//   1     3
// <img src="https://assets.leetcode.com/uploads/2021/03/12/closest1-1-tree.jpg" />
// Input: root = [4,2,5,1,3], target = 3.714286
// Output: 4

// Example 2:
// Input: root = [1], target = 4.428571
// Output: 1
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     0 <= Node.val <= 10^9
//     -10^9 <= target <= 10^9

import "fmt"
import "math"

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
func closestValue(root *TreeNode, target float64) int {
    closest, cur := root, root
    for cur != nil {
        if target == float64(cur.Val) {
            return cur.Val
        }
        if math.Abs(target - float64(cur.Val)) < math.Abs(target - float64(closest.Val)) {
            closest = cur
        } else if (math.Abs(target - float64(cur.Val)) == math.Abs(target - float64(closest.Val))) { // 处理值差值一样的情况
            if cur.Val < closest.Val { // 谁小选择谁
                closest = cur
            }
        }
        if target < float64(cur.Val) {
            cur = cur.Left
        } else {
            cur = cur.Right
        }
    }
    return closest.Val
}

// dfs 
func closestValue1(root *TreeNode, target float64) int {
    closest := root
    var dfs func(*TreeNode,float64)
    dfs = func(node *TreeNode,target float64) {
        if node == nil {
            return 
        }
        dfs(node.Left,target)
        if math.Abs(target - float64(node.Val)) < math.Abs(target - float64(closest.Val)) {
            closest = node
        } else if (math.Abs(target - float64(node.Val)) == math.Abs(target - float64(closest.Val))) { // 处理值差值一样的情况
            if node.Val < closest.Val { // 谁小选择谁
                closest = node
            }
        }
        dfs(node.Right,target)
    }
    dfs(root,target)
    return closest.Val
}

func closestValue2(root *TreeNode, target float64) int {
    diff := func ( a, b float64) float64 {
        if a > b {
            return a - b
        } else {
            return b - a
        }
    }
    closest, bl, a := root, diff(float64(root.Val), target), 0.0
    for root != nil {
        a = diff(float64(root.Val), target)
        if a < bl {
            closest = root
            bl = a
        }
        if target > float64(root.Val) {
            root = root.Right
        } else {
            root = root.Left
        }
    } 
    return closest.Val 
}


func main() {
    // Example 1:
    //         4
    //       /   \ 
    //      2     5
    //    /   \
    //   1     3
    // <img src="https://assets.leetcode.com/uploads/2021/03/12/closest1-1-tree.jpg" />
    // Input: root = [4,2,5,1,3], target = 3.714286
    // Output: 4
    tree1 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil }, &TreeNode{3, nil, nil }, },
        &TreeNode{5, nil, nil },
    }
    fmt.Println(closestValue(tree1, 3.714286)) // 4
    // Example 2:
    // Input: root = [1], target = 4.428571
    // Output: 1
    tree2 := &TreeNode{1, nil, nil }
    fmt.Println(closestValue(tree2, 4.428571)) // 1

    // tree3 := &TreeNode {
    //     4,
    //     &TreeNode{2, &TreeNode{1, nil, nil }, &TreeNode{3, nil, nil }, },
    //     &TreeNode{5, nil, nil },
    // }
    fmt.Println(closestValue(tree1, 3.5)) // 3
    fmt.Println(closestValue(tree1, 4.5)) // 5

    fmt.Println(closestValue1(tree1, 3.714286)) // 4
    fmt.Println(closestValue1(tree2, 4.428571)) // 1
    fmt.Println(closestValue1(tree1, 3.5)) // 3
    fmt.Println(closestValue1(tree1, 4.5)) // 5

    fmt.Println(closestValue2(tree1, 3.714286)) // 4
    fmt.Println(closestValue2(tree2, 4.428571)) // 1
    fmt.Println(closestValue2(tree1, 3.5)) // 3
    fmt.Println(closestValue2(tree1, 4.5)) // 5
}