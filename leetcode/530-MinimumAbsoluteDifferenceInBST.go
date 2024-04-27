package main

// 530. Minimum Absolute Difference in BST
// Given the root of a Binary Search Tree (BST), 
// return the minimum absolute difference between the values of any two different nodes in the tree.

// Example 1:
//         4 
//        /  \
//       2    6
//      /  \ 
//     1    3    
// <img src="https://assets.leetcode.com/uploads/2021/02/05/bst1.jpg" />
// Input: root = [4,2,6,1,3]
// Output: 1

// Example 2:
//         1 
//        /  \
//       0    48
//           /  \ 
//          12   49
// <img src="https://assets.leetcode.com/uploads/2021/02/05/bst2.jpg" />
// Input: root = [1,0,48,null,null,12,49]
// Output: 1
 
// Constraints:
//     The number of nodes in the tree is in the range [2, 10^4].
//     0 <= Node.val <= 10^5

import "fmt"
import "container/list"

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
func getMinimumDifference(root *TreeNode) int {
    t, res := 1 << 32 - 1, 1 << 32 - 1
    min := func (x, y int) int { if x < y { return x; }; return y; }

    var inorder func(root *TreeNode) 
    inorder = func(root *TreeNode) {
        if root == nil {
            return
        }
        inorder(root.Left)
        if root.Val > t {
            res = min(res, root.Val - t)
        } else {
            res = min(res, t - root.Val)
        }
        t = root.Val
        inorder(root.Right)
        return
    }
    inorder(root)
    return res
}

func getMinimumDifference1(root *TreeNode) int {
    valList := []int{}
    var getVal func(*TreeNode)
    getVal = func(root *TreeNode) {
        if root == nil {
            return
        }
        getVal(root.Left)
        valList = append(valList,root.Val)
        getVal(root.Right)
    }
    getVal(root)
    res := 1 << 32 - 1
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i:=0; i < len(valList) - 1; i++ {
        if v := abs(valList[i] - valList[i+1]); v < res {
            res = v
        }
    }
    return res
}

// stack
func getMinimumDifference2(root *TreeNode) int {
    res, prev := 1 << 32 - 1, -1 << 32 - 1
    stack := list.New()
    curr := root
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for curr != nil || stack.Len() > 0 {
        for curr != nil {
            stack.PushBack(curr)
            curr = curr.Left
        }
        top := stack.Remove(stack.Back()).(*TreeNode)
        res = min(res, top.Val-prev)
        prev = top.Val
        curr = top.Right
    }
    return res
}

func main() {
    // Example 1:
    //         4 
    //        /  \
    //       2    6
    //      /  \ 
    //     1    3    
    // <img src="https://assets.leetcode.com/uploads/2021/02/05/bst1.jpg" />
    // Input: root = [4,2,6,1,3]
    // Output: 1
    tree1 := &TreeNode {
        4,
        &TreeNode {
            2,
            &TreeNode{1, nil, nil},
            &TreeNode{3, nil, nil},
        },
        &TreeNode{6, nil, nil},
    }
    fmt.Println(getMinimumDifference(tree1)) // 1
    // Example 2:
    //         1 
    //        /  \
    //       0    48
    //           /  \ 
    //          12   49
    // <img src="https://assets.leetcode.com/uploads/2021/02/05/bst2.jpg" />
    // Input: root = [1,0,48,null,null,12,49]
    // Output: 1
    tree2 := &TreeNode {
        1,
        &TreeNode{0, nil, nil},
        &TreeNode {
            48,
            &TreeNode{12, nil, nil},
            &TreeNode{49, nil, nil},
        },
    }
    fmt.Println(getMinimumDifference(tree2)) // 1

    fmt.Println(getMinimumDifference1(tree1)) // 1
    fmt.Println(getMinimumDifference1(tree2)) // 1

    fmt.Println(getMinimumDifference2(tree1)) // 1
    fmt.Println(getMinimumDifference2(tree2)) // 1
}