package main

// 701. Insert into a Binary Search Tree
// You are given the root node of a binary search tree (BST) and a value to insert into the tree. 
// Return the root node of the BST after the insertion. 
// It is guaranteed that the new value does not exist in the original BST.

// Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. 
// You can return any of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/05/insertbst.jpg" />
// Input: root = [4,2,7,1,3], val = 5
// Output: [4,2,7,1,3,5]
// Explanation: Another accepted tree is:

// Example 2:
// Input: root = [40,20,60,10,30,50,70], val = 25
// Output: [40,20,60,10,30,50,70,null,null,25]

// Example 3:
// Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
// Output: [4,2,7,1,3,5]

// Constraints:
//     The number of nodes in the tree will be in the range [0, 10^4].
//     -10^8 <= Node.val <= 10^8
//     All the values Node.val are unique.
//     -10^8 <= val <= 10^8
//     It's guaranteed that val does not exist in the original BST.

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
// 递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}

// bfs
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    node := root
    for {
        if node.Val > val {
            if node.Left != nil {
                node = node.Left
            } else {
                node.Left = &TreeNode{Val: val}
                break
            }
        }
        if node.Val < val {
            if node.Right != nil {
                node = node.Right
            } else {
                node.Right = &TreeNode{Val: val}
                break
            }
        }
    }
    return root
}

// dfs
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
    var dfs func (root *TreeNode, val int) *TreeNode 
    dfs = func (root *TreeNode, val int) *TreeNode {
        if root == nil {
            return &TreeNode{Val: val}
        }
        if val < root.Val {
            root.Left = dfs(root.Left, val)
        } else {
            root.Right = dfs(root.Right, val)
        }
        return root
    }
    return dfs(root, val)
}

func insertIntoBST3(root *TreeNode, val int) *TreeNode {
    p := root
    var prev *TreeNode = nil
    for p != nil {
        prev = p
        if p.Val == val {
            return root
        } else if p.Val > val {
            p = p.Left
        } else {
            p = p.Right
        }
    }
    if prev == nil {
        prev = &TreeNode{Val: val}
        root = prev
    } else if prev.Val > val {
        prev.Left = &TreeNode{Val: val}
    } else {
        prev.Right = &TreeNode{Val: val}
    }
    return root
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/05/insertbst.jpg" />
    // Input: root = [4,2,7,1,3], val = 5
    // Output: [4,2,7,1,3,5]
    // Explanation: Another accepted tree is:
    tree1 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    fmt.Println(insertIntoBST(tree1, 5))
    // Example 2:
    // Input: root = [40,20,60,10,30,50,70], val = 25
    // Output: [40,20,60,10,30,50,70,null,null,25]
    tree2 := &TreeNode {
        40,
        &TreeNode{20, &TreeNode{10, nil, nil}, &TreeNode{30, nil, nil}, },
        &TreeNode{60, &TreeNode{50, nil, nil}, &TreeNode{70, nil, nil}, },
    }
    fmt.Println(insertIntoBST(tree2, 25))
    // Example 3:
    // Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
    // Output: [4,2,7,1,3,5]
    tree3 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    fmt.Println(insertIntoBST(tree3, 5))


    tree11 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    fmt.Println(insertIntoBST1(tree11, 5))
    tree12 := &TreeNode {
        40,
        &TreeNode{20, &TreeNode{10, nil, nil}, &TreeNode{30, nil, nil}, },
        &TreeNode{60, &TreeNode{50, nil, nil}, &TreeNode{70, nil, nil}, },
    }
    fmt.Println(insertIntoBST1(tree12, 25))
    tree13 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    fmt.Println(insertIntoBST1(tree13, 5))

    tree21 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    fmt.Println(insertIntoBST2(tree21, 5))
    tree22 := &TreeNode {
        40,
        &TreeNode{20, &TreeNode{10, nil, nil}, &TreeNode{30, nil, nil}, },
        &TreeNode{60, &TreeNode{50, nil, nil}, &TreeNode{70, nil, nil}, },
    }
    fmt.Println(insertIntoBST2(tree22, 25))
    tree23 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    fmt.Println(insertIntoBST2(tree23, 5))

    tree31 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    fmt.Println(insertIntoBST3(tree31, 5))
    tree32 := &TreeNode {
        40,
        &TreeNode{20, &TreeNode{10, nil, nil}, &TreeNode{30, nil, nil}, },
        &TreeNode{60, &TreeNode{50, nil, nil}, &TreeNode{70, nil, nil}, },
    }
    fmt.Println(insertIntoBST3(tree32, 25))
    tree33 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    fmt.Println(insertIntoBST3(tree33, 5))
}