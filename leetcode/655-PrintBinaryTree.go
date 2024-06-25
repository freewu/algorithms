package main

// 655. Print Binary Tree
// Given the root of a binary tree, construct a 0-indexed m x n string matrix res 
// that represents a formatted layout of the tree. 
// The formatted layout matrix should be constructed using the following rules:
//     The height of the tree is height and the number of rows m should be equal to height + 1.
//     The number of columns n should be equal to 2height+1 - 1.
//     Place the root node in the middle of the top row (more formally, at location res[0][(n-1)/2]).
//     For each node that has been placed in the matrix at position res[r][c], place its left child at res[r+1][c-2height-r-1] and its right child at res[r+1][c+2height-r-1].
//     Continue this process until all the nodes in the tree have been placed.
//     Any empty cells should contain the empty string "".

// Return the constructed matrix res.

// Example 1:
//      1
//     /
//    2
// <img src="https://assets.leetcode.com/uploads/2021/05/03/print1-tree.jpg" />
// Input: root = [1,2]
// Output: 
// [["","1",""],
//  ["2","",""]]

// Example 2:
//         1
//       /   \
//      2     3
//       \ 
//         4
// <img src="https://assets.leetcode.com/uploads/2021/05/03/print2-tree.jpg" />
// Input: root = [1,2,3,null,4]
// Output: 
// [["","","","1","","",""],
//  ["","2","","","","3",""],
//  ["","","4","","","",""]]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 2^10].
//     -99 <= Node.val <= 99
//     The depth of the tree will be in the range [1, 10].

import "fmt"
import "math"
import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var findHeight func(node *TreeNode) int
    findHeight = func (node *TreeNode) int {
        if node == nil {
            return -1
        }
        return max(findHeight(node.Left), findHeight(node.Right)) + 1
    }
    height := findHeight(root)
    rows := height + 1
    cols := int(math.Pow(float64(2), float64(height) + float64(1)) - 1)
    res := make([][]string, rows)
    for row := 0; row < rows; row += 1 {
        res[row] = make([]string, cols)
    }
    // res[0][(cols - 1) / 2] = string(root.Val) 
    var buildArray func(r, c int, node *TreeNode)
    buildArray = func(r, c int, node *TreeNode) {
        if node == nil {
            return
        }
        res[r][c] = strconv.Itoa(node.Val)
        buildArray(r + 1, c - int(math.Pow(float64(2), float64(height - r - 1))), node.Left) // build left
        buildArray(r + 1, c + int(math.Pow(float64(2), float64(height - r - 1))), node.Right) // build right
    }
    buildArray(0, (cols - 1) / 2, root)
    return res
}

func printTree1(root *TreeNode) [][]string {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    height := 0
    var dfs func(r *TreeNode, h int)
    dfs = func(r *TreeNode, h int) { // 获取树的高度
        if r == nil { return }
        height = max(height, h)
        dfs(r.Left, h + 1)
        dfs(r.Right, h + 1)
    }
    dfs(root, 0)
    n := math.Pow(2, float64(height+1)) - 1 // 计算树的宽度
    res := make([][]string, height + 1)
    for i := range res { 
        res[i] = make([]string, int(n)) 
    }
    var print func(p *TreeNode, r, c int)
    print = func(x *TreeNode, r, c int) {
        if x == nil { return }
        res[r][c] = strconv.Itoa(x.Val)
        p := math.Pow(2, float64(height - r - 1))
        print(x.Left, r+1, c-int(p))
        print(x.Right, r+1, c+int(p))
    }
    print(root, 0, (int(n)-1) / 2)
    return res
}

func main() {
    // Example 1:
    //      1
    //     /
    //    2
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/print1-tree.jpg" />
    // Input: root = [1,2]
    // Output: 
    // [["","1",""],
    //  ["2","",""]]
    tree1 := &TreeNode {
        1,
        &TreeNode{2, nil, nil, },
        nil,
    }
    fmt.Println(printTree(tree1)) // [[ 1 ] [2  ]]
    // Example 2:
    //         1
    //       /   \
    //      2     3
    //       \ 
    //         4
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/print2-tree.jpg" />
    // Input: root = [1,2,3,null,4]
    // Output: 
    // [["","","","1","","",""],
    //  ["","2","","","","3",""],
    //  ["","","4","","","",""]]
    tree2 := &TreeNode {
        1,
        &TreeNode{2, nil, &TreeNode{4, nil, nil, }, },
        &TreeNode{3, nil, nil,                      },
    }
    fmt.Println(printTree(tree2)) // [[   1   ] [ 2    3 ] [  4    ]]

    fmt.Println(printTree1(tree1)) // [[ 1 ] [2  ]]
    fmt.Println(printTree1(tree2)) // [[   1   ] [ 2    3 ] [  4    ]]
}