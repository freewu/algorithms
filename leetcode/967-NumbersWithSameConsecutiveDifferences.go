package main

// 967. Numbers With Same Consecutive Differences
// Given two integers n and k, 
// return an array of all the integers of length n where the difference between every two consecutive digits is k.
// You may return the answer in any order.

// Note that the integers should not have leading zeros. 
// Integers as 02 and 043 are not allowed.

// Example 1:
// Input: n = 3, k = 7
// Output: [181,292,707,818,929]
// Explanation: Note that 070 is not a valid number, because it has leading zeroes.

// Example 2:
// Input: n = 2, k = 1
// Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]

// Constraints:
//     2 <= n <= 9
//     0 <= k <= 9

import "fmt"

func numsSameConsecDiff(n int, k int) []int {
    cur := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    for i := 0; i < n - 1; i++ {
        t := []int{}
        for _, v := range cur {
            a := v % 10
            if a + k< 10 { 
                t = append(t, v * 10 + a + k) 
            }
            if a >= k && k > 0 { 
                t = append(t, v * 10 + a - k) 
            }
        }
        cur = t
    }
    return cur
}

// dfs
func numsSameConsecDiff1(n int, k int) []int {
    res := []int{}
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func(t, i int) 
    dfs = func(t, i int) {
        if i == 1 {
            res = append(res, t)
            return 
        }
        for j := 0; j < 10; j++ {
            if abs(j - t % 10) == k {
                dfs(t * 10 + j, i - 1)
            }
        }
    }
    for t := 1; t < 10; t++ {
        dfs(t,n)
    }
    return res 
}

func main() {
    // Example 1:
    // Input: n = 3, k = 7
    // Output: [181,292,707,818,929]
    // Explanation: Note that 070 is not a valid number, because it has leading zeroes.
    fmt.Println(numsSameConsecDiff(3, 7)) // [181,292,707,818,929]
    // Example 2:
    // Input: n = 2, k = 1
    // Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]
    fmt.Println(numsSameConsecDiff(2, 1)) // [181,292,707,818,929]

    fmt.Println(numsSameConsecDiff(2, 0)) // [11 22 33 44 55 66 77 88 99]
    fmt.Println(numsSameConsecDiff(9, 9)) // [909090909]

    fmt.Println(numsSameConsecDiff1(3, 7))// [181,292,707,818,929]
    fmt.Println(numsSameConsecDiff1(2, 1)) // [181,292,707,818,929]
    fmt.Println(numsSameConsecDiff1(2, 0)) // [11 22 33 44 55 66 77 88 99]
    fmt.Println(numsSameConsecDiff1(9, 9)) // [909090909]
}