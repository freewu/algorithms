package main 

// 872. Leaf-Similar Trees
// Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/16/tree.png" />
// For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
// Two binary trees are considered leaf-similar if their leaf value sequence is the same.
// Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

// Example 1:
//             3                           3
//           /   \                      /    \
//          5     1                    5      1
//        /  \   /  \                /  \    /  \
//      (6)  2 (9)  (8)            (6)  (7) (4)  2
//          / \                                 / \
//        (7)  (4)                            (9)  (8)
// <img src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-1.jpg" />
// Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-2.jpg" />
//         1          1
//        /  \       /  \
//       2    3     3    2
// Input: root1 = [1,2,3], root2 = [1,3,2]
// Output: false
 
// Constraints:
//     The number of nodes in each tree will be in the range [1, 200].
//     Both of the given trees will have values in the range [0, 200].

import "fmt"
import "reflect"

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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    list1, list2 := []int{}, []int{}
    var recursion func(node *TreeNode, leafList *[]int)
    recursion = func(node *TreeNode, leafList *[]int) { // 取出叶子节点列表
        if node == nil {
            return
        }
        if node.Left == nil && node.Right == nil {
            *leafList = append(*leafList, node.Val)
            return
        }
        recursion(node.Left, leafList)
        recursion(node.Right, leafList)
    }
    recursion(root1, &list1)
    recursion(root2, &list2)
    return reflect.DeepEqual(list1, list2)
}

func main() {
    // Example 1:
    //             3                           3
    //           /   \                      /    \
    //          5     1                    5      1
    //        /  \   /  \                /  \    /  \
    //      (6)  2 (9)  (8)            (6)  (7) (4)  2
    //          / \                                 / \
    //        (7)  (4)                            (9)  (8)
    // <img src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-1.jpg" />
    // Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
    // Output: true
    tree11 := &TreeNode {
        3,
        &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}, }, },
        &TreeNode{1, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}, },
    }
    tree12 := &TreeNode {
        3,
        &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}, },
        &TreeNode{1, &TreeNode{4, nil, nil}, &TreeNode{2, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}, }, },
    }
    fmt.Println(leafSimilar(tree11, tree12)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-2.jpg" />
    //         1          1
    //        /  \       /  \
    //       2    3     3    2
    // Input: root1 = [1,2,3], root2 = [1,3,2]
    // Output: false
    tree21 := &TreeNode {
        1,
        &TreeNode{2, nil, nil },
        &TreeNode{3, nil, nil },
    }
    tree22 := &TreeNode {
        1,
        &TreeNode{3, nil, nil },
        &TreeNode{2, nil, nil },
    }
    fmt.Println(leafSimilar(tree21, tree22)) // false
}