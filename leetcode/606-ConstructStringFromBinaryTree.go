package main

// 606. Construct String from Binary Tree
// Given the root node of a binary tree, 
// your task is to create a string representation of the tree following a specific set of formatting rules. 
// The representation should be based on a preorder traversal of the binary tree and must adhere to the following guidelines:
//     1. Node Representation: Each node in the tree should be represented by its integer value.
//     2. Parentheses for Children: If a node has at least one child (either left or right), 
//        its children should be represented inside parentheses. Specifically:
//         2.1 If a node has a left child, the value of the left child should be enclosed in parentheses immediately following the node's value.
//         2.2 If a node has a right child, the value of the right child should also be enclosed in parentheses. 
//             The parentheses for the right child should follow those of the left child.
//     3. Omitting Empty Parentheses: Any empty parentheses pairs (i.e., ()) should be omitted from the final string representation of the tree, with one specific exception: when a node has a right child but no left child. 
//        In such cases, you must include an empty pair of parentheses to indicate the absence of the left child. 
//        This ensures that the one-to-one mapping between the string representation and the original binary tree structure is maintained.

// In summary, empty parentheses pairs should be omitted when a node has only a left child or no children. 
// However, when a node has a right child but no left child, an empty pair of parentheses must precede the representation of the right child to reflect the tree's structure accurately.

// Example 1:
//         1
//       /   \
//      2     3
//     /
//    4
// <img src="https://assets.leetcode.com/uploads/2021/05/03/cons1-tree.jpg" />
// Input: root = [1,2,3,4]
// Output: "1(2(4))(3)"
// Explanation: Originally, it needs to be "1(2(4)())(3()())", but you need to omit all the empty parenthesis pairs. And it will be "1(2(4))(3)".

// Example 2:
//         1
//       /   \
//      2     3
//       \ 
//        4
// <img src="https://assets.leetcode.com/uploads/2021/05/03/cons2-tree.jpg" />
// Input: root = [1,2,3,null,4]
// Output: "1(2()(4))(3)"
// Explanation: Almost the same as the first example, except the () after 2 is necessary to indicate the absence of a left child for 2 and the presence of a right child.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -1000 <= Node.val <= 1000

import "fmt"
import "strconv"
import "strings"

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

// Recursion
func tree2str(t *TreeNode) string {
    var dfs func (node *TreeNode) string
    dfs = func (node *TreeNode) string {
        if node == nil {
            return ""
        }
        if node.Left != nil && node.Right != nil {
            return fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
        } else if node.Left == nil && node.Right == nil {
            return strconv.Itoa(node.Val)
        } else if node.Left != nil {
            return fmt.Sprintf("%d(%s)", node.Val, dfs(node.Left))
        } else {
            return fmt.Sprintf("%d()(%s)", node.Val, dfs(node.Right))
        }
    }
    return dfs(t)
}

func tree2str1(root *TreeNode) string {
    var dfs func(root *TreeNode) string
    dfs = func(root *TreeNode) string {
        switch {
        case root == nil:
            return ""
        case root.Left == nil && root.Right == nil:
            return strconv.Itoa(root.Val)
        case root.Right == nil:
            return fmt.Sprintf("%d(%s)", root.Val, dfs(root.Left))
        default:
            return fmt.Sprintf("%d(%s)(%s)", root.Val, dfs(root.Left), dfs(root.Right))
        }
    }
    return dfs(root)
}

func tree2str2(root *TreeNode) string {
    var dfs func(root *TreeNode) string
    dfs = func(root *TreeNode) string {
        r := &strings.Builder{}
        r.WriteString(strconv.Itoa(root.Val))
        if root.Left == nil && root.Right == nil {
            return r.String()
        }
        if root.Left != nil{
            r.WriteByte('(')
            r.WriteString(dfs(root.Left))
            r.WriteByte(')')
        } else {
            r.WriteString("()")
        }
        if root.Right != nil{
            r.WriteByte('(')
            r.WriteString(dfs(root.Right))
            r.WriteByte(')')
        }
        return r.String()
    }
    return dfs(root)
}

func main() {
    // Example 1:
    //         1
    //       /   \
    //      2     3
    //     /
    //    4
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/cons1-tree.jpg" />
    // Input: root = [1,2,3,4]
    // Output: "1(2(4))(3)"
    // Explanation: Originally, it needs to be "1(2(4)())(3()())", but you need to omit all the empty parenthesis pairs. And it will be "1(2(4))(3)".
    tree1 := &TreeNode{
        1, 
        &TreeNode{2, &TreeNode{5, nil, nil}, nil, },
        &TreeNode{3, nil,                    nil, },
    }
    fmt.Println(tree2str(tree1)) // "1(2(4))(3)"
    // Example 2:
    //         1
    //       /   \
    //      2     3
    //       \ 
    //        4
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/cons2-tree.jpg" />
    // Input: root = [1,2,3,null,4]
    // Output: "1(2()(4))(3)"
    // Explanation: Almost the same as the first example, except the () after 2 is necessary to indicate the absence of a left child for 2 and the presence of a right child.
    tree2 := &TreeNode{
        1, 
        &TreeNode{2, nil, &TreeNode{5, nil, nil}, },
        &TreeNode{3, nil,                    nil, },
    }
    fmt.Println(tree2str(tree2)) // "1(2()(4))(3)"

    fmt.Println(tree2str1(tree1)) // "1(2(4))(3)"
    fmt.Println(tree2str1(tree2)) // "1(2()(4))(3)"

    fmt.Println(tree2str2(tree1)) // "1(2(4))(3)"
    fmt.Println(tree2str2(tree2)) // "1(2()(4))(3)"
}