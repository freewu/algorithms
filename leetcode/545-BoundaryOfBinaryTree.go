package main

// 545. Boundary of Binary Tree
// The boundary of a binary tree is the concatenation of 
// the root, the left boundary, the leaves ordered from left-to-right, and the reverse order of the right boundary.

// The left boundary is the set of nodes defined by the following:
//     The root node's left child is in the left boundary. If the root does not have a left child, then the left boundary is empty.
//     If a node in the left boundary and has a left child, then the left child is in the left boundary.
//     If a node is in the left boundary, has no left child, but has a right child, then the right child is in the left boundary.
//     The leftmost leaf is not in the left boundary.

// The right boundary is similar to the left boundary, except it is the right side of the root's right subtree. 
// Again, the leaf is not part of the right boundary, and the right boundary is empty if the root does not have a right child.

// The leaves are nodes that do not have any children. For this problem, the root is not a leaf.
// Given the root of a binary tree, return the values of its boundary.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/11/boundary1.jpg" />
// Input: root = [1,null,2,3,4]
// Output: [1,3,4,2]
// Explanation:
// - The left boundary is empty because the root does not have a left child.
// - The right boundary follows the path starting from the root's right child 2 -> 4.
//   4 is a leaf, so the right boundary is [2].
// - The leaves from left to right are [3,4].
// Concatenating everything results in [1] + [] + [3,4] + [2] = [1,3,4,2].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/11/boundary2.jpg" />
// Input: root = [1,2,3,4,5,6,null,null,null,7,8,9,10]
// Output: [1,2,4,7,8,9,10,6,3]
// Explanation:
// - The left boundary follows the path starting from the root's left child 2 -> 4.
//   4 is a leaf, so the left boundary is [2].
// - The right boundary follows the path starting from the root's right child 3 -> 6 -> 10.
//   10 is a leaf, so the right boundary is [3,6], and in reverse order is [6,3].
// - The leaves from left to right are [4,7,8,9,10].
// Concatenating everything results in [1] + [2] + [4,7,8,9,10] + [6,3] = [1,2,4,7,8,9,10,6,3].

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -1000 <= Node.val <= 1000


import "fmt"

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) []int {
    //按题意要求，将节点分为几个类别：
    //根节点的左子为左边节点，左边节点的左子（或右子）不是叶子的话就也是左边节点；
    //同理右边界也类似。
    //然后先序遍历，遍历到的顺序就是根->左边->左边叶子->普通节点叶子->右边->右树叶子（可参考示例图）
    //把这几个部分分别存最后根据顺序拼接即可（注意部分结果需要逆序）
    leftB, rightB, normal := 1,2,3
    res, lLeaves, rBound, rLeaves :=  []int{}, []int{}, []int{}, []int{} //有可能是左叶子,右边界,和右叶子
    var traversal func(node *TreeNode, boundaryType int)
    traversal = func(node *TreeNode, boundaryType int) {
        if node == nil {
            return 
        }
        switch boundaryType {
            case leftB:
                if !(node.Left == nil && node.Right == nil) {
                    res = append(res, node.Val)
                    if node.Left!=nil {
                        traversal(node.Left, leftB)
                        traversal(node.Right, normal)
                    } else {
                        traversal(node.Right, leftB)
                    }
                } else {
                    lLeaves = append(lLeaves, node.Val)
                }
            case rightB: 
                if !(node.Left == nil && node.Right == nil) {
                    rBound = append(rBound, node.Val)
                    if node.Right!=nil {
                        traversal(node.Left, normal)
                        traversal(node.Right, rightB)
                    } else {
                        traversal(node.Left, rightB)
                    }
                } else {
                    rLeaves = append(rLeaves, node.Val)
                }
            case normal: 
                if node.Left == nil && node.Right == nil {
                    lLeaves = append(lLeaves, node.Val)
                } else {
                    traversal(node.Left, normal)
                    traversal(node.Right, normal)
                }
        }
    }
    res = append(res, root.Val) // 添加根节点
    if root.Left != nil {
        traversal(root.Left, leftB)
    }
    if root.Right != nil {
        traversal(root.Right, rightB)
    }
    for i:=0; i < len(lLeaves); i++ { // 添加 lLeaves
        res = append(res, lLeaves[i])
    }
    for i:=len(rLeaves)-1; i>=0; i-- { // 添加 rLeaves
        res = append(res, rLeaves[i])
    }
    for i := len(rBound) - 1; i >= 0; i-- { // 添加右边界
        res = append(res, rBound[i])
    }
    return res
}

func boundaryOfBinaryTree1(root *TreeNode) []int {
    if root == nil {
        return []int{0}
    }
    isLeafNode := func(node *TreeNode) bool {
        return node.Left == nil && node.Right == nil
    }
    if isLeafNode(root) {
        return []int{ root.Val }
    }
    res := []int{ root.Val }
    left, right :=make([]int,0),make([]int,0)
    var getLeftBoundary func(root *TreeNode,res []int) []int
    getLeftBoundary = func(root *TreeNode,res []int) []int {
        if root == nil{ 
            return res
        }
        if !isLeafNode(root) {
            res = append(res,root.Val)
        }
        if root.Left != nil {
            return getLeftBoundary(root.Left,res)
        } else if root.Right != nil {
            return getLeftBoundary(root.Right,res)
        }
        return res
    }
    var getRightBoundary func(root *TreeNode,res []int) []int
    getRightBoundary = func(root *TreeNode,res []int) []int {
        if root == nil {
            return res
        }
        if !isLeafNode(root) {
            res = append(res,root.Val)
        }
        if root.Right != nil {
            return getRightBoundary(root.Right,res)
        } else if root.Left != nil {
            return getRightBoundary(root.Left,res)
        }
        return res
    }
    getLeafNode := func(root *TreeNode) []int {
        res := make([]int,0)
        var dfs func(* TreeNode)
        dfs = func(node *TreeNode) {
            if node == nil {
                return
            }
            dfs(node.Left)
            if isLeafNode(node) {
                res = append(res,node.Val)
            }
            dfs(node.Right)
       }
       dfs(root)
       return res
    }
    reverse := func(src []int ) []int {
        res := make([]int,0)
        for i := len(src)-1; i >= 0; i-- {
            res = append(res, src[i])
        }
        return res
    }
    left = getLeftBoundary(root.Left,left)
    leaf := getLeafNode(root)
    right = getRightBoundary(root.Right, right)
    right = reverse(right)
    res = append(res,left...)
    res = append(res,leaf...)
    res = append(res,right...)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/11/boundary1.jpg" />
    // Input: root = [1,null,2,3,4]
    // Output: [1,3,4,2]
    // Explanation:
    // - The left boundary is empty because the root does not have a left child.
    // - The right boundary follows the path starting from the root's right child 2 -> 4.
    //   4 is a leaf, so the right boundary is [2].
    // - The leaves from left to right are [3,4].
    // Concatenating everything results in [1] + [] + [3,4] + [2] = [1,3,4,2].
    tree1 := &TreeNode {
        1,
        nil,
        &TreeNode{2, &TreeNode{3, nil, nil, }, &TreeNode{4, nil, nil, } },
    }
    fmt.Println(boundaryOfBinaryTree(tree1)) // [1,3,4,2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/11/11/boundary2.jpg" />
    // Input: root = [1,2,3,4,5,6,null,null,null,7,8,9,10]
    // Output: [1,2,4,7,8,9,10,6,3]
    // Explanation:
    // - The left boundary follows the path starting from the root's left child 2 -> 4.
    //   4 is a leaf, so the left boundary is [2].
    // - The right boundary follows the path starting from the root's right child 3 -> 6 -> 10.
    //   10 is a leaf, so the right boundary is [3,6], and in reverse order is [6,3].
    // - The leaves from left to right are [4,7,8,9,10].
    // Concatenating everything results in [1] + [2] + [4,7,8,9,10] + [6,3] = [1,2,4,7,8,9,10,6,3].
    tree2 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{7, &TreeNode{6, nil, nil, }, &TreeNode{3, nil, nil, }, }, &TreeNode{8, nil, nil, }, },
        &TreeNode{4, &TreeNode{9, nil, nil, }, &TreeNode{10, nil, nil, } },
    }
    fmt.Println(boundaryOfBinaryTree(tree2)) // [1,2,4,7,8,9,10,6,3]

    fmt.Println(boundaryOfBinaryTree1(tree1)) // [1,3,4,2]
    fmt.Println(boundaryOfBinaryTree1(tree2)) // [1,2,4,7,8,9,10,6,3]
}