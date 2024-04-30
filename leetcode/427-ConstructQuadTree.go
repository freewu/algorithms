package main

// 427. Construct Quad Tree
// Given a n * n matrix grid of 0's and 1's only. 
// We want to represent grid with a Quad-Tree.
// Return the root of the Quad-Tree representing grid.

// A Quad-Tree is a tree data structure in which each internal node has exactly four children. 
// Besides, each node has two attributes:
//     val: 
//         True if the node represents a grid of 1's or False if the node represents a grid of 0's. 
//         Notice that you can assign the val to True or False when isLeaf is False, and both are accepted in the answer.
//     isLeaf: 
//         True if the node is a leaf node on the tree or False if the node has four children.

//     class Node {
//         public boolean val;
//         public boolean isLeaf;
//         public Node topLeft;
//         public Node topRight;
//         public Node bottomLeft;
//         public Node bottomRight;
//     }

// We can construct a Quad-Tree from a two-dimensional area using the following steps:
//     If the current grid has the same value (i.e all 1's or all 0's) set isLeaf True and set val to the value of the grid and set the four children to Null and stop.
//     If the current grid has different values, set isLeaf to False and set val to any value and divide the current grid into four sub-grids as shown in the photo.
//     Recurse for each of the children with the proper sub-grid.

// If you want to know more about the Quad-Tree, you can refer to the wiki.

// Quad-Tree format:

// You don't need to read this section for solving the problem. This is only if you want to understand the output format here. The output represents the serialized format of a Quad-Tree using level order traversal, where null signifies a path terminator where no node exists below.
// It is very similar to the serialization of the binary tree. The only difference is that the node is represented as a list [isLeaf, val].
// If the value of isLeaf or val is True we represent it as 1 in the list [isLeaf, val] and if the value of isLeaf or val is False we represent it as 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/02/11/grid1.png" />
// Input: grid = [[0,1],[1,0]]
// Output: [[0,1],[1,0],[1,1],[1,1],[1,0]]
// Explanation: The explanation of this example is shown below:
// Notice that 0 represents False and 1 represents True in the photo representing the Quad-Tree.
// <img src="https://assets.leetcode.com/uploads/2020/02/12/e1tree.png" />

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/02/12/e2mat.png" />
// Input: grid = [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]
// Output: [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
// Explanation: All values in the grid are not the same. We divide the grid into four sub-grids.
// The topLeft, bottomLeft and bottomRight each has the same value.
// The topRight have different values so we divide it into 4 sub-grids where each has the same value.
// Explanation is shown in the photo below:
// <img src="https://assets.leetcode.com/uploads/2020/02/12/e2tree.png" />

// Constraints:
//     n == grid.length == grid[i].length
//     n == 2x where 0 <= x <= 6

import "fmt"

type Node struct {
    Val bool
    IsLeaf bool
    TopLeft *Node
    TopRight *Node
    BottomLeft *Node
    BottomRight *Node
}

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */
// dfs
func construct(grid [][]int) *Node {
    var dfs func(y0, x0, width int) *Node
    dfs = func(y0, x0, width int) *Node {
        if width == 1 {
            return &Node{
                Val:    grid[y0][x0] == 1,
                IsLeaf: true,
            }
        }
        w := width / 2
        topLeft := dfs(y0, x0, w)
        topRight := dfs(y0, x0+w, w)
        bottomLeft := dfs(y0+w, x0, w)
        bottomRight := dfs(y0+w, x0+w, w)
        var node *Node
        if topLeft.Val == topRight.Val && bottomLeft.Val == bottomRight.Val && topLeft.Val == bottomLeft.Val &&
            topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf {
            node = &Node{
                Val:    topLeft.Val,
                IsLeaf: true,
            }
        } else {
            node = &Node{
                Val:         true,
                IsLeaf:      false,
                TopLeft:     topLeft,
                TopRight:    topRight,
                BottomLeft:  bottomLeft,
                BottomRight: bottomRight,
            }
        }
        return node
    }
    return dfs(0, 0, len(grid))
}

func construct1(grid [][]int) *Node {
    var dfs func([][]int, int, int) *Node
    dfs = func(rows [][]int, c0, c1 int) *Node {
        for _, row := range rows {
            for _, v := range row[c0:c1] {
                if v != rows[0][c0] { // 不是叶节点
                    rMid, cMid := len(rows)/2, (c0+c1)/2
                    return &Node{
                        true,
                        false,
                        dfs(rows[:rMid], c0, cMid),
                        dfs(rows[:rMid], cMid, c1),
                        dfs(rows[rMid:], c0, cMid),
                        dfs(rows[rMid:], cMid, c1),
                    }
                }
            }
        }
        // 是叶节点
        return &Node{Val: rows[0][c0] == 1, IsLeaf: true}
    }
    return dfs(grid, 0, len(grid))
}

func main() {
    matrix1 := [][]int{{0,1},{1,0}}
    n1 := construct(matrix1)
    fmt.Println("n1.Val = ", n1.Val)
    fmt.Println("n1.IsLeaf = ", n1.IsLeaf)
    fmt.Println("n1.TopLeft = ", n1.TopLeft)
    fmt.Println("n1.TopRight = ", n1.TopRight)
    fmt.Println("n1.BottomLeft = ", n1.BottomLeft)
    fmt.Println("n1.BottomRight = ", n1.BottomRight)

    matrix2 := [][]int{
        {1,1,1,1,0,0,0,0},
        {1,1,1,1,0,0,0,0},
        {1,1,1,1,1,1,1,1},
        {1,1,1,1,1,1,1,1},
        {1,1,1,1,0,0,0,0},
        {1,1,1,1,0,0,0,0},
        {1,1,1,1,0,0,0,0},
        {1,1,1,1,0,0,0,0},
    }
    n2 := construct(matrix2)
    fmt.Println("n2.Val = ", n2.Val)
    fmt.Println("n2.IsLeaf = ", n2.IsLeaf)
    fmt.Println("n2.TopLeft = ", n2.TopLeft)
    fmt.Println("n2.TopRight = ", n2.TopRight)
    fmt.Println("n2.BottomLeft = ", n2.BottomLeft)
    fmt.Println("n2.BottomRight = ", n2.BottomRight)

    n11 := construct1(matrix1)
    fmt.Println("n11.Val = ", n11.Val)
    fmt.Println("n11.IsLeaf = ", n11.IsLeaf)
    fmt.Println("n11.TopLeft = ", n11.TopLeft)
    fmt.Println("n11.TopRight = ", n11.TopRight)
    fmt.Println("n11.BottomLeft = ", n11.BottomLeft)
    fmt.Println("n11.BottomRight = ", n11.BottomRight)
    n12 := construct1(matrix2)
    fmt.Println("n12.Val = ", n12.Val)
    fmt.Println("n12.IsLeaf = ", n12.IsLeaf)
    fmt.Println("n12.TopLeft = ", n12.TopLeft)
    fmt.Println("n12.TopRight = ", n12.TopRight)
    fmt.Println("n12.BottomLeft = ", n12.BottomLeft)
    fmt.Println("n12.BottomRight = ", n12.BottomRight)
}