package main

// 331. Verify Preorder Serialization of a Binary Tree
// One way to serialize a binary tree is to use preorder traversal. 
// When we encounter a non-null node, we record the node's value. 
// If it is a null node, we record using a sentinel value such as '#'.

//      _9_
//     /   \
//    3     2
//   / \   / \
//  4   1  #  6
// / \ / \   / \
// # # # #   # #

// For example, the above binary tree can be serialized to the string "9,3,4,#,#,1,#,#,2,#,6,#,#", 
// where '#' represents a null node.

// Given a string of comma-separated values preorder, 
// return true if it is a correct preorder traversal serialization of a binary tree.

// It is guaranteed that each comma-separated value in the string must be either an integer or a character '#' representing null pointer.

// You may assume that the input format is always valid.

// For example, it could never contain two consecutive commas, such as "1,,3".
// Note: You are not allowed to reconstruct the tree.

// Example 1:
// Input: preorder = "9,3,4,#,#,1,#,#,2,#,6,#,#"
// Output: true

// Example 2:
// Input: preorder = "1,#"
// Output: false

// Example 3:
// Input: preorder = "9,#,#,1"
// Output: false
 
// Constraints:
//     1 <= preorder.length <= 10^4
//     preorder consist of integers in the range [0, 100] and '#' separated by commas ','.

import "fmt"
import "strings"

func isValidSerialization(preorder string) bool {
    // 记录出度和度之间的差异 diff = outdegree - indegree
    nodes, diff := strings.Split(preorder, ","), 1
    for _, node := range nodes {
        //当下一个节点到来时，我们将 diff 减 1，因为这个节点提供了一个度
        diff--
        if diff < 0 {
            return false
        }
        // 如果这个节点不为 null，我们将 diff 增加 2，因为它提供两个出度
        if node != "#" {
            diff += 2
        }
    }
    // 如果序列化是正确的，则 diff 应该永远不会为负，并且 diff 在完成时将为零
    return diff == 0
}

func main() {
    fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#")) // true
    fmt.Println(isValidSerialization("1,#")) // false
    fmt.Println(isValidSerialization("9,#,#,1")) // false
}