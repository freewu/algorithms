package main

// 960. Delete Columns to Make Sorted III
// You are given an array of n strings strs, all of the same length.

// We may choose any deletion indices, and we delete all the characters in those indices for each string.

// For example, if we have strs = ["abcdef","uvwxyz"] and deletion indices {0, 2, 3}, then the final array after deletions is ["bef", "vyz"].

// Suppose we chose a set of deletion indices answer such that after deletions, the final array has every string (row) in lexicographic order. 
// (i.e., (strs[0][0] <= strs[0][1] <= ... <= strs[0][strs[0].length - 1]), and (strs[1][0] <= strs[1][1] <= ... <= strs[1][strs[1].length - 1]), and so on). 
// Return the minimum possible value of answer.length.

// Example 1:
// Input: strs = ["babca","bbazb"]
// Output: 3
// Explanation: After deleting columns 0, 1, and 4, the final array is strs = ["bc", "az"].
// Both these rows are individually in lexicographic order (ie. strs[0][0] <= strs[0][1] and strs[1][0] <= strs[1][1]).
// Note that strs[0] > strs[1] - the array strs is not necessarily in lexicographic order.

// Example 2:
// Input: strs = ["edcba"]
// Output: 4
// Explanation: If we delete less than 4 columns, the only row will not be lexicographically sorted.

// Example 3:
// Input: strs = ["ghi","def","abc"]
// Output: 0
// Explanation: All rows are already lexicographically sorted.

// Constraints:
//     n == strs.length
//     1 <= n <= 100
//     1 <= strs[i].length <= 100
//     strs[i] consists of lowercase English letters.

import "fmt"

func minDeletionSize(strs []string) int {
    deletionsToColumn := make([]int, len(strs[0]) + 1) // deletionsToColumn is 1 place longer, to accomodate for the very important End node
    // this loop simulates going from our placeholder start node to every other node (the cost is just the column index)
    // also there will always be a path from S to any other node, it translates to where we choose to start
    for i := 0; i < len(deletionsToColumn); i++ {
        deletionsToColumn[i] = i
    }
    // hasRoute checks that all the chars in the column (c) are in order with the chars in the next column (nc)
    hasRoute := func(strs []string, c, nc int) bool {
        if c == -1 || nc == len(strs[0]) { return true } // if nc is out of bounds, that means it's the end node, to which all nodes are free to reach (at a cost!)
        for si := 0; si < len(strs); si++ {
            if strs[si][c] > strs[si][nc] {
                return false
            }
        }    
        return true
    }
    // I guess there was no real need for me to split to another function, just a remainder of my previous DFS attempt
    minDeletionsRoute := func (strs []string, deletionsToColumn []int) int {
        for column := 0; column < len(strs[0]); column++ {
            for nextColumn := column + 1; nextColumn < len(deletionsToColumn); nextColumn++ {
                // it's cheaper (computationally) to check if the cost of the edge is cheaper than the previous best way to arrive to it, so that condition comes before the call to hasRoute
                if deletionsToColumn[nextColumn] >= deletionsToColumn[column] + nextColumn - column - 1 && 
                   hasRoute(strs, column, nextColumn) {
                    deletionsToColumn[nextColumn] = deletionsToColumn[column] + nextColumn - column - 1
                }
            }
        }
        return deletionsToColumn[len(deletionsToColumn)-1]
    }
    return minDeletionsRoute(strs, deletionsToColumn)
}

func main() {
    // Example 1:
    // Input: strs = ["babca","bbazb"]
    // Output: 3
    // Explanation: After deleting columns 0, 1, and 4, the final array is strs = ["bc", "az"].
    // Both these rows are individually in lexicographic order (ie. strs[0][0] <= strs[0][1] and strs[1][0] <= strs[1][1]).
    // Note that strs[0] > strs[1] - the array strs is not necessarily in lexicographic order.
    fmt.Println(minDeletionSize([]string{"babca","bbazb"})) // 3
    // Example 2:
    // Input: strs = ["edcba"]
    // Output: 4
    // Explanation: If we delete less than 4 columns, the only row will not be lexicographically sorted.
    fmt.Println(minDeletionSize([]string{"edcba"})) // 4
    // Example 3:
    // Input: strs = ["ghi","def","abc"]
    // Output: 0
    // Explanation: All rows are already lexicographically sorted.
    fmt.Println(minDeletionSize([]string{"ghi","def","abc"})) // 0
}