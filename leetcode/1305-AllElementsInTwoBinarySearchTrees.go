package main

// 1305. All Elements in Two Binary Search Trees
// Given two binary search trees root1 and root2, 
// return a list containing all the integers from both trees sorted in ascending order.

// Example 1:
//         2              1
//       /   \          /   \
//      1     4        0     3
// <img src="https://assets.leetcode.com/uploads/2019/12/18/q2-e1.png" />
// Input: root1 = [2,1,4], root2 = [1,0,3]
// Output: [0,1,1,2,3,4]

// Example 2:
//         1              8
//           \          /
//             8       1
// <img src="https://assets.leetcode.com/uploads/2019/12/18/q2-e5-.png" />
// Input: root1 = [1,null,8], root2 = [8,1]
// Output: [1,1,8,8]

// Constraints:
//     The number of nodes in each tree is in the range [0, 5000].
//     -10^5 <= Node.val <= 10^5

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
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    list1, list2 := []int{}, []int{}
    var dfs func(t *TreeNode, list *[]int)
    dfs = func(t *TreeNode, list *[]int) {
        if t == nil { return }
        dfs(t.Left, list)
        *list = append(*list, t.Val) // 中序遍历
        dfs(t.Right, list)
    }
    dfs(root1, &list1)
    dfs(root2, &list2)
    merge := func(list1, list2 []int) []int {
        res, p1, p2, n1, n2 := []int{}, 0, 0, len(list1), len(list2)
        for p1 < n1 || p2 < n2 {
            if p2 >= n2 || p1 < n1 && list1[p1] <= list2[p2] {
                res = append(res, list1[p1])
                p1 ++
            } else if p1 >= n1 || p2 < n2 && list1[p1] > list2[p2] {
                res = append(res, list2[p2])
                p2 ++
            }
        }
        return res
    }
    return merge(list1, list2)
}

func getAllElements1(root1, root2 *TreeNode) []int {
    inorderTraversal := func (root *TreeNode) []int {
        res, stack :=  []int{}, []*TreeNode{}
        curr := root
        for curr != nil || len(stack) > 0 { // bfs
            for curr != nil {
                stack = append(stack, curr)
                curr = curr.Left
            }
            curr = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res = append(res, curr.Val)
            curr = curr.Right
        }
        return res
    }
    // mergeSortedLists merges two sorted lists into one sorted list.
    mergeSortedLists := func(list1, list2 []int) []int {
        res, i, j, n, m := []int{}, 0, 0, len(list1), len(list2)
        for i < n && j < m {
            if list1[i] < list2[j] {
                res = append(res, list1[i])
                i++
            } else {
                res = append(res, list2[j])
                j++
            }
        }
        for i < n { // Add remaining elements from list1, if any.
            res = append(res, list1[i])
            i++
        }
        for j < m { // Add remaining elements from list2, if any.
            res = append(res, list2[j])
            j++
        }
        return res
    }
    return mergeSortedLists(inorderTraversal(root1), inorderTraversal(root2))
}

func main() {
    // Example 1:
    //         2              1
    //       /   \          /   \
    //      1     4        0     3
    // <img src="https://assets.leetcode.com/uploads/2019/12/18/q2-e1.png" />
    // Input: root1 = [2,1,4], root2 = [1,0,3]
    // Output: [0,1,1,2,3,4]
    tree11 := &TreeNode{
        2, 
        &TreeNode{1, nil, nil},
        &TreeNode{4, nil, nil},
    }
    tree12 := &TreeNode{
        1, 
        &TreeNode{0, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(getAllElements(tree11, tree12)) // [0,1,1,2,3,4]
    // Example 2:
    //         1              8
    //           \          /
    //             8       1
    // <img src="https://assets.leetcode.com/uploads/2019/12/18/q2-e5-.png" />
    // Input: root1 = [1,null,8], root2 = [8,1]
    // Output: [1,1,8,8]
    tree21 := &TreeNode{
        1, 
        nil,
        &TreeNode{8, nil, nil},
    }
    tree22 := &TreeNode{
        1, 
        &TreeNode{8, nil, nil},
        nil,
    }
    fmt.Println(getAllElements(tree21, tree22)) // [1,1,8,8]

    fmt.Println(getAllElements1(tree11, tree12)) // [0,1,1,2,3,4]
    fmt.Println(getAllElements1(tree21, tree22)) // [1,1,8,8]
}