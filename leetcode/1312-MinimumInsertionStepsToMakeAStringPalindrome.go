package main

// 1312. Minimum Insertion Steps to Make a String Palindrome
// Given a string s. In one step you can insert any character at any index of the string.
// Return the minimum number of steps to make s palindrome.
// A Palindrome String is one that reads the same backward as well as forward.

// Example 1:
// Input: s = "zzazz"
// Output: 0
// Explanation: The string "zzazz" is already palindrome we do not need any insertions.

// Example 2:
// Input: s = "mbadm"
// Output: 2
// Explanation: String can be "mbdadbm" or "mdbabdm".

// Example 3:
// Input: s = "leetcode"
// Output: 5
// Explanation: Inserting 5 characters the string becomes "leetcodocteel".
 
// Constraints:
//     1 <= s.length <= 500
//     s consists of lowercase English letters

import "fmt"

func minInsertions(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 2; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1]
            } else { // 如果不等于需要调整1步，取最少的步骤 + 1
                dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1 
            }
        }
    }
    return dp[0][n-1]
}

func minInsertions1(s string) int {
    reverseString := func(target string) string{
        arr := []rune(target)
        i, j := 0, len(target) - 1
        for i < j  {
            arr[i], arr[j] = arr[j], arr[i]
            i++
            j--
        }
        return string(arr)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    n, rs := len(s), reverseString(s)
    // lcs
    dp :=  make([]int, n + 1) 
    for i := 1; i <= n; i++ {
        // when go to first i subchars of s
        // look into first j subchars of rs 
        prev := 0
        for j := 1; j <= n; j++ {
            // this round dp[i][j]
            tmp := dp[j] 
            if s[i - 1] == rs[j - 1] {
                // this round dp[i][j] = dp[i - 1][j - 1] + 1
                dp[j] = prev + 1 
            } else {
                // this round dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
                dp[j] = max(dp[j], dp[j - 1]) 
            }
            // prepare next round dp[i - 1][j - 1]
            prev = tmp
        }
    }
    return n - dp[n]
}

func minInsertions2(s string) int {
    n := len(s)
    if n < 2 { return 0 }
    dp, bs := make([]int, n), []byte(s)
    pre, temp := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 2; i >= 0; i-- {
        pre = 0
        for j := i + 1; j < n; j++ {
            temp = dp[j]
            if bs[i] == bs[j] {
                dp[j] = pre
            } else {
                dp[j] = min(dp[j], dp[j-1]) + 1
            }
            pre = temp
        }
    }
    return dp[n-1]
}

func main() {
    // Example 1:
    // Input: s = "zzazz"
    // Output: 0
    // Explanation: The string "zzazz" is already palindrome we do not need any insertions.
    fmt.Println(minInsertions("zzazz")) // 0
    // Example 2:
    // Input: s = "mbadm"
    // Output: 2
    // Explanation: String can be "mbdadbm" or "mdbabdm".
    fmt.Println(minInsertions("mbadm")) // 2
    // Example 3:
    // Input: s = "leetcode"
    // Output: 5
    // Explanation: Inserting 5 characters the string becomes "leetcodocteel".
    // Explanation: The string "zzazz" is already palindrome we do not need any insertions.
    fmt.Println(minInsertions("leetcode")) // 5

    fmt.Println(minInsertions("bluefrog")) // 7

    fmt.Println(minInsertions1("zzazz")) // 0
    fmt.Println(minInsertions1("mbadm")) // 2
    fmt.Println(minInsertions1("leetcode")) // 5
    fmt.Println(minInsertions1("bluefrog")) // 7

    fmt.Println(minInsertions2("zzazz")) // 0
    fmt.Println(minInsertions2("mbadm")) // 2
    fmt.Println(minInsertions2("leetcode")) // 5
    fmt.Println(minInsertions2("bluefrog")) // 7
}