package main

// 1769. Minimum Number of Operations to Move All Balls to Each Box
// You have n boxes. 
// You are given a binary string boxes of length n, 
// where boxes[i] is '0' if the ith box is empty, and '1' if it contains one ball.

// In one operation, you can move one ball from a box to an adjacent box. 
// Box i is adjacent to box j if abs(i - j) == 1. 
// Note that after doing so, there may be more than one ball in some boxes.

// Return an array answer of size n, 
// where answer[i] is the minimum number of operations needed to move all the balls to the ith box.

// Each answer[i] is calculated considering the initial state of the boxes.

// Example 1:
// Input: boxes = "110"
// Output: [1,1,3]
// Explanation: The answer for each box is as follows:
// 1) First box: you will have to move one ball from the second box to the first box in one operation.
// 2) Second box: you will have to move one ball from the first box to the second box in one operation.
// 3) Third box: you will have to move one ball from the first box to the third box in two operations, and move one ball from the second box to the third box in one operation.

// Example 2:
// Input: boxes = "001011"
// Output: [11,8,5,4,3,4]

// Constraints:
//     n == boxes.length
//     1 <= n <= 2000
//     boxes[i] is either '0' or '1'.

import "fmt"

func minOperations(boxes string) []int {
    mp, n := make(map[int]bool), len(boxes)
    for i := 0; i < n; i++ {
        if boxes[i] == '1' { mp[i] = true }
    }
    res := make([]int, n)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < n; i++ {
        moves := 0
        for k, _ := range mp {
            if k == i { continue }
            moves += abs(k - i)
        }
        res[i] = moves
    }
    return res
}

func minOperations1(boxes string) []int {
    n := len(boxes)
    res := make([]int,n)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, v := range boxes {
        if v == '0' { continue }
        for j := 0; j < n; j++ {
            res[j] += abs(i- j)
        }
    }
    return res
}

func minOperations2(boxes string) []int {
    n, moves, count := len(boxes), 0, 0
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = moves
        if boxes[i] == '1' { count++ }
        moves += count
    }
    moves, count = 0, 0
    for i := n - 1; i >= 0; i-- {
        res[i] += moves
        if boxes[i] == '1' { count++ }
        moves += count
    }
    return res
}

func main() {
    // Example 1:
    // Input: boxes = "110"
    // Output: [1,1,3]
    // Explanation: The answer for each box is as follows:
    // 1) First box: you will have to move one ball from the second box to the first box in one operation.
    // 2) Second box: you will have to move one ball from the first box to the second box in one operation.
    // 3) Third box: you will have to move one ball from the first box to the third box in two operations, and move one ball from the second box to the third box in one operation.
    fmt.Println(minOperations("110")) // [1,1,3]
    // Example 2:
    // Input: boxes = "001011"
    // Output: [11,8,5,4,3,4]
    fmt.Println(minOperations("001011")) // [11,8,5,4,3,4]

    fmt.Println(minOperations1("110")) // [1,1,3]
    fmt.Println(minOperations1("001011")) // [11,8,5,4,3,4]

    fmt.Println(minOperations2("110")) // [1,1,3]
    fmt.Println(minOperations2("001011")) // [11,8,5,4,3,4]
}