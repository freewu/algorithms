package main

// 942. DI String Match
// A permutation perm of n + 1 integers of all the integers in the range [0, n] can be represented as a string s of length n where:
//     s[i] == 'I' if perm[i] < perm[i + 1], and
//     s[i] == 'D' if perm[i] > perm[i + 1].

// Given a string s, reconstruct the permutation perm and return it. 
// If there are multiple valid permutations perm, return any of them.

// Example 1:
// Input: s = "IDID"
// Output: [0,4,1,3,2]

// Example 2:
// Input: s = "III"
// Output: [0,1,2,3]

// Example 3:
// Input: s = "DDI"
// Output: [3,2,0,1]

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either 'I' or 'D'.

import "fmt"

func diStringMatch(s string) []int {
    res, n := []int{}, len(s)
    i, d := 0, n
    for j := 0; j < n; j++ {
        if s[j] == 'I' {
            res = append(res, i)
            i++
        } else if s[j] == 'D' {
            res = append(res, d)
            d--
        }
        if j == n - 1 {
            if s[j] == 'I' {
                res = append(res, i)
                i++
            } else if s[j] == 'D' {
                res = append(res, d)
                d--
            }
        }
    }
    return res
}

// 贪心算法
func diStringMatch1(s string) []int {
    arr, res := []int{}, []int{}
    for i := 0; i <= len(s); i++ {
        arr = append(arr, i)
    }
    start, end, final := 0, len(arr) - 1, 0
    for i := 0; i < len(s); i++ {
        if s[i] == 'I' {
            res = append(res, arr[start])
            start++
            final = start
        } else {
            res = append(res, arr[end])
            end--
            final = end
        }
    }
    res = append(res, final)
    return res
}

func main() {
    // Example 1:
    // Input: s = "IDID"
    // Output: [0,4,1,3,2] 
    fmt.Println(diStringMatch("IDID")) // [0,4,1,3,2]
    // Example 2:
    // Input: s = "III"
    // Output: [0,1,2,3]
    fmt.Println(diStringMatch("III")) // [0,1,2,3]
    // Example 3:
    // Input: s = "DDI"
    // Output: [3,2,0,1]
    fmt.Println(diStringMatch("DDI")) // [3,2,0,1] 

    fmt.Println(diStringMatch1("IDID")) // [0,4,1,3,2]
    fmt.Println(diStringMatch1("III")) // [0,1,2,3]
    fmt.Println(diStringMatch1("DDI")) // [3,2,0,1] 
}