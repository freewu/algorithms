package main

// 2005. Subtree Removal Game with Fibonacci Tree
// A Fibonacci tree is a binary tree created using the order function order(n):
//     order(0) is the empty tree.
//     order(1) is a binary tree with only one node.
//     order(n) is a binary tree that consists of a root node with the left subtree as order(n - 2) and the right subtree as order(n - 1).

// Alice and Bob are playing a game with a Fibonacci tree with Alice staring first. 
// On each turn, a player selects a node and removes that node and its subtree. 
// The player that is forced to delete root loses.

// Given the integer n, return true if Alice wins the game or false if Bob wins, assuming both players play optimally.

// A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. 
// The tree tree could also be considered as a subtree of itself.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/14/image-20210914173520-3.png" />
//         3
//       /   \
//      1     2
//             \
//              1
// Input: n = 3
// Output: true
// Explanation:
// Alice takes the node 1 in the right subtree.
// Bob takes either the 1 in the left subtree or the 2 in the right subtree.
// Alice takes whichever node Bob doesn't take.
// Bob is forced to take the root node 3, so Bob will lose.
// Return true because Alice wins.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/14/image-20210914173634-4.png" />
// Input: n = 1
// Output: false
// Explanation:
// Alice is forced to take the root node 1, so Alice will lose.
// Return false because Alice loses.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/09/14/image-20210914173425-1.png" />
//         2 
//           \
//            1
// Input: n = 2
// Output: true
// Explanation:
// Alice takes the node 1.
// Bob is forced to take the root node 2, so Bob will lose.
// Return true because Alice wins.

// Constraints:
//     1 <= n <= 100

import "fmt"

func findGameWinner(n int) bool {
    if n == 1 { return false }
    dp := make([]int, n)
    for i := 1; i < n; i++{
        if i == 1 {
            dp[i] = 1
        } else {
            dp[i] = 1 + (dp[i-1] ^ dp[i-2])
        }
    }
    return (dp[n-1] ^ dp[n-2]) != 0
}

func findGameWinner1(n int) bool {
    return n % 6 != 1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/09/14/image-20210914173520-3.png" />
    //         3
    //       /   \
    //      1     2
    //             \
    //              1
    // Input: n = 3
    // Output: true
    // Explanation:
    // Alice takes the node 1 in the right subtree.
    // Bob takes either the 1 in the left subtree or the 2 in the right subtree.
    // Alice takes whichever node Bob doesn't take.
    // Bob is forced to take the root node 3, so Bob will lose.
    // Return true because Alice wins.
    fmt.Println(findGameWinner(3)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/09/14/image-20210914173634-4.png" />
    // Input: n = 1
    // Output: false
    // Explanation:
    // Alice is forced to take the root node 1, so Alice will lose.
    // Return false because Alice loses.
    fmt.Println(findGameWinner(1)) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/09/14/image-20210914173425-1.png" />
    //         2 
    //           \
    //            1
    // Input: n = 2
    // Output: true
    // Explanation:
    // Alice takes the node 1.
    // Bob is forced to take the root node 2, so Bob will lose.
    // Return true because Alice wins.
    fmt.Println(findGameWinner(2)) // true

    fmt.Println(findGameWinner(8)) // true
    fmt.Println(findGameWinner(16)) // true
    fmt.Println(findGameWinner(32)) // true
    fmt.Println(findGameWinner(64)) // true
    fmt.Println(findGameWinner(99)) // true
    fmt.Println(findGameWinner(100)) // true
    fmt.Println(findGameWinner(37)) // false
    fmt.Println(findGameWinner(61)) // false

    fmt.Println(findGameWinner1(3)) // true
    fmt.Println(findGameWinner1(1)) // false
    fmt.Println(findGameWinner1(2)) // true
    fmt.Println(findGameWinner1(8)) // true
    fmt.Println(findGameWinner1(16)) // true
    fmt.Println(findGameWinner1(32)) // true
    fmt.Println(findGameWinner1(64)) // true
    fmt.Println(findGameWinner1(99)) // true
    fmt.Println(findGameWinner1(100)) // true
    fmt.Println(findGameWinner1(37)) // false
    fmt.Println(findGameWinner1(61)) // false
}