package main

// 2573. Find the String with LCP
// We define the lcp matrix of any 0-indexed string word of n lowercase English letters as an n x n grid such that:
//     lcp[i][j] is equal to the length of the longest common prefix between the substrings word[i,n-1] and word[j,n-1].

// Given an n x n matrix lcp, return the alphabetically smallest string word that corresponds to lcp. 
// If there is no such string, return an empty string.

// A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ, 
// string a has a letter that appears earlier in the alphabet than the corresponding letter in b. 
// For example, "aabd" is lexicographically smaller than "aaca" because the first position they differ is at the third letter, and 'b' comes before 'c'.

// Example 1:
// Input: lcp = [[4,0,2,0],[0,3,0,1],[2,0,2,0],[0,1,0,1]]
// Output: "abab"
// Explanation: lcp corresponds to any 4 letter string with two alternating letters. The lexicographically smallest of them is "abab".

// Example 2:
// Input: lcp = [[4,3,2,1],[3,3,2,1],[2,2,2,1],[1,1,1,1]]
// Output: "aaaa"
// Explanation: lcp corresponds to any 4 letter string with a single distinct letter. The lexicographically smallest of them is "aaaa". 

// Example 3:
// Input: lcp = [[4,3,2,1],[3,3,2,1],[2,2,2,1],[1,1,1,3]]
// Output: ""
// Explanation: lcp[3][3] cannot be equal to 3 since word[3,...,3] consists of only a single letter; Thus, no answer exists.

// Constraints:
//     1 <= n == lcp.length == lcp[i].length <= 1000
//     0 <= lcp[i][j] <= n

import "fmt"
import "bytes"

func findTheString(lcp [][]int) string {
    i, n := 0, len(lcp)
    res := make([]byte, n)
    for c := 'a'; c <= 'z'; c++ {
        for i < n && res[i] != 0 {
            i++
        }
        if i == n { break }
        for j := i; j < n; j++ {
            if lcp[i][j] > 0 {
                res[j] = byte(c)
            }
        }
    }
    if bytes.IndexByte(res, 0) >= 0 { return "" }
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if res[i] == res[j] {
                if i == n - 1 || j == n - 1 {
                    if lcp[i][j] != 1 { 
                        return "" 
                    }
                } else if lcp[i][j] != lcp[i + 1][j + 1] + 1 {
                    return ""
                }
            } else if lcp[i][j] > 0 {
                return ""
            }
        }
    }
    return string(res)
}

func findTheString1(lcp [][]int) string {
    n := len(lcp)
    res := make([]byte, n)
    // 因为是字典序，所以从小到大，把能填的都填了
    for c := 'a'; c <= 'z'; c++ {
        // 找第一个还没有填的位置
        i := bytes.IndexByte(res, 0)
        if i < 0 { // 说明所有的位置都有值了
            break
        }
        for j := i; j < n; j++ {
            if lcp[i][j] > 0 {
                res[j] = byte(c)
            }
        }
    }
    for _, ch := range res {
        if ch == 0 {
            return ""
        }
    }
    // 最后进行验证
    // 如果 s[i]=s[j]，那么 lcp[i][j]=lcp[i+1][j+1]+1；
    // 如果 s[i]!=s[j]，那么 lcp[i][j]=0。
    for i := n - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            cl := 0
            if res[i] == res[j] {
                cl = 1
                if i < n-1 && j < n-1 {
                    cl = lcp[i+1][j+1] + 1
                }
            }
            if lcp[i][j] != cl {
                return ""
            }
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: lcp = [[4,0,2,0],[0,3,0,1],[2,0,2,0],[0,1,0,1]]
    // Output: "abab"
    // Explanation: lcp corresponds to any 4 letter string with two alternating letters. The lexicographically smallest of them is "abab".
    fmt.Println(findTheString([][]int{{4,0,2,0},{0,3,0,1},{2,0,2,0},{0,1,0,1}})) // "abab"
    // Example 2:
    // Input: lcp = [[4,3,2,1],[3,3,2,1],[2,2,2,1],[1,1,1,1]]
    // Output: "aaaa"
    // Explanation: lcp corresponds to any 4 letter string with a single distinct letter. The lexicographically smallest of them is "aaaa". 
    fmt.Println(findTheString([][]int{{4,3,2,1},{3,3,2,1},{2,2,2,1},{1,1,1,1}})) // "aaaa"
    // Example 3:
    // Input: lcp = [[4,3,2,1],[3,3,2,1],[2,2,2,1],[1,1,1,3]]
    // Output: ""
    // Explanation: lcp[3][3] cannot be equal to 3 since word[3,...,3] consists of only a single letter; Thus, no answer exists.
    fmt.Println(findTheString([][]int{{4,3,2,1},{3,3,2,1},{2,2,2,1},{1,1,1,3}})) // ""

    fmt.Println(findTheString1([][]int{{4,0,2,0},{0,3,0,1},{2,0,2,0},{0,1,0,1}})) // "abab"
    fmt.Println(findTheString1([][]int{{4,3,2,1},{3,3,2,1},{2,2,2,1},{1,1,1,1}})) // "aaaa"
    fmt.Println(findTheString1([][]int{{4,3,2,1},{3,3,2,1},{2,2,2,1},{1,1,1,3}})) // ""
}