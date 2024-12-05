package main

// 2337. Move Pieces to Obtain a String
// You are given two strings start and target, both of length n. 
// Each string consists only of the characters 'L', 'R', and '_' where:
//     1. The characters 'L' and 'R' represent pieces, where a piece 'L' can move to the left only if there is a blank space directly to its left, 
//        and a piece 'R' can move to the right only if there is a blank space directly to its right.
//     2. The character '_' represents a blank space that can be occupied by any of the 'L' or 'R' pieces.

// Return true if it is possible to obtain the string target by moving the pieces of the string start any number of times. 
// Otherwise, return false.

// Example 1:
// Input: start = "_L__R__R_", target = "L______RR"
// Output: true
// Explanation: We can obtain the string target from start by doing the following moves:
// - Move the first piece one step to the left, start becomes equal to "L___R__R_".
// - Move the last piece one step to the right, start becomes equal to "L___R___R".
// - Move the second piece three steps to the right, start becomes equal to "L______RR".
// Since it is possible to get the string target from start, we return true.

// Example 2:
// Input: start = "R_L_", target = "__LR"
// Output: false
// Explanation: The 'R' piece in the string start can move one step to the right to obtain "_RL_".
// After that, no pieces can move anymore, so it is impossible to obtain the string target from start.

// Example 3:
// Input: start = "_R", target = "R_"
// Output: false
// Explanation: The piece in the string start can move only to the right, so it is impossible to obtain the string target from start.

// Constraints:
//     n == start.length == target.length
//     1 <= n <= 10^5
//     start and target consist of the characters 'L', 'R', and '_'.

import "fmt"

func canChange(start string, target string) bool {
    hepler := func(s string) [][]int {
        res := [][]int{}
        for i, c := range s {
            if c == 'L' {
                res = append(res, []int{1, i})
            } else if c == 'R' {
                res = append(res, []int{2, i})
            }
        }
        return res
    }
    a, b := hepler(start), hepler(target)
    if len(a) != len(b) { return false }
    for i, x := range a {
        y := b[i]
        if x[0] != y[0] { return false }
        if x[0] == 1 && x[1] < y[1] { return false }
        if x[0] == 2 && x[1] > y[1] { return false }
    }
    return true
}

func canChange1(start string, target string) bool {
    n, j := len(target), 0
    for i, v := range target {
        if v == 'L' {
            for j < n && start[j] == '_' { j++ }
            if j == n || start[j] != target[i] { return false }
            if i > j { return false }
            j++
        } else if v == 'R' {
            for j < n && start[j] == '_' { j++ }
            if j == n || start[j] != target[i] { return false }
            if i < j { return false }
            j++
        }
    }
    for ; j < n; j++ {
        if start[j] != '_' { return false }
    }
    return true
}

func canChange2(start string, target string) bool {
    n := len(target)
    for i, j := 0, 0; i < n || j < n; i, j = i + 1, j + 1 {
        for i < n && start[i] == '_' { i++ }
        for j < n && target[j] == '_' { j++ }
        if i == n || j == n {  return i == n && j == n }
        if start[i] != target[j] { return false }
        if (start[i] == 'L' && i < j) || (start[i] == 'R' && i > j) { return false }
    }
    return true
}

func main() {
    // Example 1:
    // Input: start = "_L__R__R_", target = "L______RR"
    // Output: true
    // Explanation: We can obtain the string target from start by doing the following moves:
    // - Move the first piece one step to the left, start becomes equal to "L___R__R_".
    // - Move the last piece one step to the right, start becomes equal to "L___R___R".
    // - Move the second piece three steps to the right, start becomes equal to "L______RR".
    // Since it is possible to get the string target from start, we return true.
    fmt.Println(canChange("_L__R__R_", "L______RR")) // true
    // Example 2:
    // Input: start = "R_L_", target = "__LR"
    // Output: false
    // Explanation: The 'R' piece in the string start can move one step to the right to obtain "_RL_".
    // After that, no pieces can move anymore, so it is impossible to obtain the string target from start.
    fmt.Println(canChange("R_L_", "__LR")) // false
    // Example 3:
    // Input: start = "_R", target = "R_"
    // Output: false
    // Explanation: The piece in the string start can move only to the right, so it is impossible to obtain the string target from start.
    fmt.Println(canChange("_R", "R_")) // false

    fmt.Println(canChange1("_L__R__R_", "L______RR")) // true
    fmt.Println(canChange1("R_L_", "__LR")) // false
    fmt.Println(canChange1("_R", "R_")) // false

    fmt.Println(canChange2("_L__R__R_", "L______RR")) // true
    fmt.Println(canChange2("R_L_", "__LR")) // false
    fmt.Println(canChange2("_R", "R_")) // false
}