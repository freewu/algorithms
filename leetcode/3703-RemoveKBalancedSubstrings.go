package main

// 3703. Remove K-Balanced Substrings
// You are given a string s consisting of '(' and ')', and an integer k.

// A string is k-balanced if it is exactly k consecutive '(' followed by k consecutive ')', i.e., '(' * k + ')' * k.

// For example, if k = 3, k-balanced is "((()))".

// You must repeatedly remove all non-overlapping k-balanced substrings from s, and then join the remaining parts. 
// Continue this process until no k-balanced substring exists.

// Return the final string after all possible removals.

// A substring is a contiguous non-empty sequence of characters within a string.

// ​​​​​​​Example 1:
// Input: s = "(())", k = 1
// Output: ""
// Explanation:
// k-balanced substring is "()"
// Step	Current s	k-balanced	Result s
// 1	        (())	    (())	()
// 2	        ()	        ()	    Empty
// Thus, the final string is "".

// Example 2:
// Input: s = "(()(", k = 1
// Output: "(("
// Explanation:
// k-balanced substring is "()"
// Step	Current s	k-balanced	Result s
//     1	    (()(	(()(	      ((
//     2	    ((	     -	          ((
// Thus, the final string is "((".

// Example 3:
// Input: s = "((()))()()()", k = 3
// Output: "()()()"
// Explanation:
// k-balanced substring is "((()))"
// Step	Current s	    k-balanced	    Result s
// 1	    ((()))()()()	((()))()()()	()()()
// 2	    ()()()	        -	            ()()()
// Thus, the final string is "()()()".

// Constraints:
//     2 <= s.length <= 10^5
//     s consists only of '(' and ')'.
//     1 <= k <= s.length / 2

import "fmt"

// time: O(n), space: O(n)
func removeSubstring(s string, k int) string {
    res, n := []byte{}, len(s)
    check := func(i int) byte {
        if 0 <= i && i < len(res) {
            return res[i]
        }
        return 0
    }
    left, right := 0, 0
    for si, ri := 0, 0; si <= n; si, ri = si + 1, ri + 1 {
        if si < n {
            res = append(res, s[si])
        }
        if check(ri) == ')' {
            right++
        }
        if _c := check(ri-k); _c == ')' {
            right--
        } else if _c == '(' {
            left++
        }
        if check(ri-k*2) == '(' {
            left--
        }
        if left == k && right == k {
            l := len(res) - k * 2
            res = res[:l]
            ri = l - 1
            left, right = 0, 0
            for j, e := ri, max(0, l - k); j >= e; j-- {
                if check(j) == ')' {
                    right++
                }
                if check(j - k) == '(' {
                    left++
                }
            }
        }
    }
    return string(res)
}

func removeSubstring1(s string, k int) string {
    stack, cl , cr := []byte{}, 0, 0
    for i := range s {
        x := s[i]
        if x == '('{
            if len(stack) > 0 && stack[len(stack) - 1] == '('{
                cl++
            } else {
                cl = 1
            }
            stack = append(stack,x)
        } else {
            if len(stack)>0 && stack[len(stack) - 1] == ')' {
                cr++
            } else {
                cr = 1
            }
            stack = append(stack,x)
            if cr >= k && cl >= k {
                stack = stack[:len(stack) - 2 * k]
                cl -= k
                cr -= k
                if cl == 0 {
                    for j := len(stack) - 1; j >= 0; j-- {
                        y := stack[j]
                        if y == ')' && cl == 0 {
                            cr++
                        }
                        if y == ')' && cl > 0 {
                            break
                        }
                        if y == '(' {
                            cl++
                        }
                    }
                }
            }
        }
    }
    return string(stack)
}

func main() {
    // ​​​​​​​Example 1:
    // Input: s = "(())", k = 1
    // Output: ""
    // Explanation:
    // k-balanced substring is "()"
    // Step	Current s	k-balanced	Result s
    // 1	        (())	    (())	()
    // 2	        ()	        ()	    Empty
    // Thus, the final string is "".
    fmt.Println(removeSubstring("(())", 1)) // ""
    // Example 2:
    // Input: s = "(()(", k = 1
    // Output: "(("
    // Explanation:
    // k-balanced substring is "()"
    // Step	Current s	k-balanced	Result s
    //     1	    (()(	(()(	      ((
    //     2	    ((	     -	          ((
    // Thus, the final string is "((".
    fmt.Println(removeSubstring("(()(", 1)) // "(("
    // Example 3:
    // Input: s = "((()))()()()", k = 3
    // Output: "()()()"
    // Explanation:
    // k-balanced substring is "((()))"
    // Step	Current s	    k-balanced	    Result s
    // 1	    ((()))()()()	((()))()()()	()()()
    // 2	    ()()()	        -	            ()()()
    // Thus, the final string is "()()()".   
    fmt.Println(removeSubstring("((()))()()()", 3)) // "()()()"

    fmt.Println(removeSubstring("((()))((()))", 3)) // ""
    fmt.Println(removeSubstring("()()()()()()", 3)) // "()()()()()()"
    fmt.Println(removeSubstring("((()))((()))((()))", 3)) // ""

    fmt.Println(removeSubstring1("(())", 1)) // ""
    fmt.Println(removeSubstring1("(()(", 1)) // "(("
    fmt.Println(removeSubstring1("((()))()()()", 3)) // "()()()"
    fmt.Println(removeSubstring1("((()))((()))", 3)) // ""
    fmt.Println(removeSubstring1("()()()()()()", 3)) // "()()()()()()"
    fmt.Println(removeSubstring1("((()))((()))((()))", 3)) // ""
}