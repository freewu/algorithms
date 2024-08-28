package main

// 1650. Lowest Common Ancestor of a Binary Tree III
// Given two nodes of a binary tree p and q, return their lowest common ancestor (LCA).

// Each node will have a reference to its parent node. The definition for Node is below:
//     class Node {
//         public int val;
//         public Node left;
//         public Node right;
//         public Node parent;
//     }

// According to the definition of LCA on Wikipedia: "The lowest common ancestor of two nodes p and q in a tree T is the lowest node that has both p and q as descendants (where we allow a node to be a descendant of itself)."

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// Output: 3
// Explanation: The LCA of nodes 5 and 1 is 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The LCA of nodes 5 and 4 is 5 since a node can be a descendant of itself according to the LCA definition.

// Example 3:
// Input: root = [1,2], p = 1, q = 2
// Output: 1

// Constraints:
//     The number of nodes in the tree is in the range [2, 10^5].
//     -10^9 <= Node.val <= 10^9
//     All Node.val are unique.
//     p != q
//     p and q exist in the tree.

import "fmt"

type Node struct {
    Val int
    Left *Node
    Right *Node
    Parent *Node
}

/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */
func lowestCommonAncestor(p *Node, q *Node) *Node {
    // p的循环中做的事，相当于以0为分界点，把所有节点都往数轴两端平移，移出-1e9,1e9的范围
    inf := int(1e9)
    for p != nil {
        if p.Val >= 0 {
            p.Val += (inf + 1) 
        } else {
            p.Val -= (inf + 1) 
        }     
        p = p.Parent
    }
    // q的循环中发现不在 [-1e9,1e9] 范围内的就修改回去并返回即可
    for q != nil {
        if q.Val > inf {
            q.Val -= (inf + 1) 
            return q 
        } else if q.Val < -inf {
            q.Val += (inf + 1)
            return q 
        }
        q = q.Parent
    }
    return nil
}

func main() {
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// Output: 3
// Explanation: The LCA of nodes 5 and 1 is 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The LCA of nodes 5 and 4 is 5 since a node can be a descendant of itself according to the LCA definition.

// Example 3:
// Input: root = [1,2], p = 1, q = 2
// Output: 1
    fmt.Println()
}