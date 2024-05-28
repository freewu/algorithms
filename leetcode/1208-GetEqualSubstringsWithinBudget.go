package main

// 1208. Get Equal Substrings Within Budget 
// You are given two strings s and t of the same length and an integer maxCost.

// You want to change s to t. Changing the ith character of s to ith character of t costs |s[i] - t[i]| 
// (i.e., the absolute difference between the ASCII values of the characters).

// Return the maximum length of a substring of s 
// that can be changed to be the same as the corresponding substring of t with a cost less than or equal to maxCost. 
// If there is no substring from s that can be changed to its corresponding substring from t, return 0.

// Example 1:
// Input: s = "abcd", t = "bcdf", maxCost = 3
// Output: 3
// Explanation: "abc" of s can change to "bcd".
// That costs 3, so the maximum length is 3.

// Example 2:
// Input: s = "abcd", t = "cdef", maxCost = 3
// Output: 1
// Explanation: Each character in s costs 2 to change to character in t,  so the maximum length is 1.

// Example 3:
// Input: s = "abcd", t = "acde", maxCost = 0
// Output: 1
// Explanation: You cannot make any change, so the maximum length is 1.
 
// Constraints:
//     1 <= s.length <= 10^5
//     t.length == s.length
//     0 <= maxCost <= 10^6
//     s and t consist of only lowercase English letters.

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {
    compareList := make([]int, 0)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(s); i++ {
        compareList = append(compareList, abs(int(s[i]) - int(t[i])))
    }
    res, left, right, cost := 0, 0, 0, 0
    for right != len(s) {
        if cost + compareList[right] <= maxCost {
            cost += compareList[right]
            right++
        } else {
            cost -= compareList[left]
            left++
        }
        if right-left > res {
            res = right - left
        }
    }
    return res
}

// 滑动窗口
func equalSubstring1(s string, t string, maxCost int) int {
    costs := make([]int, len(s))
    res, cost, i, j := 0, 0, 0, 0
    for i < len(s) {
        costs[i] = int(s[i]) - int(t[i])
        if costs[i] < 0 {
            costs[i] = -costs[i]
        }
        cost += costs[i]
        for cost > maxCost && j < i {
            cost -= costs[j]
            j++
        }
        if cost <= maxCost && res < i-j + 1 {
            res = i - j + 1
        }
        i++
    }
    return res
}
  
func main() {
    // Example 1:
    // Input: s = "abcd", t = "bcdf", maxCost = 3
    // Output: 3
    // Explanation: "abc" of s can change to "bcd".
    // That costs 3, so the maximum length is 3.
    fmt.Println(equalSubstring("abcd", "bcdf", 3)) // 3
    // Example 2:
    // Input: s = "abcd", t = "cdef", maxCost = 3
    // Output: 1
    // Explanation: Each character in s costs 2 to change to character in t,  so the maximum length is 1.
    fmt.Println(equalSubstring("abcd", "cdef", 3)) // 1
    // Example 3:
    // Input: s = "abcd", t = "acde", maxCost = 0
    // Output: 1
    // Explanation: You cannot make any change, so the maximum length is 1.
    fmt.Println(equalSubstring("abcd", "acde", 0)) // 1

    fmt.Println(equalSubstring1("abcd", "bcdf", 3)) // 3
    fmt.Println(equalSubstring1("abcd", "cdef", 3)) // 1
    fmt.Println(equalSubstring1("abcd", "acde", 0)) // 1
}