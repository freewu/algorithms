package main

// 3211. Generate Binary Strings Without Adjacent Zeros
// You are given a positive integer n.

// A binary string x is valid if all substrings of x of length 2 contain at least one "1".

// Return all valid strings with length n, in any order.

// Example 1:
// Input: n = 3
// Output: ["010","011","101","110","111"]
// Explanation:
// The valid strings of length 3 are: "010", "011", "101", "110", and "111".

// Example 2:
// Input: n = 1
// Output: ["0","1"]
// Explanation:
// The valid strings of length 1 are: "0" and "1".

// Constraints:
//     1 <= n <= 18

import "fmt"

func validStrings(n int) []string {
    arr := []string{ "0", "1" }
    for i := 1; i < n; i++ { // length
        tmp := []string{}
        for _, s := range arr {
            if s[len(s) - 1] == '1' {
                tmp = append(tmp, s + "0")
            }
            tmp = append(tmp, s + "1")
        } 
        arr = tmp
    }
    return arr
}

func validStrings1(n int) []string {
    res := []string{}
    elements := make([]byte, n)
    var dfs func(cur int)
    dfs = func(cur int) {
        if cur == n {
            res = append(res, string(elements))
            return
        }
        if (cur-1 < 0) || (cur-1 >= 0 && elements[cur-1] == '1') {
            elements[cur] = '0'
            dfs(cur + 1)
        }
        elements[cur] = '1'
        dfs(cur + 1)
    }
    dfs(0)
    return res
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: ["010","011","101","110","111"]
    // Explanation:
    // The valid strings of length 3 are: "010", "011", "101", "110", and "111".
    fmt.Println(validStrings(3)) // ["010","011","101","110","111"]
    // Example 2:
    // Input: n = 1
    // Output: ["0","1"]
    // Explanation:
    // The valid strings of length 1 are: "0" and "1".
    fmt.Println(validStrings(1)) // ["0","1"]

    fmt.Println(validStrings(2)) // [01 10 11]
    fmt.Println(validStrings(4)) // [0101 0110 0111 1010 1011 1101 1110 1111]
    fmt.Println(validStrings(5)) // [01010 01011 01101 01110 01111 10101 10110 10111 11010 11011 11101 11110 11111]

    fmt.Println(validStrings1(3)) // ["010","011","101","110","111"]
    fmt.Println(validStrings1(1)) // ["0","1"]
    fmt.Println(validStrings1(2)) // [01 10 11]
    fmt.Println(validStrings1(4)) // [0101 0110 0111 1010 1011 1101 1110 1111]
    fmt.Println(validStrings1(5)) // [01010 01011 01101 01110 01111 10101 10110 10111 11010 11011 11101 11110 11111]
}
