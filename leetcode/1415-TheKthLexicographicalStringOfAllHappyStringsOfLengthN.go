package main

// 1415. The k-th Lexicographical String of All Happy Strings of Length n
// A happy string is a string that:
//     consists only of letters of the set ['a', 'b', 'c'].
//     s[i] != s[i + 1] for all values of i from 1 to s.length - 1 (string is 1-indexed).

// For example, strings "abc", "ac", "b" and "abcbabcbcb" are all happy strings 
// and strings "aa", "baa" and "ababbc" are not happy strings.

// Given two integers n and k, consider a list of all happy strings of length n sorted in lexicographical order.

// Return the kth string of this list or return an empty string if there are less than k happy strings of length n.

// Example 1:
// Input: n = 1, k = 3
// Output: "c"
// Explanation: The list ["a", "b", "c"] contains all happy strings of length 1. The third string is "c".

// Example 2:
// Input: n = 1, k = 4
// Output: ""
// Explanation: There are only 3 happy strings of length 1.

// Example 3:
// Input: n = 3, k = 9
// Output: "cab"
// Explanation: There are 12 different happy string of length 3 ["aba", "abc", "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"]. You will find the 9th string = "cab"

// Constraints:
//     1 <= n <= 10
//     1 <= k <= 100

import "fmt"
import "sort"

func getHappyString(n int, k int) string {
    choices, memo := []string{"a", "b", "c"}, []string{}
    var gen func(n int, str string) 
    gen = func (n int, str string) {
        if n == 0{
            memo = append(memo, str)
            return 
        }
        for i := 0; i< len(choices); i++ {
            if str == "" || string(str[len(str)-1]) != choices[i] {
                gen(n - 1, str + choices[i])
            }
        }
    }
    gen(n, "")
    if k > len(memo) { return "" }
    sort.Strings(memo)
    return memo[k-1]
}

func getHappyString1(n int, k int) string {
    pow2 := func (i int) int {
        res := 1
        for i > 0 {
            res *= 2
            i--
        }
        return res
    }
    if pow2(n-1) * 3 < k { return "" }
    s := make([]byte, n)
    var dfs func(i int)
    dfs = func(i int) {
        if i >= n { return }
        t := pow2(n - i - 1)
        if s[i-1] == 'a' {
            if t >= k {
                s[i] = 'b'
            } else {
                s[i] = 'c'
                k -= t
            }
            dfs(i + 1)
            return
        }
        if s[i-1] == 'b' {
            if t >= k {
                s[i] = 'a'
            } else {
                s[i] = 'c'
                k -= t
            }
            dfs(i + 1)
            return
        }
        if s[i-1] == 'c' {
            if t >= k {
                s[i] = 'a'
            } else {
                s[i] = 'b'
                k -= t
            }
            dfs(i + 1)
            return
        }
    }
    if pow2(n-1) >= k {
        s[0] = 'a'
        dfs(1)
    } else if 2*pow2(n-1) >= k {
        s[0] = 'b'
        k -= pow2(n - 1)
        dfs(1)
    } else {
        s[0] = 'c'
        k -= 2 * pow2(n-1)
        dfs(1)
    }
    return string(s)
}

func main() {
    // Example 1:
    // Input: n = 1, k = 3
    // Output: "c"
    // Explanation: The list ["a", "b", "c"] contains all happy strings of length 1. The third string is "c".
    fmt.Println(getHappyString(1,3)) // "c"
    // Example 2:
    // Input: n = 1, k = 4
    // Output: ""
    // Explanation: There are only 3 happy strings of length 1.
    fmt.Println(getHappyString(1,4)) // ""
    // Example 3:
    // Input: n = 3, k = 9
    // Output: "cab"
    // Explanation: There are 12 different happy string of length 3 ["aba", "abc", "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"]. You will find the 9th string = "cab"
    fmt.Println(getHappyString(3,9)) // cab

    fmt.Println(getHappyString1(1,3)) // "c"
    fmt.Println(getHappyString1(1,4)) // ""
    fmt.Println(getHappyString1(3,9)) // cab
}