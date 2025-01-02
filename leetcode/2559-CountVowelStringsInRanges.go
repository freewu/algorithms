package main

// 2559. Count Vowel Strings in Ranges
// You are given a 0-indexed array of strings words and a 2D array of integers queries.

// Each query queries[i] = [li, ri] asks us to find the number of strings present in the range li to ri (both inclusive) of words that start and end with a vowel.

// Return an array ans of size queries.length, where ans[i] is the answer to the ith query.

// Note that the vowel letters are 'a', 'e', 'i', 'o', and 'u'.

// Example 1:
// Input: words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]
// Output: [2,3,0]
// Explanation: The strings starting and ending with a vowel are "aba", "ece", "aa" and "e".
// The answer to the query [0,2] is 2 (strings "aba" and "ece").
// to query [1,4] is 3 (strings "ece", "aa", "e").
// to query [1,1] is 0.
// We return [2,3,0].

// Example 2:
// Input: words = ["a","e","i"], queries = [[0,2],[0,1],[2,2]]
// Output: [3,2,1]
// Explanation: Every string satisfies the conditions, so we return [3,2,1].

// Constraints:
//     1 <= words.length <= 10^5
//     1 <= words[i].length <= 40
//     words[i] consists only of lowercase English letters.
//     sum(words[i].length) <= 3 * 10^5
//     1 <= queries.length <= 10^5
//     0 <= li <= ri < words.length

import "fmt"
import "sort"

func vowelStrings(words []string, queries [][]int) []int {
    vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
    count := []int{}
    for i, w := range words {
        if vowels[w[0]] && vowels[w[len(w)-1]] {
            count = append(count, i)
        }
    }
    res := make([]int, len(queries))
    for i, q := range queries {
        l, r := q[0], q[1]
        res[i] = sort.SearchInts(count, r + 1) - sort.SearchInts(count, l)
    }
    return res
}

func vowelStrings1(words []string, queries [][]int) []int {
    n, m := len(queries), len(words) + 1
    res, sum := make([]int, n), make([]int, m)
    check := func(s string) bool {
        n := len(s)
        if s[0] != 'a' && s[0] != 'e' && s[0] != 'i' && s[0] != 'o' && s[0] != 'u' { return false }
        if s[n-1] != 'a' && s[n-1] != 'e' && s[n-1] != 'i' && s[n-1] != 'o' && s[n-1] != 'u' { return false }
        return true
    }
    for i := 1; i < m; i++ {
        if check(words[i-1]) {
            sum[i] = sum[i-1] + 1
        } else {
            sum[i] = sum[i-1]
        }
    }
    for i := 0; i < n; i++ {
        res[i] = sum[queries[i][1] + 1] - sum[queries[i][0]]
    }
    return res
}



func vowelStrings2(words []string, queries [][]int) []int {
    n := len(words)
    prefix := make([]int, n)
    isVowel := func(c byte) bool { return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' }
    for i, word := range(words) {
        prefix[i] = 0
        if i > 0 { prefix[i] = prefix[i - 1]  }
        if isVowel(word[0]) && isVowel(word[len(word) - 1]) {
            prefix[i]++
        }
    }
    res := make([]int, len(queries))
    for i, query := range(queries) {
        l, r := query[0], query[1]
        if l == 0 {
            res[i] = prefix[r]
            continue
        }
        res[i] = prefix[r] - prefix[l - 1]
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]
    // Output: [2,3,0]
    // Explanation: The strings starting and ending with a vowel are "aba", "ece", "aa" and "e".
    // The answer to the query [0,2] is 2 (strings "aba" and "ece").
    // to query [1,4] is 3 (strings "ece", "aa", "e").
    // to query [1,1] is 0.
    // We return [2,3,0].
    fmt.Println(vowelStrings([]string{"aba","bcb","ece","aa","e"}, [][]int{{0,2},{1,4},{1,1}})) // [2,3,0]
    // Example 2:
    // Input: words = ["a","e","i"], queries = [[0,2],[0,1],[2,2]]
    // Output: [3,2,1]
    // Explanation: Every string satisfies the conditions, so we return [3,2,1].
    fmt.Println(vowelStrings([]string{"a","e","i"}, [][]int{{0,2},{0,1},{2,2}})) // [3,2,1]

    fmt.Println(vowelStrings1([]string{"aba","bcb","ece","aa","e"}, [][]int{{0,2},{1,4},{1,1}})) // [2,3,0]
    fmt.Println(vowelStrings1([]string{"a","e","i"}, [][]int{{0,2},{0,1},{2,2}})) // [3,2,1]

    fmt.Println(vowelStrings2([]string{"aba","bcb","ece","aa","e"}, [][]int{{0,2},{1,4},{1,1}})) // [2,3,0]
    fmt.Println(vowelStrings2([]string{"a","e","i"}, [][]int{{0,2},{0,1},{2,2}})) // [3,2,1]
}