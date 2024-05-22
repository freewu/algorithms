package main

// 336. Palindrome Pairs
// You are given a 0-indexed array of unique strings words.
// A palindrome pair is a pair of integers (i, j) such that:
//     0 <= i, j < words.length,
//     i != j, and
//     words[i] + words[j] (the concatenation of the two strings) is a palindrome.

// Return an array of all the palindrome pairs of words.
// You must write an algorithm with O(sum of words[i].length) runtime complexity.

// Example 1:
// Input: words = ["abcd","dcba","lls","s","sssll"]
// Output: [[0,1],[1,0],[3,2],[2,4]]
// Explanation: The palindromes are ["abcddcba","dcbaabcd","slls","llssssll"]

// Example 2:
// Input: words = ["bat","tab","cat"]
// Output: [[0,1],[1,0]]
// Explanation: The palindromes are ["battab","tabbat"]

// Example 3:
// Input: words = ["a",""]
// Output: [[0,1],[1,0]]
// Explanation: The palindromes are ["a","a"]

// Constraints:
//     1 <= words.length <= 5000
//     0 <= words[i].length <= 300
//     words[i] consists of lowercase English letters.

import "fmt"

// 暴力 超出时间限制 134 / 136 
func palindromePairs1(words []string) [][]int {
    res := [][]int{}
    isPalindromic := func(str string) bool { // 判断是否是回文
        if len(str) < 2 { // 只有一个字符
            return true
        }
        l, r := 0, len(str) - 1
        for l < r {
            if str[l] != str[r] {
                return false
            }
            l, r = l + 1, r - 1
        }
        return true
    }
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i != j && isPalindromic(words[i] + words[j]) {
                res = append(res, []int{i, j})
            }
        }
    }
    return res
}

func palindromePairs(words []string) [][]int {
    type Pair struct {
        a, b int
    }
    word_dict, pairs, res := map[string]int{}, map[Pair]bool{}, [][]int{}
    for i,word :=  range(words) {
        word_dict[word] = i
    }
    reverse := func (str string) string {
        res := ""
        for _, v := range str {
            res = string(v) + res
        }
        return res
    }
    isPalindromic := func(str string) bool { // 判断是否是回文
        l, r := 0, len(str) - 1
        for l < r {
            if str[l] != str[r] {
                return false
            }
            l, r = l + 1, r - 1
        }
        return true
    }
    for i,word := range(words) {
        reverse_word := reverse(word)
        length := len(word)
        for idx := 0; idx < length + 1; idx++ {
            if j,ok := word_dict[reverse_word[:idx]]; ok && j != i && isPalindromic(reverse_word[idx:]) {
                pairs[Pair{j,i}] = true
            }
            if j,ok := word_dict[reverse_word[idx:]]; ok && j != i && isPalindromic(reverse_word[:idx]) {
                pairs[Pair{i,j}] = true
            }
        }
    }
    for key,_ := range(pairs){
        res = append(res,[]int{key.a,key.b})
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["abcd","dcba","lls","s","sssll"]
    // Output: [[0,1],[1,0],[3,2],[2,4]]
    // Explanation: The palindromes are ["abcddcba","dcbaabcd","slls","llssssll"]
    fmt.Println(palindromePairs([]string{"abcd","dcba","lls","s","sssll"})) // [[0,1],[1,0],[3,2],[2,4]]
    // Example 2:
    // Input: words = ["bat","tab","cat"]
    // Output: [[0,1],[1,0]]
    // Explanation: The palindromes are ["battab","tabbat"]
    fmt.Println(palindromePairs([]string{"bat","tab","cat"})) // [[0,1],[1,0]]
    // Example 3:
    // Input: words = ["a",""]
    // Output: [[0,1],[1,0]]
    // Explanation: The palindromes are ["a","a"]
    fmt.Println(palindromePairs([]string{"a",""})) // [[0,1],[1,0]]

    fmt.Println(palindromePairs1([]string{"abcd","dcba","lls","s","sssll"})) // [[0,1],[1,0],[3,2],[2,4]]
    fmt.Println(palindromePairs1([]string{"bat","tab","cat"})) // [[0,1],[1,0]]
    fmt.Println(palindromePairs1([]string{"a",""})) // [[0,1],[1,0]]
}