package main

// 1718. Construct the Lexicographically Largest Valid Sequence
// Given an integer n, find a sequence that satisfies all of the following:
//     The integer 1 occurs once in the sequence.
//     Each integer between 2 and n occurs twice in the sequence.
//     For every integer i between 2 and n, the distance between the two occurrences of i is exactly i.

// The distance between two numbers on the sequence, a[i] and a[j], is the absolute difference of their indices, |j - i|.

// Return the lexicographically largest sequence. 
// It is guaranteed that under the given constraints, there is always a solution.

// A sequence a is lexicographically larger than a sequence b (of the same length) if in the first position 
// where a and b differ, sequence a has a number greater than the corresponding number in b. 
// For example, [0,1,9,0] is lexicographically larger than [0,1,5,6] because the first position they differ is at the third number, and 9 is greater than 5.

// Example 1:
// Input: n = 3
// Output: [3,1,2,3,2]
// Explanation: [2,3,2,1,3] is also a valid sequence, but [3,1,2,3,2] is the lexicographically largest valid sequence.

// Example 2:
// Input: n = 5
// Output: [5,3,1,4,3,5,2,4,2]

// Constraints:
//     1 <= n <= 20

import "fmt"

func constructDistancedSequence(n int) []int {
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = n - i
    }
    res := make([]int, n * 2 - 1)
    var dfs func(seq []int, arr []int) bool
    dfs = func(seq []int, arr []int) bool {
        i := 0
        for i < len(seq) {
            if seq[i] == 0 { break }
            i++
        }
        if i == len(seq) { return true }
        for j, v := range arr {
            if v != 0 && (v == 1 || (i + v < len(seq) && seq[i + v] == 0)) {
                arr[j] = 0
                if v != 1 {
                    seq[i], seq[i + v] = v, v
                } else {
                    seq[i] = v
                }
                if dfs(seq, arr) { return true }
                arr[j] = v
                if v != 1 {
                    seq[i], seq[i + v] = 0, 0
                } else {
                    seq[i] = 0
                }
            }
        }
        return false
    }
    dfs(res, arr)
    return res
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: [3,1,2,3,2]
    // Explanation: [2,3,2,1,3] is also a valid sequence, but [3,1,2,3,2] is the lexicographically largest valid sequence.
    fmt.Println(constructDistancedSequence(3)) // [3,1,2,3,2]
    // Example 2:
    // Input: n = 5
    // Output: [5,3,1,4,3,5,2,4,2]
    fmt.Println(constructDistancedSequence(5)) // [5,3,1,4,3,5,2,4,2]

    fmt.Println(constructDistancedSequence(1)) // [1]
    fmt.Println(constructDistancedSequence(2)) // [2 1 2]
    fmt.Println(constructDistancedSequence(19)) // [19 17 18 14 12 16 9 15 6 3 13 1 3 11 6 9 12 14 17 19 18 16 15 13 11 10 8 4 5 7 2 4 2 5 8 10 7]
    fmt.Println(constructDistancedSequence(20)) // [20 18 19 15 13 17 10 16 7 5 3 14 12 3 5 7 10 13 15 18 20 19 17 16 12 14 11 9 4 6 8 2 4 2 1 6 9 11 8]
}