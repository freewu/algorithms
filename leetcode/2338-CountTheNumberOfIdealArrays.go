package main

// 2338. Count the Number of Ideal Arrays
// You are given two integers n and maxValue, which are used to describe an ideal array.

// A 0-indexed integer array arr of length n is considered ideal if the following conditions hold:
//     Every arr[i] is a value from 1 to maxValue, for 0 <= i < n.
//     Every arr[i] is divisible by arr[i - 1], for 0 < i < n.

// Return the number of distinct ideal arrays of length n. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 2, maxValue = 5
// Output: 10
// Explanation: The following are the possible ideal arrays:
// - Arrays starting with the value 1 (5 arrays): [1,1], [1,2], [1,3], [1,4], [1,5]
// - Arrays starting with the value 2 (2 arrays): [2,2], [2,4]
// - Arrays starting with the value 3 (1 array): [3,3]
// - Arrays starting with the value 4 (1 array): [4,4]
// - Arrays starting with the value 5 (1 array): [5,5]
// There are a total of 5 + 2 + 1 + 1 + 1 = 10 distinct ideal arrays.

// Example 2:
// Input: n = 5, maxValue = 3
// Output: 11
// Explanation: The following are the possible ideal arrays:
// - Arrays starting with the value 1 (9 arrays): 
//    - With no other distinct values (1 array): [1,1,1,1,1] 
//    - With 2nd distinct value 2 (4 arrays): [1,1,1,1,2], [1,1,1,2,2], [1,1,2,2,2], [1,2,2,2,2]
//    - With 2nd distinct value 3 (4 arrays): [1,1,1,1,3], [1,1,1,3,3], [1,1,3,3,3], [1,3,3,3,3]
// - Arrays starting with the value 2 (1 array): [2,2,2,2,2]
// - Arrays starting with the value 3 (1 array): [3,3,3,3,3]
// There are a total of 9 + 1 + 1 = 11 distinct ideal arrays.

// Constraints:
//     2 <= n <= 10^4
//     1 <= maxValue <= 10^4

import "fmt"

func idealArrays(n int, maxValue int) int {
    res, mod := 0, 1_000_000_007
    count := make([][]int, n)
    facts := make([][]int, maxValue + 1)
    for i := range count {
        count[i] = make([]int, 16)
    }
    for i := range facts {
        facts[i] = make([]int, 16)
        for j := range facts[i] {
            facts[i][j] = -1
        }
    }
    var dfs func(i, c int) int
    dfs = func(i, c int) int {
        if facts[i][c] != -1 { return facts[i][c] }
        res := count[n-1][c - 1]
        if c < n {
            for k := 2; k * i <= maxValue; k++ {
                res = (res + dfs(k * i, c + 1)) % mod
            }
        }
        facts[i][c] = res
        return res
    }
    for i := 0; i < n; i++ {
        for j := 0; j <= i && j < 16; j++ {
            if j == 0 {
                count[i][j] = 1
            } else {
                count[i][j] = (count[i-1][j] + count[i-1][j-1]) % mod
            }
        }
    }
    for i := 1; i <= maxValue; i++ {
        res = (res + dfs(i, 1)) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, maxValue = 5
    // Output: 10
    // Explanation: The following are the possible ideal arrays:
    // - Arrays starting with the value 1 (5 arrays): [1,1], [1,2], [1,3], [1,4], [1,5]
    // - Arrays starting with the value 2 (2 arrays): [2,2], [2,4]
    // - Arrays starting with the value 3 (1 array): [3,3]
    // - Arrays starting with the value 4 (1 array): [4,4]
    // - Arrays starting with the value 5 (1 array): [5,5]
    // There are a total of 5 + 2 + 1 + 1 + 1 = 10 distinct ideal arrays.
    fmt.Println(idealArrays(2, 5)) // 10
    // Example 2:
    // Input: n = 5, maxValue = 3
    // Output: 11
    // Explanation: The following are the possible ideal arrays:
    // - Arrays starting with the value 1 (9 arrays): 
    //    - With no other distinct values (1 array): [1,1,1,1,1] 
    //    - With 2nd distinct value 2 (4 arrays): [1,1,1,1,2], [1,1,1,2,2], [1,1,2,2,2], [1,2,2,2,2]
    //    - With 2nd distinct value 3 (4 arrays): [1,1,1,1,3], [1,1,1,3,3], [1,1,3,3,3], [1,3,3,3,3]
    // - Arrays starting with the value 2 (1 array): [2,2,2,2,2]
    // - Arrays starting with the value 3 (1 array): [3,3,3,3,3]
    // There are a total of 9 + 1 + 1 = 11 distinct ideal arrays.
    fmt.Println(idealArrays(5, 3)) // 11

    fmt.Println(idealArrays(2, 1)) // 1
    fmt.Println(idealArrays(2, 10000)) // 93668
    fmt.Println(idealArrays(10000, 1)) // 1
    fmt.Println(idealArrays(10000, 10000)) // 22940607
}