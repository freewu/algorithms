package main

// 590. N-ary Tree Postorder Traversal
// Given the root of an n-ary tree, return the postorder traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal. 
// Each group of children is separated by the null value (See examples)

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [5,6,3,2,4,1]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     0 <= Node.val <= 10^4
//     The height of the n-ary tree is less than or equal to 1000.

// Follow up: Recursive solution is trivial, could you do it iteratively?

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 //  Definition for a Node.
type Node struct {
    Val      int
    Children []*Node
}

// 迭代 stack
func postorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    if len(root.Children) == 0 {
        return []int{root.Val}
    }
    res, stack, m := make([]int, 0),  make([]*Node, 0), map[*Node]int{}
    m[root]++
    for len(stack) > 0 {
        tmp := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        m[tmp]++
        if m[tmp] > 2 {
            res = append(res, tmp.Val)
            continue
        }
        if len(tmp.Children) == 0 {
            res = append(res, tmp.Val)
        } else {
            stack = append(stack, tmp)
        }
        for i := len(tmp.Children) - 1; i >= 0; i-- {
            m[tmp.Children[i]]++
            stack = append(stack, tmp.Children[i])
        }
    }
    return res
}

// best solution
func postorder1(root *Node) []int {
    res := []int{}
    if root == nil {
        return res
    }
    for _, child := range root.Children {
        res = append(res, postorder1(child)...)
    }
    res = append(res, root.Val)
    return res
}

func main() {
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png" />
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [5,6,3,2,4,1]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png" />
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]
fmt.Println()
}