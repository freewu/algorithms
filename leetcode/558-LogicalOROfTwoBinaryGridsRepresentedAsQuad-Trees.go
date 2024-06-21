package main

// 558. Logical OR of Two Binary Grids Represented as Quad-Trees
// A Binary Matrix is a matrix in which all the elements are either 0 or 1.
// Given quadTree1 and quadTree2. quadTree1 represents a n * n binary matrix and quadTree2 represents another n * n binary matrix.
// Return a Quad-Tree representing the n * n binary matrix which is the result of logical bitwise OR of the two binary matrixes represented by quadTree1 and quadTree2.
// Notice that you can assign the value of a node to True or False when isLeaf is False, and both are accepted in the answer.
// A Quad-Tree is a tree data structure in which each internal node has exactly four children. Besides, each node has two attributes:

// val: True if the node represents a grid of 1's or False if the node represents a grid of 0's.
// isLeaf: True if the node is leaf node on the tree or False if the node has the four children.
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
//     <img src="https://assets.leetcode.com/uploads/2020/02/11/new_top.png" />

// If you want to know more about the Quad-Tree, you can refer to the wiki.

// Quad-Tree format:
//     The input/output represents the serialized format of a Quad-Tree using level order traversal, where null signifies a path terminator where no node exists below.
//     It is very similar to the serialization of the binary tree. The only difference is that the node is represented as a list [isLeaf, val].
//     If the value of isLeaf or val is True we represent it as 1 in the list [isLeaf, val] and if the value of isLeaf or val is False we represent it as 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/02/11/qt1.png" />
// <img src="https://assets.leetcode.com/uploads/2020/02/11/qt2.png" />
// Input: quadTree1 = [[0,1],[1,1],[1,1],[1,0],[1,0]]
// , quadTree2 = [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
// Output: [[0,0],[1,1],[1,1],[1,1],[1,0]]
// Explanation: quadTree1 and quadTree2 are shown above. You can see the binary matrix which is represented by each Quad-Tree.
// If we apply logical bitwise OR on the two binary matrices we get the binary matrix below which is represented by the result Quad-Tree.
// Notice that the binary matrices shown are only for illustration, you don't have to construct the binary matrix to get the result tree.
// <img src="https://assets.leetcode.com/uploads/2020/02/11/qtr.png" />

// Example 2:
// Input: quadTree1 = [[1,0]], quadTree2 = [[1,0]]
// Output: [[1,0]]
// Explanation: Each tree represents a binary matrix of size 1*1. Each matrix contains only zero.
// The resulting matrix is of size 1*1 with also zero.
 
// Constraints:
//     quadTree1 and quadTree2 are both valid Quad-Trees each representing a n * n grid.
//     n == 2x where 0 <= x <= 9.

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
// func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
//     if quadTree1.IsLeaf { if quadTree1.Val { return quadTree1 } else { return quadTree2 } }
//     if quadTree2.IsLeaf { if quadTree2.Val { return quadTree1 } else { return quadTree1 } }

//     topLeft := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
//     topRight := intersect(quadTree1.TopRight, quadTree2.TopRight)
//     bottomLeft := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
//     bottomRight := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
//     topLeftIsLeaf,     topLeftVal     := topLeft.IsLeaf,     topLeft.Val
//     topRightIsLeaf,    topRightVal    := topRight.IsLeaf,    topRight.Val
//     bottomLeftIsLeaf,  bottomLeftVal  := bottomLeft.IsLeaf,  bottomLeft.Val
//     bottomRightIsLeaf, bottomRightVal := bottomRight.IsLeaf, bottomRight.Val

//     if  topLeftIsLeaf && topRightIsLeaf && bottomLeftIsLeaf && bottomRightIsLeaf && 
//         (topLeftVal == topRightVal && topRightVal == bottomLeftVal && bottomLeftVal == bottomRightVal) {
//         return &Node{topLeftIsLeaf, topLeftVal, nil, nil, nil, nil}
//     }
//     return &Node{false, false, topLeft, topRight, bottomLeft, bottomRight }
// }

func intersect(quadTree1, quadTree2 *Node) *Node {
    if quadTree1.IsLeaf { 
        if quadTree1.Val {//若1为叶子结点且值为1 直接返回全为1的结点
            return &Node{Val: true, IsLeaf: true}
        }
        return quadTree2
    }
    if quadTree2.IsLeaf {//若2为叶子结点 则转换为上面那种情况
        return intersect(quadTree2, quadTree1)
    }
    n1 := intersect(quadTree1.TopLeft, quadTree2.TopLeft) //分治
    n2 := intersect(quadTree1.TopRight, quadTree2.TopRight)
    n3 := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
    n4 := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
    if n1.IsLeaf && n2.IsLeaf && n3.IsLeaf && n4.IsLeaf && 
       n1.Val == n2.Val && n1.Val == n3.Val && n1.Val == n4.Val { // 为了分区或运算以后如果4个分区均相同 则变为一个整体 视为叶子结点
        return &Node{Val: n1.Val, IsLeaf: true}
    }
    return &Node{false, false, n1, n2, n3, n4}
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/02/11/qt1.png" />
    // <img src="https://assets.leetcode.com/uploads/2020/02/11/qt2.png" />
    // Input: quadTree1 = [[0,1],[1,1],[1,1],[1,0],[1,0]]
    // , quadTree2 = [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
    // Output: [[0,0],[1,1],[1,1],[1,1],[1,0]]
    // Explanation: quadTree1 and quadTree2 are shown above. You can see the binary matrix which is represented by each Quad-Tree.
    // If we apply logical bitwise OR on the two binary matrices we get the binary matrix below which is represented by the result Quad-Tree.
    // Notice that the binary matrices shown are only for illustration, you don't have to construct the binary matrix to get the result tree.
    // <img src="https://assets.leetcode.com/uploads/2020/02/11/qtr.png" />

    // Example 2:
    // Input: quadTree1 = [[1,0]], quadTree2 = [[1,0]]
    // Output: [[1,0]]
    // Explanation: Each tree represents a binary matrix of size 1*1. Each matrix contains only zero.
    // The resulting matrix is of size 1*1 with also zero.
    fmt.Println()
}