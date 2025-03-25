package main

// 2689. Extract Kth Character From The Rope Tree
// You are given the root of a binary tree and an integer k. 
// Besides the left and right children, every node of this tree has two other properties, a string node.val containing only lowercase English letters (possibly empty) and a non-negative integer node.len. 
// There are two types of nodes in this tree:
//     Leaf: These nodes have no children, node.len = 0, and node.val is some non-empty string.
//     Internal: These nodes have at least one child (also at most two children), node.len > 0, and node.val is an empty string.

// The tree described above is called a Rope binary tree. Now we define S[node] recursively as follows:
//     If node is some leaf node, S[node] = node.val,
//     Otherwise if node is some internal node, S[node] = concat(S[node.left], S[node.right]) and S[node].length = node.len.

// Return k-th character of the string S[root].
// Note: If s and p are two strings, concat(s, p) is a string obtained by concatenating p to s. For example, concat("ab", "zz") = "abzz".

// Example 1:
// Input: root = [10,4,"abcpoe","g","rta"], k = 6
// Output: "b"
// Explanation: In the picture below, we put an integer on internal nodes that represents node.len, and a string on leaf nodes that represents node.val.
// You can see that S[root] = concat(concat("g", "rta"), "abcpoe") = "grtaabcpoe". So S[root][5], which represents 6th character of it, is equal to "b".
// <img src="https://assets.leetcode.com/uploads/2023/05/14/example1.png" />

// Example 2:
// Input: root = [12,6,6,"abc","efg","hij","klm"], k = 3
// Output: "c"
// Explanation: In the picture below, we put an integer on internal nodes that represents node.len, and a string on leaf nodes that represents node.val.
// You can see that S[root] = concat(concat("abc", "efg"), concat("hij", "klm")) = "abcefghijklm". So S[root][2], which represents the 3rd character of it, is equal to "c".
// <img src="https://assets.leetcode.com/uploads/2023/05/14/example2.png" />

// Example 3:
// Input: root = ["ropetree"], k = 8
// Output: "e"
// Explanation: In the picture below, we put an integer on internal nodes that represents node.len, and a string on leaf nodes that represents node.val.
// You can see that S[root] = "ropetree". So S[root][7], which represents 8th character of it, is equal to "e".
// <img src="https://assets.leetcode.com/uploads/2023/05/14/example3.png" />

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^3]
//     node.val contains only lowercase English letters
//     0 <= node.val.length <= 50
//     0 <= node.len <= 10^4
//     for leaf nodes, node.len = 0 and node.val is non-empty
//     for internal nodes, node.len > 0 and node.val is empty
//     1 <= k <= S[root].length

import "fmt"

type RopeTreeNode struct {
    len   int
    val   string
    left  *RopeTreeNode
    right *RopeTreeNode
}

/**
 * Definition for a rope tree node.
 * type RopeTreeNode struct {
 * 	   len   int
 * 	   val   string
 * 	   left  *RopeTreeNode
 * 	   right *RopeTreeNode
 * }
 */
func getKthCharacter(root *RopeTreeNode, k int) byte {
    var helper func(node *RopeTreeNode, k int) byte
    helper = func(node *RopeTreeNode, k int) byte {
        if node.len == 0 {  return node.val[k-1] }
        left := 0
        if node.left != nil {
            left = node.left.len
            if left == 0{
                left = len(node.left.val)
            }
        }
        if left >= k { // 只需跟踪左分支的大小，来判断
            return helper(node.left, k)
        } else {
            return helper(node.right, k - left)
        }
    }
    return helper(root, k)
}

func getKthCharacter1(root *RopeTreeNode, k int) byte {
    var helper func(node *RopeTreeNode, k int) byte
    helper = func(node *RopeTreeNode, k int) byte {
        if node.len == 0 { return node.val[k - 1] }
        left, right := 0, 0
        if node.left != nil {
            left = node.left.len
            if left == 0 {
                left = len(node.left.val)
            }
        }
        if node.right != nil {
            right = node.right.len
            if right == 0 {
                right = len(node.right.val)
            }
        }
        if k <= left {
            return helper(node.left, k)
        }
        return helper(node.right, k - left)
    }
    return helper(root, k)
}

func main() {
    // Example 1:
    // Input: root = [10,4,"abcpoe","g","rta"], k = 6
    // Output: "b"
    // Explanation: In the picture below, we put an integer on internal nodes that represents node.len, and a string on leaf nodes that represents node.val.
    // You can see that S[root] = concat(concat("g", "rta"), "abcpoe") = "grtaabcpoe". So S[root][5], which represents 6th character of it, is equal to "b".
    // <img src="https://assets.leetcode.com/uploads/2023/05/14/example1.png" />
    tree1 := &RopeTreeNode{
        10,
        "",
        &RopeTreeNode{4, "",        &RopeTreeNode{0, "g", nil, nil, }, &RopeTreeNode{0, "rta", nil, nil, }, },
        &RopeTreeNode{0, "abcpoe",  nil, nil, }, 
    }
    fmt.Printf("%c\n", getKthCharacter(tree1, 6)) // b
    // Example 2:
    // Input: root = [12,6,6,"abc","efg","hij","klm"], k = 3
    // Output: "c"
    // Explanation: In the picture below, we put an integer on internal nodes that represents node.len, and a string on leaf nodes that represents node.val.
    // You can see that S[root] = concat(concat("abc", "efg"), concat("hij", "klm")) = "abcefghijklm". So S[root][2], which represents the 3rd character of it, is equal to "c".
    // <img src="https://assets.leetcode.com/uploads/2023/05/14/example2.png" />
    tree2 := &RopeTreeNode{
        12,
        "",
        &RopeTreeNode{6, "", &RopeTreeNode{0, "abc", nil, nil, }, &RopeTreeNode{0, "efg", nil, nil, }, },
        &RopeTreeNode{6, "", &RopeTreeNode{0, "hij", nil, nil, }, &RopeTreeNode{0, "klm", nil, nil, }, },
    }
    fmt.Printf("%c\n", getKthCharacter(tree2, 3)) // c
    // Example 3:
    // Input: root = ["ropetree"], k = 8
    // Output: "e"
    // Explanation: In the picture below, we put an integer on internal nodes that represents node.len, and a string on leaf nodes that represents node.val.
    // You can see that S[root] = "ropetree". So S[root][7], which represents 8th character of it, is equal to "e".
    // <img src="https://assets.leetcode.com/uploads/2023/05/14/example3.png" />
    tree3 := &RopeTreeNode{0, "ropetree", nil, nil, }
    fmt.Printf("%c\n", getKthCharacter(tree3, 8)) // e

    fmt.Printf("%c\n", getKthCharacter1(tree1, 6)) // b
    fmt.Printf("%c\n", getKthCharacter1(tree2, 3)) // c
    fmt.Printf("%c\n", getKthCharacter1(tree3, 8)) // e
}