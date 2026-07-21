package main

// 3998. Transform Binary String Using Subsequence Sort
// You are given a binary string s.

// You are also given an array of strings strs, where each strs[i] has the same length as s and consists of characters '0', '1', and '?'. Each '?' can be replaced by either '0' or '1'.

// You may perform the following operation any number of times (including zero):
//     1. Choose any subsequence sub of s.
//     2. Sort sub in non-decreasing order.
//     3. Replace the chosen subsequence in s with the sorted sub, keeping all other characters unchanged.

// Return a boolean array ans, where ans[i] is true if it's possible to replace all '?' in strs[i] with '0' or '1' and transform s into the resulting string using the allowed operation above, otherwise return false.

// Example 1:
// Input: s = "101", strs = ["1?1","0?1","0?0"]
// Output: [true,true,false]
// Explanation:
// i   | Replacement | Result        | Result strs[i] | Operation(s)                                                                               | Result
// 0   | "1?1"       | ? → 0	      | "101"          | Matches s.                                                                                 | true                  
// 1   | "0?1"       | ? → 1	      | "011"          | Select the subsequence at indices [0..2] of s → "101". Sort "101" to get "011" = strs[i].	| true
// 2   | "0?0"       | ? → 0 or 1    | "000" or "010" | Not feasible.                                                                              | false                 
// Thus, ans = [true, true, false].

// Example 2:
// Input: s = "1100", strs = ["0011","11?1","1?1?"]
// Output: [true,false,true]
// Explanation:
// i   | Replacement | Result                        | Result strs[i] | Operation(s)                                                                                  | Result
// 0	|1"	  |1"	  | -	                          | "0011"	       | Select the subsequence at indices [0..3] of s → "1100". Sort "1100" to get "0011" = strs[i].  | true
// 1	| "11?1"	  | ? → 0	                      | "1101"	       | Not feasible.                                                                                 | false
// 2	| "1?1?"	  | First ? → 0 Second ? → 0	  | "1010"	       | Select the subsequence at indices [1, 2] of s → "10". Sort "10" to get "01", so s = "1010".   | true
// Thus, ans = [true, false, true].

// Example 3:
// Input: s = "1010", strs = ["0011"]
// Output: [true]
// Explanation:
// i   | Replacement | Result                        | Result strs[i] | Operation(s)                                                                                                 | Result
// 0	| "0011"	  | -	                          | "0011"	       | Select the subsequence at indices [0, 2, 3] of s → "110". Sort "110" to get "011", so s = "0011" = strs[i].  | true    
// Thus, ans = [true].

// Constraints:
//     1 <= n == s.length <= 2000
//     s[i] is either '0' or '1'.
//     1 <= strs.length <= 2000
//     strs[i].length == n
//     strs[i] is either '0', '1', or '?'​​​​​​​.

import "fmt"
import "strings"

func transformStr(s string, strs []string) []bool {
    m := len(strs)
    res := make([]bool, m)
    zero := strings.Count(s, "0")

next:
    for k, t := range strs {
        a := strings.Count(t, "0") 
        b := strings.Count(t, "?") 
        if a > zero || a + b < zero {
            continue
        }
        x, y := 0, 0
        for j := range t {
            if t[j] == '?' {
                if a < zero {
                    x++
                    a++
                }
            } else if t[j] == '0' {
                x++
            }
            if s[j] == '0' {
                y++
            }
            if x < y {
                continue next
            }
        }
        res[k] = true
    }
    return res
}

func transformStr1(s string, strs []string) []bool {
    n := len(s)
    c1 := 0
    arr1 := make([]int, n)
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            c1++
        }
        arr1[i] = c1
    }
    c0 := 0
    arr0 := make([]int, n)
    for i := n-1; i >= 0; i-- {
        if s[i] == '0' {
            c0++
        }
        arr0[i] = c0
    }
    res := make([]bool, 0, len(strs))
    for _, query := range strs {
        flag := true
        c0, c1 = 0, 0
        for i := 0; i < n; i++ {
            if query[i] == '1' {
                c1++
            }
            if c1 > arr1[i] {
                flag = false
                break
            }
        }
        if flag {
            for i := n-1; i >= 0; i-- {
                if query[i] == '0' {
                    c0++
                }
                if c0 > arr0[i] {
                    flag = false
                    break
                }
            }
        }
        res = append(res, flag)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "101", strs = ["1?1","0?1","0?0"]
    // Output: [true,true,false]
    // Explanation:
    // i   | Replacement | Result        | Result strs[i] | Operation(s)                                                                               | Result
    // 0   | "1?1"       | ? → 0	      | "101"          | Matches s.                                                                                 | true                  
    // 1   | "0?1"       | ? → 1	      | "011"          | Select the subsequence at indices [0..2] of s → "101". Sort "101" to get "011" = strs[i].	| true
    // 2   | "0?0"       | ? → 0 or 1    | "000" or "010" | Not feasible.                                                                              | false                 
    // Thus, ans = [true, true, false].
    fmt.Println(transformStr("101", []string{"1?1","0?1","0?0"})) // [true,true,false]
    // Example 2:
    // Input: s = "1100", strs = ["0011","11?1","1?1?"]
    // Output: [true,false,true]
    // Explanation:
    // i   | Replacement | Result                        | Result strs[i] | Operation(s)                                                                                  | Result
    // 0	|1"	  |1"	  | -	                          | "0011"	       | Select the subsequence at indices [0..3] of s → "1100". Sort "1100" to get "0011" = strs[i].  | true
    // 1	| "11?1"	  | ? → 0	                      | "1101"	       | Not feasible.                                                                                 | false
    // 2	| "1?1?"	  | First ? → 0 Second ? → 0	  | "1010"	       | Select the subsequence at indices [1, 2] of s → "10". Sort "10" to get "01", so s = "1010".   | true
    // Thus, ans = [true, false, true].
    fmt.Println(transformStr("1100", []string{"0011","11?1","1?1?"})) // [true,false,true]
    // Example 3:
    // Input: s = "1010", strs = ["0011"]
    // Output: [true]
    // Explanation:
    // i   | Replacement | Result                        | Result strs[i] | Operation(s)                                                                                                 | Result
    // 0	| "0011"	  | -	                          | "0011"	       | Select the subsequence at indices [0, 2, 3] of s → "110". Sort "110" to get "011", so s = "0011" = strs[i].  | true    
    // Thus, ans = [true].
    fmt.Println(transformStr("1010", []string{"0011"})) // [true]

    fmt.Println(transformStr1("101", []string{"1?1","0?1","0?0"})) // [true,true,false]
    fmt.Println(transformStr1("1100", []string{"0011","11?1","1?1?"})) // [true,false,true]
    fmt.Println(transformStr1("1010", []string{"0011"})) // [true]
}