package main

// 652. Find Duplicate Subtrees
// Given the root of a binary tree, return all duplicate subtrees.
// For each kind of duplicate subtrees, you only need to return the root node of any one of them.
// Two trees are duplicate if they have the same structure with the same node values.

// Example 1:
//             1
//           /   \
//         (2)     3
//         /     /   \
//       (4)   (2)     4
//             /
//           (4)
// <img src="https://assets.leetcode.com/uploads/2020/08/16/e1.jpg" />
// Input: root = [1,2,3,4,null,2,4,null,null,4]
// Output: [[2,4],[4]]

// Example 2:
//         2
//       /   \
//     (1)    (1)
// <img src="https://assets.leetcode.com/uploads/2020/08/16/e2.jpg" />
// Input: root = [2,1,1]
// Output: [[1]]

// Example 3:
//         2
//       /   \
//     (2)   (2)
//     /     /
//   (4)   (4)
// <img src="https://assets.leetcode.com/uploads/2020/08/16/e33.jpg" />
// Input: root = [2,2,2,3,null,3,null]
// Output: [[2,3],[3]]
 
// Constraints:
//     The number of the nodes in the tree will be in the range [1, 5000]
//     -200 <= Node.val <= 200

import "fmt"
import "strconv"

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    exist := map[string]int{}
    res := []*TreeNode{}
    var dfs func(root *TreeNode) string
    dfs = func(root *TreeNode) string {
        if root == nil {
            return "nil"
        }
        tree := strconv.Itoa(root.Val)
        tree += "," + dfs(root.Left)
        tree += "," + dfs(root.Right)
        if v, _ := exist[tree]; v == 1 { // 存在一样的key 说有重复的子树
            res = append(res, root)
        }
        exist[tree]++
        return tree
    }
    dfs(root)
    return res
}

func findDuplicateSubtrees1(root *TreeNode) []*TreeNode {
    type pair struct {
        node   *TreeNode
        index  int
    }
    repeat, seen, index := map[*TreeNode]struct{}{}, map[[3]int]pair{}, 0
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        tri := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
        if p, ok := seen[tri]; ok {
            repeat[p.node] = struct{}{}
            return p.index
        }
        index++
        seen[tri] = pair{node, index}
        return index
    }
    dfs(root)
    res := make([]*TreeNode, 0, len(repeat))
    for node := range repeat {
        res = append(res, node)
    }
    return res
}

func main() {
    // Example 1:
    //             1
    //           /   \
    //         (2)     3
    //         /     /   \
    //       (4)   (2)     4
    //             /
    //           (4)
    // <img src="https://assets.leetcode.com/uploads/2020/08/16/e1.jpg" />
    // Input: root = [1,2,3,4,null,2,4,null,null,4]
    // Output: [[2,4],[4]]
    tree1 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{4,  nil,                    nil}, nil,                    },
        &TreeNode{3, &TreeNode{2,  &TreeNode{4, nil, nil}, nil}, &TreeNode{4, nil, nil}, },
    }
    fmt.Println(findDuplicateSubtrees(tree1)) //[[2,4],[4]]
    // Example 2:
    //         2
    //       /   \
    //     (1)    (1)
    // <img src="https://assets.leetcode.com/uploads/2020/08/16/e2.jpg" />
    // Input: root = [2,1,1]
    // Output: [[1]]
    tree2 := &TreeNode {
        2,
        &TreeNode{1, nil, nil, },
        &TreeNode{1, nil, nil, },
    }
    fmt.Println(findDuplicateSubtrees(tree2)) //[[1]]
    // Example 3:
    //         2
    //       /   \
    //     (2)   (2)
    //     /     /
    //   (3)   (3)
    // <img src="https://assets.leetcode.com/uploads/2020/08/16/e33.jpg" />
    // Input: root = [2,2,2,3,null,3,null]
    // Output: [[2,3],[3]]
    tree3 := &TreeNode {
        2,
        &TreeNode{2, &TreeNode{3, nil, nil, }, nil, },
        &TreeNode{2, &TreeNode{3, nil, nil, }, nil, },
    }
    fmt.Println(findDuplicateSubtrees(tree3)) // [[2,3],[3]]

    fmt.Println(findDuplicateSubtrees1(tree1)) //[[2,4],[4]]
    fmt.Println(findDuplicateSubtrees1(tree2)) //[[1]]
    fmt.Println(findDuplicateSubtrees1(tree3)) // [[2,3],[3]]
}