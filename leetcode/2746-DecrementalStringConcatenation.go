package main

// 2746. Decremental String Concatenation
// You are given a 0-indexed array words containing n strings.

// Let's define a join operation join(x, y) between two strings x and y as concatenating them into xy. 
// However, if the last character of x is equal to the first character of y, one of them is deleted.

// For example join("ab", "ba") = "aba" and join("ab", "cde") = "abcde".

// You are to perform n - 1 join operations. Let str0 = words[0]. 
// Starting from i = 1 up to i = n - 1, for the ith operation, you can do one of the following:
//     1. Make stri = join(stri - 1, words[i])
//     2. Make stri = join(words[i], stri - 1)

// Your task is to minimize the length of strn - 1.

// Return an integer denoting the minimum possible length of strn - 1.

// Example 1:
// Input: words = ["aa","ab","bc"]
// Output: 4
// Explanation: In this example, we can perform join operations in the following order to minimize the length of str2: 
// str0 = "aa"
// str1 = join(str0, "ab") = "aab"
// str2 = join(str1, "bc") = "aabc" 
// It can be shown that the minimum possible length of str2 is 4.

// Example 2:
// Input: words = ["ab","b"]
// Output: 2
// Explanation: In this example, str0 = "ab", there are two ways to get str1: 
// join(str0, "b") = "ab" or join("b", str0) = "bab". 
// The first string, "ab", has the minimum length. Hence, the answer is 2.

// Example 3:
// Input: words = ["aaa","c","aba"]
// Output: 6
// Explanation: In this example, we can perform join operations in the following order to minimize the length of str2: 
// str0 = "aaa"
// str1 = join(str0, "c") = "aaac"
// str2 = join("aba", str1) = "abaaac"
// It can be shown that the minimum possible length of str2 is 6.

// Constraints:
//     1 <= words.length <= 1000
//     1 <= words[i].length <= 50
//     Each character in words[i] is an English lowercase letter

import "fmt"

func minimizeConcatenatedLength(words []string) int {
    type Group struct {
        index int
        left byte
        right byte
    }
    memo := make(map[Group]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(index int, left ,right byte) int
    dfs = func(index int, left ,right byte) int {
        if index == len(words) { return 0 }
        if v, ok := memo[Group{ index, left, right }]; ok { return v }
        res := 1 << 31
        if words[index][len(words[index]) - 1] == left {
            res = min(res, dfs(index + 1, words[index][0], right) + (len(words[index]) - 1))
        } else {
            res = min(res, dfs(index + 1, words[index][0], right) + len(words[index]))
        }
        if words[index][0] == right {
            res = min(res, dfs(index + 1,left, words[index][len(words[index]) - 1]) + (len(words[index]) - 1))
        }else{
            res = min(res, dfs(index + 1,left, words[index][len(words[index]) - 1]) + len(words[index]))
        }
        memo[Group{ index, left, right }] = res
        return res
    }
    return dfs(1,words[0][0],words[0][len(words[0]) - 1]) + len(words[0])
}

func minimizeConcatenatedLength1(words []string) int {
    n := len(words)
    memo := map[int]int{}
    var dfs func(index int, b, e byte) int
    dfs = func(index int, b, e byte) int {
        if index == n { return 0 }
        key := index * 10000 + int(b) * 100 + int(e)
        if v, ok := memo[key]; ok { return v }
        str := words[index]
        n := len(str)
        b1, e1, mn, mm  := str[0] - 'a', str[n - 1] - 'a', 0, 0
        if b == e1 { mn = 1 }
        if e == b1 { mm = 1 }
        res := dfs(index + 1, b1, e) + n - mn
        res = min(res, dfs(index + 1, b, e1) + n - mm)
        memo[key] = res
        return res
    }
    start := words[0]
    return dfs(1, start[0] - 'a', start[len(start) - 1] - 'a') + len(start)
}

func main() {
    // Example 1:
    // Input: words = ["aa","ab","bc"]
    // Output: 4
    // Explanation: In this example, we can perform join operations in the following order to minimize the length of str2: 
    // str0 = "aa"
    // str1 = join(str0, "ab") = "aab"
    // str2 = join(str1, "bc") = "aabc" 
    // It can be shown that the minimum possible length of str2 is 4.
    fmt.Println(minimizeConcatenatedLength([]string{"aa","ab","bc"})) // 4
    // Example 2:
    // Input: words = ["ab","b"]
    // Output: 2
    // Explanation: In this example, str0 = "ab", there are two ways to get str1: 
    // join(str0, "b") = "ab" or join("b", str0) = "bab". 
    // The first string, "ab", has the minimum length. Hence, the answer is 2.
    fmt.Println(minimizeConcatenatedLength([]string{"ab","b"})) // 2
    // Example 3:
    // Input: words = ["aaa","c","aba"]
    // Output: 6
    // Explanation: In this example, we can perform join operations in the following order to minimize the length of str2: 
    // str0 = "aaa"
    // str1 = join(str0, "c") = "aaac"
    // str2 = join("aba", str1) = "abaaac"
    // It can be shown that the minimum possible length of str2 is 6.
    fmt.Println(minimizeConcatenatedLength([]string{"aaa","c","aba"})) // 6

    fmt.Println(minimizeConcatenatedLength([]string{"bluefrog","leetcode"})) // 16

    fmt.Println(minimizeConcatenatedLength1([]string{"aa","ab","bc"})) // 4
    fmt.Println(minimizeConcatenatedLength1([]string{"ab","b"})) // 2
    fmt.Println(minimizeConcatenatedLength1([]string{"aaa","c","aba"})) // 6
    fmt.Println(minimizeConcatenatedLength1([]string{"bluefrog","leetcode"})) // 16
}