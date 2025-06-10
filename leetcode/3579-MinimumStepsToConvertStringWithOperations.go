package main

// 3579. Minimum Steps to Convert String with Operations
// You are given two strings, word1 and word2, of equal length. 
// You need to transform word1 into word2.

// For this, divide word1 into one or more contiguous substrings. 
// For each substring substr you can perform the following operations:

//     1. Replace: Replace the character at any one index of substr with another lowercase English letter.

//     2. Swap: Swap any two characters in substr.

//     3. Reverse Substring: Reverse substr.

// Each of these counts as one operation 
// and each character of each substring can be used in each type of operation at most once (i.e. no single index may be involved in more than one replace, one swap, or one reverse).

// Return the minimum number of operations required to transform word1 into word2.

// Example 1:
// Input: word1 = "abcdf", word2 = "dacbe"
// Output: 4
// Explanation:
// Divide word1 into "ab", "c", and "df". The operations are:
// For the substring "ab",
// Perform operation of type 3 on "ab" -> "ba".
// Perform operation of type 1 on "ba" -> "da".
// For the substring "c" do no operations.
// For the substring "df",
// Perform operation of type 1 on "df" -> "bf".
// Perform operation of type 1 on "bf" -> "be".

// Example 2:
// Input: word1 = "abceded", word2 = "baecfef"
// Output: 4
// Explanation:
// Divide word1 into "ab", "ce", and "ded". The operations are:
// For the substring "ab",
// Perform operation of type 2 on "ab" -> "ba".
// For the substring "ce",
// Perform operation of type 2 on "ce" -> "ec".
// For the substring "ded",
// Perform operation of type 1 on "ded" -> "fed".
// Perform operation of type 1 on "fed" -> "fef".

// Example 3:
// Input: word1 = "abcdef", word2 = "fedabc"
// Output: 2
// Explanation:
// Divide word1 into "abcdef". The operations are:
// For the substring "abcdef",
// Perform operation of type 3 on "abcdef" -> "fedcba".
// Perform operation of type 2 on "fedcba" -> "fedabc".

// Constraints:
//     1 <= word1.length == word2.length <= 100
//     word1 and word2 consist only of lowercase English letters.

import "fmt"

func minOperations(word1, word2 string) int {
    count := [26][26]int{}
    op, n := 0, len(word1)
    update := func(x, y byte) {
        if x == y { return }
        x -= 'a'
        y -= 'a'
        if count[y][x] > 0 {
            count[y][x]--
        } else {
            count[x][y]++
            op++
        }
    }
    // 预处理 revOp
    revOp := make([][]int, n)
    for i := range revOp {
        revOp[i] = make([]int, n)
    }
    // 中心扩展法
    // i 为偶数表示奇长度子串，i 为奇数表示偶长度子串
    for i := 0; i < 2 * n - 1; i++ {
        op, count = 1, [26][26]int{}
        // 从闭区间 [l,r] 开始向左右扩展
        l, r := i / 2, (i+1) / 2
        for l >= 0 && r < n {
            update(word1[l], word2[r])
            if l != r {
                update(word1[r], word2[l])
            }
            revOp[l][r] = op
            l--
            r++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    f := make([]int, n + 1)
    for i := 0; i < n; i++ {
        res := 1 << 31
        op, count = 0, [26][26]int{} // 不反转时的最小操作次数
        for j := i; j >= 0; j-- {
            update(word1[j], word2[j])
            res = min(res, f[j] + min(op, revOp[j][i]))
        }
        f[i + 1] = res
    }
    return f[n]
}

func main() {
    // Example 1:
    // Input: word1 = "abcdf", word2 = "dacbe"
    // Output: 4
    // Explanation:
    // Divide word1 into "ab", "c", and "df". The operations are:
    // For the substring "ab",
    // Perform operation of type 3 on "ab" -> "ba".
    // Perform operation of type 1 on "ba" -> "da".
    // For the substring "c" do no operations.
    // For the substring "df",
    // Perform operation of type 1 on "df" -> "bf".
    // Perform operation of type 1 on "bf" -> "be".
    fmt.Println(minOperations("abcdf", "dacbe")) // 4
    // Example 2:
    // Input: word1 = "abceded", word2 = "baecfef"
    // Output: 4
    // Explanation:
    // Divide word1 into "ab", "ce", and "ded". The operations are:
    // For the substring "ab",
    // Perform operation of type 2 on "ab" -> "ba".
    // For the substring "ce",
    // Perform operation of type 2 on "ce" -> "ec".
    // For the substring "ded",
    // Perform operation of type 1 on "ded" -> "fed".
    // Perform operation of type 1 on "fed" -> "fef".
    fmt.Println(minOperations("abceded", "baecfef")) // 4
    // Example 3:
    // Input: word1 = "abcdef", word2 = "fedabc"
    // Output: 2
    // Explanation:
    // Divide word1 into "abcdef". The operations are:
    // For the substring "abcdef",
    // Perform operation of type 3 on "abcdef" -> "fedcba".
    // Perform operation of type 2 on "fedcba" -> "fedabc".
    fmt.Println(minOperations("abcdef", "fedabc")) // 2

    fmt.Println(minOperations("leetcode", "bluefrog")) // 8
    fmt.Println(minOperations("bluefrog", "leetcode")) // 8
}