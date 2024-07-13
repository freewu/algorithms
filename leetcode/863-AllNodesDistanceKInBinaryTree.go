package main

// 863. All Nodes Distance K in Binary Tree
// Given the root of a binary tree, the value of a target node target, and an integer k, 
// return an array of the values of all nodes that have a distance k from the target node.

// You can return the answer in any order.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
// Output: [7,4,1]
// Explanation: The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.

// Example 2:
// Input: root = [1], target = 1, k = 3
// Output: []
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 500].
//     0 <= Node.val <= 500
//     All the values Node.val are unique.
//     target is the value of one of the nodes in the tree.
//     0 <= k <= 1000

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
// 并查集
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    res, mp := []int{}, make(map[*TreeNode]int)
    var find func(root, target *TreeNode)
    find = func(root, target *TreeNode) {
        if root == nil { return }
        if root == target {
            mp[root] = 0
            return
        }
        find(root.Left, target)
        if val, ok := mp[root.Left]; ok {
            mp[root] = val + 1
            return
        }
        find(root.Right, target)
        if val, ok := mp[root.Right]; ok {
            mp[root] = val + 1
            return
        }
    }
    var search func(root *TreeNode, distance int)
    search = func(root *TreeNode, distance int) {
        if root == nil { return }
        if val, ok := mp[root]; ok {
            distance = val
        }
        if distance == k {
            res = append(res, root.Val)
        }
        search(root.Left, distance+1)
        search(root.Right, distance+1)
    }
    find(root, target)
    search(root, 0)
    return res
}

// dfs
func distanceK1(root *TreeNode, target *TreeNode, k int) []int {
    res := []int{}
    var dfs func(node *TreeNode, tmp int)
    dfs = func(node *TreeNode, tmp int) {
        if node == nil { return }
        if tmp == k {
            res = append(res, node.Val)
            return
        }
        dfs(node.Left, tmp+1)
        dfs(node.Right, tmp+1)
    }
    var dfs2 func(node *TreeNode) int 
    dfs2 = func(node *TreeNode) int {
        if node == nil { return -1 }
        if node == target {
            dfs(node, 0)
            return 1
        } else {
            l, r := dfs2(node.Left), dfs2(node.Right)
            if l > 0 {
                if l == k {
                    res = append(res, node.Val)
                }
                dfs(node.Right, l+1)
                return l+1
            }
            if r > 0 {
                if r == k {
                    res = append(res, node.Val)
                }
                dfs(node.Left, r+1)
                return r+1
            }
        }
        return -1
    }
    dfs2(root)
    return res
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
    // Output: [7,4,1]
    // Explanation: The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.
    tree1 := &TreeNode {
        3,
        &TreeNode{5, &TreeNode{6, nil, nil, }, &TreeNode{2, &TreeNode{7, nil, nil, }, &TreeNode{4, nil, nil, }, }, },
        &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{8, nil,                      nil,                      }, },
    }
    fmt.Println(distanceK(tree1,&TreeNode{5, nil, nil, }, 2)) // [7,4,1]
    // Example 2:
    // Input: root = [1], target = 1, k = 3
    // Output: []
    tree2 := &TreeNode{1, nil, nil, }
    fmt.Println(distanceK(tree2, &TreeNode{1, nil, nil, }, 3)) // []

    tree11 := &TreeNode {
        3,
        &TreeNode{5, &TreeNode{6, nil, nil, }, &TreeNode{2, &TreeNode{7, nil, nil, }, &TreeNode{4, nil, nil, }, }, },
        &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{8, nil,                      nil,                      }, },
    }
    fmt.Println(distanceK1(tree11, &TreeNode{5, nil, nil, }, 2)) // [7,4,1]
    tree12 := &TreeNode{1, nil, nil, }
    fmt.Println(distanceK1(tree12, &TreeNode{1, nil, nil, }, 3)) // []
}