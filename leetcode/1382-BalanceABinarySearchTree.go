package main

// 1382. Balance a Binary Search Tree
// Given the root of a binary search tree, return a balanced binary search tree with the same node values. 
// If there is more than one answer, return any of them.

// A binary search tree is balanced if the depth of the two subtrees of every node never differs by more than 1.

// Example 1:
//     1
//       \
//        2             =>      2                  3
//         \                  /    \      or     /    \
//          3                1      3           1      4
//            \                      \           \
//             4                      4            2
// <img src="https://assets.leetcode.com/uploads/2021/08/10/balance1-tree.jpg" />
// Input: root = [1,null,2,null,3,null,4,null,null]
// Output: [2,1,3,null,null,null,4]
// Explanation: This is not the only correct answer, [3,1,4,null,2] is also correct.

// Example 2:
//         2
//       /   \
//      1     3
// <img src="" />
// Input: root = [2,1,3]
// Output: [2,1,3]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     1 <= Node.val <= 10^5

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
    sorted := []int{}
    var inorder func(node *TreeNode, sorted *[]int)
    inorder = func(node *TreeNode, sorted *[]int) {
        if node == nil { return }
        inorder(node.Left, sorted)
        *sorted = append(*sorted, node.Val)
        inorder(node.Right, sorted)
    }
    inorder(root, &sorted)
    var build func(sorted []int) *TreeNode
    build = func(sorted []int) *TreeNode {
        if len(sorted) == 0 { return nil }
        index := len(sorted) / 2
        return &TreeNode{Val: sorted[index], Left: build(sorted[:index]), Right: build(sorted[index + 1:])}
    }
    return build(sorted)
}

// stack
func balanceBST1(root *TreeNode) *TreeNode {
    nodes, stack := []*TreeNode{}, []*TreeNode{}
    curr := root
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        nodes = append(nodes, curr)
        curr = curr.Right
    }
    var build func(nodes []*TreeNode) *TreeNode
    build = func(nodes []*TreeNode) *TreeNode {
        if len(nodes) == 0 {
            return nil
        }
        mid := len(nodes) / 2
        root := nodes[mid]
        root.Left, root.Right = build(nodes[:mid]), build(nodes[mid+1:])
        return root
    }
    return build(nodes)
}


func main() {
    // Example 1:
    //     1
    //       \
    //        2             =>      2                  3
    //         \                  /    \      or     /    \
    //          3                1      3           1      4
    //            \                      \           \
    //             4                      4            2
    // <img src="https://assets.leetcode.com/uploads/2021/08/10/balance1-tree.jpg" />
    // Input: root = [1,null,2,null,3,null,4,null,null]
    // Output: [2,1,3,null,null,null,4]
    // Explanation: This is not the only correct answer, [3,1,4,null,2] is also correct.
    tree1 := &TreeNode {
        1,
        nil,
        &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, nil, }, }, },
    }
    fmt.Println(balanceBST(tree1)) // &{3 0xc0000940c0 0xc0000940d8}
    // Example 2:
    //         2
    //       /   \
    //      1     3
    // <img src="" />
    // Input: root = [2,1,3]
    // Output: [2,1,3]
    tree2 := &TreeNode {
        2,
        &TreeNode{1, nil, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(balanceBST(tree2)) // &{2 0xc000094180 0xc000094198}

    tree11 := &TreeNode {
        1,
        nil,
        &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, nil, }, }, },
    }
    fmt.Println(balanceBST1(tree11)) // &{3 0xc0000940c0 0xc0000940d8}
    tree12 := &TreeNode {
        2,
        &TreeNode{1, nil, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(balanceBST1(tree12)) // &{2 0xc000094180 0xc000094198}
}