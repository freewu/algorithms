package main

// 2998. Minimum Number of Operations to Make X and Y Equal
// You are given two positive integers x and y.

// In one operation, you can do one of the four following operations:
//     1. Divide x by 11 if x is a multiple of 11.
//     2. Divide x by 5 if x is a multiple of 5.
//     3. Decrement x by 1.
//     4. Increment x by 1.

// Return the minimum number of operations required to make x and y equal.

// Example 1:
// Input: x = 26, y = 1
// Output: 3
// Explanation: We can make 26 equal to 1 by applying the following operations: 
// 1. Decrement x by 1
// 2. Divide x by 5
// 3. Divide x by 5
// It can be shown that 3 is the minimum number of operations required to make 26 equal to 1.

// Example 2:
// Input: x = 54, y = 2
// Output: 4
// Explanation: We can make 54 equal to 2 by applying the following operations: 
// 1. Increment x by 1
// 2. Divide x by 11 
// 3. Divide x by 5
// 4. Increment x by 1
// It can be shown that 4 is the minimum number of operations required to make 54 equal to 2.

// Example 3:
// Input: x = 25, y = 30
// Output: 5
// Explanation: We can make 25 equal to 30 by applying the following operations: 
// 1. Increment x by 1
// 2. Increment x by 1
// 3. Increment x by 1
// 4. Increment x by 1
// 5. Increment x by 1
// It can be shown that 5 is the minimum number of operations required to make 25 equal to 30.

// Constraints:
//     1 <= x, y <= 10^4

import "fmt"

// bfs
func minimumOperationsToMakeEqual(x int, y int) int {
    queue, visited := [][]int{ { x, 0 } }, make(map[int]bool)
    for len(queue) > 0 {
        pair := queue[0]
        queue = queue[1:]
        val, level := pair[0], pair[1]
        if val == y {  return level }
        if visited[val] { continue }
        visited[val] = true
        if val % 11 == 0 {
            queue = append(queue, []int{ val / 11,level + 1 })
        }
        if val % 5 == 0 {
            queue = append(queue, []int{ val / 5, level + 1 })
        }
        queue = append(queue, []int{ val + 1, level + 1 })
        queue = append(queue, []int{ val - 1, level + 1 })
    }
    return -1
}

// dfs
func minimumOperationsToMakeEqual1(x int, y int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(v int) int
    dfs = func(v int) int {
        if v <= y { return y - v }
        res, p, q := v - y, v % 5, v % 11
        res = min(res, p  + 1 + dfs((v - p) / 5))
        res = min(res, q  + 1 + dfs((v - q) / 11))
        res = min(res, 5  - p + 1 + dfs((v + 5  - p) / 5))
        res = min(res, 11 - q + 1 + dfs((v + 11 - q) / 11))
        return res
    }
    return dfs(x)
}


func main() {
    // Example 1:
    // Input: x = 26, y = 1
    // Output: 3
    // Explanation: We can make 26 equal to 1 by applying the following operations: 
    // 1. Decrement x by 1
    // 2. Divide x by 5
    // 3. Divide x by 5
    // It can be shown that 3 is the minimum number of operations required to make 26 equal to 1.
    fmt.Println(minimumOperationsToMakeEqual(26, 1)) // 3
    // Example 2:
    // Input: x = 54, y = 2
    // Output: 4
    // Explanation: We can make 54 equal to 2 by applying the following operations: 
    // 1. Increment x by 1
    // 2. Divide x by 11 
    // 3. Divide x by 5
    // 4. Increment x by 1
    // It can be shown that 4 is the minimum number of operations required to make 54 equal to 2.
    fmt.Println(minimumOperationsToMakeEqual(54, 2)) // 4
    // Example 3:
    // Input: x = 25, y = 30
    // Output: 5
    // Explanation: We can make 25 equal to 30 by applying the following operations: 
    // 1. Increment x by 1
    // 2. Increment x by 1
    // 3. Increment x by 1
    // 4. Increment x by 1
    // 5. Increment x by 1
    // It can be shown that 5 is the minimum number of operations required to make 25 equal to 30.
    fmt.Println(minimumOperationsToMakeEqual(25, 30)) // 5

    fmt.Println(minimumOperationsToMakeEqual(1, 1)) // 0
    fmt.Println(minimumOperationsToMakeEqual(1, 10000)) // 9999
    fmt.Println(minimumOperationsToMakeEqual(10000, 1)) // 8
    fmt.Println(minimumOperationsToMakeEqual(10000, 10000)) // 0
    
    fmt.Println(minimumOperationsToMakeEqual1(26, 1)) // 3
    fmt.Println(minimumOperationsToMakeEqual1(54, 2)) // 4
    fmt.Println(minimumOperationsToMakeEqual1(25, 30)) // 5
    fmt.Println(minimumOperationsToMakeEqual1(1, 1)) // 0
    fmt.Println(minimumOperationsToMakeEqual1(1, 10000)) // 9999
    fmt.Println(minimumOperationsToMakeEqual1(10000, 1)) // 8
    fmt.Println(minimumOperationsToMakeEqual1(10000, 10000)) // 0
}